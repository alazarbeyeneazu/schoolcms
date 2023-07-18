package errors

import (
	"net/http"

	"github.com/joomcode/errorx"
)

type ErrorType struct {
	StatusCode int
	ErrorType  *errorx.Type
}

var Error = []ErrorType{
	{
		StatusCode: http.StatusBadRequest,
		ErrorType:  ErrWriteError,
	},
	{
		StatusCode: http.StatusBadRequest,
		ErrorType:  ErrValidationError,
	},
}

// list of error namespaces
var (
	databaseError   = errorx.NewNamespace("database error").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	validationError = errorx.NewNamespace("validation error ").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
)

// list of errors types in all of the above namespaces

var (
	ErrWriteError      = errorx.NewType(databaseError, "unable to create")
	ErrReadError       = errorx.NewType(databaseError, "unable to read")
	ErrValidationError = errorx.NewType(validationError, "validation error")
)
