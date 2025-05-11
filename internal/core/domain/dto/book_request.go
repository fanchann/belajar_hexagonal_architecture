package dto

type CreateBookRequest struct {
	Title       string `json:"title" validate:"required"`
	Author      string `json:"author" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateBookRequest struct {
	ID          string `json:"-" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Author      string `json:"author" validate:"required"`
	Description string `json:"description" validate:"required"`
}
