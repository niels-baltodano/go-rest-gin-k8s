package handler

import (
	"github.com/gin-gonic/gin"
	"go-rest-gin-k8s/platform/bulletin"
	"go-rest-gin-k8s/platform/repo"
	"net/http"
	"time"
)

func NewsBulletinPost() gin.HandlerFunc {
	return func( c *gin.Context) {
		var b bulletin.Bulletin
		if c.Bind(&b) == nil {
			b.CreatedAt = time.Now()
			if err := repo.AddBulletin(b); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "internal error" + err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	}

}