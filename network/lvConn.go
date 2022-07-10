package network

import (
	"errors"
	"fmt"
	"github.com/redeslab/go-simple/util"
	"io"
	"net"
)

const (
	MTU     = 1 << 24
	LenSize = 4
)

type LVConn struct {
	net.Conn
	lenBuf [LenSize]byte
}

func NewLVConn(conn net.Conn) net.Conn {
	return &LVConn{Conn: conn}
}

func (lc *LVConn) Read(buf []byte) (int, error) {

	if _, err := io.ReadFull(lc.Conn, lc.lenBuf[:]); err != nil {
		return 0, err
	}

	dataLen := int(util.ByteToUint(lc.lenBuf[:]))
	if dataLen == 0 || dataLen >= MTU {
		return 0, fmt.Errorf("MTU overflow:%d", dataLen)
	}
	return io.ReadFull(lc.Conn, buf[:dataLen])
}

func (lc *LVConn) Write(buf []byte) (n int, err error) {
	if len(buf) == 0 {
		err = fmt.Errorf("write empty data to sock client")
		fmt.Println(err)
		return
	}
	dataLen := uint32(len(buf))
	headerBuf := util.UintToByte(dataLen)

	n, err = lc.Conn.Write(headerBuf)
	if err != nil {
		return 0, err
	}
	if n != len(headerBuf) {
		return 0, errors.New("write header buf error, system buffer fulled")
	}
	n, err = lc.Conn.Write(buf)

	if err != nil {
		return 0, err
	}
	if n != len(buf) {
		return 0, errors.New("write buf error, system buffer fulled")
	}

	return int(dataLen), err
}
