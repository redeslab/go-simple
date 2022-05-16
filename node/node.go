package node

import (
	"encoding/json"
	"fmt"
	"github.com/btcsuite/goleveldb/leveldb"
	"github.com/btcsuite/goleveldb/leveldb/filter"
	"github.com/btcsuite/goleveldb/leveldb/opt"
	"github.com/op/go-logging"
	"github.com/redeslab/go-simple/account"
	"github.com/redeslab/go-simple/conn"
	"net"
	"sync"
	"time"
)

var (
	instance *Node = nil
	once     sync.Once
	nLog, _  = logging.GetLogger("node")
)

type Node struct {
	subAddr  account.ID
	srvConn  net.Listener
	ctrlChan *net.UDPConn
	database *leveldb.DB
	quit     chan struct{}
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

func (n *Node) ctrlChanReceive(req *MsgReq) *MsgAck {
	ack := &MsgAck{}
	ack.Typ = req.Typ
	ack.Msg = "failure"
	ack.Code = 1
	nLog.Debug("Control Channel Receive:", req.String())

	switch req.Typ {
	case MsgPingTest:
		ack.Code = 0
		ack.Msg = "success"
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
	req := &MsgReq{}
	err := json.Unmarshal(buf, req)
	if err != nil {
		nLog.Warning("control channel bad request ", err)
		return
	}
	nLog.Debug("CtrlService raw data:", string(buf))
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
		go n.newWorker(conn)
	}
}

func (n *Node) Stop() {
	_ = n.srvConn.Close()
	_ = n.database.Close()
	close(n.quit)
}

func (n *Node) newWorker(conn net.Conn) {
	nLog.Debug("======>>>new conn:", conn.RemoteAddr().String())
	_ = conn.(*net.TCPConn).SetKeepAlive(true)
	defer conn.SetDeadline(time.Now().Add(_conf.TimeOut))

	lvConn := network.NewLVConn(conn)
	jsonConn := &network.JsonConn{Conn: lvConn}
	req := &SetupReq{}
	if err := jsonConn.ReadJsonMsg(req); err != nil {
		nLog.Errorf("read setup msg err:%s", err)
		return
	}
	jsonConn.WriteAck(nil)

	var aesKey account.PipeCryptKey
	if err := account.GenerateAesKey(&aesKey, req.SubAddr.ToPubKey(), WInst().CryptKey()); err != nil {
		nLog.Errorf("generate aes key err:%s", err)
		return
	}

	aesConn, err := network.NewAesConn(lvConn, aesKey[:], req.IV)
	if err != nil {
		nLog.Errorf("create aes connection err:%s", err)
		return
	}
	jsonConn = &network.JsonConn{Conn: aesConn}
	prob := &ProbeReq{}
	if err := jsonConn.ReadJsonMsg(prob); err != nil {
		nLog.Errorf("read probe msg err:%s", err)
		return
	}

	tgtConn, err := net.Dial("tcp", prob.Target)
	if err != nil {
		nLog.Errorf("dial target[%s] err:%s", prob.Target, err)
		return
	}
	_ = tgtConn.(*net.TCPConn).SetKeepAlive(true)
	defer tgtConn.SetDeadline(time.Now().Add(_conf.TimeOut))
	jsonConn.WriteAck(nil)

	var peerMaxPacketSize = prob.MaxPacketSize
	if peerMaxPacketSize == 0 {
		peerMaxPacketSize = ConnectionBufSize
	}
	nLog.Debugf("Setup pipe for:[%s] from:%s with peer max size=%d",
		prob.Target,
		aesConn.RemoteAddr().String(),
		peerMaxPacketSize)

	go func() { //TODO::
		buffer := make([]byte, peerMaxPacketSize)
		for {
			no, err := aesConn.Read(buffer)
			if no == 0 {
				nLog.Warning("read from client failed", err, no)
				return
			}
			_, err = tgtConn.Write(buffer[:no])
			if err != nil {
				nLog.Warning("write to target failed", err)
				return
			}
		}
	}()

	buffer := make([]byte, peerMaxPacketSize)
	for {
		no, err := tgtConn.Read(buffer)
		if no == 0 {
			nLog.Warning("Target->Proxy read err:", err)
			_ = tgtConn.SetDeadline(time.Now().Add(time.Second * 10))
			break
		}

		_, err = aesConn.Write(buffer[:no])
		if err != nil {
			nLog.Warning("Proxy->Client write  err:", err, no)
			break
		}
	}
}
