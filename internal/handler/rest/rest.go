package rest

import (
	"github.com/gin-gonic/gin"
)

type User interface {
	CreatUser(c *gin.Context)
}

type School interface {
	CreateSchool(c *gin.Context)
	AssignStudentToSchool(c *gin.Context)
}

type Teacher interface {
	CreateTeacher(c *gin.Context)
	AssignTeachersToSchool(c *gin.Context)
}

type Grade interface {
	CreateGrade(c *gin.Context)
}

type Student interface {
	CreateStudent(c *gin.Context)
}

type Family interface {
	CreateFamily(C *gin.Context)
	AssignFamilyToStudent(c *gin.Context)
}
