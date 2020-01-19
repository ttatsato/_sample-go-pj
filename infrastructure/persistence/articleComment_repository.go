package persistence

import (
	"app/domain"
	"app/domain/repository"
	"github.com/jinzhu/gorm"
)

// ArticleCommentRepositoryImpl Implements repository.ArticleCommentRepository
type ArticleCommentRepositoryImpl struct {
	Conn *gorm.DB
}

// ArticleCommentRepositoryWithRDB returns initialized ArticleCommentRepositoryImpl
func ArticleCommentRepositoryWithRDB(conn *gorm.DB) repository.ArticleCommentRepository {
	return &ArticleCommentRepositoryImpl{Conn: conn}
}

// Get ArticleComment by id return domain.ArticleComment
func (r *ArticleCommentRepositoryImpl) Get(id int) (*domain.ArticleComment, error) {
	ArticleComment := &domain.ArticleComment{}
	if err := r.Conn.First(&ArticleComment, id).Error; err != nil {
		return nil, err
	}
	return ArticleComment, nil
}

// GetAll ArticleComment return all domain.ArticleComment
func (r *ArticleCommentRepositoryImpl) GetAll() ([]domain.ArticleComment, error) {
	ArticleComment := []domain.ArticleComment{}

	if err := r.Conn.Find(&ArticleComment).Error; err != nil {
		return nil, err
	}

	return ArticleComment, nil
}

// Save to add ArticleComment
func (r *ArticleCommentRepositoryImpl) Save(ArticleComment *domain.ArticleComment) error {
	if err := r.Conn.Save(&ArticleComment).Error; err != nil {
		return err
	}

	return nil
}

// Remove to delete ArticleComment by id
func (r *ArticleCommentRepositoryImpl) Remove(id int) error {
	tx := r.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	ArticleComment := domain.ArticleComment{}
	if err := tx.First(&ArticleComment, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	ArticleComment.Status = "deleted"
	if err := tx.Save(&ArticleComment).Error; err != nil {
		tx.Rollback()
		return err
	}

	// gorm Delete()
	// モデルにDeletedAt フィールドがある場合、自動的にソフトデリート機能。
	// Unscopedを使うことでレコードを完全に消去することができます。
	// tx.Unscoped().Delete(&ArticleComment)
	if err := tx.Delete(&ArticleComment).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Update is update ArticleComment
func (r *ArticleCommentRepositoryImpl) Update(ArticleComment *domain.ArticleComment) error {
	if err := r.Conn.Model(&ArticleComment).UpdateColumns(ArticleComment).Error; err != nil {
		return err
	}

	return nil
}
