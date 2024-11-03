package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/techerpierre/kasa-api/internal/application/dto"
	"github.com/techerpierre/kasa-api/internal/domain/entities"
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

func (h *BookingHTTPHandler) Create(c *gin.Context) {
	var body dto.BookingInputDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": "Cannot parse body."}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	var bookingData entities.Booking
	dto.PipeInputDTOInBooking(&body, &bookingData)

	booking, exception := h.api.Create(bookingData)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData dto.BookingDTO
	dto.PipeBookingInDTO(&booking, &responseData)

	response := dto.CreateResponse(http.StatusOK, responseData, nil)

	c.JSON(response.StatusCode, response)
}

func (h *BookingHTTPHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var body dto.BookingInputDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": "Cannot parse body."}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	var bookingData entities.Booking
	dto.PipeInputDTOInBooking(&body, &bookingData)

	booking, exception := h.api.Update(id, bookingData)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData dto.BookingDTO
	dto.PipeBookingInDTO(&booking, &responseData)

	response := dto.CreateResponse(http.StatusOK, responseData, nil)

	c.JSON(response.StatusCode, response)
}

func (h *BookingHTTPHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	exception := h.api.Delete(id)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	response := dto.CreateResponse(http.StatusOK, gin.H{"message": "Booking deletion success."}, nil)

	c.JSON(response.StatusCode, response)
}

func (h *BookingHTTPHandler) List(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))

	if err != nil {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": `The "page" query must be an integer.`}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	pagesize, err := strconv.Atoi(c.DefaultQuery("pagesize", "10"))

	if err != nil {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": `The "pagesize" query must be an integer.`}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	var filters dto.BookingFiltersDTO

	if err := c.ShouldBindQuery(&filters); err != nil {
		response := dto.CreateResponse(http.StatusInternalServerError, gin.H{"error": "Cannot parse queries"}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	listing := entities.Listing{
		Page:     page,
		Pagesize: pagesize,
		Filters:  dto.MakeBookingFilters(filters),
	}

	bookings, count, exception := h.api.List(listing)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData []dto.BookingDTO

	for _, booking := range bookings {
		var result dto.BookingDTO
		dto.PipeBookingInDTO(&booking, &result)
		responseData = append(responseData, result)
	}

	response := dto.CreateResponse(http.StatusOK, responseData, &count)

	c.JSON(response.StatusCode, response)
}

func (h *BookingHTTPHandler) FindOne(c *gin.Context) {
	id := c.Param("id")

	booking, exception := h.api.FindOne(id)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData dto.BookingDTO
	dto.PipeBookingInDTO(&booking, &responseData)

	response := dto.CreateResponse(http.StatusOK, responseData, nil)

	c.JSON(response.StatusCode, response)
}
