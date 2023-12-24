package blobs

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleList
// @summary List All Blobs
// @router /blobs [get]
// @produce json
// @success 200 {object} SuccessResponse[[]Blob]
// @failure 500 {object} FailureResponse
func HandleList(c *gin.Context) {
	blobs := make([]Blob, 0)

	pager := client.NewListBlobsFlatPager(container, nil)

	for pager.More() {
		// advance to the next page
		page, err := pager.NextPage(context.Background())
		if err != nil {
			Failure(c, http.StatusInternalServerError, err)
			return
		}

		// print the blob names for this page
		for _, blob := range page.Segment.BlobItems {
			blobs = append(blobs, Blob{
				Name:      *blob.Name,
				VersionID: blob.VersionID,
			})
		}
	}

	Success[[]Blob](c, http.StatusOK, blobs)
}
