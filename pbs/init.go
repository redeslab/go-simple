package pbs

import (
	"fmt"
	"github.com/redeslab/go-simple/account"
	"github.com/redeslab/go-simple/node"
	"github.com/redeslab/go-simple/util"
	"github.com/spf13/cobra"
	"os"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "init miner node",
	Long:  `TODO::.`,
	Run:   initMiner,
}

func init() {
	InitCmd.Flags().StringVarP(&param.password, "password", "p", "", "Password to create account.")
}

func initMiner(_ *cobra.Command, _ []string) {

	baseDir := util.BaseDir()
	if _, ok := util.FileExists(baseDir); ok {
		fmt.Println("======>>>Duplicate init operation")
		return
	}
	if len(param.password) == 0 {
		pwd, err := util.ReadPassWord2()
		if err != nil {
			panic(err)
		}
		param.password = pwd
	}

	if err := os.Mkdir(baseDir, os.ModePerm); err != nil {
		panic(err)
	}

	defaultCfg := node.InitDefaultConfig()

	w, err := account.NewWallet(param.password, false)
	if err != nil {
		panic(err)
	}

	if err := w.SaveToPath(defaultCfg.WalletPath); err != nil {
		panic(err)
	}
	fmt.Printf("======>>>Create wallet success!\nmain=>[%s] \nsub=>[%s]\n<<<======\n", w.MainAddress(), w.SubAddress())
}
