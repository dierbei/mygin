package goft

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Goft struct {
	*gin.Engine
	group *gin.RouterGroup
}

func Ignite() *Goft {
	return &Goft{Engine: gin.New()}
}

func (g *Goft) Launch() {
	g.Run(":8080")
}

func (g *Goft) Mount(group string, iclasses ...IClass) *Goft {
	g.group = g.Group(group)
	for _, class := range iclasses {
		class.Build(g)
	}
	return g
}

func (g *Goft) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) {
	g.group.Handle(httpMethod, relativePath, handlers...)
}

func (g *Goft) Attach(f Fairing) *Goft {
	g.Use(func(ctx *gin.Context) {
		if err := f.OnRequest(ctx); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.Next()
		}
	})
	return g
}
