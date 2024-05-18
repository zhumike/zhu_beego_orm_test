package main

import (
	"github.com/astaxie/beego"
	_ "zhu_orm_test/models"
	_ "zhu_orm_test/routers"
)

func main() {
	beego.Run()
}
