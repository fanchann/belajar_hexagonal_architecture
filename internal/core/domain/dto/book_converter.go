package dto

import (
	"github.com/fanchann/grpc-crud/internal/core/domain/entity"
	"github.com/fanchann/grpc-crud/internal/pkg/uid"
)

func CreateBookRequestToEntity(req *CreateBookRequest) *entity.Books {
	return &entity.Books{
		ID:          uid.GenerateUUID(),
		Title:       req.Title,
		Author:      req.Author,
		Description: req.Description,
	}
}

func UpdateBookRequestToEntity(req *UpdateBookRequest) *entity.Books {
	return &entity.Books{
		ID:          req.ID,
		Title:       req.Title,
		Author:      req.Author,
		Description: req.Description,
	}
}

func BookEntityToResponse(book *entity.Books) *BookResponse {
	return &BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
	}
}
