#!/bin/sh
cd ./eip2539go
rm go.sum
# go mod edit --replace github.com/ethereum/go-ethereum=github.com/kilic/go-ethereum@bls_precompiles
go get -u
make static_external_asm
