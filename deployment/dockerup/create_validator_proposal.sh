#!/bin/bash

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)

source $SCRIPT_DIR/.env

TEMPLATE_FILE="$SCRIPT_DIR/create_validator_proposal.json"
OUTPUT_FILE="$SCRIPT_DIR/proposal.json"
MECHAIND_CMD="mechaind"

function generate() {
    home=$1
    echo "add keys..."
    $MECHAIND_CMD keys add validator --keyring-backend test --home "$home" >${home}/validator_info 2>&1
    $MECHAIND_CMD keys add delegator --keyring-backend test --home "$home" >${home}/delegator_info 2>&1
    $MECHAIND_CMD keys add validator_bls --keyring-backend test --home "$home" --algo eth_bls >${home}/bls_info 2>&1
    $MECHAIND_CMD keys add validator_relayer --keyring-backend test --home "$home" >${home}/relayer_info 2>&1
    $MECHAIND_CMD keys add validator_challenger --keyring-backend test --home "$home" >${home}/challenger_info 2>&1

    echo "show keys..."
    VALIDATOR_ADDR=$($MECHAIND_CMD keys show validator -a --keyring-backend test --home "$home")
    DELEGATOR_ADDR=$($MECHAIND_CMD keys show delegator -a --keyring-backend test --home "$home")
    RELAYER_ADDR=$($MECHAIND_CMD keys show validator_relayer -a --keyring-backend test --home "$home")
    CHALLENGER_ADDR=$($MECHAIND_CMD keys show validator_challenger -a --keyring-backend test --home "$home")
    VALIDATOR_BLS=$($MECHAIND_CMD keys show validator_bls --keyring-backend test --home "$home" --output json | jq -r '.pubkey_hex')
    VALIDATOR_BLS_PROOF=$($MECHAIND_CMD keys sign ${VALIDATOR_BLS} --keyring-backend test --home "$home" --from validator_bls)
    VALIDATOR_NODE_PUB_KEY=$(cat ${home}/config/priv_validator_key.json | jq -r '.pub_key.value')

    if [ -z "$VALIDATOR_ADDR" ] || [ -z "$DELEGATOR_ADDR" ] || [ -z "$RELAYER_ADDR" ] || [ -z "$CHALLENGER_ADDR" ] || [ -z "$VALIDATOR_BLS" ] || [ -z "$VALIDATOR_BLS_PROOF" ] || [ -z "$VALIDATOR_NODE_PUB_KEY" ]; then
        echo "Error: Failed to generate necessary keys or addresses."
        exit 1
    fi

    echo "generated validator proposal file..."
    sed -e "s|\${MONIKER_NAME}|$MONIKER_NAME|g" \
        -e "s|\${VALIDATOR_NODE_PUB_KEY}|$VALIDATOR_NODE_PUB_KEY|g" \
        -e "s|\${VALIDATOR_ADDR}|$VALIDATOR_ADDR|g" \
        -e "s|\${DELEGATOR_ADDR}|$DELEGATOR_ADDR|g" \
        -e "s|\${VALIDATOR_BLS}|$VALIDATOR_BLS|g" \
        -e "s|\${VALIDATOR_BLS_PROOF}|$VALIDATOR_BLS_PROOF|g" \
        -e "s|\${RELAYER_ADDR}|$RELAYER_ADDR|g" \
        -e "s|\${CHALLENGER_ADDR}|$CHALLENGER_ADDR|g" \
        "$TEMPLATE_FILE" >"$OUTPUT_FILE"

    if [ $? -eq 0 ]; then
        echo "create_validator_proposal.json has been generated successfully at $OUTPUT_FILE."
    else
        echo "Error: Failed to create create_validator_proposal.json."
    fi
    echo VALIDATOR_ADDR: $VALIDATOR_ADDR
    echo DELEGATOR_ADDR: $DELEGATOR_ADDR

    echo "send tokens..."
    echo '#!/bin/bash' >${SCRIPT_DIR}/send_tokens.sh
    echo "docker exec mechaind-validator-0 mechaind tx bank send validator0 $VALIDATOR_ADDR 10000000000000000000000000azkme --keyring-backend test --node http://localhost:26657 -y --fees 6000000azkme" >>${SCRIPT_DIR}/send_tokens.sh
    echo "docker exec mechaind-validator-0 mechaind tx bank send validator0 $DELEGATOR_ADDR 10000000000000000000000000azkme --keyring-backend test --node http://localhost:26657 -y --fees 6000000azkme" >>${SCRIPT_DIR}/send_tokens.sh
}

