package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web/pkg/resp"
)

func HandleError(c *gin.Context) {
	c.Next()
	es := c.Errors.Errors()
	if len(es) != 0 {
		data := ""
		for i := 0; i < len(es); i++ {
			data += fmt.Sprintf("[err%d]:%s\n", i+1, es[i])
		}
		resp.ErrResp(c, data)
	}
}
