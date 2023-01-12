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
		u.POST("/login", Login)          //登录
		u.POST("/register", Register)    //注册
		u.POST("/logout", JWT(), Logout) //退出登录
		u.POST("/query", QueryUser)
		u.POST("/update", JWT(), UpdateUser)
		u.POST("/queryAll", QueryAllUsers)
	}

	g := r.Group("/goods") //商品模块
	{
		g.POST("/add", JWT(), AddGoods) //新增商品
		g.POST("/update", JWT(), UpdateGoods)
		g.POST("/delete", JWT(), DeleteGoods) //删除商品
		g.POST("/query", QueryGoods)          //查找商品
		g.POST("/queryAll", QueryAllGoods)
	}

	s := r.Group("/shop") //商店模块
	{
		s.POST("/add", JWT(), AddShop)
		s.POST("/delete", JWT(), DeleteShop)
		s.POST("/update", JWT(), UpdateShop)
		s.POST("/query", QueryShops)
		s.POST("/queryAll", QueryAllShops)
	}

	o := r.Group("/order")
	{
		o.POST("/add", JWT(), AddOrder)
		o.POST("/query", QueryOrders)
		o.POST("/queryAll", QueryAllOrders)
	}

	r.Run(conf.BackEndPort)
}
func initFrontEndRouters() {
	r := gin.Default()
	r.Static("templates/", "./templates")
	r.LoadHTMLGlob("templates/html/*")

	r.GET("/register", FRegister)
	r.GET("/login", FLogin)
	r.POST("/jwt", JWT(), FLogin)

	r.Run(conf.FrontEndPort)
}
