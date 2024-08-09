package ethRPCUtils

import (
	"EthereumRPCShell/common/wLog"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strings"
)

func reqFromJSON(c *gin.Context, reqData interface{}) bool {
	commReq, err := c.GetRawData()
	if err != nil {
		wLog.Log.Error("get common in request failed: ", err.Error())
		return false
	}

	err = json.Unmarshal(commReq, reqData)
	if err != nil {
		wLog.Log.Warn("reqFromJSON unmarshal data failed: ", err.Error())
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
