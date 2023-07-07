package user

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
	user rest.User,
	log logger.Logger, authDomains state.AuthDomains) {
	userRoutes := []routing.Router{
		{
			Method:      http.MethodPost,
			Path:        "users",
			Handler:     user.CreatUser,
			Middlewares: []gin.HandlerFunc{},
			Domain: []state.Domain{
				authDomains.System,
			},
		},
	}
	routing.RegisterRoute(group, userRoutes, log, authDomains)
}
