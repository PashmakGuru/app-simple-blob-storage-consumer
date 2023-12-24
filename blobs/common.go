package blobs

import (
	"github.com/gin-gonic/gin"
)

type Blob struct {
	Name      string  `json:"name"`
	VersionID *string `json:"version_id,omitempty"`
}

type SuccessResponse[T any] struct {
	Data T `json:"data"`
}

type FailureResponse struct {
	Error string `json:"error"`
}

func Failure(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{"error": err.Error()})
}

func Success[T any](c *gin.Context, code int, data T) {
	c.JSON(code, SuccessResponse[T]{
		Data: data,
	})
}
