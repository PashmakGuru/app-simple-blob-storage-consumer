package blobs

import (
	"context"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShowResponse struct {
	Blob    Blob   `json:"blob"`
	Content string `json:"content"`
}

// HandleShowBlob
// @summary Get a Blob with Content
// @router /blobs/{blob} [get]
// @param blob path string true "Blob Name"
// @produce json
// @success 200 {object} SuccessResponse[ShowResponse] "Blob is fetched"
// @failure 400 {object} FailureResponse "There was an error fetching the blob"
// @failure 500 {object} FailureResponse "There was an error fetching the blob"
func HandleShowBlob(c *gin.Context) {
	blobName := c.Param("blob")

	blobDownloadResponse, err := client.DownloadStream(context.TODO(), container, blobName, nil)
	if err != nil {
		Failure(c, http.StatusBadRequest, err)
		return
	}

	reader := blobDownloadResponse.Body
	content, err := io.ReadAll(reader)
	if err != nil {
		Failure(c, http.StatusInternalServerError, err)
		return
	}

	Success[ShowResponse](c, http.StatusOK, ShowResponse{
		Blob: Blob{
			Name: blobName,
		},
		Content: string(content),
	})
}
