package middlewares

import (
	"net/http"
	"strings"

	"github.com/Stardome-Team/Stardome-API/libraries/go/jwt"
	"github.com/Stardome-Team/Stardome-API/services/player-service/models"
	"github.com/Stardome-Team/Stardome-API/services/player-service/utils"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

// AuthHandlerMiddleware :
func AuthHandlerMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := extractToken(c)

		if token == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.Result{
				Error: models.Error{
					Error: &models.ErrorObject{
						Code:    http.StatusUnauthorized,
						Message: utils.ErrorAuthorizationTokenNotFound.Error(),
						Errors: []models.ErrorsObject{{
							Domain:  c.Request.URL.Path,
							Message: utils.ErrorAuthorizationTokenNotFound.Error(),
							Reason:  utils.ReasonAuthorizationFailed,
						}},
					},
				},
			})
			return
		}

		err := jwt.VerifyToken(*token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.Result{
				Error: models.Error{
					Error: &models.ErrorObject{
						Code:    http.StatusUnauthorized,
						Message: utils.ErrorAuthorizationVerificationFailed.Error(),
						Errors: []models.ErrorsObject{{
							Domain:  c.Request.URL.Path,
							Message: utils.ErrorAuthorizationVerificationFailed.Error(),
							Reason:  utils.ReasonAuthorizationFailed,
						}},
					}},
			})
			return
		}

		hasExpired := jwt.HasTokenExpired(*token)

		if hasExpired {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.Result{
				Error: models.Error{
					Error: &models.ErrorObject{
						Code:    http.StatusUnauthorized,
						Message: utils.ErrorAuthorizationVerificationFailed.Error(),
						Errors: []models.ErrorsObject{{
							Domain:  c.Request.URL.Path,
							Message: utils.ErrorAuthorizationVerificationFailed.Error(),
							Reason:  utils.ReasonAuthorizationFailed,
						}},
					}},
			})
			return
		}

		c.Next()
	}
}

func extractToken(c *gin.Context) *string {

	authHeader := c.GetHeader(authorizationHeader)

	authArr := strings.Split(authHeader, " ")

	if len(authArr) == 2 {
		return &authArr[1]
	}

	return nil
}
