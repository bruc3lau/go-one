package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	ID       int64  `json:"id" form:"id"`
	UserName string `json:"user_name" form:"user_name"`
}

func UserInit(r *gin.Engine) {

	//r.GET("/getUser", func(c *gin.Context) {
	//	user := &UserInfo{}
	//	if err := c.ShouldBind(&user); err == nil {
	//		c.JSON(http.StatusOK, user)
	//	} else {
	//		c.JSON(http.StatusOK, gin.H{
	//			//"code":    500,
	//			"message": err.Error(),
	//		})
	//	}
	//})

	userRouter := r.Group("/user")
	{
		userRouter.GET("/get", func(c *gin.Context) {
			user := &UserInfo{}
			if err := c.ShouldBind(&user); err == nil {
				c.JSON(http.StatusOK, user)
			} else {
				c.JSON(http.StatusOK, gin.H{
					//"code":    500,
					"message": err.Error(),
				})
			}
		})
	}
}
