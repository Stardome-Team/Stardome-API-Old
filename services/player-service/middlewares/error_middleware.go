package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Stardome-Team/Stardome-API/services/player-service/utils"

	"github.com/Stardome-Team/Stardome-API/services/player-service/models"

	"github.com/stoewer/go-strcase"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// ErrorHandlerMiddleware :
func ErrorHandlerMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			var errorsList []models.ErrorsObject

			for _, e := range c.Errors {
				switch e.Type {
				case gin.ErrorTypePublic:
					errorsList = append(errorsList, PublicErrorToObject(e, c))
				case gin.ErrorTypeBind:
					errs := e.Err.(validator.ValidationErrors)

					for _, err := range errs {
						errorsList = append(errorsList, ValidationErrorToObject(err, c))
					}
				default:
					// Report Error
				}

			}

			if len(errorsList) != 0 {
				if !c.Writer.Written() {
					c.AbortWithStatusJSON(c.Writer.Status(), models.Result{

						Error: models.Error{
							Error: &models.ErrorObject{
								Code:    c.Writer.Status(),
								Message: errorsList[0].Message,
								Errors:  errorsList,
							},
						},
					})
				}
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, models.Result{
					Error: models.Error{
						Error: &models.ErrorObject{
							Code:    http.StatusInternalServerError,
							Message: utils.ErrorInternalError.Error(),
							Errors: []models.ErrorsObject{{
								Domain:  c.Request.URL.Path,
								Message: utils.ErrorInternalError.Error(),
								Reason:  utils.ReasonInternalServer,
							}},
						},
					},
				})
			}
		}
	}
}

// ValidationErrorToObject :
func ValidationErrorToObject(e validator.FieldError, c *gin.Context) models.ErrorsObject {
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
	}

	if len(message) == 0 {
		message = fmt.Sprintf("%s is not valid", field)
	}

	return models.ErrorsObject{
		Domain:  c.Request.URL.Path,
		Message: message,
		Reason:  utils.ReasonFieldValidationError,
	}
}

// PublicErrorToObject :
func PublicErrorToObject(e *gin.Error, c *gin.Context) models.ErrorsObject {

	return models.ErrorsObject{
		Domain:  c.Request.URL.Path,
		Message: e.Error(),
		Reason:  e.Meta.(string),
	}
}
