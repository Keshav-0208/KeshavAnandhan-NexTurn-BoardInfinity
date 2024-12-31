package service

import (
	"ecommerce-inventory/model"
	"ecommerce-inventory/repository"
	"errors"
	"fmt"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (service *ProductService) AddProduct(product *model.Product) error {
	// Validate product data
	if err := validateProduct(product); err != nil {
		return err
	}

	if err := service.repo.AddProduct(product); err != nil {
		return fmt.Errorf("error adding product: %v", err)
	}
	return nil
}

func (service *ProductService) GetProductByID(id int) (*model.Product, error) {
	product, err := service.repo.GetProductByID(id)
	if err != nil {
		return nil, fmt.Errorf("product not found: %v", err)
	}
	return product, nil
}

func (service *ProductService) UpdateProduct(product *model.Product) error {
	// Validate product data
	if err := validateProduct(product); err != nil {
		return err
	}

	if err := service.repo.UpdateProduct(product); err != nil {
		return fmt.Errorf("error updating product: %v", err)
	}
	return nil
}

func (service *ProductService) DeleteProduct(id int) error {
	if err := service.repo.DeleteProduct(id); err != nil {
		return fmt.Errorf("error deleting product: %v", err)
	}
	return nil
}

func (service *ProductService) GetAllProducts(page, limit int) ([]model.Product, error) {
	products, err := service.repo.GetAllProducts(page, limit)
	if err != nil {
		return nil, fmt.Errorf("error retrieving products: %v", err)
	}
	return products, nil
}

func validateProduct(product *model.Product) error {
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price <= 0 {
		return errors.New("product price must be greater than 0")
	}
	if product.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}
	return nil
}
