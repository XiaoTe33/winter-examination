package main

import (
	"winter-examination/src/app"
	"winter-examination/src/dao"
)

func main() {
	app.InitRouters()
	dao.InitDb()
}
