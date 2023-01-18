package app

import (
	"time"

	"winter-examination/src/conf"
	"winter-examination/src/utils"

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
	r.MaxMultipartMemory = 3 << 20
	r.Use(Cors())
	r.LoadHTMLFiles("./src/static/managePage/manage.html")

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

	o := r.Group("/order") //订单模块
	{
		o.POST("/add", JWT(), AddOrder)                     //新增订单
		o.POST("/updateStatus", JWT(), UpdateOrderStatus)   //修改订单状态
		o.POST("/updateAddress", JWT(), UpdateOrderAddress) //修改订单地址
		o.POST("/query", QueryOrders)                       //查询订单
		o.POST("/queryAll", QueryAllOrders)                 //查看所有订单
	}

	st := r.Group("/star") //收藏模块
	{
		st.POST("/add", JWT(), AddStar)          //收藏商品
		st.POST("/query", JWT(), QueryUserStars) //我的收藏
		st.POST("/queryAll", QueryAllStars)      //查看所有
		st.POST("/delete", JWT(), DeleteStar)    //取消收藏
	}

	e := r.Group("/evaluation") //评价模块
	{
		e.POST("/add", JWT(), AddEvaluation)        //新增评价
		e.POST("/delete", JWT(), DeleteEvaluations) //删除评价
		e.POST("/query", QueryGoodsEvaluations)     //查询某个商品的评价
		e.POST("/queryAll", QueryAllEvaluations)    //查看所有
	}

	r.GET("/RESTART/:yes", utils.Restart) //数据库重开

	r.GET("/manage", func(c *gin.Context) {
		c.HTML(200, "manage.html", nil)
	})

	r.Run(conf.BackEndPort)
}
func initFrontEndRouters() {
	r := gin.Default()
	r.Use(Cors())
	r.Static("js/", "./templates/js")
	r.Static("css/", "./templates/css")
	r.Static("images", "./templates/images")
	r.Static(conf.GinPathOfUserPhoto, "./src/static/user/photos")
	r.Static(conf.GinPathOfGoodsPicture, "./src/static/goods/pictures")
	r.Static(conf.GinPathOfEvaluationPictures, "./src/static/evaluation/pictures")
	r.LoadHTMLGlob("templates/html/*")

	r.GET("/register", FRegister)
	r.GET("/login", FLogin)
	r.POST("/jwt", JWT(), FLogin)

	r.Run(conf.FrontEndPort)
}
