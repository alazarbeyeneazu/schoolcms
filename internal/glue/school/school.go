package school

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
	school rest.School,
	log logger.Logger, authDomains state.AuthDomains) {
	schoolRoutes := []routing.Router{
		{
			Method:      http.MethodPost,
			Path:        "schools",
			Handler:     school.CreateSchool,
			Middlewares: []gin.HandlerFunc{},
			Domain: []state.Domain{
				authDomains.System,
			},
		}, {
			Method:      http.MethodPost,
			Path:        "school/students",
			Handler:     school.AssignStudentToSchool,
			Middlewares: []gin.HandlerFunc{},
			Domain: []state.Domain{
				authDomains.System,
				authDomains.Corporate,
			},
		},
		{
			Method:      http.MethodGet,
			Path:        "schools",
			Handler:     school.GetAllSchools,
			Middlewares: []gin.HandlerFunc{},
			Domain: []state.Domain{
				authDomains.System,
				authDomains.Corporate,
			},
		}, {
			Method:      http.MethodGet,
			Path:        "schools/:id",
			Handler:     school.GetSchoolByID,
			Middlewares: []gin.HandlerFunc{},
			Domain: []state.Domain{
				authDomains.System,
				authDomains.Corporate,
			},
		},
	}
	routing.RegisterRoute(group, schoolRoutes, log, authDomains)
}
