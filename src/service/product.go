package service

import (
	"context"
	"synapsis-ecommerce/src/repository"
)

func (s *service) GetProducts(ctx context.Context) ([]repository.Product, error) {

	products, err := s.repo.GetProducts(ctx)
	if err != nil {
		return []repository.Product{}, err
	}

	return products, nil
}

func (s *service) GetProductByCategory(ctx context.Context, category string) ([]repository.Product, error) {

	products, err := s.repo.GetProductByCategory(ctx, category)
	if err != nil {
		return []repository.Product{}, err
	}

	return products, nil
}
