package rest

import (
	"github.com/gin-gonic/gin"
)

type User interface {
	CreatUser(c *gin.Context)
}
