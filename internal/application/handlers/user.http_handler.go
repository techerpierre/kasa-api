package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/techerpierre/kasa-api/internal/application/dto"
	"github.com/techerpierre/kasa-api/internal/domain/entities"
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
	h.app.PATCH("/users/:id", h.Update)
	h.app.DELETE("/users/:id", h.Delete)
	h.app.GET("/users", h.List)
	h.app.GET("/users/profile", h.Profile)
	h.app.GET("/users/:id", h.FindOne)
}

func (h *UserHTTPHandler) Create(c *gin.Context) {
	var body dto.UserInputDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": "Cannot parse body"}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	var userData entities.User
	dto.PipeInputDTOInUser(&body, &userData)

	user, exception := h.api.Create(userData)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData dto.UserDTO
	dto.PipeUserInDTO(&user, &responseData)

	response := dto.CreateResponse(http.StatusOK, responseData, nil)

	c.JSON(response.StatusCode, response)
}

func (h *UserHTTPHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var body dto.UserInputDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": "Cannot parse body"}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	var userData entities.User
	dto.PipeInputDTOInUser(&body, &userData)

	user, exception := h.api.Update(id, userData)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData dto.UserDTO
	dto.PipeUserInDTO(&user, &responseData)

	response := dto.CreateResponse(http.StatusOK, responseData, nil)

	c.JSON(http.StatusOK, response)
}

func (h *UserHTTPHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	exception := h.api.Delete(id)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	response := dto.CreateResponse(http.StatusOK, gin.H{"message": "User deletion success."}, nil)

	c.JSON(response.StatusCode, response)
}

func (h *UserHTTPHandler) List(c *gin.Context) {
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

	var filters dto.UserFiltersDTO

	if err := c.ShouldBindQuery(&filters); err != nil {
		response := dto.CreateResponse(http.StatusInternalServerError, gin.H{"error": "Cannot parse queries"}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	listing := entities.Listing{
		Page:     page,
		Pagesize: pagesize,
		Filters:  dto.MakeUserFilters(filters),
	}

	users, count, exception := h.api.List(listing)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData []dto.UserDTO

	for _, user := range users {
		var result dto.UserDTO
		dto.PipeUserInDTO(&user, &result)
		responseData = append(responseData, result)
	}

	response := dto.CreateResponse(http.StatusOK, responseData, &count)

	c.JSON(http.StatusOK, response)
}

func (h *UserHTTPHandler) FindOne(c *gin.Context) {
	id := c.Param("id")

	user, exception := h.api.FindOne(id)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData dto.UserDTO
	dto.PipeUserInDTO(&user, &responseData)

	response := dto.CreateResponse(http.StatusOK, responseData, nil)

	c.JSON(http.StatusOK, response)
}

func (*UserHTTPHandler) Profile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Method not implemented."})
}
