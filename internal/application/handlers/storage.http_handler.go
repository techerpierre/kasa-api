package handlers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/techerpierre/kasa-api/internal/application/dto"
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type StorageHTTPHandler struct {
	app *gin.Engine
	api ports.StorageInput
}

func CreateStorageHTTPHandler(app *gin.Engine, api ports.StorageInput) *StorageHTTPHandler {
	return &StorageHTTPHandler{
		app: app,
		api: api,
	}
}

func (h *StorageHTTPHandler) RegisterRoutes() {
	h.app.POST("/upload", h.Upload)
	h.app.GET("/stream", h.Stream)
}

func (h *StorageHTTPHandler) Upload(c *gin.Context) {
	var body dto.CreateStorageInput

	if err := c.ShouldBindJSON(&body); err != nil {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": "Cannot parse body."}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	if !dto.IsValidStorageEndpoint(string(body.Endpoint)) {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": "This storage endpoint is not valid."}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	if strings.Contains(body.Indentifier, "/") || strings.Contains(body.Indentifier, ".") {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": `The specified ID can not contains "." or "/"`}, nil)
		c.JSON(response.StatusCode, response)
		return
	}

	if body.UploadMode == dto.UPLOAD_MODE_SINGLE {
		handler, err := c.FormFile("file")
		if err != nil {
			response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": "Unable to get file in the request body."}, nil)
			c.JSON(response.StatusCode, response)
			return
		}

		path, exception := h.uploadSignle(body, handler)

		if exception != nil {
			httpException, statusCode := dto.HTTPExceptionFromException(exception)
			response := dto.CreateResponse(statusCode, httpException, nil)
			c.JSON(statusCode, response)
			return
		}

		response := dto.CreateResponse(http.StatusAccepted, gin.H{"path": path}, nil)
		c.JSON(response.StatusCode, response)
	} else if body.UploadMode == dto.UPLOAD_MODE_MULTIPLE {
		form, err := c.MultipartForm()
		if err != nil {
			response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": "Unable to get files in the request body."}, nil)
			c.JSON(response.StatusCode, response)
			return
		}

		handlers := form.File["files"]
		if len(handlers) == 0 {
			response := dto.CreateResponse(http.StatusOK, gin.H{"message": "No files provided."}, nil)
			c.JSON(response.StatusCode, response)
			return
		}

		var paths []string

		for _, handler := range handlers {
			path, exception := h.uploadSignle(body, handler)

			if exception != nil {
				httpException, statusCode := dto.HTTPExceptionFromException(exception)
				response := dto.CreateResponse(statusCode, httpException, nil)
				c.JSON(statusCode, response)
				return
			}

			paths = append(paths, path)
		}

		response := dto.CreateResponse(http.StatusAccepted, gin.H{"path": paths}, nil)
		c.JSON(response.StatusCode, response)
	} else {
		response := dto.CreateResponse(http.StatusBadRequest, gin.H{"error": `Upload mode can only be one of "signle" or "multiple"`}, nil)
		c.JSON(response.StatusCode, response)
	}
}

func (h *StorageHTTPHandler) Stream(c *gin.Context) {
	path := c.Query("internalPath")
	fileData, exception := h.api.Read(path)

	if exception != nil {
		httpException, statusCode := dto.HTTPExceptionFromException(exception)
		response := dto.CreateResponse(statusCode, httpException, nil)
		c.JSON(statusCode, response)
		return
	}

	reader := bytes.NewReader(fileData)

	c.Stream(func(w io.Writer) bool {
		buffer := make([]byte, 8)
		n, err := reader.Read(fileData)
		if err != nil {
			return false
		}

		w.Write(buffer[:n])
		return true
	})
}

func (h *StorageHTTPHandler) uploadSignle(params dto.CreateStorageInput, handler *multipart.FileHeader) (string, *entities.Exception) {
	file, err := handler.Open()
	if err != nil {
		return "", entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			"Unable to open the specified file.",
		)
	}

	path := fmt.Sprintf("/%s/%s-%d", params.Endpoint, params.Indentifier, time.Now().UnixNano())

	return h.api.Write(path, file, handler)
}
