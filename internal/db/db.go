package db

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var once sync.Once

func init() {
	once.Do(func() { initDb() })
}

func initDb() {

	config := &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix: "sf_", // 表名前缀
		// SingularTable: true,  // 使用单数表名
	}}

	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/starfire_cloud?charset=utf8mb4&parseTime=True&loc=Local"), config)

	if err != nil {
		fmt.Println("连接数据库失败")
		panic(err)
	}

	DB = db
}
