package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestGin(t *testing.T) {
	// 1.创建路由
	egine := gin.Default()
	// 2.绑定路由规则，执行的函数
	egine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world gin!")
	})
	//3.监听端口：默认8080
	egine.Run(":8000")
}
