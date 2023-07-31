package managedsecret

import (
	"smanager/internal/httputils"
	httputil "smanager/internal/httputils"
	"smanager/internal/middleware"
)

func NewHandlerConfigProvider(loginService IManagedSecretService) httputils.HandlerProvider {
	mh := &ManagedSecretHandler{loginService}
	middlewares := []middleware.MiddlewareType{middleware.Secured}
	return &httputils.HandlerProviderBase{
		HandlersConfigs: []httputil.HandlerConfig{
			{Route: "/managedSecret", Method: httputil.POST, Handler: mh.CreateManagedSecret, MiddlewareTypes: middlewares},
			{Route: "/managedSecret", Method: httputil.GET, Handler: mh.ListManagedSecret, MiddlewareTypes: middlewares},
			{Route: "/managedSecret", Method: httputil.PUT, Handler: mh.EditManagedSecret, MiddlewareTypes: middlewares},
			{Route: "/managedSecret/:name", Method: httputil.GET, Handler: mh.GetSecret, MiddlewareTypes: middlewares},
		},
	}
}
