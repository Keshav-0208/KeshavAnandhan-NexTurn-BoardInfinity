package repository

import (
	"ecommerce-inventory/model"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (repo *ProductRepository) AddProduct(product *model.Product) error {
	_, err := repo.DB.Exec(`INSERT INTO products (name, description, price, stock, category_id) 
	                         VALUES (?, ?, ?, ?, ?)`, product.Name, product.Description, product.Price, product.Stock, product.CategoryID)
	return err
}

func (repo *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	row := repo.DB.QueryRow(`SELECT id, name, description, price, stock, category_id FROM products WHERE id = ?`, id)

	var product model.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID); err != nil {
		return nil, fmt.Errorf("product not found")
	}

	return &product, nil
}
