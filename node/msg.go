package node

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/redeslab/go-simple/account"
	"github.com/redeslab/go-simple/network"
)

const (
	_ = iota
	MsgPingTest
)

type SetupReq struct {
	IV      network.Salt
	SubAddr account.ID
}

type ProbeReq struct {
	Target        string `json:"Target"`
	MaxPacketSize int    `json:"MaxPacketSize,omitempty"`
}

func (sr *SetupReq) String() string {
	return fmt.Sprintf("\n@@@@@@@@@@@@@@@@@@@[Setup Request]@@@@@@@@@@@@@@@@@"+
		"\nIV:\t%s"+
		"\nSubAddr:\t%s"+
		"\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@",
		hexutil.Encode(sr.IV[:]),
		sr.SubAddr.String())
}

type PingTest struct {
	PayLoad string
}

func (pt *PingTest) String() string {
	if pt == nil {
		return "PingTest is nil"
	}
	return fmt.Sprintf("PayLoad: %s", pt.PayLoad)
}

type CtrlMsg struct {
	Typ int       `json:"typ"`
	PT  *PingTest `json:"pt,omitempty"`
}

func (mr *CtrlMsg) String() string {
	return fmt.Sprintf("type :%d nPingTest:%s\r\n",
		mr.Typ,
		mr.PT.String())
}

type MsgAck struct {
	Typ  int         `json:"typ"`
	Code int         `json:"code"` //0 success 1 failure
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (ack *MsgAck) String() string {
	if ack == nil {
		return "ack is nil"
	}
	j, err := json.Marshal(*ack)
	if err != nil {
		return err.Error()
	}

	return string(j)
}
