package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequestResponse(ctx *gin.Context, payload interface{}) {
	WriteJsonResponse(ctx, http.StatusBadRequest, gin.H{
		"message": payload,
	})
}

func InternalServerJsonResponse(ctx *gin.Context, payload interface{}) {
	WriteJsonResponse(ctx, http.StatusInternalServerError, gin.H{
		"message": payload,
	})
}

func NotFoundResponse(ctx *gin.Context, payload interface{}) {
	WriteJsonResponse(ctx, http.StatusNotFound, gin.H{
		"message": payload,
	})
}

func WriteJsonResponse(ctx *gin.Context, status int, payload interface{}) {
	ctx.JSON(status, payload)
}
