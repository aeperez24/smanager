package common

import "github.com/gin-gonic/gin"

type HandlerConfig struct {
	Route   string
	Method  HttpMethod
	Handler func(*gin.Context)
}

type HandlerProvider interface {
	GetHandlers() []HandlerConfig
}

func RegisterRoutes(engine *gin.Engine, handlers []HandlerConfig) {
	methodMap := toMethodMap(engine)
	for _, handler := range handlers {
		method, ok := methodMap[handler.Method]
		if ok {
			method(handler.Route, handler.Handler)
		}
	}
}

type handlerType func(string, ...gin.HandlerFunc) gin.IRoutes

func toMethodMap(engine *gin.Engine) map[HttpMethod]handlerType {
	resultMap := make(map[HttpMethod]handlerType)
	resultMap[GET] = engine.GET
	resultMap[POST] = engine.POST
	resultMap[PUT] = engine.PUT
	resultMap[DELETE] = engine.DELETE
	return resultMap
}
