package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type AccomodationHTTPHandler struct {
	app *gin.Engine
	api ports.AccommodationInput
}

func CreateAccomodationHTTPHandler(app *gin.Engine, api ports.AccommodationInput) *AccomodationHTTPHandler {
	return &AccomodationHTTPHandler{
		app: app,
		api: api,
	}
}

func (h *AccomodationHTTPHandler) RegisterRoutes() {
	h.app.POST("/accommodations", h.Create)
	h.app.PATCH("/accommodations", h.Update)
	h.app.DELETE("/accommodations", h.Delete)
	h.app.GET("/accommodations", h.List)
	h.app.GET("/accommodations/:id", h.FindOne)
}

func (*AccomodationHTTPHandler) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create new accommodation."})
}

func (*AccomodationHTTPHandler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update a accommodation."})
}

func (*AccomodationHTTPHandler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete a accommodation."})
}

func (*AccomodationHTTPHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List accommodations."})
}

func (*AccomodationHTTPHandler) FindOne(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Find a accommodation."})
}
