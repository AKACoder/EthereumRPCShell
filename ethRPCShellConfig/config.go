package ethRPCShellConfig

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/AKACoder/EthereumRPCShell/common/rpcLog"
	"github.com/AKACoder/EthereumRPCShell/network/http"
	"os"
)

type Config struct {
	HTTP http.Config   `json:"http"`
	LOG  rpcLog.Config `json:"w_log"`
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

var DefaultLogCfg = rpcLog.Config{
	Level:   5,
	LogFile: "shell.rpcLog",
}
