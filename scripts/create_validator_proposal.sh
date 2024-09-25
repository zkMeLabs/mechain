#!/bin/bash

SCRIPT_DIR=$(dirname "$(realpath "$0")")

TEMPLATE_FILE="$SCRIPT_DIR/create_validator_proposal.json-e"
OUTPUT_FILE="$SCRIPT_DIR/create_validator_proposal.json"

# 定义 mechaind 变量
MECHAIND_CMD="build/mechaind"
CONFIG_PATH="manual/"

# 生成密钥
echo "add keys..."
$MECHAIND_CMD keys add validator --keyring-backend test --home "$CONFIG_PATH" >${CONFIG_PATH}/validator_info 2>&1
$MECHAIND_CMD keys add delegator --keyring-backend test --home "$CONFIG_PATH" >${CONFIG_PATH}/delegator_info 2>&1
$MECHAIND_CMD keys add validator_bls --keyring-backend test --home "$CONFIG_PATH" --algo eth_bls >${CONFIG_PATH}/bls_info 2>&1
$MECHAIND_CMD keys add validator_relayer --keyring-backend test --home "$CONFIG_PATH" >${CONFIG_PATH}/relayer_info 2>&1
$MECHAIND_CMD keys add validator_challenger --keyring-backend test --home "$CONFIG_PATH" >${CONFIG_PATH}/challenger_info 2>&1

# 获取地址和密钥
echo "show keys..."
VALIDATOR_ADDR=$($MECHAIND_CMD keys show validator -a --keyring-backend test --home "$CONFIG_PATH")
DELEGATOR_ADDR=$($MECHAIND_CMD keys show delegator -a --keyring-backend test --home "$CONFIG_PATH")
RELAYER_ADDR=$($MECHAIND_CMD keys show validator_relayer -a --keyring-backend test --home "$CONFIG_PATH")
CHALLENGER_ADDR=$($MECHAIND_CMD keys show validator_challenger -a --keyring-backend test --home "$CONFIG_PATH")
VALIDATOR_BLS=$($MECHAIND_CMD keys show validator_bls --keyring-backend test --home "$CONFIG_PATH" --output json --output-document a.json | jq -r '.pubkey_hex')
VALIDATOR_BLS_PROOF=$($MECHAIND_CMD tx sign ${VALIDATOR_BLS} --keyring-backend test --home "$CONFIG_PATH" --from validator_bls --output json | jq -r '.signature')
VALIDATOR_NODE_PUB_KEY=$(cat ${CONFIG_PATH}/config/priv_validator_key.json | jq -r '.pub_key.value')

if [ -z "$VALIDATOR_ADDR" ] || [ -z "$DELEGATOR_ADDR" ] || [ -z "$RELAYER_ADDR" ] || [ -z "$CHALLENGER_ADDR" ] || [ -z "$VALIDATOR_BLS" ] || [ -z "$VALIDATOR_BLS_PROOF" ] || [ -z "$VALIDATOR_NODE_PUB_KEY" ]; then
    echo "Error: Failed to generate necessary keys or addresses."
    exit 1
fi

# echo "generated validator proposal file..."
# sed -e "s|\${NODE_NAME}|$NODE_NAME|g" \
#     -e "s|\${VALIDATOR_NODE_PUB_KEY}|$VALIDATOR_NODE_PUB_KEY|g" \
#     -e "s|\${VALIDATOR_ADDR}|$VALIDATOR_ADDR|g" \
#     -e "s|\${DELEGATOR_ADDR}|$DELEGATOR_ADDR|g" \
#     -e "s|\${VALIDATOR_BLS}|$VALIDATOR_BLS|g" \
#     -e "s|\${VALIDATOR_BLS_PROOF}|$VALIDATOR_BLS_PROOF|g" \
#     -e "s|\${RELAYER_ADDR}|$RELAYER_ADDR|g" \
#     -e "s|\${CHALLENGER_ADDR}|$CHALLENGER_ADDR|g" \
#     "$TEMPLATE_FILE" >"$OUTPUT_FILE"

# if [ $? -eq 0 ]; then
#     echo "create_validator_proposal.json has been generated successfully at $OUTPUT_FILE."
# else
#     echo "Error: Failed to create create_validator_proposal.json."
# fi
