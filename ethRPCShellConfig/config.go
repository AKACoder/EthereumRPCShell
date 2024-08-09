package ethRPCShellConfig

import (
	"EthereumRPCShell/common/wLog"
	"EthereumRPCShell/network/http"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	HTTP http.Config `json:"http"`
	LOG  wLog.Config `json:"wLog"`
}

func (c *Config) LoadConfig(cfgFile string) (err error) {
	cfgStr, err := os.ReadFile(cfgFile)
	if nil != err {
		fmt.Println("load ethRPCShellConfig file error: ", err)
		return
	}

	cfgStr = bytes.TrimPrefix(cfgStr, []byte("\xef\xbb\xbf"))
	err = json.Unmarshal(cfgStr, c)

	if nil != err {
		fmt.Println("invalid ethRPCShellConfig file: ", err)
		return
	}

	return
}

var DefaultLogCfg = wLog.Config{
	Level:   5,
	LogFile: "shell.wLog",
}
