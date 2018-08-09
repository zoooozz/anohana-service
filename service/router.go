package service

import (
	"github.com/labstack/echo"
)

func initRouter(e *echo.Echo) {
	e.GET("/", index)
	e.POST("/x/login", login)
	admin := e.Group("/x/", Adminmiddleware())
	{
		admin.GET("user", user)
		admin.POST("user/edit", userEdit)
		admin.POST("user/add", index)
		admin.POST("upload", index)
	}

}
