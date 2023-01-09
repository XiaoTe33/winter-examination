package app

import (
	"time"
	"winter-examination/src/conf"

	"github.com/gin-gonic/gin"
)

func InitRouters() {
	go initBackEndRouters()
	gin.SetMode(gin.ReleaseMode)
	time.Sleep(3 * time.Second)
	go initFrontEndRouters()
	forever := make(chan int)
	<-forever
}

func initBackEndRouters() {
	r := gin.Default()
	r.Use(Cors())
	u := r.Group("/user") //用户模块
	{
		u.POST("/login", Login)       //登录
		u.POST("/register", Register) //注册
		u.GET("/logout", Logout)      //退出登录
	}

	g := r.Group("/goods") //商品模块
	{
		g.POST("/add", AddGoods)       //新增商品
		g.POST("/delete", DeleteGoods) //删除商品
		g.POST("/query", QueryGoods)   //查找商品
	}
	r.Run(conf.BackEndPort)
}
func initFrontEndRouters() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/html/*")

	r.GET("/register", FRegister)
	r.GET("/login", FLogin)
	r.POST("/jwt", JWT(), FLogin)

	r.Run(conf.FrontEndPort)
}
