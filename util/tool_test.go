package util

import (
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
