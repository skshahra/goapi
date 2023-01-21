package database

import (
	"bitbucket.org/skshahriarahmed/sh_ra/handler"
	"gorm.io/gorm"
)

func DatabaseInitialization(Mysqldb *gorm.DB) handler.DatabaseCollections {
	return  handler.DatabaseCollections{
		MySqlDB: Mysqldb,
	}
}