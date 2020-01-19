package config

import (
	"app/domain"
	"log"
	"github.com/jinzhu/gorm"
)

// DBMigrate will create & migrate the tables, then make the some relationships if necessary
func DBMigrate() (*gorm.DB, error) {
	conn, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if conn.HasTable(domain.Article{}) {
		conn.DropTable(domain.Article{})
	}
	conn.AutoMigrate(domain.Article{})
	log.Println("Migration has been processed")
	return conn, nil
}