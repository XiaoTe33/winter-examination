package conf

const (
	FrontEndPort = ":9091"
	BackEndPort  = ":9090"

	Database = "mysql"
	MysqlDNS = "root:root@tcp(127.0.0.1:3306)/winter_examination_database"

	JWTLastTime = 3600 //jwt持续时间（秒）

	OKMsg = "200" //成功返回消息

	LocalSavePathOfGoodsPictures      = ".\\src\\static\\goods\\pictures\\"      //商品图片保存路径
	LocalSavePathOfUserPhotos         = ".\\src\\static\\user\\photos\\"         //用户头像保存地址
	LocalSavePathOfEvaluationPictures = ".\\src\\static\\evaluation\\pictures\\" //评价图片保存地址

	OrderIdLeftShiftNumber = 32 //生成全局OrderId左移位数
	GoodsIdLeftShiftNumber = 20 //生成全局GoodsId左移位数

	OrderIdBaseTimeStamp int64 = 1672531200 //2023-1-1 00:00:00
	GoodsIdBaseTimeStamp int64 = 1672531200 //2023-1-1 00:00:00
)

const (
	IP = "http://localhost"

	GinPathOfUserPhoto          = "user/photo/" //路由地址
	GinPathOfGoodsPicture       = "goods/picture/"
	GinPathOfEvaluationPictures = "evaluation/picture/"

	WebLinkPathOfUserPhoto          = IP + FrontEndPort + "/" + GinPathOfUserPhoto //链接地址
	WebLinkPathOfGoodsPicture       = IP + FrontEndPort + "/" + GinPathOfGoodsPicture
	WebLinkPathOfEvaluationPictures = IP + FrontEndPort + "/" + GinPathOfEvaluationPictures
)
