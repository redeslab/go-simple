package util

import (
	"log"
	"os"
	"time"
)

func GetNowMsTime() int64 {
	return time.Now().UnixNano() / 1e6
}

func Int64Time2String(t int64) string {
	tm := time.Unix(t/1000, 0)
	return tm.Format("2006-01-02/15:04:05")
}

func Save2File(data []byte, filename string) error {

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		f.Close()
		log.Fatal(err)
	}

	return nil
}
