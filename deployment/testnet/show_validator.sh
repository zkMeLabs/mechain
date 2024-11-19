#!/bin/bash
curl -s https://testnet-api.mechain.tech/cosmos/staking/v1beta1/validators | jq '.pagination.total'
