package node

import (
	"github.com/redeslab/go-simple/account"
	"github.com/redeslab/go-simple/network"
	"io"
	"net"
	"time"
)

type worker struct {
	wid   int
	local net.Conn
}

func (w *worker) startWork() {
	conn := w.local
	nLog.Debug("======>>>new network:", w.wid, conn.RemoteAddr().String())
	_ = conn.(*net.TCPConn).SetKeepAlive(true)
	defer conn.SetDeadline(time.Now().Add(_conf.TimeOut))

	lvConn := network.NewLVConn(conn)
	jsonConn := &network.JsonConn{Conn: lvConn}
	req := &SetupReq{}
	if err := jsonConn.ReadJsonMsg(req); err != nil {
		nLog.Errorf("[%d]read setup msg err:%s", w.wid, err)
		return
	}
	jsonConn.WriteAck(nil)

	var aesKey account.PipeCryptKey
	if err := account.GenerateAesKey(&aesKey, req.SubAddr.ToPubKey(), WInst().CryptKey()); err != nil {
		nLog.Errorf("[%d]generate aes key err:%s", w.wid, err)
		return
	}

	aesConn, err := network.NewAesConn(lvConn, aesKey[:], req.IV)
	if err != nil {
		nLog.Errorf("[%d]create aes connection err:%s", w.wid, err)
		return
	}
	jsonConn = &network.JsonConn{Conn: aesConn}
	prob := &ProbeReq{}
	if err := jsonConn.ReadJsonMsg(prob); err != nil {
		nLog.Errorf("[%d]read probe msg err:%s", w.wid, err)
		return
	}

	tgtConn, err := net.Dial("tcp", prob.Target)
	if err != nil {
		nLog.Errorf("[%d]dial target[%s] err:%s", w.wid, prob.Target, err)
		return
	}
	_ = tgtConn.(*net.TCPConn).SetKeepAlive(true)
	defer tgtConn.SetDeadline(time.Now().Add(_conf.TimeOut))
	jsonConn.WriteAck(nil)

	var peerMaxPacketSize = prob.MaxPacketSize
	if peerMaxPacketSize == 0 {
		peerMaxPacketSize = ConnectionBufSize
	}
	nLog.Debugf("Setup pipe[%d] for:[%s] from:%s with peer max size=%d",
		w.wid,
		prob.Target,
		aesConn.RemoteAddr().String(),
		peerMaxPacketSize)

	go w.upStream(aesConn, tgtConn)

	w.downStream(aesConn, tgtConn, peerMaxPacketSize)
}

func (w *worker) upStream(aesConn, tgtConn net.Conn) {
	buffer := make([]byte, ConnectionBufSize)
	for {
		no, err := aesConn.Read(buffer)
		if no == 0 {
			if err != io.EOF {
				nLog.Warningf("[%d]read:client--xxx-->proxy---->target err=>%s left:%d", w.wid, err, no)
			} else {
				nLog.Debugf("[%d]read: client--xxx-->proxy---->target EOF ", w.wid)
			}
			return
		}
		_, err = tgtConn.Write(buffer[:no])
		if err != nil {
			nLog.Warningf("[%d]write: client---->proxy--xxx-->target err=>%s", w.wid, err)
			return
		}
		nLog.Debugf("[%d]read: client---->proxy---->target data:%d ", w.wid, no)
	}
}

func (w *worker) downStream(aesConn, tgtConn net.Conn, peerMaxPacketSize int) {
	buffer := make([]byte, ConnectionBufSize)
	for {
		no, err := tgtConn.Read(buffer)
		if no == 0 {
			if err != io.EOF {
				nLog.Warningf("[%d]read: client<----proxy<--xxx--target err=>%s", w.wid, err)
			} else {
				nLog.Debugf("[%d]read: client<----proxy<--xxx--target EOF ", w.wid)
			}
			_ = tgtConn.SetDeadline(time.Now().Add(_conf.TimeOut))
			break
		}

		writeNo, err := aesConn.Write(buffer[:no])
		if err != nil {
			nLog.Warningf("[%d]write client<--xxx--proxy<----target err:%s left=%d", w.wid, err, no)
			break
		}

		nLog.Debugf("[%d]read: client<----proxy<--xxx--target data:%d", w.wid, no, writeNo)

	//	var idx = 0
	//	var data []byte
	//writeToCli:
	//	if no > peerMaxPacketSize {
	//		data = buffer[idx : idx+peerMaxPacketSize]
	//		nLog.Debugf("[%d]big data need to split no=%d idx=%d", w.wid, no, idx)
	//	} else {
	//		data = buffer[idx : idx+no]
	//	}
	//	writeNo, err := aesConn.Write(data)
	//	if err != nil {
	//		nLog.Warningf("[%d]write client<--xxx--proxy<----target err:%s left=%d", w.wid, err, no)
	//		break
	//	}
	//	no = no - peerMaxPacketSize
	//	if no > 0 {
	//		idx = idx + writeNo
	//		goto writeToCli
	//	}
	//}
}
