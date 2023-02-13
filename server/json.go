package server

import "github.com/gin-gonic/gin"

func json(c *gin.Context) {
	data, _ := c.Get("packetData")
	c.JSON(200, data)
}
