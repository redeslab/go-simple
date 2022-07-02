package network

import (
	"crypto/aes"
	"crypto/cipher"
	"net"
)

type AesConn struct {
	net.Conn
	salt  Salt
	block cipher.Block
	//encStream cipher.Stream
	//decStream cipher.Stream
}

func NewAesConn(conn net.Conn, key []byte, iv Salt) (net.Conn, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ac := &AesConn{
		Conn:  conn,
		salt:  iv,
		block: block,
	}

	//ac.decStream = cipher.NewCFBDecrypter(block, ac.salt[:])
	//ac.encStream = cipher.NewCFBEncrypter(block, ac.salt[:])

	return ac, nil
}

func (ac *AesConn) Read(buf []byte) (n int, err error) {
	n, err = ac.Conn.Read(buf)
	if err != nil {
		return
	}
	//ac.decStream.XORKeyStream(buf[:n], buf[:n])
	cipher.NewCFBDecrypter(ac.block, ac.salt[:]).XORKeyStream(buf[:n], buf[:n])
	return
}

func (ac *AesConn) Write(buf []byte) (n int, err error) {
	//ac.encStream.XORKeyStream(buf, buf)
	cipher.NewCFBEncrypter(ac.block, ac.salt[:]).XORKeyStream(buf, buf)
	return ac.Conn.Write(buf)
}
