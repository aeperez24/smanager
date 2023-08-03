package user

import "smanager/internal/httputils"

func NewLoginHandlerConfigProvider(service IUserService) httputils.HandlerProvider {
	uh := NewUserHandler(service)
	handlersConfigs := []httputils.HandlerConfig{
		{Route: "/user", Method: httputils.POST, Handler: uh.Create}}

	return &httputils.HandlerProviderBase{HandlersConfigs: handlersConfigs}
}
