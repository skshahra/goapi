package database

import (
	"fmt"
	"os"

	"bitbucket.org/skshahriarahmed/sh_ra/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlDBConnection() *gorm.DB {
	fmt.Println(os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_USER_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_IP") + ":" + os.Getenv("MYSQL_PORT") + ")/" + os.Getenv("MYSQL_DATABASE") + "?charset=utf8&parseTime=True&loc=Local")
	fmt.Println(os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_USER_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_IP") + ":" + os.Getenv("MYSQL_PORT") + ")/" + os.Getenv("MYSQL_DATABASE") + "?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_USER_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_IP") + ":" + os.Getenv("MYSQL_PORT") + ")/" + os.Getenv("MYSQL_DATABASE") + "?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                                                                                                                                                            // default size for string fields
		DisableDatetimePrecision:  true,                                                                                                                                                                                                           // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                                                                                                                                                           // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                                                                                                                                                           // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                                                                                                                                                          // auto configure based on currently MySQL version
	}), &gorm.Config{})

	logs.ERROR("Error in gorm.Open() ", err)
	if err == nil {
		fmt.Println("✨Mysql Connected Successfully ✨")

	}
	return db
}
