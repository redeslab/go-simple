package pbs

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
	Run:   configReg,
}

func init() {
	RegCmd.Flags().StringVarP(&param.minerIP, "minerIP",
		"m", "", "Simple reg -m [MY IP Address]")

	RegCmd.Flags().StringVarP(&param.password, "password",
		"p", "", "Simple reg -p [PASSWORD]")

	RegCmd.Flags().StringVarP(&param.priKey, "key", "k", "", "config contract admin private key")
}

func configReg(_ *cobra.Command, _ []string) {

	if len(param.priKey) == 0 || len(param.password) == 0 || len(param.minerIP) == 0 {
		fmt.Println("parameter needed: [ETH Admin Key], [Node Password], [Node Host]")
		return
	}

	pk, err := hex.DecodeString(param.priKey)
	if err != nil {
		fmt.Println("======>>>invalid contract private key", err)
		return
	}
	ethPk := cryptoEth.ToECDSAUnsafe(pk)

	node.InitNodeConfig(param.password, "")
	tx, err := ethapi.RegNewMiner(node.WInst().SubAddress().String(), param.minerIP, ethPk)
	if err != nil {
		fmt.Println("======> RegNewMiner err:", err)
		return
	}
	fmt.Println("======>>> reg new miner success tx=>", tx)
}
