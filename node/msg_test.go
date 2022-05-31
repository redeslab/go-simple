package node

import (
	"encoding/json"
	"fmt"
	"github.com/redeslab/go-simple/network"
	"net"
	"testing"
	"time"
)

func TestSetupMsg(t *testing.T) {
	sr := &SetupReq{
		IV:      *network.NewSalt(),
		SubAddr: "SV4a9kD9PdTLihJhD3CgZbTkMSoc528sXG1Tupz2PCwDZ3",
	}

	bts, _ := json.Marshal(sr)
	fmt.Println(string(bts))
}

func TestTargetConn(t *testing.T) {
	tgtConn, err := net.Dial("tcp", "gw.line.naver.jp:443")
	if err != nil {
		t.Fatal(err)
	}
	buffer := make([]byte, 0, 1<<20)
	no, err := tgtConn.Read(buffer)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("read no:", no)

	time.Sleep(time.Second * 5)
}
