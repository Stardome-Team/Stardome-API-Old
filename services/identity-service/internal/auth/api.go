package auth

import (
	"github.com/Stardome-Team/Stardome-API/services/identity-service/pkg/log"
	"github.com/gin-gonic/gin"
)

const (
	authorizationEndpoint = "/auth/player"
)

type handler struct {
	service Service
	logger  log.Logger
}

// AuthenticationRequest request model sent to authenticate players
type AuthenticationRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Controller contains interface for authentication handlers
type Controller interface {
	AuthorizeUser(c *gin.Context)
}

// CreateHandlers sets up routing to the HTTP request
func CreateHandlers(r *gin.RouterGroup, s Service, l log.Logger) {
	h := &handler{service: s, logger: l}

	registerHandlers(r, h)
}

func registerHandlers(r *gin.RouterGroup, ctr Controller) {
	r.POST(authorizationEndpoint, ctr.AuthorizeUser)
}

// AuthorizeUser handler function to authorize the user
func (h *handler) AuthorizeUser(c *gin.Context) {

	var request AuthenticationRequest

	if err := c.BindJSON(&request); err != nil {
		h.logger.Errorf("invalid request: %v", err)
		return
	}

}
