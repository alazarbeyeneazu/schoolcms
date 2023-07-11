package rest

import (
	"github.com/gin-gonic/gin"
)

type User interface {
	CreatUser(c *gin.Context)
}

type School interface {
	CreateSchool(c *gin.Context)
}

type Teacher interface {
	CreateTeacher(c *gin.Context)
	AssignTeachersToSchool(c *gin.Context)
}

type Grade interface {
	CreateGrade(c *gin.Context)
}
