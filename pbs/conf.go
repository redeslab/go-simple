package pbs

import (
	"encoding/hex"
	"fmt"
	cryptoEth "github.com/ethereum/go-ethereum/crypto"
	"github.com/redeslab/go-simple/contract/ethapi"
	"github.com/redeslab/go-simple/node"
	"github.com/spf13/cobra"
)

var ConfCmd = &cobra.Command{
	Use:   "conf",
	Short: "config self to block chain service",
	Long:  `TODO::.`,
	Run:   configReg,
}

func init() {
	ConfCmd.Flags().StringVarP(&param.minerIP, "minerIP",
		"m", "", "Simple conf -m [MY IP Address]")

	ConfCmd.Flags().StringVarP(&param.password, "password",
		"p", "", "Simple conf -p [PASSWORD]")

	ConfCmd.Flags().StringVarP(&param.priKey, "key", "k", "", "config contract admin private key")

	ConfCmd.Flags().Int8VarP(&param.confOp, "op", "o", 0, "config operations 0:reg, 1:update, 2:delete")

	ConfCmd.Flags().BoolVarP(&param.one, "one", "n", false, "one one miner's config")
	ConfCmd.Flags().StringVarP(&param.id, "id",
		"i", "", "Simple conf -n -i [MINER's ADDR]")
	ConfCmd.Flags().BoolVarP(&param.all, "all", "a", false, "one all miner's config")
}

func configReg(_ *cobra.Command, _ []string) {

	if param.all {
		data := ethapi.SyncServerList()
		if len(data) == 0 {
			fmt.Println("no valid config")
		}
		for idx, datum := range data {
			fmt.Println(idx, datum.Addr, datum.Host)
		}
		return
	}

	if param.one {
		data := ethapi.RefreshHostByAddr(param.id)
		fmt.Println(data)
		return
	}

	if len(param.priKey) == 0 || len(param.password) == 0 {
		fmt.Println("parameter needed: [ETH Admin Key], [Node Password]")
		return
	}

	pk, err := hex.DecodeString(param.priKey)
	if err != nil {
		fmt.Println("======>>>invalid contract private key", err)
		return
	}
	ethPk := cryptoEth.ToECDSAUnsafe(pk)

	node.PrepareConfig(param.password, "")
	var tx = ""
	switch param.confOp {
	case 0:
		if len(param.minerIP) == 0 {
			fmt.Println("Node host parameter needed")
			return
		}
		tx, err = ethapi.RegNewMiner(node.WInst().SubAddress().String(), param.minerIP, ethPk)
	case 1:
		if len(param.minerIP) == 0 {
			fmt.Println("Node host parameter needed")
			return
		}
		tx, err = ethapi.UpdateNewMiner(node.WInst().SubAddress().String(), param.minerIP, ethPk)
	case 2:
		tx, err = ethapi.DelNewMiner(node.WInst().SubAddress().String(), param.minerIP, ethPk)
	default:
		fmt.Println("======>>>unknown config operations!")
		return

	}
	if err != nil {
		fmt.Println("======> miner operation err:", err)
		return
	}
	fmt.Println("======>>>miner operation success tx=>", tx)
}
