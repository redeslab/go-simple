package pbs

import (
	"encoding/hex"
	"encoding/json"
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
		"i", "", "Simple adv -n -i [ADVERTISE NAME]")
	AdvertiseCmd.Flags().StringVarP(&param.priKey, "eth-key",
		"k", "", "Simple adv -n -i [ADVERTISE NAME]")
	AdvertiseCmd.Flags().BoolVarP(&param.all, "all", "a", false, "one advertisements")
	AdvertiseCmd.Flags().StringVarP(&param.contractAddr, "address", "d", "", "smart contract address")

	img = AdvertiseCmd.Flags().String("img", "", "--img")
	link = AdvertiseCmd.Flags().String("link", "", "--link")
	typ = AdvertiseCmd.Flags().Int("typ", 0, "--typ")
}

var link *string
var img *string
var typ *int

func advertiseOp(_ *cobra.Command, _ []string) {

	contractAddr := param.contractAddr
	if len(contractAddr) == 0 {
		contractAddr = ethapi.AdvertiseAddr
	}
	if param.all {
		data := ethapi.AdvertiseList(contractAddr)
		if len(data) == 0 {
			fmt.Println("no valid config")
		}
		for idx, datum := range data {
			fmt.Println(idx, datum.Name, datum.ConfigInJson)
		}
		return
	}

	if param.one {
		data := ethapi.QueryOneAdItem(param.id, contractAddr)
		fmt.Println(data)
		return
	}
	if len(param.priKey) == 0 {
		fmt.Println("======>>>empty private key")
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
	case 0, 1:
		if len(*link) == 0 {
			fmt.Println("=====>>> invalid ad lik")
			return
		}
		if len(*img) == 0 {
			fmt.Println("=====>>> invalid ad img")
			return
		}
		if len(param.id) == 0 {
			fmt.Println("=====>>> invalid ad name")
			return
		}
		var adInst = &ethapi.AdvertiseConfig{
			ImgUrl:  *img,
			LinkUrl: *link,
			Typ:     *typ,
		}

		bts, _ := json.Marshal(adInst)
		if param.confOp == 0 {
			tx, err = ethapi.RegNewAD(param.id, string(bts), contractAddr, ethPk)
		} else {
			tx, err = ethapi.UpdateAd(param.id, string(bts), contractAddr, ethPk)
		}
	case 2:
		if len(param.id) == 0 {
			fmt.Println("=====>>> invalid ad name")
			return
		}
		tx, err = ethapi.DelAD(param.id, contractAddr, ethPk)
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
