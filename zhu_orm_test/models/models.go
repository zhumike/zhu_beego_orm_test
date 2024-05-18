package models

/**
 * @Author: zhenzhongzhu
 * @Description:
 * @File: models
 * @Date: 2024/5/18 10:13
 */

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	Id   int
	Name string
	Pwd  string
}

func init() {
	// 注册数据库
	orm.RegisterDataBase("default", "mysql", "root:123456zhu@@tcp(127.0.0.1:3306)/zhutest?charset=utf8")

	//注册表
	orm.RegisterModel(new(UserInfo))

	//生成表
	orm.RunSyncdb("default", false, true)

}
