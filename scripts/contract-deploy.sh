#!/usr/bin/env bash
export NETWORK_URL=https://eth-goerli.g.alchemy.com/v2/FgzHk_iQstzYDBe1-Oa0O6jmlubg5PG7
export OWNER_ADDRESS=0x2423c82FD27Cc64AfA56D693c786e0b95cB03C24
export OWNER_PASSWORD=password
export OWNER_BALANCE=100000000000000
export KEY_DIR=./accounts

go run cmd/contract-deploy/main.go