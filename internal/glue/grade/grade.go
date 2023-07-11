package grade

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
	gradeHandler rest.Grade,
	log logger.Logger,
	authDomain state.AuthDomains,

) {
	gradeRoute := []routing.Router{
		{
			Method:      http.MethodPost,
			Path:        "grades",
			Handler:     gradeHandler.CreateGrade,
			Middlewares: []gin.HandlerFunc{},
			Domain: []state.Domain{
				authDomain.Corporate,
				authDomain.System,
			},
		},
	}

	routing.RegisterRoute(group, gradeRoute, log, authDomain)

}
