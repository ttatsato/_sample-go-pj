package usecase

import (
	"app/config"
	"app/domain"
	"app/infrastructure/persistence"
)

func CreateNewArticle(article *domain.Article) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()
	repo := persistence.ArticleRepositoryWithRDB(conn)
	return repo.Save(article)
}

func FetchAllArticles() ([]domain.Article, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	repo := persistence.ArticleRepositoryWithRDB(conn)
	return repo.GetAll()
}

func UpdateArticle(article *domain.Article) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()
	repo := persistence.ArticleRepositoryWithRDB(conn)
	return repo.Update(article)
}
