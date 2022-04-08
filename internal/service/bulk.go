package service

import (
	"bufio"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hengfeiyang/lsmdb/internal/pkg/lsm"
)

func Bulk(c *gin.Context) {
	prefix := c.Query("prefix")
	buf := bufio.NewReader(c.Request.Body)
	n := 0
	var key string
	for {
		n++
		val, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			c.JSON(http.StatusOK, gin.H{"status": 1, "message": err.Error()})
			return
		}
		if prefix != "" {
			key = fmt.Sprintf("%s-%d", prefix, n)
		} else {
			key = uuid.New().String()
		}
		lsm.DB.Set(key, string(val))
	}
	c.JSON(http.StatusOK, gin.H{"status": 0, "message": "ok", "count": n})
}
