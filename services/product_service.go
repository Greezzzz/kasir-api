package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll() ([]models.Products, error) {
	return s.repo.GetAll()
}

func (s *ProductService) Create(product *models.Products) error {
	return s.repo.Create(product)
}

func (s *ProductService) GetByID(id string) (*models.Products, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Update(product *models.Products) error {
	return s.repo.Update(product)
}

func (s *ProductService) Delete(id string) error {
	return s.repo.Delete(id)
}
