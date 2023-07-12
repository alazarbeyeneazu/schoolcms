package initiator

import (
	"schoolcms/internal/constant"
	"schoolcms/internal/constant/state"

	"github.com/spf13/viper"
)

type State struct {
	AuthDomains state.AuthDomains
}

func InitState() State {
	authDomains := state.AuthDomains{
		Corporate: state.Domain{
			ID:   viper.GetString("service.authorization.domain.corporate"),
			Name: string(constant.Corporate),
		},
		System: state.Domain{
			ID:   viper.GetString("service.authorization.domain.system"),
			Name: string(constant.System),
		},
		User: state.Domain{
			ID:   viper.GetString("service.authorization.domain.user"),
			Name: string(constant.User),
		},
	}
	return State{
		AuthDomains: authDomains,
	}
}
