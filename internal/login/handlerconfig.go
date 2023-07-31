package login

import (
	"smanager/internal/httputils"
)

func NewLoginHandlerConfigProvider(loginService ILoginService) httputils.HandlerProvider {
	lh := &LoginHandler{loginService}
	handlersConfigs := []httputils.HandlerConfig{
		{Route: "/login", Method: httputils.POST, Handler: lh.Login}}

	return &httputils.HandlerProviderBase{HandlersConfigs: handlersConfigs}
}
