package family

import (
	"net/http"
	"schoolcms/internal/constant/state"
	"schoolcms/internal/glue/routing.go"
	"schoolcms/internal/handler/rest"
	"schoolcms/platform/logger"

	"github.com/gin-gonic/gin"
)

func InitRoute(group *gin.RouterGroup, familyHandler rest.Family, log logger.Logger, authDomain state.AuthDomains) {
	familyRoutes := []routing.Router{
		{
			Method:      http.MethodPost,
			Path:        "/families",
			Handler:     familyHandler.CreateFamily,
			Middlewares: []gin.HandlerFunc{},
			Domain: []state.Domain{
				authDomain.Corporate,
				authDomain.System,
			},
		},
	}
	routing.RegisterRoute(group, familyRoutes, log, authDomain)
}
