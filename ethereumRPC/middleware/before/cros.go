package before

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func cros(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")

    if c.Request.Method == "OPTIONS" {
        if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
            c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
        }
        c.AbortWithStatus(http.StatusNoContent)
    } else {
        c.Next()
    }
}