function balance() {
    home=$1
    VALIDATOR_ADDR=$($MECHAIND_CMD keys show validator -a --keyring-backend test --home "$home")
    echo "validator($VALIDATOR_ADDR) balance..."
    $MECHAIND_CMD query bank balances $VALIDATOR_ADDR --home $home
    DELEGATOR_ADDR=$($MECHAIND_CMD keys show delegator -a --keyring-backend test --home "$home")
    echo "delegator($DELEGATOR_ADDR) balance..."
    $MECHAIND_CMD query bank balances $DELEGATOR_ADDR --home $home
}

function grant() {
    echo "grant..."
    home=$1
    VALIDATOR_ADDR=$($MECHAIND_CMD keys show validator -a --keyring-backend test --home "$home")
    DELEGATOR_ADDR=$($MECHAIND_CMD keys show delegator -a --keyring-backend test --home "$home")
    $MECHAIND_CMD tx authz grant 0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2 generic \
        --msg-type=/cosmos.staking.v1beta1.MsgDelegate --gas="600000" --gas-prices="10000000000azkme" \
        --from=${DELEGATOR_ADDR} --home=$home --keyring-backend=test --broadcast-mode sync -y
}

function proposal() {
    echo "create proposal..."
    home=$1
    VALIDATOR_ADDR=$($MECHAIND_CMD keys show validator -a --keyring-backend test --home "$home")
    DELEGATOR_ADDR=$($MECHAIND_CMD keys show delegator -a --keyring-backend test --home "$home")
    # echo $MECHAIND_CMD tx gov submit-proposal $OUTPUT_FILE --gas="600000" --gas-prices="10000000000azkme" \
    # --from=${DELEGATOR_ADDR} --home=$home --keyring-backend=test --broadcast-mode sync -y
    $MECHAIND_CMD tx gov submit-proposal $OUTPUT_FILE --gas="600000" --gas-prices="10000000000azkme" \
        --from=${VALIDATOR_ADDR} --home=$home --keyring-backend=test --broadcast-mode sync -y
}

function create() {
    echo "create proposal..."
    home=$1
    VALIDATOR_ADDR=$($MECHAIND_CMD keys show validator -a --keyring-backend test --home "$home")
    DELEGATOR_ADDR=$($MECHAIND_CMD keys show delegator -a --keyring-backend test --home "$home")
    $MECHAIND_CMD tx staking create-validator $OUTPUT_FILE --home $home --keyring-backend test \
        --chain-id $CHAIN_ID --from ${DELEGATOR_ADDR} --node tcp://localhost:26657 -b sync \
        --gas "200000000" --fees "100000000000000000000azkme" --yes >${home}/create.log 2>&1
}

function query_proposal() {
    echo "query proposal..."
    tx_hash=$(grep 'txhash' ${home}/create.log | awk '{print $2}')
    $MECHAIND_CMD q tx $tx_hash --home /app --output json >${SCRIPT_DIR}/tx.json
    proposal_id=$(jq '.logs[] | .events[] | select(.type == "submit_proposal") | .attributes[] | select(.key == "proposal_id") | .value | tonumber' ${SCRIPT_DIR}/tx.json)
    echo "docker exec mechaind-validator-0curl -s http://127.0.0.1:1317/cosmos/gov/v1/proposals/$proposal_id | jq"
}

function clean() {
    rm $OUTPUT_FILE
    rm -r $home/keyring-test
}

function vote() {
    size=$1
    vote_id=$2
    for ((i = 0; i < ${size}; i++)); do
        docker exec mechaind-validator-${i} mechaind tx gov vote $vote_id yes --from=validator$i --chain-id="mechain_5151-1" --keyring-backend=test --gas-prices=10000azkme -y
    done
}

function show_validator() {
    docker exec mechaind-validator-0 curl -s http://localhost:1317/cosmos/staking/v1beta1/validators | jq '.pagination.total'
}

CMD=$1
home=$2

case ${CMD} in
generate)
    echo "===== generate ===="
    clean $home
    generate $home
    echo "===== end ===="
    ;;
balance)
    echo "===== balance ===="
    balance $home
    echo "===== end ===="
    ;;
grant)
    echo "===== grant ===="
    grant $home
    echo "===== end ===="
    ;;
proposal)
    echo "===== init ===="
    proposal $home
    echo "===== end ===="
    ;;
create)
    echo "===== create ===="
    create $home
    echo "===== end ===="
    ;;
query_proposal)
    echo "===== query proposal ===="
    query_proposal
    echo "===== end ===="
    ;;
clean)
    echo "===== clean ===="
    clean
    echo "===== end ===="
    ;;
vote)
    echo "===== vote ===="
    vote $2 $3
    echo "===== end ===="
    ;;
show_validator)
    echo "===== show_validator ===="
    show_validator
    echo "===== end ===="
    ;;
*)
    echo "Usage: $0 proposal | vote | clean | config"
    ;;
esac
