package database

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
	"github.com/jmoiron/sqlx"
)

// DB 数据库连接实例
var DB *sqlx.DB

// InitDB 初始化数据库连接
func InitDB() {
	var err error
	DB, err = sqlx.Open("mysql", "root:Qq958299102#@tcp(127.0.0.1:3308)/cehui?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 测试数据库连接
	err = DB.Ping()
	if err != nil {
		log.Fatalf("数据库连接测试失败: %v", err)
	}

	fmt.Println("数据库连接成功")
}
