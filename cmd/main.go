package main

import (
	"github.com/gin-gonic/gin"
	"mygin/classes"
	"mygin/goft"
	"mygin/middleware"
)

func main() {
	r := gin.Default()

	goft.Ignite().Attach(
		middleware.NewUserMiddle(),
	).Mount(
		"/v1",
		classes.NewIndexClass(),
	).Mount(
		"/v2",
		classes.NewUserClass(),
	).Launch()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
