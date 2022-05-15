package pbs

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/redeslab/go-simple/account"
	"github.com/redeslab/go-simple/node"
	"github.com/spf13/cobra"
)

var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "show miner's basic info",
	Long:  `TODO::.`,
}

var ShowAddrCmd = &cobra.Command{
	Use:   "address",
	Short: "miner's network layer address",
	Long:  `TODO::.`,
	Run:   showAddr,
}

func init() {
	ShowCmd.AddCommand(ShowAddrCmd)
	ShowAddrCmd.Flags().StringVarP(&param.path, "wallet.path",
		"w", "", "Simple show -w [Wallet Path]")
	ShowAddrCmd.Flags().StringVarP(&param.password, "password", "p", "", "Password to create account.")

}
func showAddr(_ *cobra.Command, _ []string) {
	if len(param.password) == 0 {
		print("======>>> wallet password needed")
		return
	}
	node.InitNodeConfig(param.password, param.path)

	w, err := account.LoadWallet(param.path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(w.MainAddress().String())
	fmt.Println(w.SubAddress().String())
	fmt.Println(hexutil.Encode(w.SubAddress().ToPubKey()))

}
