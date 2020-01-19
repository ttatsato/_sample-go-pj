package repository

import "ibp/domain"

// ArticleRepository represent repository of  the Article
// Expect implementation by the infrastructure layer
type ArticleRepository interface {
	Get(id int) (*domain.Article, error)
	GetAll() ([]domain.Article, error)
	Save(*domain.Article) error
	Remove(id int) error
	Update(*domain.Article) error
}