#!/usr/bin/env bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}
source ${workspace}/.env
source ${workspace}/utils.sh

bin_name=mechaind
bin=${workspace}/../../build/${bin_name}

function init() {
    size=$1
    rm -rf ${workspace}/.local
    mkdir -p ${workspace}/.local
    for ((i=0;i<${size};i++));do
        mkdir -p ${workspace}/.local/validator${i}
        mkdir -p ${workspace}/.local/relayer${i}
        mkdir -p ${workspace}/.local/challenger${i}

        # init chain
        ${bin} init validator${i} --chain-id ${CHAIN_ID} --default-denom ${STAKING_BOND_DENOM} --home ${workspace}/.local/validator${i}

        # create genesis accounts
        ${bin} keys add validator${i} --keyring-backend test --home ${workspace}/.local/validator${i} > ${workspace}/.local/validator${i}/info 2>&1
        ${bin} keys add validator_delegator${i} --keyring-backend test --home ${workspace}/.local/validator${i} > ${workspace}/.local/validator${i}/delegator_info 2>&1
        ${bin} keys add validator_bls${i} --keyring-backend test --home ${workspace}/.local/validator${i} --algo eth_bls > ${workspace}/.local/validator${i}/bls_info 2>&1
        ${bin} keys add relayer${i} --keyring-backend test --home ${workspace}/.local/relayer${i} > ${workspace}/.local/relayer${i}/relayer_info 2>&1
        ${bin} keys add challenger${i} --keyring-backend test --home ${workspace}/.local/challenger${i} > ${workspace}/.local/challenger${i}/challenger_info 2>&1
    done
}

