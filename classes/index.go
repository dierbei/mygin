package classes

import (
	"github.com/gin-gonic/gin"
	"mygin/goft"
)

type IndexClass struct {
}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}

func (i *IndexClass) GetIndex(ctx *gin.Context) string {
	return "index"
}

func (i *IndexClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/", i.GetIndex)
}
