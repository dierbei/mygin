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

func (g *Goft) Handle(httpMethod, relativePath string, handler interface{}) {
	//if h, ok := handler.(func(ctx *gin.Context) string); ok {
	//	g.group.Handle(httpMethod, relativePath, func(ctx *gin.Context) {
	//		ctx.String(http.StatusOK, h(ctx))
	//	})
	//}
	//g.group.Handle(httpMethod, relativePath, handlers...)
	if h := Convert(handler); h != nil {
		g.group.Handle(httpMethod, relativePath, Convert(handler))
	}
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
