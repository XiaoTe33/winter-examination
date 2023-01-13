package main

import (
	"winter-examination/src/app"
	"winter-examination/src/dao"
	"winter-examination/src/utils"
)

func main() {
	dao.InitDb()
	utils.InitOrderIdGenerator()
	app.InitRouters()
}
