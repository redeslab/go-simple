package node

import (
	"github.com/redeslab/go-simple/account"
	"sync"
)

var (
	wInstance *MinerWallet = nil
	wOnce     sync.Once
)

type MinerWallet struct {
	account.Wallet
}

func WInst() *MinerWallet {
	wOnce.Do(func() {
		wInstance = loadWallet()
	})
	return wInstance
}
func loadWallet() *MinerWallet {
	w, err := account.LoadWallet(_conf.WalletPath)
	if err != nil {
		panic(err)
	}
	return &MinerWallet{Wallet: w}
}
