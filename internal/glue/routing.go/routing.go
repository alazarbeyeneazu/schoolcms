package routing

import (
	"context"
	"fmt"
	"loyalty/internal/constant"
	"loyalty/platform/logger"
	"path"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Method      string
	Path        string
	Handler     gin.HandlerFunc
	Middlewares []gin.HandlerFunc
	Domain      []constant.Domain
}

func RegisterRoute(
	grp *gin.RouterGroup,
	routes []Router,
	zapLogger logger.Logger,
) {
	for _, route := range routes {
		for _, domain := range route.Domain {
			var handler []gin.HandlerFunc

			var endpoint string

			switch domain {
			case constant.Client:
				endpoint = path.Join(string(constant.Client), route.Path)

			default:
				zapLogger.Fatal(context.Background(), fmt.Sprintf("Invalid Domain %s Registered on Route", domain))
			}

			handler = append(handler, route.Middlewares...)
			handler = append(handler, route.Handler)
			grp.Handle(route.Method, endpoint, handler...)
		}
	}
}
