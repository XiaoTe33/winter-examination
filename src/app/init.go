package app

import (
	"winter-examination/src/conf"

	"github.com/gin-gonic/gin"
)

func InitRouters() {
	r := gin.Default()
	r.GET("/index", Index)

	u := r.Group("/user")
	{
		u.POST("/login", Login)
		u.POST("/register", Register)
		u.GET("/logout", Logout)
	}

	g := r.Group("/goods")
	{
		g.POST("/add")
		g.POST("/delete")
		g.GET("/query/id/:id")
		g.GET("/query/name/:name")
		g.GET("/query/all")
	}

	r.Run(conf.Port)
}
