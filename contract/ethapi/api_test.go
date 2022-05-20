package ethapi

import (
	"encoding/hex"
	"flag"
	"fmt"
	cryptoEth "github.com/ethereum/go-ethereum/crypto"
	"testing"
)

var (
	userAddr = ""
	priKey   = ""
)

func init() {
	flag.StringVar(&userAddr, "addr", "", "--addr")
	flag.StringVar(&priKey, "key", "", "--key")
}

//go test -run  TestSetAdvertiseAddr --key --addr
func TestSetAdvertiseAddr(t *testing.T) {
	pk, err := hex.DecodeString(priKey)
	if err != nil {
		fmt.Println("======>>>invalid contract private key", err)
		return
	}
	ethPk := cryptoEth.ToECDSAUnsafe(pk)
	tx, err := SetAdvertiseAddress(userAddr, ethPk)
	fmt.Println("tx:", tx, err)
}

//go test -run  TestQueryAdvertiseAddr
func TestQueryAdvertiseAddr(t *testing.T) {
	ret, err := GetAdvertiseAddress()
	fmt.Println("result:", ret, err)
}

func TestLoadAdList(t *testing.T) {
	ads := AdvertiseList("")
	fmt.Println(ads)
}
