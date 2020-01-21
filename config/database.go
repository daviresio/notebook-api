package config

import (
	"fmt"
	"github.com/daviresio/financeiro_api/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	connectionString = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "127.0.0.1:3306", "financeiro_db")
	DB *gorm.DB
)

func init() {
	var err error
	DB, err = gorm.Open("mysql", connectionString)

	if err != nil {
		println(err)
		panic(err)
	}

	DB.SingularTable(true)

	DB.LogMode(true)

	DB.AutoMigrate(&model.Book{}, &model.Notebook{}, &model.User{}, &model.Address{}, &model.Tag{})
}

