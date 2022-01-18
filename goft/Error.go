package goft

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": e})
			}
		}()
		context.Next()
	}
}

func Error(err error, errMsg ...string) {
	if err != nil {
		if len(errMsg) > 0 {
			panic(errMsg[0])
		}
		panic(err)
	}
}
