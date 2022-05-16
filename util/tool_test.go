package util

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeOut(t *testing.T) {
	defer func() {
		println("======>>> defer", time.Now().Add(time.Second*5).String())
	}()

	println("======>>> start", time.Now().String())
	time.Sleep(time.Second * 10)
	println("======>>> end", time.Now().String())

}
func TestBufSize(t *testing.T) {
	buf := make([]byte, 1024)
	fmt.Println(len(buf), cap(buf))
}
