package config

import (
	"app/domain"
	"github.com/jinzhu/gorm"
	"log"
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
		conn.DropTable(domain.ArticleComment{})
	}
	conn.AutoMigrate(domain.Article{})
	conn.AutoMigrate(domain.ArticleComment{}).AddForeignKey("article_id", "articles(id)", "RESTRICT", "RESTRICT")
	log.Println("Migration has been processed")
	return conn, nil
}
