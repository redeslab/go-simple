package network

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
)

const (
	BuffSize = 1 << 15
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

func (conn *JsonConn) SynBuffer(buff []byte, v interface{}) error {
	if err := conn.WriteJsonMsg(v); err != nil {
		return err
	}

	ack := &ACK{}
	if err := conn.ReadJsonBuffer(buff, ack); err != nil {
		return err
	}

	if !ack.Success {
		return fmt.Errorf("create payment channel failed:%s", ack.Message)
	}

	return nil
}

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

func (conn *JsonConn) ReadJsonMsg(v interface{}) error {
	buffer := make([]byte, BuffSize)
	n, err := conn.Read(buffer)
	if err != nil && err != io.EOF {
		return err
	}
	if n == 0 {
		return fmt.Errorf("read empty data")
	}
	//fmt.Println("======>>>remove me later:=>", string(buffer))
	if err = json.Unmarshal(buffer[:n], v); err != nil {
		return err
	}
	return nil
}

func (conn *JsonConn) ReadJsonBuffer(buff []byte, v interface{}) error {
	n, err := conn.Read(buff)
	if err != nil && err != io.EOF {
		return err
	}
	if n == 0 {
		return fmt.Errorf("read empty data")
	}
	//fmt.Println("======>>>remove me later:=>", string(buffer))
	if err = json.Unmarshal(buff[:n], v); err != nil {
		return err
	}
	return nil
}

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
	_, _ = conn.Write(data)
}
