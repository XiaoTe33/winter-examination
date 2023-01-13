package app

import (
	"time"

	"winter-examination/src/conf"

	"github.com/gin-gonic/gin"
)

func InitRouters() {
	go initBackEndRouters()
	gin.SetMode(gin.ReleaseMode)
	time.Sleep(1 * time.Second)
	go initFrontEndRouters()
	forever := make(chan int)
	<-forever
}

func initBackEndRouters() {

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.Use(Cors())
	u := r.Group("/user") //用户模块
	{
		u.POST("/login", Login)              //登录
		u.POST("/register", Register)        //注册
		u.POST("/logout", JWT(), Logout)     //退出登录
		u.POST("/query", QueryUser)          //查用户
		u.POST("/update", JWT(), UpdateUser) //修改用户信息
		u.POST("/queryAll", QueryAllUsers)   //查看当前所有用户
		u.POST("/photo", AddUserPhoto)       //加头像
	}

	g := r.Group("/goods") //商品模块
	{
		g.POST("/add", JWT(), AddGoods)          //新增商品
		g.POST("/update", JWT(), UpdateGoods)    //修改商品信息
		g.POST("/delete", JWT(), DeleteGoods)    //删除商品
		g.POST("/query", QueryGoods)             //查找商品
		g.POST("/queryAll", QueryAllGoods)       //查看所有商品信息
		g.POST("/shoppingCar", GoodsShoppingCar) //购物车操作
	}

	s := r.Group("/shop") //商店模块
	{
		s.POST("/add", JWT(), AddShop)       //新增商店
		s.POST("/delete", JWT(), DeleteShop) //删除商店
		s.POST("/update", JWT(), UpdateShop) //更新商店信息
		s.POST("/query", QueryShops)         //查询商店
		s.POST("/queryAll", QueryAllShops)   //查看当前所有商店信息
	}

	o := r.Group("/order")
	{
		o.POST("/add", JWT(), AddOrder)                     //新增订单
		o.POST("/updateStatus", JWT(), UpdateOrderStatus)   //修改订单状态
		o.POST("/updateAddress", JWT(), UpdateOrderAddress) //修改订单地址
		o.POST("/query", QueryOrders)                       //查询订单
		o.POST("/queryAll", QueryAllOrders)                 //查看所有订单
	}

	r.Run(conf.BackEndPort)
}
func initFrontEndRouters() {
	r := gin.Default()
	r.Static("templates/", "./templates")
	r.Static("photos/", "./src/photos")
	r.LoadHTMLGlob("templates/html/*")

	r.GET("/register", FRegister)
	r.GET("/login", FLogin)
	r.POST("/jwt", JWT(), FLogin)

	r.Run(conf.FrontEndPort)
}
