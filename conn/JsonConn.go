package network

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
)

const (
	BuffSize      = 1 << 21
	StructBufSize = 1 << 16
)

type ACK struct {
	Success bool
	Message string
}
type JsonConn struct {
	net.Conn
}

func DialJson(network, address string) (*JsonConn, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &JsonConn{Conn: conn}, nil
}

func (conn *JsonConn) Syn(v interface{}) error {
	if err := conn.WriteJsonMsg(v); err != nil {
		return err
	}

	ack := &ACK{}
	if err := conn.ReadJsonMsg(ack); err != nil {
		return err
	}

	if !ack.Success {
		return fmt.Errorf("create payment channel failed:%s", ack.Message)
	}

	return nil
}

//func (conn *JsonConn) SynTCP(v interface{}) error {
//	if err := conn.WriteJsonMsgTCP(v); err != nil {
//		return err
//	}
//
//	ack := &ACK{}
//	if err := conn.ReadJsonMsgTCP(ack); err != nil {
//		return err
//	}
//
//	if !ack.Success {
//		return fmt.Errorf("create payment channel failed:%s", ack.Message)
//	}
//
//	return nil
//}

func (conn *JsonConn) WriteJsonMsg(v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	if _, err := conn.Write(data); err != nil {
		return err
	}
	return nil
}

//func (conn *JsonConn) WriteJsonMsgTCP(v interface{}) error {
//	data, err := json.Marshal(v)
//	if err != nil {
//		return err
//	}
//
//	l := len(data)
//
//	lbuf := UintToByte(uint32(l))
//
//	var n int
//	n, err = conn.Write(lbuf)
//
//	if err != nil {
//		return err
//	}
//	if n != len(lbuf) {
//		return errors.New("write length buf error, system buffer fulled")
//	}
//
//	n, err = conn.Write(data)
//
//	if err != nil {
//		return err
//	}
//	if n != len(data) {
//		return errors.New("write data error, system buffer fulled")
//	}
//	return nil
//}

func (conn *JsonConn) ReadJsonMsg(v interface{}) error {
	buffer := make([]byte, BuffSize)
	n, err := conn.Read(buffer)
	if err != nil && err != io.EOF {
		return err
	}
	if n == 0 {
		return fmt.Errorf("read empty data")
	}

	if err = json.Unmarshal(buffer[:n], v); err != nil {
		return err
	}
	return nil
}

//func (conn *JsonConn) ReadJsonMsgTCP(v interface{}) error {
//	lbuf := make([]byte, 4)
//	n, err := io.ReadFull(conn, lbuf)
//	if err != nil {
//		return err
//	}
//	if n != 4 {
//		return errors.New("Read buffer error")
//	}
//
//	l := ByteToUint(lbuf)
//
//	if l > StructBufSize {
//		return errors.New("Content size too large")
//	}
//
//	buffer := make([]byte, int(l))
//
//	n, err = io.ReadFull(conn, buffer)
//	if err != nil {
//		return err
//	}
//	if n != len(buffer) {
//		return errors.New("Read buffer error")
//	}
//	if err = json.Unmarshal(buffer, v); err != nil {
//		return err
//	}
//	return nil
//}

func (conn *JsonConn) WriteAck(err error) {
	var data []byte
	if err == nil {
		data, _ = json.Marshal(&ACK{
			Success: true,
			Message: "Success",
		})
	} else {
		data, _ = json.Marshal(&ACK{
			Success: false,
			Message: err.Error(),
		})
	}
	conn.Write(data)
}

//
//func (conn *JsonConn) WriteAckTCP(err error) {
//	ack := &ACK{
//		Success: true,
//		Message: "Success",
//	}
//
//	if err != nil {
//		ack.Success = false
//		ack.Message = err.Error()
//	}
//
//	conn.WriteJsonMsgTCP(ack)
//}
