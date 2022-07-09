package network

import (
	"fmt"
	"net"
	"testing"
)

func initTcpSrv() {
	conn, err := net.ListenTCP("tcp4", &net.TCPAddr{
		Port: 1111,
	})
	if err != nil {
		panic(err)
	}
	for {
		newConn, err := conn.AcceptTCP()
		if err != nil {
			panic(err)
		}

		go work(newConn)
	}
}

func work(conn *net.TCPConn) {
	buff := make([]byte, 1<<20)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			panic(err)
		}
		n, err = conn.Write(buff[:n])
		if err != nil {
			panic(err)
		}
	}
}

func TestLVConn(t *testing.T) {
	go initTcpSrv()

	test_data := make([]byte, 100)
	for i := uint8(0); i < 100; i++ {
		test_data[i] = i
	}

	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 1111,
	})

	if err != nil {
		t.Fatal(err)
	}
	var lvCnn = NewLVConn(conn)
	_, err = lvCnn.Write(test_data)
	buff := make([]byte, 20)
	for {

		n, err := lvCnn.Read(buff)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("compare data:", buff[:n])
	}
}
