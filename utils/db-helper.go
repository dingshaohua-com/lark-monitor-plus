package utils

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Pool *pgxpool.Pool
	once sync.Once // 确保初始化只执行一次
)

// InitDB 初始化连接池
func InitDB() {
	once.Do(func() {
		// 1. 获取 DSN，建议从环境变量获取，方便部署
		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			// 备用：本地开发默认配置
			dsn = "postgres://postgres:dshvv@103.110.80.247:5432/school"
		}

		// 2. 解析并配置连接池参数
		config, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			panic(fmt.Sprintf("无法解析数据库 DSN: %v", err))
		}

		// --- 生产环境参数优化 ---
		config.MaxConns = 20                      // 最大连接数
		config.MinConns = 5                       // 最小空闲连接数
		config.MaxConnLifetime = 30 * time.Minute // 连接最长存活时间
		config.MaxConnIdleTime = 10 * time.Minute // 空闲连接回收时间

		// 3. 创建连接池
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		Pool, err = pgxpool.NewWithConfig(ctx, config)
		if err != nil {
			panic(fmt.Sprintf("无法创建连接池: %v", err))
		}

		// 4. 尝试 Ping 一下，确保数据库是真的活着的
		if err := Pool.Ping(ctx); err != nil {
			panic(fmt.Sprintf("数据库连接测试失败: %v", err))
		}

		fmt.Println("🐘 PostgreSQL 连接池初始化成功 (pgxpool)")
	})
}

// CloseDB 优雅关闭
func CloseDB() {
	if Pool != nil {
		Pool.Close()
		fmt.Println("🐘 PostgreSQL 连接池已安全关闭")
	}
}
