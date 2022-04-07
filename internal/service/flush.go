package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hengfeiyang/lsmdb/internal/pkg/db"
)

func Flush(c *gin.Context) {
	if err := db.DB.Flush(); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1, "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0, "message": "flush ok"})
	}
}
