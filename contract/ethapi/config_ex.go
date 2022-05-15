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
	InfuraUrl    = "https://kovan.infura.io/v3/56bf070cd1714103b3bd40e1da1edf86"
	ContractAddr = "0xBF945030192a61E5f725Cee7fc7cc097fF76dc65"
)

func ethApi(priKey *ecdsa.PrivateKey) (*ChainConfig, *bind.TransactOpts, error) {
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
	confAddr := common.HexToAddress(ContractAddr)
	sysConfig, err = NewChainConfig(confAddr, client)
	if err != nil {
		client.Close()
		return nil, nil, err
	}
	return sysConfig, transactor, nil
}

func SyncServerList() []ConfigServerItem {
	sysConf, _, err := ethApi(nil)
	if err != nil {
		fmt.Println("======> ethApi err:", err)
		return nil
	}
	items, err := sysConf.ServerList(nil)
	if err != nil {
		fmt.Println("======> get server list err:", err)
		return nil
	}
	items = items[1:]
	return items
}

func RefreshHostByAddr(addr string) string {
	sysConf, _, err := ethApi(nil)
	if err != nil {
		fmt.Println("======> ethApi err:", err)
		return ""
	}

	host, err := sysConf.QueryByOne(nil, addr)
	if err != nil {
		fmt.Println("======> QueryByOne  err:", err)
		return ""
	}

	return host
}

func RegNewMiner(addr, host string, key *ecdsa.PrivateKey) (string, error) {
	sysConf, option, err := ethApi(key)
	if err != nil {
		return "", err
	}

	tx, err := sysConf.AddServer(option, addr, host)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}
