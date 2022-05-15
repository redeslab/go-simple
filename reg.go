package main

import (
	"encoding/hex"
	"fmt"
	cryptoEth "github.com/ethereum/go-ethereum/crypto"
	"github.com/redeslab/go-simple/contract/ethapi"
	"github.com/redeslab/go-simple/node"
	"github.com/spf13/cobra"
)

var RegCmd = &cobra.Command{
	Use:   "reg",
	Short: "register self to block chain service",
	Long:  `TODO::.`,
	Run:   basReg,
}

func init() {
	RegCmd.Flags().StringVarP(&param.minerIP, "minerIP",
		"m", "", "Simple reg -m [MY IP Address]")

	RegCmd.Flags().StringVarP(&param.password, "password",
		"p", "", "Simple reg -p [PASSWORD]")

	RegCmd.Flags().StringVarP(&param.priKey, "key", "k", "", "config contract admin private key")
}

func basReg(_ *cobra.Command, _ []string) {

	node.PathSetting.WalletPath = node.WalletDir(node.BaseDir())

	if err := node.WInst().Open(param.password); err != nil {
		fmt.Println("======>>>password not correct, can't open wallet")
		return
	}
	pk, err := hex.DecodeString(param.priKey)
	if err != nil {
		fmt.Println("======>>>invalid contract private key", err)
		return
	}
	ethPk := cryptoEth.ToECDSAUnsafe(pk)
	tx, err := ethapi.RegNewMiner(node.WInst().SubAddress().String(), param.minerIP, ethPk)
	if err != nil {
		fmt.Println("======> RegNewMiner err:", err)
		return
	}
	fmt.Println("======>>> reg new miner success tx=>", tx)
}
