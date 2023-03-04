package http

import "github.com/gin-gonic/gin"

func Response(ctx *gin.Context, statusCode int, body any) {
	ctx.JSON(statusCode, gin.H{
		"data": body,
	})
}

func Error(ctx *gin.Context, statusCode int, body any) {
	ctx.JSON(statusCode, gin.H{
		"error": body,
	})
}
