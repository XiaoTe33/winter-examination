package conf

const (
	FrontEndPort = ":9091"
	BackEndPort  = ":9090"
)

const (
	Database = "mysql"
	MysqlDNS = "root:root@tcp(127.0.0.1:3306)/winter_examination_database"
)

// JWTLastTime jwt持续时间（秒）
const JWTLastTime = 3600
