package mysql

import (
	"fmt"
	"web_app/settings"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Init(mysqlConfig *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.DBName,
	)
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	DB.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))
	DB.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))
	return
}

func Close() {
	_ = DB.Close()
}
