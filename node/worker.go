package node

import (
	"github.com/redeslab/go-simple/account"
	"github.com/redeslab/go-simple/network"
	"io"
	"net"
	"time"
)

const (
	DefaultBufSize = 1 << 20
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
	ctrlBuf := make([]byte, 1<<11)
	if err := jsonConn.ReadJsonBuffer(ctrlBuf, req); err != nil {
		nLog.Errorf("[%d]read setup msg err:%s", w.wid, err)
		return
	}
	jsonConn.WriteAck(nil)

	var aesKey account.PipeCryptKey
	if err := account.GenerateAesKey(&aesKey, req.SubAddr.ToPubKey(), WInst().CryptKey()); err != nil {
		nLog.Errorf("[%d]generate aes key err:%s", w.wid, err)
		return
	}

	aesConn, err := network.NewAesConn(conn, aesKey[:], req.IV)
	if err != nil {
		nLog.Errorf("[%d]create aes connection err:%s", w.wid, err)
		return
	}
	lvConn = network.NewLVConn(aesConn)

	jsonConn = &network.JsonConn{Conn: lvConn}
	prob := &ProbeReq{}
	if err := jsonConn.ReadJsonBuffer(ctrlBuf, prob); err != nil {
		nLog.Errorf("[%d]read probe msg err:%s", w.wid, err)
		return
	}

	tgtConn, err := net.Dial("tcp4", prob.Target)
	if err != nil {
		nLog.Errorf("[%d]dial target[%s] err:%s", w.wid, prob.Target, err)
		return
	}
	_ = tgtConn.(*net.TCPConn).SetKeepAlive(true)
	defer tgtConn.SetDeadline(time.Now().Add(_conf.TimeOut))
	jsonConn.WriteAck(nil)

	nLog.Debugf("Setup pipe[%d] for:[%s] from:%s ",
		w.wid,
		prob.Target,
		aesConn.RemoteAddr().String())

	go w.upStream(lvConn, tgtConn)
	w.downStream(lvConn, tgtConn)
	_ = tgtConn.Close()
}

func relay(src, dst net.Conn) {
	buf := make([]byte, DefaultBufSize)
	defer src.Close()
	defer dst.Close()

	_, err := io.CopyBuffer(src, dst, buf)
	if err != nil {
		nLog.Warningf("relay finalized by err:%s", err)
		return
	}

	nLog.Debugf("relay finished:[%s--->%s]===>[%s--->%s]",
		src.LocalAddr().String(),
		src.RemoteAddr().String(),
		dst.LocalAddr().String(),
		dst.RemoteAddr().String())
}

func (w *worker) upStream(lvConn, tgtConn net.Conn) {
	buffer := make([]byte, DefaultBufSize)
	for {
		no, err := lvConn.Read(buffer)
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
		nLog.Debugf("[%d]upStream: client---->proxy---->target data:%d ", w.wid, no)
	}
}

func (w *worker) downStream(lvConn, tgtConn net.Conn) {
	buffer := make([]byte, DefaultBufSize)
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

		writeNo, err := lvConn.Write(buffer[:no])
		if err != nil {
			nLog.Warningf("[%d]write client<--xxx--proxy<----target err:%s left=%d", w.wid, err, no)
			break
		}

		nLog.Debugf("[%d]downStream: client<----proxy<--xxx--target data:%d written:%d", w.wid, no, writeNo)
	}
}
