package teacher

import (
	"net/http"
	"schoolcms/internal/constant/state"
	"schoolcms/internal/glue/routing.go"
	"schoolcms/internal/handler/rest"
	"schoolcms/platform/logger"

	"github.com/gin-gonic/gin"
)

func InitRoute(
	group *gin.RouterGroup,
	teacherHandler rest.Teacher,
	log logger.Logger,
	authDomain state.AuthDomains,
) {
	teacherRoute := []routing.Router{
		{
			Method:      http.MethodPost,
			Path:        "teachers",
			Handler:     teacherHandler.CreateTeacher,
			Middlewares: []gin.HandlerFunc{},
			Domain: []state.Domain{
				authDomain.System,
			},
		},
		{
			Method:      http.MethodPost,
			Path:        "assign/teachers",
			Handler:     teacherHandler.AssignTeachersToSchool,
			Middlewares: []gin.HandlerFunc{},
			Domain: []state.Domain{
				authDomain.Corporate,
			},
		},
	}
	routing.RegisterRoute(group, teacherRoute, log, authDomain)
}
