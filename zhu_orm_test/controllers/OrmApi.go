package controllers

/**
 * @Author: zhenzhongzhu
 * @Description:
 * @File: OrmApi
 * @Date: 2024/5/18 12:25
 */

import (
	"fmt"
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

// InsertOne
// @Author ZhenZhen Zhu
// @Description: 插入数据
// @receiver
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
		rp := tools.Customize(200, "成功插入数据")
		c.Data["json"] = rp
		c.ServeJSON()
	}

	c.Data["Website"] = "测试orm beego api"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "test.html"
}

// GetUserOne
// @Author ZhenZhen Zhu
// @Description: 根据id查找用户
// @receiver c
func (c *OrmApiController) GetUserOne() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info(err)
		// 处理错误，例如返回404 Not Found
		c.Data["json"] = map[string]interface{}{"error": "id not found"}
		c.Abort("404")
		return
	}
	fmt.Printf("值：%v 类型：%T", id, id)
	//c.Ctx.WriteString("orm test")

	o := orm.NewOrm()
	userinfo := models.UserInfo{}

	userinfo.Id = id
	err2 := o.Read(&userinfo)
	if err2 != nil {
		logs.Error("查询失败")
		return
	}

	logs.Info("查询成功", userinfo)
	//c.Ctx.WriteString("用户名：" + userinfo.Name + "\n" + userinfo.Pwd)
	//rp := tools.Success
	// 返回用户信息

	rp := tools.Successed(userinfo)
	c.Data["json"] = rp
	c.ServeJSON()

}

// UpdateUserOne
// @Author ZhenZhen Zhu
// @Description: 根据id更新用户信息
// @receiver
func (c *OrmApiController) UpdateUserOne() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info(err)
		// 处理错误，例如返回404 Not Found
		c.Data["json"] = map[string]interface{}{"error": "id not found"}
		c.Abort("404")
		return
	}
	fmt.Printf("值：%v 类型：%T", id, id)
	//c.Ctx.WriteString("orm test")
	newName := c.GetString("name")
	newPwd := c.GetString("pwd")

	o := orm.NewOrm()
	userinfo := models.UserInfo{}

	userinfo.Id = id
	err2 := o.Read(&userinfo)
	if err2 != nil {
		logs.Error("查询失败")
		return
	}
	logs.Info("查询成功", userinfo)

	err3 := o.Read(&userinfo)
	if err3 == nil {
		userinfo.Name = newName
		userinfo.Pwd = newPwd
		_, err := o.Update(&userinfo)
		if err != nil {
			logs.Error("更新失败")
		}
		// 返回用户信息
		c.Data["json"] = "更新成功"
		c.ServeJSON()
	} else {
		c.Data["json"] = "更新失败"
		c.ServeJSON()
	}

}

// DeleteUserOne
// @Author ZhenZhen Zhu
// @Description: 根据id删除用户数据
// @receiver
func (c *OrmApiController) DeleteUserOne() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info(err)
		// 处理错误，例如返回404 Not Found
		c.Data["json"] = map[string]interface{}{"error": "id not found"}
		c.Abort("404")
		return
	}
	fmt.Printf("值：%v 类型：%T", id, id)

	o := orm.NewOrm()
	userinfo := models.UserInfo{}

	userinfo.Id = id
	err2 := o.Read(&userinfo)
	if err2 != nil {
		logs.Error("查询失败")
		rp := tools.Customize(500, "删除失败，无对应id")
		c.Data["json"] = rp
		c.ServeJSON()
		return
	}
	logs.Info("查询成功", userinfo)

	_, err3 := o.Delete(&userinfo)
	if err3 != nil {
		logs.Error("删除失败")
	} else {
		rp := tools.Customize(200, "删除成功")
		c.Data["json"] = rp
		c.ServeJSON()
	}
}
