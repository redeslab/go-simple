package network

import (
	"errors"
	"fmt"
	"github.com/redeslab/go-simple/util"
	"io"
	"net"
)

const (
	MaxBuffer = 1 << 30
	LenSize   = 4
)

type LVConn struct {
	net.Conn
	bufCache []byte
	lenBuf   [LenSize]byte
}

func NewLVConn(conn net.Conn) net.Conn {
	return &LVConn{Conn: conn}
}

func (lc *LVConn) Read(buf []byte) (int, error) {
	leftLen := len(lc.bufCache)
	if leftLen > 0 {
		cpLen := copy(buf, lc.bufCache)
		if cpLen == leftLen {
			lc.bufCache = nil
		} else {
			lc.bufCache = lc.bufCache[cpLen:]
		}
		//fmt.Println("=============> buf cached", leftLen, cpLen)
		return cpLen, nil
	}

	if _, err := io.ReadFull(lc.Conn, lc.lenBuf[:]); err != nil {
		return 0, err
	}

	dataLen := int(util.ByteToUint(lc.lenBuf[:]))
	if dataLen == 0 || dataLen >= MaxBuffer {
		return 0, fmt.Errorf("wrong buffer size:%d", dataLen)
	}
	fmt.Println("=============> dataLen=>", dataLen)

	bufLen := len(buf)
	if bufLen >= dataLen {
		buf = buf[:dataLen]
		//fmt.Println("=============> buf is enough", bufLen, dataLen)
		return io.ReadFull(lc.Conn, buf)
	}

	//fmt.Println("=============> buf is small", bufLen, dataLen)
	lc.bufCache = make([]byte, dataLen)
	n, err := io.ReadFull(lc.Conn, lc.bufCache)
	if err != nil {
		return n, err
	}
	//fmt.Println("=============> read to cache first", n)
	cpLen := copy(buf, lc.bufCache)
	lc.bufCache = lc.bufCache[cpLen:]
	//fmt.Println("=============> read left", cpLen)
	//fmt.Println(buf)
	//fmt.Println(lc.bufCache)
	return bufLen, nil
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
