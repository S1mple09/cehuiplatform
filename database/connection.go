package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"

	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 数据库连接实例
var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	// 使用 sqlx 获取数据库连接
	dsn := "root:Qq958299102#@tcp(127.0.0.1:3308)/cehui?charset=utf8mb4&parseTime=True&loc=Local"
	sqlxDB, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 测试数据库连接
	err = sqlxDB.Ping()
	if err != nil {
		log.Fatalf("数据库连接测试失败: %v", err)
	}

	// 将 sqlx.DB 转换为 gorm.DB
	DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlxDB.DB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("转换数据库连接失败: %v", err)
	}

	fmt.Println("数据库连接成功")
}
