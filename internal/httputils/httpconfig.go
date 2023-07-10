package httputil

import (
	"smanager/internal/middleware"

	"github.com/gin-gonic/gin"
)

type HandlerConfig struct {
	Route           string
	Method          HttpMethod
	Handler         func(*gin.Context)
	MiddlewareTypes []middleware.MiddlewareType
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
func RegisterRoutesWithMiddleware(engine *gin.Engine, handlers []HandlerConfig, middlewares map[middleware.MiddlewareType]gin.HandlerFunc) {
	methodMap := toMethodMap(engine)
	for _, handler := range handlers {
		method, ok := methodMap[handler.Method]
		if ok {
			handleFuns := make([]gin.HandlerFunc, 0)
			for _, middlewareType := range handler.MiddlewareTypes {
				middlewareToApply, ok := middlewares[middlewareType]
				if ok {
					handleFuns = append(handleFuns, middlewareToApply)
				}
			}
			handleFuns = append(handleFuns, handler.Handler)
			method(handler.Route, handleFuns...)
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
