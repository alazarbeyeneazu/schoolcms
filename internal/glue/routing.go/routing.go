package routing

import (
	"context"
	"fmt"
	"path"
	"schoolcms/internal/constant"
	"schoolcms/internal/constant/state"
	"schoolcms/platform/logger"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Method      string
	Path        string
	Handler     gin.HandlerFunc
	Middlewares []gin.HandlerFunc
	Domain      []state.Domain
}

func RegisterRoute(
	grp *gin.RouterGroup,
	routes []Router,
	zapLogger logger.Logger,
	authDomains state.AuthDomains,
) {
	for _, route := range routes {
		for _, domain := range route.Domain {
			var handler []gin.HandlerFunc

			var endpoint string

			switch domain.Name {
			case authDomains.System.Name:
				endpoint = path.Join(string(constant.System), route.Path)
			case authDomains.Corporate.Name:
				endpoint = path.Join(string(constant.Corporate), route.Path)
			case authDomains.User.Name:
				endpoint = path.Join(string(constant.User), route.Path)

			default:
				zapLogger.Fatal(context.Background(), fmt.Sprintf("Invalid Domain %s Registered on Route", domain))
			}

			handler = append(handler, route.Middlewares...)
			handler = append(handler, route.Handler)
			grp.Handle(route.Method, endpoint, handler...)
		}
	}
}
