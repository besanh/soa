package api

import (
	"github.com/besanh/soa/services"
	"github.com/gin-gonic/gin"
)

func Validate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Header.Get("Authorization") != services.SECRET_KEY {
			ctx.JSON(401, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
