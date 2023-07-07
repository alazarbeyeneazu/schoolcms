package initiator

import (
	"schoolcms/internal/constant/state"
	"schoolcms/internal/glue/user"
	"schoolcms/platform/logger"

	"github.com/gin-gonic/gin"
)

func InitRouter(
	group *gin.RouterGroup,
	handler Handler,
	module Module,
	log logger.Logger,
	authDomains state.AuthDomains,
) {
	user.InitRoute(group, handler.User, log.Named("user route"), authDomains)
}
