package handler

import (
	"github.com/gin-gonic/gin"
	"go-rest-gin-k8s/platform/repo"
	"net/http"
)

func BulletinsGet() gin.HandlerFunc  {
	return func(c *gin.Context){
		results,_ := repo.GetBulletins()
		c.JSON(http.StatusOK, results)
	}
}