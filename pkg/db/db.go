package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// MySqlOptions 定义 Mysql 数据库的选项
type MySqlOptions struct {
	Host                  string
	Username              string
	Password              string
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	LogLevel              int
}

func (o *MySqlOptions) DSN() string {
	return fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parsetTime=%t&loc=%s`,
		o.Username,
		o.Password,
		o.Host,
		o.Database,
		true,
		"Local",
	)
}

// NewMySQL 使用 options 创建一个新的 gorm 数据库实例
func NewMySQL(opts *MySqlOptions) (*gorm.DB, error) {
	logLevel := logger.Silent
	if opts.LogLevel != 0 {
		logLevel = logger.LogLevel(opts.LogLevel)
	}

	db, err := gorm.Open(mysql.Open(opts.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)

	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

	return db, nil
}
