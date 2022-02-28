package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var DB *gorm.DB
func DataBase(connString string)  {
	fmt.Println(connString)
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	DB = db
	if err != nil {
		panic("数据库连接错误")
	}
	fmt.Println("数据库连接成功")
	sqlDb, _ := db.DB()
	sqlDb.SetConnMaxIdleTime(25) // 设置连接池
	sqlDb.SetMaxOpenConns(100) // 设置最大连接数
	sqlDb.SetConnMaxLifetime(30 * time.Second)
	db.AutoMigrate(&User{}, &Todo{})
}