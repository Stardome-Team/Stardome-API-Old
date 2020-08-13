package auth

import (
	"github.com/Blac-Panda/Stardome-API/services/identity-service/pkg/log"
	"github.com/gin-gonic/gin"
)

const (
	authorizationEndpoint = "/authorize"
)

// Request request model sent to authenticate players
type Request struct {
	UserName string
	Password string
}

// Controller contains interface for authentication handlers
type Controller interface {
	AuthorizeUser(c *gin.Context)
}

type handler struct {
	service Service
	logger  log.Logger
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
	c.JSON(200, map[string]interface{}{"hello": "world"})

	return
}
