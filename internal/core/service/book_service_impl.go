package service

import (
	"context"

	"github.com/go-playground/validator"

	"github.com/fanchann/grpc-crud/internal/core/domain/dto"
	"github.com/fanchann/grpc-crud/internal/port/inbound"
	"github.com/fanchann/grpc-crud/internal/port/outbound"
)

type newBookServiceImpl struct {
	repo      outbound.IBookRepoPorts
	validator *validator.Validate
}

func NewBookService(repo outbound.IBookRepoPorts, validator *validator.Validate) inbound.IBookServicePorts {
	return &newBookServiceImpl{
		repo:      repo,
		validator: validator,
	}
}

func (n *newBookServiceImpl) CreateNewBook(ctx context.Context, req *dto.CreateBookRequest) error {
	if err := n.validator.Struct(req); err != nil {
		return err
	}

	book := dto.CreateBookRequestToEntity(req)
	if err := n.repo.CreateBook(ctx, book); err != nil {
		return err
	}
	return nil
}

func (n *newBookServiceImpl) GetAllBooks(ctx context.Context) []dto.BookResponse {
	books := n.repo.GetAllBooks(ctx)
	var bookResponses []dto.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, *dto.BookEntityToResponse(&book))
	}
	return bookResponses
}

func (n *newBookServiceImpl) DeleteBook(ctx context.Context, id string) error {
	if err := n.repo.DeleteBook(ctx, id); err != nil {
		return err
	}
	return nil
}

func (n *newBookServiceImpl) UpdateBook(ctx context.Context, id string, req *dto.UpdateBookRequest) error {
	if err := n.validator.Struct(req); err != nil {
		return err
	}

	book := dto.UpdateBookRequestToEntity(req)
	if err := n.repo.UpdateBook(ctx, book); err != nil {
		return err
	}
	return nil
}

func (n *newBookServiceImpl) GetBookByID(ctx context.Context, id string) (*dto.BookResponse, error) {
	book, err := n.repo.GetBookByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.BookEntityToResponse(book), nil
}
