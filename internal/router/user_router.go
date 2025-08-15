package router

import (
	"fmt"
	"go-one/internal/router/controller/admin"
	"time"

	"github.com/gin-gonic/gin"
)

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
		//userRouter.GET("/get", func(c *gin.Context) {
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
		//userRouter.GET("/add", func(c *gin.Context) {
		//	c.JSON(http.StatusOK, gin.H{"message": "success"})
		//})

		userRouter.GET("/get", admin.UserController{}.Get)
		userRouter.GET("/add", initMiddleware, admin.UserController{}.Add)
	}

}

func initMiddleware(c *gin.Context) {
	//fmt.Println("initMiddleware", c.FullPath())
	fmt.Println("before Middleware")
	begin := time.Now()
	c.Next()
	after := time.Now()
	fmt.Println("after Middleware")
	duration := after.Sub(begin)
	fmt.Println("duration: ", duration)
}
