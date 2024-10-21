package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type RatingHTTPHandler struct {
	app *gin.Engine
	api ports.RatingInput
}

func CreateRatingHTTPHandler(app *gin.Engine, api ports.RatingInput) *RatingHTTPHandler {
	return &RatingHTTPHandler{
		app: app,
		api: api,
	}
}

func (h *RatingHTTPHandler) RegisterRoutes() {
	h.app.POST("/ratings", h.Create)
	h.app.PATCH("/ratings", h.Update)
	h.app.DELETE("/ratings", h.Delete)
	h.app.GET("/ratings", h.List)
	h.app.GET("/ratings/:id", h.FindOne)
}

func (*RatingHTTPHandler) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create new rating."})
}

func (*RatingHTTPHandler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update a rating."})
}

func (*RatingHTTPHandler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete a rating."})
}

func (*RatingHTTPHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List ratings."})
}

func (*RatingHTTPHandler) FindOne(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Find a rating."})
}
