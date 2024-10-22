package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/techerpierre/kasa-api/internal/application/dto"
	"github.com/techerpierre/kasa-api/internal/domain/entities"
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
	h.app.PATCH("/ratings/:id", h.Update)
	h.app.DELETE("/ratings/:id", h.Delete)
	h.app.GET("/ratings", h.List)
	h.app.GET("/ratings/:id", h.FindOne)
}

func (h *RatingHTTPHandler) Create(c *gin.Context) {
	var body dto.RatingInputDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": "Cannot parse body."}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	var ratingData entities.Rating
	dto.PipeInputDTOInRating(&body, &ratingData)

	rating, exception := h.api.Create(ratingData)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData dto.RatingDTO
	dto.PipeRatingInDTO(&rating, &responseData)

	response := dto.CreateResponse(http.StatusOK, responseData, nil)

	c.JSON(response.StatusCode, response)
}

func (h *RatingHTTPHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var body dto.RatingInputDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": "Cannot parse body."}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	var ratingData entities.Rating
	dto.PipeInputDTOInRating(&body, &ratingData)

	rating, exception := h.api.Update(id, ratingData)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData dto.RatingDTO
	dto.PipeRatingInDTO(&rating, &responseData)

	response := dto.CreateResponse(http.StatusOK, responseData, nil)

	c.JSON(response.StatusCode, response)
}

func (h *RatingHTTPHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	exception := h.api.Delete(id)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	response := dto.CreateResponse(http.StatusOK, gin.H{"message": "Rating deletion success."}, nil)

	c.JSON(response.StatusCode, response)
}

func (h *RatingHTTPHandler) List(c *gin.Context) {
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

	listing := entities.Listing{
		Page:     page,
		Pagesize: pagesize,
	}

	ratings, count, exception := h.api.List(listing)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData []dto.RatingDTO

	for _, rating := range ratings {
		var result dto.RatingDTO
		dto.PipeRatingInDTO(&rating, &result)
		responseData = append(responseData, result)
	}

	response := dto.CreateResponse(http.StatusOK, responseData, &count)

	c.JSON(response.StatusCode, response)
}

func (h *RatingHTTPHandler) FindOne(c *gin.Context) {
	id := c.Param("id")

	rating, exception := h.api.FindOne(id)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData dto.RatingDTO
	dto.PipeRatingInDTO(&rating, &responseData)

	response := dto.CreateResponse(http.StatusOK, responseData, nil)

	c.JSON(response.StatusCode, response)
}
