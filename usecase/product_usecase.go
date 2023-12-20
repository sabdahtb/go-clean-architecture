package usecase

import (
	"context"
	"time"

	"sabda.go/learn/domain"
)

type productUsecase struct {
	productRepository domain.ProductRepository
	contextTimeout    time.Duration
}

func NewProductUsecase(productRepository domain.ProductRepository, timeout time.Duration) domain.ProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
		contextTimeout:    timeout,
	}
}

func (pu *productUsecase) Create(c context.Context, product domain.Product) (message string, err error) {
	pr := pu.productRepository
	return pr.Create(c, product)
}

func (pu *productUsecase) GetAll(c context.Context) (product []domain.Product, err error) {
	pr := pu.productRepository
	return pr.GetAll(c)
}

func (pu *productUsecase) GetById(c context.Context, id string) (product domain.Product, err error) {
	pr := pu.productRepository
	return pr.GetById(c, id)
}

func (pu *productUsecase) Delete(c context.Context, id string) (message string, err error) {
	pr := pu.productRepository
	return pr.Delete(c, id)
}

func (pu *productUsecase) Update(c context.Context, id string, product domain.Product) (message string, err error) {
	pr := pu.productRepository
	return pr.Update(c, id, product)
}
