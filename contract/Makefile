SHELL=PATH='$(PATH)' /bin/sh

.PHONY: all
all: abi
	abigen --abi eth_abis/Config.abi --pkg ethapi --type ChainConfig --out ./ethapi/confg.go

abi:
	solc --abi --overwrite *.sol -o ./eth_abis

clean:
#	rm -f *.abi
	rm -f *.go
