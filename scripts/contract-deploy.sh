#!/usr/bin/env bash
export NETWORK_URL=https://eth-goerli.g.alchemy.com/v2/FgzHk_iQstzYDBe1-Oa0O6jmlubg5PG7
export OWNER_ADDRESS=0x2CB965269FBb3B5c4f4F3cD9268444e2B8557cDd
export OWNER_PASSWORD=password
export OWNER_BALANCE=100000000000000
export KEY_DIR=./accounts

go run cmd/contract-deploy/main.go