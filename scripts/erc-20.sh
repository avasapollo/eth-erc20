#!/usr/bin/env bash
export NETWORK_URL=https://eth-goerli.g.alchemy.com/v2/FgzHk_iQstzYDBe1-Oa0O6jmlubg5PG7
export OWNER_ADDRESS=0xA58b122cea87Bb37f42FB28A9a42F60DDbcB9446
export OWNER_PASSWORD=password
export KEY_DIR=./accounts
export CONTRACT_ADDRESS=0x1C3b6F6F4A53e778601D65544fb4A121eFC69dbf

go run cmd/erc-20/main.go