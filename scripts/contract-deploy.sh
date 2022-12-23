#!/usr/bin/env bash
export NETWORK_URL=https://eth-goerli.g.alchemy.com/v2/FgzHk_iQstzYDBe1-Oa0O6jmlubg5PG7
export OWNER_ADDRESS=0x3F15cb553FAA92aD4fBde47E6CA727A4A0d49d85
export OWNER_PASSWORD=password
export OWNER_BALANCE=1000000000000
export KEY_DIR=./accounts

go run cmd/contract-deploy/main.go