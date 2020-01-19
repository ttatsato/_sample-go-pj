package persistence

import (
	"app/domain"
	"app/domain/repository"
	"github.com/jinzhu/gorm"
)

// ArticleRepositoryImpl Implements repository.ArticleRepository
type ArticleRepositoryImpl struct {
	Conn *gorm.DB
}

// ArticleRepositoryWithRDB returns initialized ArticleRepositoryImpl
func ArticleRepositoryWithRDB(conn *gorm.DB) repository.ArticleRepository {
	return &ArticleRepositoryImpl{Conn: conn}
}

// Get Article by id return domain.Article
func (r *ArticleRepositoryImpl) Get(id int) (*domain.Article, error) {
	Article := &domain.Article{}
	if err := r.Conn.Preload("Comment").First(&Article, id).Error; err != nil {
		return nil, err
	}
	return Article, nil
}

// GetAll Article return all domain.Article
func (r *ArticleRepositoryImpl) GetAll() ([]domain.Article, error) {
	Article := []domain.Article{}
	//if err := r.Conn.Preload("Comment").Find(&Article).Error; err != nil {
	//	return nil, err
	//}
	if err := r.Conn.Find(&Article).Error; err != nil {
		return nil, err
	}

	return Article, nil
}

// Save to add Article
func (r *ArticleRepositoryImpl) Save(Article *domain.Article) error {
	if err := r.Conn.Save(&Article).Error; err != nil {
		return err
	}

	return nil
}

// Remove to delete Article by id
func (r *ArticleRepositoryImpl) Remove(id int) error {
	tx := r.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	Article := domain.Article{}
	if err := tx.First(&Article, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	Article.Status = "deleted"
	if err := tx.Save(&Article).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&Article).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Update is update Article
func (r *ArticleRepositoryImpl) Update(Article *domain.Article) error {
	if err := r.Conn.Model(&Article).UpdateColumns(Article).Error; err != nil {
		return err
	}

	return nil
}
