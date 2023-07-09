package login

import httputil "smanager/internal/httputils"

type LoginHandlerConfigProvider struct {
	handlersConfigs []httputil.HandlerConfig
}

func (lp *LoginHandlerConfigProvider) GetHandlers() []httputil.HandlerConfig {
	return lp.handlersConfigs
}

func NewLoginHandlerConfigProvider(loginService ILoginService) *LoginHandlerConfigProvider {
	lh := &LoginHandler{loginService}
	handlersConfigs := []httputil.HandlerConfig{
		{Route: "/login", Method: httputil.POST, Handler: lh.Login}}

	return &LoginHandlerConfigProvider{handlersConfigs: handlersConfigs}
}
