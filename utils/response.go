package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义标准返回格式
func standardResp(context *gin.Context, code int, msg string, data any) {
	context.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// SuccessResp 定义标准成功返回
func SuccessResp(context *gin.Context, data any) {
	standardResp(context, 0, "ok", data)
}
