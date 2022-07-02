package node

import (
	"encoding/json"
	"fmt"
	"github.com/btcsuite/goleveldb/leveldb"
	"github.com/btcsuite/goleveldb/leveldb/filter"
	"github.com/btcsuite/goleveldb/leveldb/opt"
	"github.com/op/go-logging"
	"github.com/redeslab/go-simple/account"
	"net"
	"sync"
)

var (
	instance *Node = nil
	once     sync.Once
	nLog, _  = logging.GetLogger(LogModuleName)
)

type Node struct {
	subAddr  account.ID
	srvConn  net.Listener
	ctrlChan *net.UDPConn
	database *leveldb.DB
	quit     chan struct{}
	pipeID   int
}

func Inst() *Node {
	once.Do(func() {
		instance = newNode()
	})
	return instance
}

func newNode() *Node {
	sa := WInst().SubAddress()

	opts := opt.Options{
		Strict:      opt.DefaultStrict,
		Compression: opt.NoCompression,
		Filter:      filter.NewBloomFilter(10),
	}

	db, err := leveldb.OpenFile(_conf.DBPath, &opts)
	if err != nil {
		panic(err)
	}
	srvPort := sa.ToServerPort()
	c, err := net.Listen("tcp4", fmt.Sprintf(":%d", srvPort))
	if err != nil {
		panic(err)
	}
	p, err := net.ListenUDP("udp4", &net.UDPAddr{Port: int(srvPort)})
	if err != nil {
		panic(err)
	}

	n := &Node{
		subAddr:  sa,
		srvConn:  c,
		ctrlChan: p,
		database: db,
		quit:     make(chan struct{}, 16),
	}

	return n
}

func (n *Node) StartUp() {
	go n.Mining()
	go n.CtrlService()
}

func (n *Node) ctrlChanReceive(req *CtrlMsg) *MsgAck {
	ack := &MsgAck{}
	ack.Typ = req.Typ
	ack.Msg = "failure"
	ack.Code = 1
	nLog.Debug("Control Channel Receive:", req.String())

	switch req.Typ {
	case MsgPingTest:
		ack.Code = 0
		ack.Msg = "success"
	default:
		ack.Code = -1
		ack.Msg = "unknown message type"
	}
	return ack
}

func (n *Node) CtrlService() {
	nLog.Info("control channel working", n.ctrlChan.LocalAddr().String())
	for {
		buf := make([]byte, 10240)

		nr, addr, err := n.ctrlChan.ReadFrom(buf)
		if err != nil {
			nLog.Warning("control channel error ", err)
			continue
		}
		go n.ctrlMsg(buf[:nr], addr)
	}
}

func (n *Node) ctrlMsg(buf []byte, addr net.Addr) {
	req := &CtrlMsg{}
	err := json.Unmarshal(buf, req)
	if err != nil {
		nLog.Warning("control channel bad request ", err)
		return
	}
	data := n.ctrlChanReceive(req)
	if j, err := json.Marshal(*data); err != nil {
		nLog.Debug("Marshal ctrlMsg data failed", data.String())
		return
	} else {
		n.ctrlChan.WriteTo(j, addr)
	}
}

func (n *Node) Mining() {
	defer n.srvConn.Close()
	nLog.Info("service thread working", n.srvConn.Addr().String())
	for {
		conn, err := n.srvConn.Accept()
		if err != nil {
			panic(err)
		}
		n.pipeID++
		w := &worker{
			wid:   n.pipeID,
			local: conn,
		}
		go w.startWork()
	}
}

func (n *Node) Stop() {
	_ = n.srvConn.Close()
	_ = n.database.Close()
	close(n.quit)
}
