package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": 0, "message": "ok"})
}
