package ethapi

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

const (
	InfuraUrl     = "https://kovan.infura.io/v3/56bf070cd1714103b3bd40e1da1edf86"
	ConfigAddr    = "0xA9de05401C30f9E87910DF5ae83Ab591bE9EA296"
	AdvertiseAddr = "0xfA1159141f9fE553B18d5E67ee58e7ABE53cCa71"
)

func configApi(priKey *ecdsa.PrivateKey, contract string) (*ChainConfig, *bind.TransactOpts, error) {
	client, err := ethclient.Dial(InfuraUrl)
	if err != nil {
		fmt.Println("======> eth client Dial err:", err)
		return nil, nil, err
	}
	var transactor *bind.TransactOpts = nil
	if priKey != nil {
		var nid *big.Int
		nid, err = client.ChainID(context.TODO())
		if err != nil {
			return nil, nil, err
		}
		transactor, err = bind.NewKeyedTransactorWithChainID(priKey, nid)
		if err != nil {
			return nil, nil, err
		}
	}

	var sysConfig *ChainConfig
	confAddr := common.HexToAddress(contract)
	sysConfig, err = NewChainConfig(confAddr, client)
	if err != nil {
		client.Close()
		return nil, nil, err
	}
	return sysConfig, transactor, nil
}

func advertiseApi(priKey *ecdsa.PrivateKey, contract string) (*Advertise, *bind.TransactOpts, error) {
	client, err := ethclient.Dial(InfuraUrl)
	if err != nil {
		fmt.Println("======> eth client Dial err:", err)
		return nil, nil, err
	}
	var transactor *bind.TransactOpts = nil
	if priKey != nil {
		var nid *big.Int
		nid, err = client.ChainID(context.TODO())
		if err != nil {
			return nil, nil, err
		}
		transactor, err = bind.NewKeyedTransactorWithChainID(priKey, nid)
		if err != nil {
			return nil, nil, err
		}
	}

	var ad *Advertise
	confAddr := common.HexToAddress(contract)
	ad, err = NewAdvertise(confAddr, client)
	if err != nil {
		client.Close()
		return nil, nil, err
	}
	return ad, transactor, nil
}
