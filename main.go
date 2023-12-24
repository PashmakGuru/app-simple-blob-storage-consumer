//	@title			Simple Blob Storage Consumer API
//  @description	A simple implementation example for Azure blob storage

package main

import (
	"log"
	"os"

	"github.com/PashmakGuru/project-name/blobs"
	_ "github.com/PashmakGuru/project-name/docs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func mustPrepareEnvs() {
	keys := []string{
		"AZURE_CLIENT_ID",
		"AZURE_TENANT_ID",
		"AZURE_CLIENT_SECRET",
		"AZURE_STORAGE_ACCOUNT",
		"AZURE_STORAGE_CONTAINER",
	}

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	for _, key := range keys {
		value, exists := os.LookupEnv(key)
		if exists {
			os.Setenv(key, value)
		}
	}
}

func main() {
	mustPrepareEnvs()

	blobs.Prepare(os.Getenv("AZURE_STORAGE_CONTAINER"))

	log.Println("setting up web server (http://0.0.0.0:8080) for Simple Blob Storage Consumer...")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/blobs", blobs.HandleListBlobs)
	r.POST("/blobs", blobs.HandleCreateBlob)
	r.DELETE("/blobs/:blob", blobs.HandleDeleteBlob)

	r.Run(":8080")
}
