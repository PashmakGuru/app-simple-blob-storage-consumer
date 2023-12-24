package blobs

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Blob struct {
	Name string `json:"name"`
}

// HandleListBlobs
// @summary List Blob Storage Items
// @success 200 {object} []Blob
// @router /blobs [get]
func HandleListBlobs(c *gin.Context) {
	log.Println("Indexing blobs...")

	var blobs []Blob

	pager := client.NewListBlobsFlatPager(container, nil)

	for pager.More() {
		// advance to the next page
		page, err := pager.NextPage(context.Background())
		if err != nil {
			panic(err)
		}

		// print the blob names for this page
		for _, blob := range page.Segment.BlobItems {
			blobs = append(blobs, Blob{
				Name: *blob.Name,
			})
		}
	}

	c.JSON(http.StatusOK, blobs)
}
