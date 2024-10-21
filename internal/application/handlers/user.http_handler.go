package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type UserHTTPHandler struct {
	app *gin.Engine
	api ports.UserInput
}

func CreateUserHTTPHandler(app *gin.Engine, api ports.UserInput) *UserHTTPHandler {
	return &UserHTTPHandler{
		app: app,
		api: api,
	}
}

func (h *UserHTTPHandler) RegisterRoutes() {
	h.app.POST("/users", h.Create)
	h.app.PATCH("/users", h.Update)
	h.app.DELETE("/users", h.Delete)
	h.app.GET("/users", h.List)
	h.app.GET("/users/profile", h.Profile)
	h.app.GET("/users/:id", h.FindOne)
}

func (*UserHTTPHandler) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create new user."})
}

func (*UserHTTPHandler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update a user."})
}

func (*UserHTTPHandler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete a user."})
}

func (*UserHTTPHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List users."})
}

func (*UserHTTPHandler) FindOne(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Find a user."})
}

func (*UserHTTPHandler) Profile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get your profile."})
}
