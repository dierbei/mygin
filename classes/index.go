package classes

import (
	"github.com/gin-gonic/gin"
	"mygin/goft"
	"net/http"
)

type IndexClass struct {
}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}

func (i *IndexClass) GetIndex() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"result": "index ok.",
		})
	}
}

func (i *IndexClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/", i.GetIndex())
}
