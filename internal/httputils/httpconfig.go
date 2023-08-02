package httputils

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

type HandlerProviderBase struct {
	HandlersConfigs []HandlerConfig
}

func (lp *HandlerProviderBase) GetHandlers() []HandlerConfig {
	return lp.HandlersConfigs
}

func RegisterRoutes(engine *gin.Engine, handlers []HandlerConfig) {
	for _, handler := range handlers {
		engine.Handle(handler.Method.String(), handler.Route, handler.Handler)
	}
}

func RegisterRoutesWithMiddleware(engine *gin.Engine, handlers []HandlerConfig, middlewares map[middleware.MiddlewareType]gin.HandlerFunc) {
	for _, handler := range handlers {

		handleFuns := make([]gin.HandlerFunc, 0)
		for _, middlewareType := range handler.MiddlewareTypes {
			middlewareToApply, ok := middlewares[middlewareType]
			if ok {
				handleFuns = append(handleFuns, middlewareToApply)
			}
		}
		handleFuns = append(handleFuns, handler.Handler)
		engine.Handle(handler.Method.String(), handler.Route, handleFuns...)
	}
}

type handlerType func(string, ...gin.HandlerFunc) gin.IRoutes
