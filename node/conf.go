package node

import (
	"encoding/json"
	"fmt"
	"github.com/redeslab/go-simple/util"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	WalletPath string        `json:"wallet.path"`
	DBPath     string        `json:"database"`
	LogPath    string        `json:"log_file"`
	TimeOut    time.Duration `json:"conn_timeout"`
}

var Version = "1.0.1"

const (
	WalletFile        = "wallet.json"
	LogFile           = "log.svn"
	ConfFile          = "conf.svn"
	ConnTimeOut       = 5
	ConnectionBufSize = 65535
)

var _conf = &Config{}

func (pc *Config) String() string {
	return fmt.Sprintf("\n++++++++++++++++++++++++++++++++++++++++++++++++++++\n"+
		"+WalletPath:\t%s+\n"+
		"+DBPath:\t%s+\n"+
		"+LogPath:\t%s+\n"+
		"+TimeOut:\t%d+\n"+
		"++++++++++++++++++++++++++++++++++++++++++++++++++++\n",
		pc.WalletPath,
		pc.DBPath,
		pc.LogPath,
		pc.TimeOut)
}

func InitDefaultConfig() *Config {
	base := util.BaseDir()
	if _, ok := util.FileExists(base); !ok {
		panic("Init node first, please!' HOP init -p [PASSWORD]'")
	}
	cfg := &Config{}
	cfg.WalletPath = filepath.Join(base, string(filepath.Separator), WalletFile)
	cfg.LogPath = filepath.Join(base, string(filepath.Separator), LogFile)
	cfg.TimeOut = ConnTimeOut * time.Second
	confPath := filepath.Join(base, string(filepath.Separator), ConfFile)

	fmt.Println(cfg.String())
	byt, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(confPath, byt, 0644); err != nil {
		panic(err)
	}

	return cfg
}
func initConfig(confPath string) {
	base := util.BaseDir()
	_, exist := util.FileExists(base)
	if !exist {
		panic("init service first please!")
	}
	if len(confPath) == 0 {
		confPath = filepath.Join(base, string(filepath.Separator), ConfFile)
	}
	confData, err := ioutil.ReadFile(confPath)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(confData, _conf); err != nil {
		panic(err)
	}
}

func InitNodeConfig(auth, confPath string) {

	initConfig(confPath)

	util.InitLog(_conf.LogPath)

	if auth == "" {
		fmt.Println("Password=>")
		pw, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}
		auth = string(pw)
	}

	if err := WInst().Open(auth); err != nil {
		panic(err)
	}
}

func ChangeConnCloseTimeOut(toInSeconds int) {
	_conf.TimeOut = time.Duration(toInSeconds) * time.Second
}
