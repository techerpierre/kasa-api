package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type AuthorizationHTTPHandler struct {
	app *gin.Engine
	api ports.AuthorizationsInput
}

func CreateAuthorizationHTTPHandler(app *gin.Engine, api ports.AuthorizationsInput) *AuthorizationHTTPHandler {
	return &AuthorizationHTTPHandler{
		app: app,
		api: api,
	}
}

func (h *AuthorizationHTTPHandler) RegisterRoutes() {
	h.app.POST("/authorizations", h.Create)
	h.app.PATCH("/authorizations", h.Update)
	h.app.DELETE("/authorizations", h.Delete)
	h.app.GET("/authorizations", h.List)
	h.app.GET("/authorizations/:id", h.FindOne)
}

func (*AuthorizationHTTPHandler) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create new authorization."})
}

func (*AuthorizationHTTPHandler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update a authorization."})
}

func (*AuthorizationHTTPHandler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete a authorization."})
}

func (*AuthorizationHTTPHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List authorizations."})
}

func (*AuthorizationHTTPHandler) FindOne(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Find a authorization."})
}
