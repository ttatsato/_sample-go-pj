package usecase

import (
	"app/config"
	"app/domain"
	"app/infrastructure/persistence"
)

func CreateNewArticle (article *domain.Article) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil
	}
	defer conn.Close()
	repo := persistence.ArticleRepositoryWithRDB(conn)
	return repo.Save(article)
}