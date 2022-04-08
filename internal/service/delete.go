package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hengfeiyang/lsmdb/internal/pkg/lsm"
)

func Delete(c *gin.Context) {
	key := c.Param("key")
	lsm.DB.Delete(key)
	c.JSON(http.StatusOK, gin.H{"status": 0, "message": "ok"})
}
