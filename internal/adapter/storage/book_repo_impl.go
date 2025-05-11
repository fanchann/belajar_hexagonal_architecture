package storage

import (
	"context"

	"gorm.io/gorm"

	"github.com/fanchann/grpc-crud/internal/core/domain/entity"
	"github.com/fanchann/grpc-crud/internal/port/outbound"
)

type newBookRepoImpl struct {
	db *gorm.DB
}

func NewBookRepoImpl(db *gorm.DB) outbound.IBookRepoPorts {
	return &newBookRepoImpl{
		db: db,
	}
}

func (n *newBookRepoImpl) GetAllBooks(ctx context.Context) []entity.Books {
	var books []entity.Books
	if err := n.db.WithContext(ctx).Find(&books).Error; err != nil {
		return nil
	}
	return books
}

func (n *newBookRepoImpl) GetBookByID(ctx context.Context, id string) (*entity.Books, error) {
	var book entity.Books
	if err := n.db.WithContext(ctx).First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (n *newBookRepoImpl) CreateBook(ctx context.Context, book *entity.Books) error {
	if err := n.db.WithContext(ctx).Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (n *newBookRepoImpl) UpdateBook(ctx context.Context, book *entity.Books) error {
	if err := n.db.WithContext(ctx).Save(book).Error; err != nil {
		return err
	}
	return nil
}

func (n *newBookRepoImpl) DeleteBook(ctx context.Context, id string) error {
	if err := n.db.WithContext(ctx).Delete(&entity.Books{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
