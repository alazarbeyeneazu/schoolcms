package initiator

import (
	"schoolcms/internal/constant/state"
	"schoolcms/internal/glue/school"
	"schoolcms/internal/glue/teacher"
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
	school.InitRoute(group, handler.School, log.Named("school route"), authDomains)
	teacher.InitRoute(group, handler.Teacher, log.Named("school route"), authDomains)
}
