package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/towelong/vgo/errors"
	customValidator "github.com/towelong/vgo/pkg/validate"
	"strings"
)

func Error(ctx *gin.Context) {
	ctx.Next()
	if len(ctx.Errors) > 0 {
		e := ctx.Errors.Last()
		switch err := e.Err.(type) {
		case *errors.Error:
			ctx.JSON(err.Code, err)
		case validator.ValidationErrors:
			wrapError(ctx, err)
		case *validator.ValidationErrors:
			wrapError(ctx, *err)
		default:
			defaultError := errors.Unknown
			ctx.JSON(defaultError.Code, defaultError)
		}
	}
}

func wrapError(ctx *gin.Context, err validator.ValidationErrors) {
	mapErrors := make(map[string]string)
	var (
		errString string
		ce        *errors.Error
	)
	for _, v := range err {
		errString = v.Translate(customValidator.Trans)
		filedName := strings.ToLower(v.StructField())
		mapErrors[filedName] = errString
	}
	ce = errors.ParamsErr
	if len(err) > 1 {
		ce.Message = mapErrors
	} else {
		ce.Message = errString
	}
	ctx.JSON(ce.Code, ce)
}
