package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库
func InitDB() (*gorm.DB, error) {
	host := "localhost"
	port := 3306
	user := "root"
	password := "longyue0511"
	dbname := "short_link_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
