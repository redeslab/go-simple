package ethapi

import (
	"crypto/ecdsa"
	"fmt"
)

type AdvertiseConfig struct {
	ImgUrl  string `json:"img_url"`
	LinkUrl string `json:"link_url"`
	Typ     int    `json:"typ"`
}

func AdvertiseList(contractAddr string) []AdvertiseAdItem {
	if len(contractAddr) == 0 {
		contractAddr = AdvertiseAddr
	}
	ad, _, err := advertiseApi(nil, contractAddr)
	if err != nil {
		fmt.Println("======> advertiseApi err:", err)
		return nil
	}
	items, err := ad.AdList(nil)
	if err != nil {
		fmt.Println("======> get server list err:", err)
		return nil
	}
	items = items[1:]
	return items
}

func QueryOneAdItem(addr, contractAddr string) string {
	ad, _, err := advertiseApi(nil, contractAddr)
	if err != nil {
		fmt.Println("======> advertiseApi err:", err)
		return ""
	}

	conf, err := ad.QueryByOne(nil, addr)
	if err != nil {
		fmt.Println("======> QueryByOne  err:", err)
		return ""
	}

	return conf
}

func RegNewAD(name, jsonConfig, contractAddr string, key *ecdsa.PrivateKey) (string, error) {
	ad, option, err := advertiseApi(key, contractAddr)
	if err != nil {
		return "", err
	}

	tx, err := ad.AddItem(option, name, jsonConfig)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

func DelAD(name, contractAddr string, key *ecdsa.PrivateKey) (string, error) {
	ad, option, err := advertiseApi(key, contractAddr)
	if err != nil {
		return "", err
	}

	tx, err := ad.RemoveItem(option, name)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}
