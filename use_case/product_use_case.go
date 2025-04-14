package use_case

import (
	"crud-go/model"
	"crud-go/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(productRepository repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: productRepository,
	}
}

func (p ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := p.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId
	return product, nil
}

func (p *ProductUseCase) GetAllProducts() ([]model.Product, error) {
	return p.repository.GetAllProducts()
}

func (p ProductUseCase) GetProductById(product_id int) (*model.Product, error) {
	product, err := p.repository.GetProductById(product_id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p ProductUseCase) UpdateProduct(product model.Product) (model.Product, error) {
	err := p.repository.UpdateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}
