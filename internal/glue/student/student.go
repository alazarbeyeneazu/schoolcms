package student

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
	studentHandler rest.Student,
	log logger.Logger,
	authDomains state.AuthDomains,
) {
	studentRoute := []routing.Router{
		{
			Method:      http.MethodPost,
			Path:        "students",
			Handler:     studentHandler.CreateStudent,
			Middlewares: []gin.HandlerFunc{},
			Domain: []state.Domain{
				authDomains.Corporate,
				authDomains.System,
			},
		},
	}
	routing.RegisterRoute(group, studentRoute, log, authDomains)
}
