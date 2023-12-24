package blobs

import (
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	_ "github.com/PashmakGuru/app-simple-blob-storage-consumer/docs"
)

var client *azblob.Client
var container string

func Prepare(containerName string) {
	container = containerName

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Panicln("unable to prepare Azure credentials", err)
	}

	client, err = azblob.NewClient("https://"+os.Getenv("AZURE_STORAGE_ACCOUNT")+".blob.core.windows.net/", cred, nil)
	if err != nil {
		log.Panicln("unable to create a Azure Blob client instance", err)
	}
}
