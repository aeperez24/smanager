package managedsecret

import (
	httputil "smanager/internal/httputils"
	"smanager/internal/middleware"
)

type HandlerConfigProvider struct {
	handlersConfigs []httputil.HandlerConfig
}

func (hconfigProviders *HandlerConfigProvider) GetHandlers() []httputil.HandlerConfig {
	return hconfigProviders.handlersConfigs
}
func NewHandlerConfigProvider(loginService IManagedSecretService) *HandlerConfigProvider {
	mh := &ManagedSecretHandler{loginService}
	middlewares := []middleware.MiddlewareType{middleware.Secured}
	return &HandlerConfigProvider{
		handlersConfigs: []httputil.HandlerConfig{
			{Route: "/managedSecret", Method: httputil.POST, Handler: mh.CreateManagedSecret, MiddlewareTypes: middlewares},
			{Route: "/managedSecret", Method: httputil.GET, Handler: mh.ListManagedSecret, MiddlewareTypes: middlewares},
			{Route: "/managedSecret", Method: httputil.PUT, Handler: mh.EditManagedSecret, MiddlewareTypes: middlewares},
			{Route: "/managedSecret/:name", Method: httputil.GET, Handler: mh.GetSecret, MiddlewareTypes: middlewares},
		},
	}
}
