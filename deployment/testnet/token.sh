#!/bin/bash
mechaind tx bank send v0 $1 1000000000000000000000000azkme --keyring-backend test --node https://testnet-lcd.mechain.tech:443 -y --fees 6000000000000azkme

mechaind --node https://testnet-lcd.mechain.tech:443 q bank balances $1
