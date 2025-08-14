package admin

import "github.com/gin-gonic/gin"

type BaseController struct {
}

func (con BaseController) success(c *gin.Context) {
	c.JSON(200, gin.H{"message": "success"})
}

func (con BaseController) error(c *gin.Context) {
	c.JSON(500, gin.H{"message": "error"})
}
