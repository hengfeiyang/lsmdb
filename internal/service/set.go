package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Set(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": 0, "message": "ok"})
}
