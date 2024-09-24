package ethRPCUtils

import (
	"encoding/json"
	"github.com/AKACoder/EthereumRPCShell/common/rpcLog"
	"github.com/gin-gonic/gin"
	"strings"
)

func reqFromJSON(c *gin.Context, reqData interface{}) bool {
	commReq, err := c.GetRawData()
	if err != nil {
		rpcLog.Log.Error("get common in request failed: ", err.Error())
		return false
	}

	err = json.Unmarshal(commReq, reqData)
	if err != nil {
		rpcLog.Log.Warn("reqFromJSON unmarshal data failed: ", err.Error())
		return false
	}

	return true
}

func GetReqData(c *gin.Context, reqData interface{}) bool {
	if strings.ToLower(c.ContentType()) == "application/json" {
		return reqFromJSON(c, reqData)
	}

	return false
}
