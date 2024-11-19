#!/bin/bash
mechaind q tx $1 --node "https://testnet-lcd.mechain.tech:443" --output json | jq '.logs[] | .events[] | select(.type == "submit_proposal") | .attributes[] | select(.key == "proposal_id") | .value | tonumber'
