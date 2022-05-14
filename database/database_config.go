package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = ""
const DB_NAME = "loan_mgt"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var Db *gorm.DB

func InitDB() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_HOST + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
        fmt.Println("Error Connecting Database:",err)
		return nil
	}
	return db
}



// type DBConfig struct {
// 	Host     string
// 	Port     int
// 	User     string
// 	Password string
// 	DBName   string
// }

// func InitDatabase() *DBConfig {
// 	db_config := DBConfig{
// 		Host:     "localhost",
// 		Port:     3306,
// 		User:     "root",
// 		Password: "",
// 		DBName:   "loan_mgt",
// 	}
// 	return &db_config
// }

// func DB_url(db_config *DBConfig) string {
// 	return fmt.Sprintf(
// 		"#{db_config.User}:#{db_config.Password}@tcp(#{db_config.Host}:#{db_config.Port})/#{db_config.DBName}?charset=utf8&parseTime=true&loc=Local",
// 	)
// }
