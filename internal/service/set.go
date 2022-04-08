package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hengfeiyang/lsmdb/internal/pkg/lsm"
)

func Set(c *gin.Context) {
	key := c.Param("key")
	val, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 1, "message": err.Error()})
		return
	}
	if key == "" {
		key = uuid.New().String()
	}
	lsm.DB.Set(key, string(val))
	c.JSON(http.StatusOK, gin.H{"status": 0, "message": "ok"})
}
