package managedsecret

import (
	"smanager/internal/httputils"
	"smanager/internal/middleware"
)

func NewHandlerConfigProvider(loginService IManagedSecretService) httputils.HandlerProvider {
	mh := &ManagedSecretHandler{loginService}
	middlewares := []middleware.MiddlewareType{middleware.Secured}
	return &httputils.HandlerProviderBase{
		HandlersConfigs: []httputils.HandlerConfig{
			{Route: "/managedSecret", Method: httputils.POST, Handler: mh.CreateManagedSecret, MiddlewareTypes: middlewares},
			{Route: "/managedSecret", Method: httputils.GET, Handler: mh.ListManagedSecret, MiddlewareTypes: middlewares},
			{Route: "/managedSecret", Method: httputils.PUT, Handler: mh.EditManagedSecret, MiddlewareTypes: middlewares},
			{Route: "/managedSecret/:name", Method: httputils.GET, Handler: mh.GetSecret, MiddlewareTypes: middlewares},
		},
	}
}
