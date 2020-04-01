package public

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 公共结构体返回
func JSONSuccess(c *gin.Context, code int, msg string, data interface{}) {
	if code == 0 {
		code = 200
	}
	if msg == "" {
		msg = "success"
	}
	c.JSON(http.StatusOK, Resp{
		Success: true,
		Code:    code,
		Msg:     msg,
		Data:    data,
	})
}

func JSONFail(c *gin.Context, code int, msg string) {
	if code == 0 {
		code = 400
	}
	if msg == "" {
		msg = "unknown"
	}

	c.JSON(http.StatusOK, Resp{
		Success: false,
		Code:    code,
		Msg:     msg,
	})
}

func JSONErr(c *gin.Context, code int, biz string, err error) {
	if code == 0 {
		code = 400
	}
	if biz == "" {
		biz = "unknown"
	}
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			Success: false,
			Code:    code,
			Msg:     fmt.Sprintf(`biz: %v error: %v`, biz, err.Error()),
		})
	} else {
		c.JSON(http.StatusOK, Resp{
			Success: false,
			Code:    code,
			Msg:     fmt.Sprintf(`biz: %v error: unknown must fix`, biz),
		})
	}
}
