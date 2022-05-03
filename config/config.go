package config

import (
	"fmt"
	"goctr/structs"

	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:toor@tcp(db:3306)/godb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(fmt.Sprintf("Can't connect to database: %s", err))
	}

	db.AutoMigrate(structs.Person{})

	return db
}
