package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func init() {
	Connect()
	//CreateDb()
	SyncTable()
}

var dsn string
var o orm.Ormer
var db_type string = beego.AppConfig.String("db_type")
var db_host string = beego.AppConfig.String("db_host")
var db_port string = beego.AppConfig.String("db_port")
var db_user string = beego.AppConfig.String("db_user")
var db_pass string = beego.AppConfig.String("db_pass")
var db_name string = beego.AppConfig.String("db_name")
var db_path string = beego.AppConfig.String("db_path")
var db_sslmode string = beego.AppConfig.String("db_sslmode")

//数据库连接
func Connect() {
	switch db_type {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DRMySQL)
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", db_user, db_pass, db_host, db_port, db_name)
	case "postgres":
		orm.RegisterDriver("postgres", orm.DRPostgres)
		dsn = fmt.Sprintf("dbname=%s host=%s user=%s password=%s port=%s sslmode=%s", db_name, db_host, db_user, db_pass, db_port, db_sslmode)
	case "sqlite3":
		orm.RegisterDriver("sqlite3", orm.DRSqlite)
		if db_path == "" {
			db_path = "./"
		}
		dsn = fmt.Sprintf("%s%s.db", db_path, db_name)
	default:
		beego.Critical("数据库类型不支持:", db_type)
	}
	orm.RegisterDataBase("default", db_type, dsn)
}

func CreateDb() {
	var sqlstring string
	switch db_type {
	case "mysql":
		sqlstring = fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARSET utf8 COLLATE utf8_general_ci", db_name)
	case "postgres":
		sqlstring = fmt.Sprintf("CREATE DATABASE %s", db_name)
	case "sqlite3":
		os.Remove(dsn)
		sqlstring = fmt.Sprintf("create table int (n varchar(32));drop table int;")
	default:
		beego.Critical("数据库类型不支持:", db_type)
	}
	db, err := sql.Open(db_type, dsn)
	if err != nil {
		panic(err.Error())
	}
	log.Println(sqlstring)
	log.Println(dsn)
	res, err := db.Exec(sqlstring)
	if err != nil {
		log.Println(err, res)
	} else {
		log.Println("数据库", db_name, "创建成功")
	}
	defer db.Close()
}

func SyncTable() {
	// 数据库别名
	name := "default"
	// drop table 后再建表
	force := false
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	orm.RegisterModel(new(Menus), new(Sites), new(Categories), new(Tags), new(Files), new(Articles), new(Links), new(Users), new(Zans))
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
