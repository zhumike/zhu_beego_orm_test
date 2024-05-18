package controllers

/**
 * @Author: zhenzhongzhu
 * @Description:
 * @File: OrmApi
 * @Date: 2024/5/18 12:25
 */

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
	"zhu_orm_test/models"
	"zhu_orm_test/tools"
)

type OrmApiController struct {
	beego.Controller
}

func (c *OrmApiController) InsertOne() {
	//生成orm对象
	o := orm.NewOrm()

	//被插入的数据结构对象
	userinfo := models.UserInfo{}

	randomStr := tools.GenerateRandomString(2)
	ranNum := tools.GenerateRandomNum(1, 223333391)
	numStr := strconv.Itoa(ranNum)

	//给结构体对象赋值
	userinfo.Name = "李四" + randomStr
	userinfo.Pwd = "odododdo123" + numStr

	//插入
	_, err := o.Insert(&userinfo)
	if err != nil {
		logs.Info("插入失败", err.Error())
		return
	} else {
		logs.Info("插入的用户名为：%s", userinfo.Name)
	}

	c.Data["Website"] = "测试orm beego api"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "test.html"
}
