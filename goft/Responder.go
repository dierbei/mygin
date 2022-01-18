package goft

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

var (
	ResponderList = []Responder{new(StringResponder), new(ModelResponder), new(ModelsResponder)}
)

type Responder interface {
	RespondTo() gin.HandlerFunc
}

func Convert(handler interface{}) gin.HandlerFunc {
	handlerRefVal := reflect.ValueOf(handler)
	for _, responder := range ResponderList {
		rpRefVal := reflect.ValueOf(responder).Elem()
		if handlerRefVal.Type().ConvertibleTo(rpRefVal.Type()) {
			rpRefVal.Set(handlerRefVal)
			return rpRefVal.Interface().(Responder).RespondTo()
		}
	}
	return nil
}

type StringResponder func(ctx *gin.Context) string

func (r StringResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, r(ctx))
	}
}

type ModelResponder func(ctx *gin.Context) Model

func (r ModelResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, r(ctx))
	}
}

type ModelsResponder func(ctx *gin.Context) Models

func (r ModelsResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
	}
}