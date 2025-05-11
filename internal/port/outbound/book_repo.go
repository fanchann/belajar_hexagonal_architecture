package outbound

import (
	"context"

	"github.com/fanchann/grpc-crud/internal/core/domain/entity"
)

type IBookRepoPorts interface {
	GetAllBooks(ctx context.Context) []entity.Books
	GetBookByID(ctx context.Context, id string) (*entity.Books, error)
	CreateBook(ctx context.Context, book *entity.Books) error
	UpdateBook(ctx context.Context, book *entity.Books) error
	DeleteBook(ctx context.Context, id string) error
}
