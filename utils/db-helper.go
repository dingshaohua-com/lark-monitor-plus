package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 定义一个全局变量，以便其他地方调用
var DB *gorm.DB

func InitDB() {
	var err error
	databaseUrl := "host=103.110.80.247 user=postgres password=dshvv dbname=school port=5432 sslmode=disable"

	// 注意：这里使用 = 而不是 :=，因为 DB 已经在外部声明了
	DB, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		panic("无法连接数据库: " + err.Error())
	}
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("获取 sql.DB 失败:", err)
		return
	}
	err = sqlDB.Close()
	if err != nil {
		fmt.Println("关闭数据库连接失败:", err)
	} else {
		fmt.Println("--- 数据库连接已安全关闭 ---")
	}
}
