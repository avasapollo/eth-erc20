#!/usr/bin/env bash
export NETWORK_URL=https://eth-goerli.g.alchemy.com/v2/FgzHk_iQstzYDBe1-Oa0O6jmlubg5PG7
export OWNER_ADDRESS=0x2CB965269FBb3B5c4f4F3cD9268444e2B8557cDd
export OWNER_PASSWORD=password
export KEY_DIR=./accounts
export CONTRACT_ADDRESS=0x5974dE36fd99E8D3bD24CEB44cb66e0f294006C3

go run cmd/erc-20/main.go