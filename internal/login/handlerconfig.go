package login

import "smanager/internal/common"

type LoginHandlerConfigProvider struct {
	handlersConfigs []common.HandlerConfig
}

func (lp *LoginHandlerConfigProvider) GetHandlers() []common.HandlerConfig {
	return lp.handlersConfigs
}

func NewLoginHandlerConfigProvider(loginService ILoginService) *LoginHandlerConfigProvider {
	lh := &LoginHandler{loginService}
	handlersConfigs := []common.HandlerConfig{
		{Route: "/login", Method: common.POST, Handler: lh.Login}}

	return &LoginHandlerConfigProvider{handlersConfigs: handlersConfigs}
}
