package dto

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindReq(c *gin.Context, req interface{}) error {
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	return nil
}

func Response(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}
