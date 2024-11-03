package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/techerpierre/kasa-api/internal/application/dto"
	"github.com/techerpierre/kasa-api/internal/domain/entities"
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

func (h *AccomodationHTTPHandler) Create(c *gin.Context) {
	var body dto.AccommodationInputDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": "Cannot parse body."}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	var accommodationData entities.Accommodation
	dto.PipeInputDTOInAccommodation(&body, &accommodationData)

	accommodation, exception := h.api.Create(accommodationData)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData dto.AccommodationDTO
	dto.PipeAccommodationInDTO(&accommodation, &responseData)

	response := dto.CreateResponse(http.StatusOK, responseData, nil)

	c.JSON(response.StatusCode, response)
}

func (h *AccomodationHTTPHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var body dto.AccommodationInputDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": "Cannot parse body."}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	var accommodationData entities.Accommodation
	dto.PipeInputDTOInAccommodation(&body, &accommodationData)

	accommodation, exception := h.api.Update(id, accommodationData)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData dto.AccommodationDTO
	dto.PipeAccommodationInDTO(&accommodation, &responseData)

	response := dto.CreateResponse(http.StatusOK, responseData, nil)

	c.JSON(response.StatusCode, response)
}

func (h *AccomodationHTTPHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	exception := h.api.Delete(id)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	response := dto.CreateResponse(http.StatusOK, gin.H{"message": "Accommodation deletion success."}, nil)

	c.JSON(response.StatusCode, response)
}

func (h *AccomodationHTTPHandler) List(c *gin.Context) {
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

	var filters dto.AccommodationFiltersDTO

	if err := c.ShouldBindQuery(&filters); err != nil {
		response := dto.CreateResponse(http.StatusInternalServerError, gin.H{"error": "Cannot parse queries"}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	listing := entities.Listing{
		Page:     page,
		Pagesize: pagesize,
		Filters:  dto.MakeAccommodationFilters(filters),
	}

	accommodations, count, exception := h.api.List(listing)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData []dto.AccommodationDTO

	for _, accommodation := range accommodations {
		var result dto.AccommodationDTO
		dto.PipeAccommodationInDTO(&accommodation, &result)
		responseData = append(responseData, result)
	}

	response := dto.CreateResponse(http.StatusOK, responseData, &count)

	c.JSON(response.StatusCode, response)
}

func (h *AccomodationHTTPHandler) FindOne(c *gin.Context) {
	id := c.Param("id")

	accommodation, exception := h.api.FindOne(id)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData dto.AccommodationDTO
	dto.PipeAccommodationInDTO(&accommodation, &responseData)

	response := dto.CreateResponse(http.StatusOK, responseData, nil)

	c.JSON(response.StatusCode, response)
}
