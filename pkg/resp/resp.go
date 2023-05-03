package resp

import "github.com/gin-gonic/gin"

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func ErrResp(c *gin.Context, data string) {
	c.JSON(200, Resp{
		Code: 0,
		Msg:  "failed",
		Data: data,
	})
}

func CommonResp(c *gin.Context, data any) {
	c.JSON(200, Resp{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}
