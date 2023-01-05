#!/usr/bin/env bash
export OWNER_PASSWORD=password
export RECEIVER_PASSWORD=password
export DELEGATE_PASSWORD=password
export BUYER_PASSWORD=password
export KEY_DIR=./accounts
go run cmd/eoa-creator/main.go