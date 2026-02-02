package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll() ([]models.Category, error) {
	query := "SELECT id, name, description FROM categories"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]models.Category, 0)

	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *CategoryRepository) Create(category *models.Category) error {
	query := "INSERT INTO categories (id, name, description) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(query, category.ID, category.Name, category.Description).Scan(&category.ID)
	return err
}

func (r *CategoryRepository) GetByID(id string) (*models.Category, error) {
	query := "SELECT id, name, description FROM categories WHERE id = $1"

	var category models.Category
	err := r.db.QueryRow(query, id).Scan(&category.ID, &category.Name, &category.Description)
	if err == sql.ErrNoRows {
		return nil, errors.New("category not found")
	}

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) Update(category *models.Category) error {
	query := "UPDATE categories SET name = $1, description = $2 WHERE id = $3"
	result, err := r.db.Exec(query, category.Name, category.Description, category.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("category not found")
	}
	return nil
}

func (r *CategoryRepository) Delete(id string) error {
	query := "DELETE FROM categories WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("category not found")
	}
	return err
}
