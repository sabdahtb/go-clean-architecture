package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	Name      string
	Stock     uint16
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type CreateProductRequest struct {
	Name  string `json:"name" binding:"required"`
	Stock uint16 `json:"stock" binding:"required"`
}

type ProductRepository interface {
	Create(c context.Context, product Product) (message string, err error)
	Update(c context.Context, id string, product Product) (message string, err error)
	Delete(c context.Context, id string) (message string, err error)
	GetAll(c context.Context) (product []Product, err error)
	GetById(c context.Context, id string) (product Product, err error)
}

type ProductUsecase interface {
	Create(c context.Context, product Product) (message string, err error)
	Update(c context.Context, id string, product Product) (message string, err error)
	Delete(c context.Context, id string) (message string, err error)
	GetAll(c context.Context) (product []Product, err error)
	GetById(c context.Context, id string) (product Product, err error)
}
