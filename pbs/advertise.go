package pbs

import (
	"encoding/hex"
	"fmt"
	cryptoEth "github.com/ethereum/go-ethereum/crypto"
	"github.com/redeslab/go-simple/contract/ethapi"
	"github.com/spf13/cobra"
)

var AdvertiseCmd = &cobra.Command{
	Use:   "adv",
	Short: "miner's network layer address",
	Long:  `TODO::.`,
	Run:   advertiseOp,
}

func init() {
	AdvertiseCmd.Flags().Int8VarP(&param.confOp, "op", "o", 0, "advertisements operations 0:reg, 1:update, 2:delete")

	AdvertiseCmd.Flags().BoolVarP(&param.one, "one", "n", false, "one one advertisements")
	AdvertiseCmd.Flags().StringVarP(&param.id, "id",
		"i", "", "Simple adv -n -i [ADVERTISE NAMEW]")
	AdvertiseCmd.Flags().BoolVarP(&param.all, "all", "a", false, "one advertisements")
	AdvertiseCmd.Flags().StringVarP(&param.contractAddr, "address", "d", "", "smart contract address")
}

func advertiseOp(_ *cobra.Command, _ []string) {
	if param.all {
		data := ethapi.AdvertiseList(param.contractAddr)
		if len(data) == 0 {
			fmt.Println("no valid config")
		}
		for idx, datum := range data {
			fmt.Println(idx, datum.Name, datum.ConfigInJson)
		}
		return
	}

	if param.one {
		data := ethapi.QueryOneAdItem(param.id, param.contractAddr)
		fmt.Println(data)
		return
	}

	pk, err := hex.DecodeString(param.priKey)
	if err != nil {
		fmt.Println("======>>>invalid contract private key", err)
		return
	}
	ethPk := cryptoEth.ToECDSAUnsafe(pk)
	var tx = ""
	switch param.confOp {
	case 0:
		tx, err = ethapi.RegNewAD(param.id, "", param.contractAddr, ethPk)
	default:
		fmt.Println("======>>>unknown advertise operations!")
		return

	}
	if err != nil {
		fmt.Println("======> advertise operation err:", err)
		return
	}
	fmt.Println("======>>>advertise operation success tx=>", tx)
}
