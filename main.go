package main

import (
	"github.com/gin-gonic/gin"
	"go-rest-gin-k8s/httpd/handler"
	"go-rest-gin-k8s/platform/repo"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/board", handler.BulletinsGet())
	r.POST("/board", handler.NewsBulletinPost())
	repo.Migrations()

	if err := r.Run(":8080"); err != nil {
		log.Println("Error: ", err)
		panic(err)
	}
}
