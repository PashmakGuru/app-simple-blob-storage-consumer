package blobs

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleDelete
// @summary Delete a Blob
// @router /blobs/{blob} [delete]
// @param blob path string true "Blob Name"
// @produce json
// @success 200 {object} SuccessResponse[Blob] "Blob is deleted successfully"
// @failure 400 {object} FailureResponse "There was an error deleting the blob"
func HandleDelete(c *gin.Context) {
	blobName := c.Param("blob")

	_, err := client.DeleteBlob(context.Background(), container, blobName, nil)
	if err != nil {
		Failure(c, http.StatusBadRequest, err)
		return
	}

	Success[Blob](c, http.StatusOK, Blob{
		Name: blobName,
	})
}
