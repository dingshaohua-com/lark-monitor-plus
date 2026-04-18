package main

import (
	"fmt"
	"lark-monitor-plus/utils"
)

type Teacher struct {
	// 必须大写开头！后面反引号里的才是数据库真实的列名
	ID   string `gorm:"column:id"`
	Name string
	Age  int
}

// TableName 核心：手动指定表名，覆盖 GORM 的复数约定
func (Teacher) TableName() string {
	return "teacher"
}

func main() {
	// 1. 初始化连接
	utils.InitDB()

	// 2. 延迟关闭（确保在 main 函数退出前，连接池被清理）
	defer utils.CloseDB()

	// 3. 查询数据
	var users []Teacher
	utils.DB.Find(&users)
	fmt.Printf("%+v\n", users)
}

//conn, _ := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
