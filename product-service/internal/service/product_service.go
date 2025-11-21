package service

import (
	"context"
	"errors"

	"github.com/davidcm146/shopee-microservice/product-service/internal/dto"
	"github.com/davidcm146/shopee-microservice/product-service/internal/models"
	"github.com/davidcm146/shopee-microservice/product-service/internal/repository"
)

type ProductService interface {
	FindAll(ctx context.Context) ([]*models.Product, error)
	FindByID(ctx context.Context, id string) (*models.Product, error)
	Create(ctx context.Context, input *dto.CreateProductInput) (*models.Product, error)
	Update(ctx context.Context, id string, input *dto.UpdateProductInput) (*models.Product, error)
	SoftDelete(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) FindAll(ctx context.Context) ([]*models.Product, error) {
	return s.repo.FindAll(ctx)
}

func (s *productService) FindByID(ctx context.Context, id string) (*models.Product, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *productService) Create(ctx context.Context, input *dto.CreateProductInput) (*models.Product, error) {
	product := models.NewProduct(input.Name, input.SellerID, input.Category, input.Description, input.Price, input.Quantity, input.Features)
	return s.repo.Create(ctx, product)
}

func (s *productService) Update(ctx context.Context, id string, input *dto.UpdateProductInput) (*models.Product, error) {
	existingProduct, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existingProduct == nil || existingProduct.IsDeleted {
		return nil, errors.New("product not found or has been deleted")
	}

	if input.Name != nil {
		existingProduct.Name = *input.Name
	}

	if input.Category != nil {
		existingProduct.Category = *input.Category
	}

	if input.Description != nil {
		existingProduct.Description = *input.Description
	}

	if input.Price != nil {
		existingProduct.Price = *input.Price
	}

	if input.Quantity != nil {
		existingProduct.Quantity = *input.Quantity
	}

	if input.Features != nil {
		existingProduct.Features = *input.Features
	}

	updated, err := s.repo.Update(ctx, existingProduct)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *productService) SoftDelete(ctx context.Context, id string) error {
	existingProduct, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if existingProduct == nil || existingProduct.IsDeleted {
		return errors.New("product not found or already deleted")
	}

	err = s.repo.SoftDelete(ctx, id)
	return err
}

func (s *productService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
