#!/bin/bash

# 设置变量
MONIKER_NAME="validator-lsl"

# create accounts
echo "Creating validator, delegator, validator BLS, relayer, and challenger accounts..."
mechaind keys add validator --keyring-backend test
mechaind keys add delegator --keyring-backend test
mechaind keys add validator_bls --keyring-backend test --algo eth_bls
mechaind keys add validator_relayer --keyring-backend test
mechaind keys add validator_challenger --keyring-backend test

# get accouts and private key
echo "Obtaining addresses and keys..."
VALIDATOR_ADDR=$(mechaind keys show validator -a --keyring-backend test)
DELEGATOR_ADDR=$(mechaind keys show delegator -a --keyring-backend test)
RELAYER_ADDR=$(mechaind keys show validator_relayer -a --keyring-backend test)
CHALLENGER_ADDR=$(mechaind keys show validator_challenger -a --keyring-backend test)
VALIDATOR_BLS=$(mechaind keys show validator_bls --keyring-backend test --output json | jq -r '.pubkey_hex')
VALIDATOR_BLS_PROOF=$(mechaind keys sign ${VALIDATOR_BLS} --keyring-backend test --from validator_bls)
VALIDATOR_NODE_PUB_KEY=$(cat ~/.mechaind/config/priv_validator_key.json | jq -r '.pub_key.value')

echo VALIDATOR_ADDR: $VALIDATOR_ADDR
echo DELEGATOR_ADDR: $DELEGATOR_ADDR
echo RELAYER_ADDR: $RELAYER_ADDR
echo CHALLENGER_ADDR: $CHALLENGER_ADDR
echo VALIDATOR_BLS: $VALIDATOR_BLS
echo VALIDATOR_BLS_PROOF: $VALIDATOR_BLS_PROOF
echo VALIDATOR_NODE_PUB_KEY: $VALIDATOR_NODE_PUB_KEY

# generate the JSON file
echo "Generating proposal.json..."
cat <<EOF >proposal.json
{
  "messages": [
    {
      "@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
      "description": {
        "moniker": "${MONIKER_NAME}",
        "identity": "",
        "website": "https://${MONIKER_NAME}",
        "security_contact": "",
        "details": "${MONIKER_NAME} details"
      },
      "commission": {
        "rate": "0.070000000000000000",
        "max_rate": "1.000000000000000000",
        "max_change_rate": "0.010000000000000000"
      },
      "min_self_delegation": "1",
      "delegator_address": "${DELEGATOR_ADDR}",
      "validator_address": "${VALIDATOR_ADDR}",
      "pubkey": {
        "@type": "/cosmos.crypto.ed25519.PubKey",
        "key": "${VALIDATOR_NODE_PUB_KEY}"
      },
      "value": {
        "denom": "azkme",
        "amount": "10000000000000000000000"
      },
      "from": "0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2",
      "relayer_address": "${RELAYER_ADDR}",
      "challenger_address": "${CHALLENGER_ADDR}",
      "bls_key": "${VALIDATOR_BLS}",
      "bls_proof": "${VALIDATOR_BLS_PROOF}"
    }
  ],
  "metadata": "",
  "title": "Create ${MONIKER_NAME} Validator",
  "summary": "create ${MONIKER_NAME} validator",
  "deposit": "10000000000000000000000azkme"
}
EOF

echo "proposal.json has been created."
