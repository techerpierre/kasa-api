package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type BookingHTTPHandler struct {
	app *gin.Engine
	api ports.BookingInput
}

func CreateBookingHTTPHandler(app *gin.Engine, api ports.BookingInput) *BookingHTTPHandler {
	return &BookingHTTPHandler{
		app: app,
		api: api,
	}
}

func (h *BookingHTTPHandler) RegisterRoutes() {
	h.app.POST("/bookings", h.Create)
	h.app.PATCH("/bookings", h.Update)
	h.app.DELETE("/bookings", h.Delete)
	h.app.GET("/bookings", h.List)
	h.app.GET("/bookings/:id", h.FindOne)
}

func (*BookingHTTPHandler) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create new booking."})
}

func (*BookingHTTPHandler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update a booking."})
}

func (*BookingHTTPHandler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete a booking."})
}

func (*BookingHTTPHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List bookings."})
}

func (*BookingHTTPHandler) FindOne(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Find a booking."})
}
