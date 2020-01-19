package repository

import "app/domain"

type ArticleCommentRepository interface {
	Get(id int) (*domain.ArticleComment, error)
	GetAll() ([]domain.ArticleComment, error)
	Save(*domain.ArticleComment) error
	Remove(id int) error
	Update(*domain.ArticleComment) error
}
