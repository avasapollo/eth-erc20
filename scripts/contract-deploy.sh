#!/usr/bin/env bash
export NETWORK_URL=https://eth-goerli.g.alchemy.com/v2/FgzHk_iQstzYDBe1-Oa0O6jmlubg5PG7
export OWNER_ADDRESS=0xC7D8352831E53750A7013B171B8Ad93D8F670B43
export OWNER_PASSWORD=password
export OWNER_BALANCE=1000000000000
export KEY_DIR=./accounts

go run cmd/contract-deploy/main.go