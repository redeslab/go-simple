package network

import (
	"errors"
	"fmt"
	"github.com/redeslab/go-simple/util"
	"io"
	"net"
)

type LVConn struct {
	net.Conn
}

func NewLVConn(conn net.Conn) net.Conn {
	return &LVConn{Conn: conn}
}

func (lc *LVConn) Read(buf []byte) (n int, err error) {

	lenBuf := make([]byte, 4)
	if _, err = io.ReadFull(lc.Conn, lenBuf); err != nil {
		//if err != io.EOF {
		//	logger.Notice("\nRead length of data err:", err)
		//}
		return
	}

	dataLen := util.ByteToUint(lenBuf)
	if dataLen == 0 || dataLen > BuffSize {
		err = fmt.Errorf("wrong buffer size:%d", dataLen)
		return
	}

	if len(buf) < int(dataLen) {
		return 0, fmt.Errorf("buffer is too small(buf:%d, data:%d)", len(buf), dataLen)
	}

	buf = buf[:dataLen]
	if n, err = io.ReadFull(lc.Conn, buf); err != nil {
		return
	}
	return int(dataLen), err
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
