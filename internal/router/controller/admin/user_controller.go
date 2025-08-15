package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	ID       int64  `json:"id" form:"id"`
	UserName string `json:"user_name" form:"user_name"`
}

type UserController struct {
	BaseController
}

func (u UserController) Add(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{"message": "success"})
	fmt.Println("Add success")
	u.success(c)
}

func (u UserController) Get(c *gin.Context) {
	user := &UserInfo{}
	if err := c.ShouldBind(&user); err == nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusOK, gin.H{
			//"code":    500,
			"message": err.Error(),
		})
	}
}
