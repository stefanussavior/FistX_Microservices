package test_database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/* testing untuk database */

func TestDb() *gorm.DB {
	User := "root"
	Password := ""
	Host := "localhost"
	Port := "3306"
	DBNAME := "db_fistx_users"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", User, Password, Host, Port, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err)
	}
	return db
}
