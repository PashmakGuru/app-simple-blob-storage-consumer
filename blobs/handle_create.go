package blobs

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// HandleCreateBlob
// @summary Create a new Storage Blob
// @router /blobs [post]
// @accept json
// @produce json
// @param json body CreateRequest true "Details of the blob"
// @success 201 {object} SuccessResponse[Blob] "Blob is created successfully"
// @failure 400 {object} FailureResponse "There was an error creating the blob"
func HandleCreateBlob(c *gin.Context) {
	var input CreateRequest

	if err := c.BindJSON(&input); err != nil {
		Failure(c, http.StatusBadRequest, err)
		return
	}

	_, err := client.UploadBuffer(context.Background(), container, input.Name, []byte(input.Content), nil)
	if err != nil {
		Failure(c, http.StatusBadRequest, err)
		return
	}

	Success[Blob](c, http.StatusCreated, Blob{
		Name: input.Name,
	})
}
