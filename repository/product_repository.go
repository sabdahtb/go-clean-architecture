package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"sabda.go/learn/domain"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{
		db,
	}
}

func (pr *productRepository) Create(c context.Context, product domain.Product) (message string, err error) {
	db := pr.db
	uid := uuid.New().String()

	save_product := domain.Product{
		ID:    uid,
		Name:  product.Name,
		Stock: product.Stock,
	}

	result := db.Create(&save_product)
	return "success cretae product", result.Error
}

func (pr *productRepository) GetAll(c context.Context) (product []domain.Product, err error) {
	db := pr.db
	var products []domain.Product

	db.Find(&products)
	return products, nil
}

func (pr *productRepository) GetById(c context.Context, id string) (product domain.Product, err error) {
	db := pr.db
	var resultProduct = domain.Product{ID: id}

	db.First(&resultProduct)
	return resultProduct, nil
}

func (pr *productRepository) Delete(c context.Context, id string) (message string, err error) {
	db := pr.db
	var resultProduct = domain.Product{ID: id}

	db.First(&resultProduct)
	db.Delete(&resultProduct)
	return "succes delete product", nil
}

func (pr *productRepository) Update(c context.Context, id string, product domain.Product) (message string, err error) {
	db := pr.db
	var resultProduct = domain.Product{ID: id, Name: product.Name, Stock: product.Stock}

	db.Save(&resultProduct)
	return "succes update product", nil
}
