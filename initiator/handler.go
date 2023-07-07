package initiator

import (
	"context"
	"schoolcms/internal/handler/rest"
	"schoolcms/internal/handler/rest/user"
	"schoolcms/platform/logger"
	"time"
)

type Handler struct {
	User rest.User
}

func InitHandler(ctx context.Context, module Module, log logger.Logger, timeout time.Duration) Handler {
	return Handler{
		User: user.Init(ctx, module.User, log.Named("user handler"), timeout),
	}
}
