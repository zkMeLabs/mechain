#!/bin/bash
curl -s https://testnet-api.mechain.tech/cosmos/gov/v1/proposals/$1 | jq -r '.proposal.status'
