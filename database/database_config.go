package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func InitDatabase() *DBConfig {
	db_config := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "",
		DBName:   "loan_mgt",
	}
	return &db_config
}

func DB_url(db_config *DBConfig) string {
	return fmt.Sprintf(
		"#{db_config.User}:#{db_config.Password}@tcp(#{db_config.Host}:#{db_config.Port})/#{db_config.DBName}?charset=utf8&parseTime=true&loc=Local",
	)
}
