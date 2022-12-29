#!/usr/bin/env bash
export OWNER_PASSWORD=password
export OWNER_BALANCE=1000000000000
export RECEIVER_PASSWORD=password
export KEY_DIR=./accounts
go run cmd/owner-creator/main.go