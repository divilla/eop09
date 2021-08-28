#!/bin/bash
go build -o ./cmd/build/server ./cmd/server/main.go
./cmd/build/server
