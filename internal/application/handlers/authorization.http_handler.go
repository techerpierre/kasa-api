package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/techerpierre/kasa-api/internal/application/dto"
	"github.com/techerpierre/kasa-api/internal/domain/entities"
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

func (h *AuthorizationHTTPHandler) Create(c *gin.Context) {
	var body dto.AuthorizationsInputDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": "Cannot parse body."}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	var authorizationData entities.Authorizations
	dto.PipeInputDTOInAuthorizations(&body, &authorizationData)

	authorization, exception := h.api.Create(authorizationData)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData dto.AuthorizationsDTO
	dto.PipeAuthorizationsInDTO(&authorization, &responseData)

	response := dto.CreateResponse(http.StatusOK, responseData, nil)

	c.JSON(response.StatusCode, response)
}

func (h *AuthorizationHTTPHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var body dto.AuthorizationsInputDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": "Cannot parse body."}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	var authorizationData entities.Authorizations
	dto.PipeInputDTOInAuthorizations(&body, &authorizationData)

	authorization, exception := h.api.Update(id, authorizationData)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData dto.AuthorizationsDTO
	dto.PipeAuthorizationsInDTO(&authorization, &responseData)

	response := dto.CreateResponse(http.StatusOK, responseData, nil)

	c.JSON(response.StatusCode, response)
}

func (h *AuthorizationHTTPHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	exception := h.api.Delete(id)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	response := dto.CreateResponse(http.StatusOK, gin.H{"error": "Cannot parse body."}, nil)

	c.JSON(response.StatusCode, response)
}

func (h *AuthorizationHTTPHandler) List(c *gin.Context) {
	authorizations, count, exception := h.api.List()

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData []dto.AuthorizationsDTO

	for _, authorization := range authorizations {
		var result dto.AuthorizationsDTO
		dto.PipeAuthorizationsInDTO(&authorization, &result)
		responseData = append(responseData, result)
	}

	response := dto.CreateResponse(http.StatusOK, responseData, &count)

	c.JSON(response.StatusCode, responseData)
}

func (h *AuthorizationHTTPHandler) FindOne(c *gin.Context) {
	id := c.Param("id")

	authorizations, exception := h.api.FindOne(id)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	var responseData dto.AuthorizationsDTO
	dto.PipeAuthorizationsInDTO(&authorizations, &responseData)

	response := dto.CreateResponse(http.StatusOK, responseData, nil)

	c.JSON(response.StatusCode, response)
}
