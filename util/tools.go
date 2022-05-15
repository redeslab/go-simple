package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/op/go-logging"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

const (
	DefaultBaseDir = ".simpleVPN"
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

func UintToByte(val uint32) []byte {
	lenBuf := make([]byte, 4, 4)
	binary.BigEndian.PutUint32(lenBuf, val)
	return lenBuf
}
func ByteToUint(buff []byte) uint32 {
	return binary.BigEndian.Uint32(buff)
}

func FileExists(fileName string) (os.FileInfo, bool) {

	fileInfo, err := os.Lstat(fileName)

	if fileInfo != nil || (err != nil && !os.IsNotExist(err)) {
		return fileInfo, true
	}
	return nil, false
}

func TouchDir(dir string) error {
	if _, ok := FileExists(dir); ok {
		return nil
	}

	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		return err
	}
	return nil
}
func ReadPassWord2() (string, error) {
	fmt.Println("Password=>")
	pw, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}

	fmt.Println("Password Again=>")
	pw2, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}

	if bytes.Compare(pw, pw2) != 0 {
		return "", fmt.Errorf("not same of 2 inputs")
	}

	return string(pw), nil
}

func InitLog(logPath string) {
	logFile, err := os.OpenFile(logPath,
		os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	fileBackend := logging.NewLogBackend(logFile, "-->", 0)
	fileFormat := logging.MustStringFormatter(
		`{time:01-02/15:04:05} %{longfunc:-30s} %{shortfile:-22.20s} > %{level:.4s} %{message}`,
	)
	fileFormatBackend := logging.NewBackendFormatter(fileBackend, fileFormat)

	leveledFileBackend := logging.AddModuleLevel(fileFormatBackend)

	cmdFormat := logging.MustStringFormatter(
		`%{color}%{time:01-02/15:04:05} %{shortfile:-20.18s} %{shortfunc:-20.20s} [%{level:.4s}] %{message}%{color:reset}`,
	)
	cmdBackend := logging.NewLogBackend(os.Stderr, "\n>>>", 0)
	formattedCmdBackend := logging.NewBackendFormatter(cmdBackend, cmdFormat)

	logging.SetBackend(leveledFileBackend, formattedCmdBackend)
	fmt.Println("log init success")
}

func BaseDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	baseDir := filepath.Join(usr.HomeDir, string(filepath.Separator), DefaultBaseDir)
	return baseDir
}
