package db

import (
	"database/sql"
	"gim/config"

	_ "github.com/go-sql-driver/mysql"
)

var DBCli *sql.DB

//初始化mysql
func init() {
	var err error
	DBCli, err = sql.Open("mysql", config.LogicConf.MySQL)
	//此处需要提供设置MySQL的参数的接口
	if err != nil {
		panic(err)
	}
}
