package grpcHandler

import (
	"context"

	"github.com/fanchann/grpc-crud/internal/core/domain/dto"
	"github.com/fanchann/grpc-crud/internal/port/inbound"
	BooksProto "github.com/fanchann/grpc-crud/proto/protogen"
)

type newBookGrpcImpl struct {
	bookServicePort inbound.IBookServicePorts
	BooksProto.BookServiceServer
}

func NewBookGrpcImpl(bookServicePort inbound.IBookServicePorts) BooksProto.BookServiceServer {
	return &newBookGrpcImpl{
		bookServicePort: bookServicePort,
	}
}

func (n *newBookGrpcImpl) CreateBook(p0 context.Context, p1 *BooksProto.BookCreateRequest) (*BooksProto.Empty, error) {
	bookRequest := dto.CreateBookRequest{
		Title:       p1.Title,
		Author:      p1.Author,
		Description: p1.Description,
	}
	if err := n.bookServicePort.CreateNewBook(p0, &bookRequest); err != nil {
		return nil, err
	}
	return &BooksProto.Empty{}, nil
}

func (n *newBookGrpcImpl) UpdateBook(p0 context.Context, p1 *BooksProto.BookUpdateRequest) (*BooksProto.Empty, error) {
	bookRequest := dto.UpdateBookRequest{
		ID:          p1.Id,
		Title:       p1.Title,
		Description: p1.Description,
		Author:      p1.Author,
	}
	if err := n.bookServicePort.UpdateBook(p0, p1.Id, &bookRequest); err != nil {
		return nil, err
	}
	return &BooksProto.Empty{}, nil
}

func (n *newBookGrpcImpl) GetBook(p0 context.Context, p1 *BooksProto.BookGetRequest) (*BooksProto.BookResponse, error) {
	bookResponse, err := n.bookServicePort.GetBookByID(p0, p1.Id)
	if err != nil {
		return nil, err
	}
	return &BooksProto.BookResponse{
		Id:          bookResponse.ID,
		Title:       bookResponse.Title,
		Author:      bookResponse.Author,
		Description: bookResponse.Description,
	}, nil
}

func (n *newBookGrpcImpl) DeleteBook(p0 context.Context, p1 *BooksProto.BookDeleteRequest) (*BooksProto.Empty, error) {
	if err := n.bookServicePort.DeleteBook(p0, p1.Id); err != nil {
		return nil, err
	}
	return &BooksProto.Empty{}, nil
}

func (n *newBookGrpcImpl) GetAllBooks(p0 context.Context, p1 *BooksProto.Empty) (*BooksProto.BookListResponse, error) {
	books := n.bookServicePort.GetAllBooks(p0)
	var bookResponses []*BooksProto.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, &BooksProto.BookResponse{
			Id:          book.ID,
			Title:       book.Title,
			Author:      book.Author,
			Description: book.Description,
		})
	}
	return &BooksProto.BookListResponse{Books: bookResponses}, nil
}
