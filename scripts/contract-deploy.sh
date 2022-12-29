#!/usr/bin/env bash
export NETWORK_URL=https://eth-goerli.g.alchemy.com/v2/FgzHk_iQstzYDBe1-Oa0O6jmlubg5PG7
export OWNER_ADDRESS=0xA58b122cea87Bb37f42FB28A9a42F60DDbcB9446
export OWNER_PASSWORD=password
export OWNER_BALANCE=100000000000000
export KEY_DIR=./accounts

go run cmd/contract-deploy/main.go