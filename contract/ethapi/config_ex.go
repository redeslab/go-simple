package ethapi

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

func SyncServerList() []ConfigServerItem {
	sysConf, _, err := configApi(nil, ConfigAddr)
	if err != nil {
		fmt.Println("======> configApi err:", err)
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
	sysConf, _, err := configApi(nil, ConfigAddr)
	if err != nil {
		fmt.Println("======> configApi err:", err)
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
	sysConf, option, err := configApi(key, ConfigAddr)
	if err != nil {
		return "", err
	}

	tx, err := sysConf.AddServer(option, addr, host)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func UpdateNewMiner(addr, host string, key *ecdsa.PrivateKey) (string, error) {
	sysConf, option, err := configApi(key, ConfigAddr)
	if err != nil {
		return "", err
	}

	tx, err := sysConf.ChangeServer(option, addr, host)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func DelNewMiner(addr, host string, key *ecdsa.PrivateKey) (string, error) {
	sysConf, option, err := configApi(key, ConfigAddr)
	if err != nil {
		return "", err
	}

	tx, err := sysConf.RemoveServer(option, addr)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func GetAdvertiseAddress() (string, error) {
	sysConf, _, err := configApi(nil, ConfigAddr)
	if err != nil {
		fmt.Println("======> configApi err:", err)
		return "", err
	}
	addr, err := sysConf.AdvertiseAddr(nil)
	if err != nil {
		return "", err
	}
	strAddr := addr.String()
	if strAddr == "0x0000000000000000000000000000000000000000" {
		return "", fmt.Errorf("no valid contract address")
	}
	return strAddr, nil
}

func SetAdvertiseAddress(addr string, key *ecdsa.PrivateKey) (string, error) {
	sysConf, option, err := configApi(key, ConfigAddr)
	if err != nil {
		return "", err
	}
	ethAddr := common.HexToAddress(addr)
	tx, err := sysConf.SetAdvertisAddr(option, ethAddr)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}
