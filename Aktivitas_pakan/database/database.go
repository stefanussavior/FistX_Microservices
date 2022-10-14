package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ConnectDatabase() *gorm.DB {
	User := "root"
	Pass := ""
	Host := "localhost"
	Port := "3306"
	Dbname := "db_fistx_aktivitas_pakan"

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", User, Pass, Host, Port, Dbname)

	db, err := gorm.Open("mysql", URL)

	if err != nil {
		panic(err.Error())
	}
	return db
}
