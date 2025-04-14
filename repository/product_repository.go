package repository

import (
	"crud-go/model"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return ProductRepository{
		connection: db,
	}
}

func (p *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := p.connection.Prepare("INSERT INTO products (product_name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (p *ProductRepository) GetAllProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM products"
	rows, err := p.connection.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return []model.Product{}, err
	}
	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price)
		if err != nil {
			fmt.Println(err.Error())
			return []model.Product{}, err
		}
		productList = append(productList, productObj)
	}

	rows.Close()
	return productList, err
}

func (p *ProductRepository) GetProductById(product_id int) (*model.Product, error) {
	query, err := p.connection.Prepare("SELECT * FROM products WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product model.Product

	err = query.QueryRow(product_id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	query.Close()
	return &product, nil
}

func (p *ProductRepository) UpdateProduct(product model.Product) error {
	query, err := p.connection.Prepare("UPDATE products SET product_name = $1, price = $2 WHERE id = $3")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = query.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	query.Close()
	return nil
}
