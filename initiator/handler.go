package initiator

import (
	"loyalty/platform/logger"
	"time"
)

type Handler struct {
}

func InitHandler(module Module, log logger.Logger, timeout time.Duration) Handler {
	return Handler{}
}
