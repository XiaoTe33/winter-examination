package conf

const (
	FrontEndPort = ":9091"
	BackEndPort  = ":9090"

	Database = "mysql"
	MysqlDNS = "root:root@tcp(127.0.0.1:3306)/winter_examination_database"

	JWTLastTime = 3600 //jwt持续时间（秒）

	OKMsg = "200" //成功返回消息

	SavePathOfGoodsPictures = ".\\src\\static\\goods\\pictures\\" //商品图片保存路径

	OrderIdLeftShiftNumber = 32 //生成全局OrderId左移位数
	GoodsIdLeftShiftNumber = 20 //生成全局GoodsId左移位数

	OrderIdBaseTimeStamp int64 = 1672531200 //2023-1-1 00:00:00
	GoodsIdBaseTimeStamp int64 = 1672531200 //2023-1-1 00:00:00
)

const (
	IP = "http://localhost"

	LocalPathOfUserPhoto    = "../../src/static/user/photos/"
	LocalPathOfGoodsPicture = "../../src/static/goods/pictures/"

	GinPathOfUserPhoto    = "user/photo/"
	GinPathOfGoodsPicture = "goods/picture/"

	WebLinkPathOfUserPhoto    = IP + FrontEndPort + "/" + GinPathOfUserPhoto
	WebLinkPathOfGoodsPicture = IP + FrontEndPort + "/" + GinPathOfGoodsPicture
)
