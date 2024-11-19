#!/bin/bash

mechaind tx staking create-validator ./proposal.json --keyring-backend test --chain-id "mechain_5151-1" --from $1 --node "https://testnet-lcd.mechain.tech:443" -b sync --gas "200000000" --fees "100000000000000000000azkme" --yes
