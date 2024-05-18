package routers

import (
	"github.com/astaxie/beego"
	"zhu_orm_test/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//	orm controller  /insert
	beego.Router("/insert", &controllers.OrmApiController{}, "get:InsertOne")
	beego.Router("/getOne", &controllers.OrmApiController{}, "get:GetUserOne")
	beego.Router("/updateOne", &controllers.OrmApiController{}, "get:UpdateUserOne")
	beego.Router("/deleteOne", &controllers.OrmApiController{}, "get:DeleteUserOne")

}
