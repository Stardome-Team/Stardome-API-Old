package errors

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/stoewer/go-strcase"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// ErrorHandlerMiddleware :
func ErrorHandlerMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			var errorsList []ErrorObject

			for _, e := range c.Errors {
				switch e.Type {
				case gin.ErrorTypePublic:
					errorsList = append(errorsList, publicErrorToObject(e, c))
				case gin.ErrorTypeBind:
					if (reflect.TypeOf(e.Err) == reflect.TypeOf(validator.ValidationErrors{})) {
						errs := e.Err.(validator.ValidationErrors)

						for _, err := range errs {
							errorsList = append(errorsList, validationErrorToObject(err, c))
						}
					} else {
						errorsList = append(errorsList, ErrorObject{
							Domain:  c.Request.URL.Path,
							Message: ValidationError.Message,
							Reason:  ValidationError.Reason,
						})
					}
				default:
				}
			}

			if len(errorsList) != 0 {
				c.AbortWithStatusJSON(
					c.Writer.Status(),
					Error{
						Error: &ErrorResponse{
							StatusCode: c.Writer.Status(),
							Message:    errorsList[0].Message,
							Errors:     errorsList,
						},
					},
				)
			} else {
				c.AbortWithStatusJSON(
					http.StatusInternalServerError,
					Error{
						Error: &ErrorResponse{
							StatusCode: http.StatusInternalServerError,
							Message:    InternalError.Message,
							Errors: []ErrorObject{
								{
									Domain:  c.Request.URL.Path,
									Message: InternalError.Message,
									Reason:  InternalError.Reason,
								},
							},
						},
					},
				)
			}
		}
	}
}

func validationErrorToObject(e validator.FieldError, c *gin.Context) ErrorObject {
	var field string = strcase.LowerCamelCase(e.Field())
	var message string
	switch e.Tag() {
	case "required":
		message = fmt.Sprintf("%s is required", field)
	case "max":
		message = fmt.Sprintf("%s cannot be longer than %s", field, e.Param())
	case "min":
		message = fmt.Sprintf("%s must be longer than %s", field, e.Param())
	case "email":
		message = fmt.Sprintf("Invalid email format")
	case "len":
		message = fmt.Sprintf("%s must be %s characters long", field, e.Param())
	default:
		message = ValidationError.Message
	}

	if len(message) == 0 {
		message = fmt.Sprintf("%s is not valid", field)
	}

	return ErrorObject{
		Domain:  c.Request.URL.Path,
		Message: message,
		Reason:  ValidationError.Reason,
	}
}

func publicErrorToObject(e *gin.Error, c *gin.Context) ErrorObject {

	return ErrorObject{
		Domain:  c.Request.URL.Path,
		Message: e.Error(),
		Reason:  e.Meta.(string),
	}
}
