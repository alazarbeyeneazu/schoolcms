package school

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
	school rest.School,
	log logger.Logger, authDomains state.AuthDomains) {
	schoolRoutes := []routing.Router{
		{
			Method:      http.MethodPost,
			Path:        "schools",
			Handler:     school.CreateSchool,
			Middlewares: []gin.HandlerFunc{},
			Domain: []state.Domain{
				authDomains.System,
			},
		},
	}
	routing.RegisterRoute(group, schoolRoutes, log, authDomains)
}
