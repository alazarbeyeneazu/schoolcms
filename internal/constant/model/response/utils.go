package response

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func SendSuccessResponse(ctx *gin.Context, statusCode int, data interface{},
	metaData *MetaData) {

	ctx.JSON(
		statusCode,
		Response{
			OK:       true,
			MetaData: metaData,
			Data:     data,
		},
	)
}

func SendErrorResponse(ctx *gin.Context, err *ErrorResponse) {
	ctx.AbortWithStatusJSON(err.Code, Response{
		OK:    false,
		Error: err,
	})
}

func ErrorFields(err error) []FieldError {
	var errs []FieldError

	if data, ok := err.(validation.Errors); ok {
		for i, v := range data {
			errs = append(errs, FieldError{
				Name:        i,
				Description: v.Error(),
			},
			)
		}

		return errs
	}

	return nil
}
