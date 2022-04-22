package databases

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learngo/goland/yangpanping/bloggin/utils"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var Db *gorm.DB
var Err error

func DbInit() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	Db, Err = gorm.Open(mysql.Open(dns), &gorm.Config{})

	if Err != nil {
		panic("failed to connect database")
	}
	sqlDB, _ := Db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

}