function generate_genesis() {
    size=$1

    declare -a validator_addrs=()
    for ((i=0;i<${size};i++));do
        # export validator addresses
        validator_addrs+=("$(${bin} keys show validator${i} -a --keyring-backend test --home ${workspace}/.local/validator${i})")
    done

    declare -a deletgator_addrs=()
    for ((i=0;i<${size};i++));do
        # export delegator addresses
        deletgator_addrs+=("$(${bin} keys show validator_delegator${i} -a --keyring-backend test --home ${workspace}/.local/validator${i})")
    done

    declare -a relayer_addrs=()
    for ((i=0;i<${size};i++));do
        # export validator addresses
        relayer_addrs+=("$(${bin} keys show relayer${i} -a --keyring-backend test --home ${workspace}/.local/relayer${i})")
    done

    declare -a challenger_addrs=()
    for ((i=0;i<${size};i++));do
        # export validator addresses
        challenger_addrs+=("$(${bin} keys show challenger${i} -a --keyring-backend test --home ${workspace}/.local/challenger${i})")
    done

    mkdir -p ${workspace}/.local/gentx
    for ((i=0;i<${size};i++));do
        for validator_addr in "${validator_addrs[@]}";do
            # init genesis account in genesis state
            ${bin} add-genesis-account $validator_addr ${GENESIS_ACCOUNT_BALANCE}${STAKING_BOND_DENOM} --home ${workspace}/.local/validator${i}
        done

        for deletgator_addr in "${deletgator_addrs[@]}";do
            # init genesis account in genesis state
            ${bin} add-genesis-account $deletgator_addr ${GENESIS_ACCOUNT_BALANCE}${STAKING_BOND_DENOM} --home ${workspace}/.local/validator${i}
        done

        for relayer_addr in "${relayer_addrs[@]}";do
            # init genesis account in genesis state
            ${bin} add-genesis-account $relayer_addr ${GENESIS_ACCOUNT_BALANCE}${STAKING_BOND_DENOM} --home ${workspace}/.local/validator${i}
        done

        for challenger_addr in "${challenger_addrs[@]}";do
            # init genesis account in genesis state
            ${bin} add-genesis-account $challenger_addr ${GENESIS_ACCOUNT_BALANCE}${STAKING_BOND_DENOM} --home ${workspace}/.local/validator${i}
        done

        rm -rf ${workspace}/.local/validator${i}/config/gentx/

        validatorAddr=${validator_addrs[$i]}
        deletgatorAddr=${deletgator_addrs[$i]}
        relayerAddr="$(${bin} keys show relayer${i} -a --keyring-backend test --home ${workspace}/.local/relayer${i})"
        challengerAddr="$(${bin} keys show challenger${i} -a --keyring-backend test --home ${workspace}/.local/challenger${i})"
        blsKey="$(${bin} keys show validator_bls${i} --keyring-backend test --home ${workspace}/.local/validator${i} --output json | jq -r .pubkey_hex)"
        blsProof="$(${bin} keys sign "${blsKey}" --from validator_bls${i} --keyring-backend test --home ${workspace}/.local/validator${i})"

        # create bond validator tx
        ${bin} gentx "${STAKING_BOND_AMOUNT}${STAKING_BOND_DENOM}" "$validatorAddr" "$deletgatorAddr" "$relayerAddr" "$challengerAddr" "$blsKey" "$blsProof" \
            --home ${workspace}/.local/validator${i} \
            --keyring-backend=test \
            --chain-id=${CHAIN_ID} \
            --moniker="validator${i}" \
            --commission-max-change-rate=${COMMISSION_MAX_CHANGE_RATE} \
            --commission-max-rate=${COMMISSION_MAX_RATE} \
            --commission-rate=${COMMISSION_RATE} \
            --details="validator${i}" \
            --website="http://website" \
            --node tcp://localhost:$((${VALIDATOR_RPC_PORT_START}+${i})) \
            --node-id "validator${i}" \
            --ip 127.0.0.1 \
            --gas ""
        cp ${workspace}/.local/validator${i}/config/gentx/gentx-validator${i}.json ${workspace}/.local/gentx/
    done

    node_ids=""
    # bond validator tx in genesis state
    for ((i=0;i<${size};i++));do
        cp ${workspace}/.local/gentx/* ${workspace}/.local/validator${i}/config/gentx/
        ${bin} collect-gentxs --home ${workspace}/.local/validator${i}
        node_ids="$(${bin} tendermint show-node-id --home ${workspace}/.local/validator${i})@127.0.0.1:$((${VALIDATOR_P2P_PORT_START}+${i})) ${node_ids}"
    done

    persistent_peers=$(joinByString ',' ${node_ids})
    for ((i=0;i<${size};i++));do
        if [ "$i" -gt 0 ]; then
            cp ${workspace}/.local/validator0/config/genesis.json ${workspace}/.local/validator${i}/config/
        fi
        sed -i -e "s/minimum-gas-prices = \"0stake\"/minimum-gas-prices = \"0${BASIC_DENOM}\"/g" ${workspace}/.local/validator${i}/config/app.toml
        sed -i -e "s/\"stake\"/\"${BASIC_DENOM}\"/g" ${workspace}/.local/validator${i}/config/genesis.json
        sed -i -e "s/\"denom_metadata\": \[\]/\"denom_metadata\": \[${NATIVE_COIN_DESC}\]/g" ${workspace}/.local/validator${i}/config/genesis.json
	sed -i -e "s/seeds = \"[^\"]*\"/seeds = \"\"/g" ${workspace}/.local/validator${i}/config/config.toml
        sed -i -e "s/persistent_peers = \".*\"/persistent_peers = \"${persistent_peers}\"/g" ${workspace}/.local/validator${i}/config/config.toml
        sed -i -e "s/timeout_commit = \"5s\"/timeout_commit = \"500ms\"/g" ${workspace}/.local/validator${i}/config/config.toml
        sed -i -e "s/addr_book_strict = true/addr_book_strict = false/g" ${workspace}/.local/validator${i}/config/config.toml
        sed -i -e "s/allow_duplicate_ip = false/allow_duplicate_ip = true/g" ${workspace}/.local/validator${i}/config/config.toml
        sed -i -e "s/snapshot-interval = 0/snapshot-interval = ${SNAPSHOT_INTERVAL}/g" ${workspace}/.local/validator${i}/config/app.toml
        sed -i -e "s/snapshot-keep-recent = 2/snapshot-keep-recent = ${SNAPSHOT_KEEP_RECENT}/g" ${workspace}/.local/validator${i}/config/app.toml
        sed -i -e "s/\"reserve_time\": \"15552000\"/\"reserve_time\": \"60\"/g" ${workspace}/.local/validator${i}/config/genesis.json
        sed -i -e "s/\"forced_settle_time\": \"86400\"/\"forced_settle_time\": \"30\"/g" ${workspace}/.local/validator${i}/config/genesis.json
        sed -i -e "s/172800s/${DEPOSIT_VOTE_PERIOD}/g" ${workspace}/.local/validator${i}/config/genesis.json
        sed -i -e "s/\"10000000\"/\"${GOV_MIN_DEPOSIT_AMOUNT}\"/g" ${workspace}/.local/validator${i}/config/genesis.json
        sed -i -e "s/\"max_bytes\": \"22020096\"/\"max_bytes\": \"1048576\"/g" ${workspace}/.local/validator${i}/config/genesis.json
        sed -i -e "s/\"challenge_count_per_block\": \"1\"/\"challenge_count_per_block\": \"5\"/g" ${workspace}/.local/validator${i}/config/genesis.json
        sed -i -e "s/\"challenge_keep_alive_period\": \"300\"/\"challenge_keep_alive_period\": \"10\"/g" ${workspace}/.local/validator${i}/config/genesis.json
        sed -i -e "s/\"heartbeat_interval\": \"1000\"/\"heartbeat_interval\": \"100\"/g" ${workspace}/.local/validator${i}/config/genesis.json
        sed -i -e "s/\"attestation_inturn_interval\": \"120\"/\"attestation_inturn_interval\": \"10\"/g" ${workspace}/.local/validator${i}/config/genesis.json
        sed -i -e "s/\"discontinue_confirm_period\": \"604800\"/\"discontinue_confirm_period\": \"5\"/g" ${workspace}/.local/validator${i}/config/genesis.json
        sed -i -e "s/\"discontinue_deletion_max\": \"100\"/\"discontinue_deletion_max\": \"2\"/g" ${workspace}/.local/validator${i}/config/genesis.json
        sed -i -e "s/\"voting_period\": \"30s\"/\"voting_period\": \"5s\"/g" ${workspace}/.local/validator${i}/config/genesis.json
        sed -i -e "s/\"update_global_price_interval\": \"0\"/\"update_global_price_interval\": \"1\"/g" ${workspace}/.local/validator${i}/config/genesis.json
        sed -i -e "s/\"update_price_disallowed_days\": 2/\"update_price_disallowed_days\": 0/g" ${workspace}/.local/validator${i}/config/genesis.json
        #sed -i -e "s/\"community_tax\": \"0.020000000000000000\"/\"community_tax\": \"0\"/g" ${workspace}/.local/validator${i}/config/genesis.json
        sed -i -e "s/log_level = \"info\"/\log_level= \"debug\"/g" ${workspace}/.local/validator${i}/config/config.toml
        sed -i -e "s#node = \"tcp://localhost:26657\"#node = \"tcp://localhost:$((${VALIDATOR_RPC_PORT_START}+${i}))\"#g" ${workspace}/.local/validator${i}/config/client.toml
	sed -i -e "/Address defines the gRPC server address to bind to/{N;s/address = \"localhost:9090\"/address = \"localhost:$((${VALIDATOR_GRPC_PORT_START}+${i}))\"/;}" ${workspace}/.local/validator${i}/config/app.toml
	sed -i -e "/Address defines the gRPC-web server address to bind to/{N;s/address = \"localhost:9091\"/address = \"localhost:$((${VALIDATOR_GRPC_PORT_START}-1-${i}))\"/;}" ${workspace}/.local/validator${i}/config/app.toml
	sed -i -e "/Address defines the EVM RPC HTTP server address to bind to/{N;s/address = \"127.0.0.1:8545\"/address = \"127.0.0.1:$((${EVM_SERVER_PORT_START}+${i}*2))\"/;}" ${workspace}/.local/validator${i}/config/app.toml
	sed -i -e "/Address defines the EVM WebSocket server address to bind to/{N;s/address = \"127.0.0.1:8546\"/address = \"127.0.0.1:$((${EVM_SERVER_PORT_START}+1+${i}*2))\"/;}" ${workspace}/.local/validator${i}/config/app.toml
    done

    # enable swagger API for validator0
    sed -i -e "/Enable defines if the API server should be enabled/{N;s/enable = false/enable = true/;}" ${workspace}/.local/validator0/config/app.toml
    sed -i -e 's/swagger = false/swagger = true/' ${workspace}/.local/validator0/config/app.toml

    # enable telemetry for validator0
    sed -i -e "/other sinks such as Prometheus/{N;s/enable = false/enable = true/;}" ${workspace}/.local/validator0/config/app.toml
}

function start() {
    size=$1
    for ((i=0;i<${size};i++));do
        mkdir -p ${workspace}/.local/validator${i}/logs
        nohup ${bin} start --home ${workspace}/.local/validator${i} \
            --address 0.0.0.0:$((${VALIDATOR_ADDRESS_PORT_START}+${i})) \
            --grpc.address 0.0.0.0:$((${VALIDATOR_GRPC_PORT_START}+${i})) \
            --p2p.laddr tcp://0.0.0.0:$((${VALIDATOR_P2P_PORT_START}+${i})) \
            --p2p.external-address 127.0.0.1:$((${VALIDATOR_P2P_PORT_START}+${i})) \
            --rpc.laddr tcp://0.0.0.0:$((${VALIDATOR_RPC_PORT_START}+${i})) \
            --log_format json > ${workspace}/.local/validator${i}/logs/node.log &
    done
}

function stop() {
    ps -ef | grep ${bin_name} | grep validator | awk '{print $2}' | xargs kill
}

function export_validator {
    size=$1

    for ((i = 0; i < ${size}; i++)); do
        bls_priv_key=("$(echo "y" | ${bin} keys export validator_bls${i} --unarmored-hex --unsafe --keyring-backend test --home ${workspace}/.local/validator${i})")
        relayer_key=("$(echo "y" | ${bin} keys export relayer${i}  --unarmored-hex --unsafe --keyring-backend test --home ${workspace}/.local/relayer${i})")

        echo "validator_bls${i} bls_priv_key: ${bls_priv_key}"
        echo "relayer${i} relayer_key: ${relayer_key}"
    done
}

CMD=$1
SIZE=3
if [ ! -z $2 ] && [ "$2" -gt "0" ]; then
    SIZE=$2
fi

case ${CMD} in
init)
    echo "===== init ===="
    init $SIZE
    echo "===== end ===="
    ;;
generate)
    echo "===== generate genesis ===="
    generate_genesis $SIZE
    echo "===== end ===="
    ;;

export_validator)
    export_validator $SIZE
    ;;
start)
    echo "===== start ===="
    start $SIZE
    echo "===== end ===="
    ;;
stop)
    echo "===== stop ===="
    stop
    echo "===== end ===="
    ;;
all)
    echo "===== stop ===="
    stop
    echo "===== init ===="
    init $SIZE
    echo "===== generate genesis ===="
    generate_genesis $SIZE
    echo "===== start ===="
    start $SIZE
    echo "===== end ===="
    ;;
*)
    echo "Usage: localup.sh all | init | generate | start | stop | export_sps"
    ;;
esac
