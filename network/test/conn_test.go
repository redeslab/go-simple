package test

import (
	"fmt"
	"github.com/redeslab/go-simple/network"
	"github.com/redeslab/go-simple/node"
	"net"
	"testing"
	"time"
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
	var lvCnn = network.NewLVConn(conn)
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
	key := network.NewSalt()
	iv := network.NewSalt()

	aesConn, err := network.NewAesConn(conn, (*key)[:], *iv)
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

func initTcpSrv2(k []byte, iv network.Salt) {
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
		lv := &network.LVConn{Conn: newConn}
		aesConn, err := network.NewAesConn(lv, k, iv)
		if err != nil {
			panic(err)
		}
		go work(aesConn)
	}
}
func TestAesConn2(t *testing.T) {

	key := network.NewSalt()
	iv := network.NewSalt()
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
	lvConn := &network.LVConn{
		Conn: conn,
	}

	aesConn, err := network.NewAesConn(lvConn, (*key)[:], *iv)
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

func initTcpSrv3(k []byte, iv network.Salt) {
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
		aesConn, err := network.NewAesConn(newConn, k, iv)
		if err != nil {
			panic(err)
		}
		lv := &network.LVConn{Conn: aesConn}
		go work(lv)
	}
}

func TestAesConn3(t *testing.T) {

	key := network.NewSalt()
	iv := network.NewSalt()
	go initTcpSrv3((*key)[:], *iv)
	time.Sleep(time.Second * 3)

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

	aesConn, err := network.NewAesConn(conn, (*key)[:], *iv)
	if err != nil {
		t.Fatal(err)
	}
	lvConn := &network.LVConn{
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

func initTcpSrv4(k []byte, iv network.Salt) {
	conn, err := net.ListenTCP("tcp4", &net.TCPAddr{
		Port: 1114,
	})
	if err != nil {
		panic(err)
	}
	for {

		newConn, err := conn.AcceptTCP()
		if err != nil {
			panic(err)
		}
		go func() {
			fmt.Println("---------xxxx---_>")
			lvConn := network.NewLVConn(newConn)
			jsonConn := &network.JsonConn{Conn: lvConn}
			req := &node.SetupReq{}
			ctrlBuf := make([]byte, 2048)
			if err := jsonConn.ReadJsonBuffer(ctrlBuf, req); err != nil {
				panic(err)
			}
			jsonConn.WriteAck(nil)
			fmt.Println("---------33333---_>")

			aesConn, err := network.NewAesConn(newConn, k, iv)
			if err != nil {
				panic(err)
			}
			lvConn = network.NewLVConn(aesConn)

			jsonConn = &network.JsonConn{Conn: lvConn}
			prob := &node.ProbeReq{}
			if err := jsonConn.ReadJsonBuffer(ctrlBuf, prob); err != nil {
				panic(err)
			}
			fmt.Println("---------4444444---_>", prob.Target)
			jsonConn.WriteAck(nil)

			work(lvConn)
		}()

	}
}

func TestJsonConn4(t *testing.T) {

	key := network.NewSalt()
	iv := network.NewSalt()
	go initTcpSrv4((*key)[:], *iv)
	time.Sleep(time.Second * 3)
	fmt.Println("---------11111---_>")

	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 1114,
	})
	if err != nil {
		t.Fatal(err)
	}

	lvConn := network.NewLVConn(conn)

	req := &node.SetupReq{
		IV:      *iv,
		SubAddr: "1111111",
	}
	jsonConn := &network.JsonConn{Conn: lvConn}
	buf := make([]byte, 2048)
	if err := jsonConn.SynBuffer(buf, req); err != nil {
		t.Fatal(err)
	}
	fmt.Println("---------2222222---_>")

	aesConn, err := network.NewAesConn(conn, (*key)[:], *iv)
	if err != nil {
		t.Fatal(err)
	}
	lvConn = network.NewLVConn(aesConn)

	jsonConn = &network.JsonConn{Conn: lvConn}
	if err := jsonConn.SynBuffer(buf, &node.ProbeReq{
		Target: "www.google.com",
	}); err != nil {
		t.Fatal(err)
	}
	fmt.Println("---------5555555---_>")

	test_data := make([]byte, 1000000)
	for i := 0; i < 1000000; i++ {
		test_data[i] = uint8(i)
	}
	_, err = lvConn.Write(test_data)
	if err != nil {
		t.Fatal(err)
	}

	buff := make([]byte, 255)
	for {
		n, err := lvConn.Read(buff)
		if err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second)
		fmt.Println("compare data:", buff[:n])
	}
}
