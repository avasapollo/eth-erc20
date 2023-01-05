#!/usr/bin/env bash
export NETWORK_URL=https://eth-goerli.g.alchemy.com/v2/FgzHk_iQstzYDBe1-Oa0O6jmlubg5PG7
export OWNER_ADDRESS=0x2423c82FD27Cc64AfA56D693c786e0b95cB03C24
export OWNER_PASSWORD=password
export KEY_DIR=./accounts
export CONTRACT_ADDRESS=0xd138304cef7524A334E2cEEc30bBAf558069a1F1

go run cmd/erc-20/main.go