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
func work(conn net.Conn) {
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
	if err != nil {
		t.Fatal(err)
	}
	buff := make([]byte, 20)
	for {

		n, err := lvCnn.Read(buff)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("compare data:", buff[:n])
	}
}
func TestAesConn1(t *testing.T) {
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
	key := NewSalt()
	iv := NewSalt()

	aesConn, err := NewAesConn(conn, (*key)[:], *iv)
	if err != nil {
		t.Fatal(err)
	}

	_, err = aesConn.Write(test_data)
	if err != nil {
		t.Fatal(err)
	}
	buff := make([]byte, 20)
	for {

		n, err := aesConn.Read(buff)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("compare data:", buff[:n])
	}
}

func initTcpSrv2(k []byte, iv Salt) {
	conn, err := net.ListenTCP("tcp4", &net.TCPAddr{
		Port: 1112,
	})
	if err != nil {
		panic(err)
	}
	for {
		newConn, err := conn.AcceptTCP()
		if err != nil {
			panic(err)
		}
		lv := &LVConn{Conn: newConn}
		aesConn, err := NewAesConn(lv, k, iv)
		if err != nil {
			panic(err)
		}
		go work(aesConn)
	}
}
func TestAesConn2(t *testing.T) {

	key := NewSalt()
	iv := NewSalt()
	go initTcpSrv2((*key)[:], *iv)

	test_data := make([]byte, 100)
	for i := uint8(0); i < 100; i++ {
		test_data[i] = i
	}
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 1112,
	})
	if err != nil {
		t.Fatal(err)
	}
	lvConn := &LVConn{
		Conn: conn,
	}

	aesConn, err := NewAesConn(lvConn, (*key)[:], *iv)
	if err != nil {
		t.Fatal(err)
	}

	_, err = aesConn.Write(test_data)
	if err != nil {
		t.Fatal(err)
	}

	buff := make([]byte, 20)
	for {
		n, err := aesConn.Read(buff)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("compare data:", buff[:n])
	}
}

func initTcpSrv3(k []byte, iv Salt) {
	conn, err := net.ListenTCP("tcp4", &net.TCPAddr{
		Port: 1113,
	})
	if err != nil {
		panic(err)
	}
	for {
		newConn, err := conn.AcceptTCP()
		if err != nil {
			panic(err)
		}
		aesConn, err := NewAesConn(newConn, k, iv)
		if err != nil {
			panic(err)
		}
		lv := &LVConn{Conn: aesConn}
		go work(lv)
	}
}

func TestAesConn3(t *testing.T) {

	key := NewSalt()
	iv := NewSalt()
	go initTcpSrv3((*key)[:], *iv)

	test_data := make([]byte, 100)
	for i := uint8(0); i < 100; i++ {
		test_data[i] = i
	}
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 1113,
	})
	if err != nil {
		t.Fatal(err)
	}

	aesConn, err := NewAesConn(conn, (*key)[:], *iv)
	if err != nil {
		t.Fatal(err)
	}
	lvConn := &LVConn{
		Conn: aesConn,
	}

	_, err = lvConn.Write(test_data)
	if err != nil {
		t.Fatal(err)
	}

	buff := make([]byte, 21)
	for {
		n, err := lvConn.Read(buff)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("compare data:", buff[:n])
	}
}
