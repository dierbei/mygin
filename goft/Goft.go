package goft

import "github.com/gin-gonic/gin"

type Goft struct {
	*gin.Engine
}

func Ignite() *Goft {
	return &Goft{Engine: gin.New()}
}

func (g *Goft) Launch() {
	g.Run(":8080")
}

func (g *Goft) Mount(iclasses ...IClass) *Goft {
	for _, class := range iclasses {
		class.Build(g)
	}
	return g
}
