package main

import (
	"context"
	"github.com/EDEN-NN/hydra-api/infra/database/mongodb"
	"github.com/gin-gonic/gin"
)

func main() {

	client, err := mongodb.Connect()
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(context.Background())

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
