package routers

import (
	"github.com/astaxie/beego"
	"zhu_orm_test/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//	orm controller
	beego.Router("/insert", &controllers.OrmApiController{}, "get:InsertOne")
}
