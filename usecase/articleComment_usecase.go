package usecase

import (
	"app/config"
	"app/domain"
	"app/infrastructure/persistence"
)

func CreateNewArticleComment(articleComment *domain.ArticleComment) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()
	repo := persistence.ArticleCommentRepositoryWithRDB(conn)
	return repo.Save(articleComment)
}

func UpdateArticleComment(articleComment *domain.ArticleComment) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()
	repo := persistence.ArticleCommentRepositoryWithRDB(conn)
	return repo.Update(articleComment)
}

// RemoveArticle soft delete
func RemoveArticleComment(articleCommentId int) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()
	repo := persistence.ArticleCommentRepositoryWithRDB(conn)
	return repo.Remove(articleCommentId)
}
