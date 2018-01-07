package tsDb

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
)

func ConnectDb() error {

	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=" + url.QueryEscape("Local")
	fmt.Println("数据库地址:", dsn)
	err := orm.RegisterDataBase("default", "mysql", dsn)
	if err != nil {
		beego.Debug("数据库服务器链接失败；", err)
	} else {
		beego.Debug("数据库服务器链接成功")
	}

	return err
}
