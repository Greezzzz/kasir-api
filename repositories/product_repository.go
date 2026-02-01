package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]models.Products, error) {
	query := "SELECT id, name, price, stock, created_at, id_category FROM products"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Products, 0)

	for rows.Next() {
		var product models.Products
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CreatedAt, &product.IDCategory)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) Create(product *models.Products) error {
	query := "INSERT INTO products (id, name, price, stock, id_category) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := r.db.QueryRow(query, product.ID, product.Name, product.Price, product.Stock, product.IDCategory).Scan(&product.ID)
	return err
}

func (r *ProductRepository) GetByID(id string) (*models.Products, error) {
	query := "SELECT id, name, price, stock, created_at, id_category FROM products WHERE id = $1"

	var product models.Products
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CreatedAt, &product.IDCategory)
	if err == sql.ErrNoRows {
		return nil, errors.New("product not found")
	}
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) Update(product *models.Products) error {
	query := "UPDATE products SET name = $1, price = $2, stock = $3, id_category = $4 WHERE id = $5"
	result, err := r.db.Exec(query, product.Name, product.Price, product.Stock, product.IDCategory, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (r *ProductRepository) Delete(id string) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}
	return err
}
