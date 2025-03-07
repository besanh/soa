package services

import (
	"context"
	"errors"
	"time"

	"github.com/besanh/soa/common/log"
	"github.com/besanh/soa/models"
	"github.com/besanh/soa/repositories"
	"github.com/google/uuid"
)

type (
	IProducts interface {
		Insert(ctx context.Context, data *models.ProductsRequest) error
		Update(ctx context.Context, id string, data *models.ProductsRequest) error
		Delete(ctx context.Context, id string) error
		Select(ctx context.Context, query *models.ProductsQuery) (total int, result []models.ProductsResponse, err error)
		SelectScroll(ctx context.Context, query *models.ProductsQuery) (result []models.ProductsResponse, err error)
	}

	Products struct {
	}
)

var ProductsService IProducts

func NewProducts() IProducts {
	return &Products{}
}

func (p *Products) Insert(ctx context.Context, data *models.ProductsRequest) error {
	total, _, err := repositories.ProductRepo.Select(ctx, &models.ProductsQuery{
		ProductName: data.ProductName,
		Limit:       1,
		Offset:      0,
	})
	if err != nil {
		log.Errorf("failed to get products: %v", err)
		return err
	} else if total > 0 {
		return errors.New("product already exists")
	}

	product := &models.Products{
		ProductId:         uuid.NewString(),
		ProductName:       data.ProductName,
		Price:             data.Price,
		Status:            data.Status,
		DateCreated:       time.Now().Format("2006-01-02"),
		Quantity:          1,
		StockLocation:     data.StockLocation,
		SupplierId:        data.SupplierId,
		ProductCategoryId: data.ProductCategoryId,
	}

	if err := repositories.ProductRepo.Insert(ctx, product); err != nil {
		log.Errorf("failed to insert product: %v", err)
		return err
	}
	return nil
}

func (p *Products) Update(ctx context.Context, id string, data *models.ProductsRequest) error {
	total, _, err := repositories.ProductRepo.Select(ctx, &models.ProductsQuery{
		ProductId: id,
		Limit:     1,
		Offset:    0,
	})
	if err != nil {
		log.Errorf("failed to get products: %v", err)
		return err
	} else if total <= 0 {
		return errors.New("product not found")
	}

	product := &models.Products{
		ProductId:         id,
		ProductName:       data.ProductName,
		Price:             data.Price,
		Status:            data.Status,
		DateCreated:       time.Now().Format("2006-01-02"),
		Quantity:          1,
		StockLocation:     data.StockLocation,
		SupplierId:        data.SupplierId,
		ProductCategoryId: data.ProductCategoryId,
	}

	if err := repositories.ProductRepo.Update(ctx, id, product); err != nil {
		log.Errorf("failed to update product: %v", err)
		return err
	}
	return nil
}

func (p *Products) Delete(ctx context.Context, id string) error {
	total, _, err := repositories.ProductRepo.Select(ctx, &models.ProductsQuery{
		ProductId: id,
		Limit:     1,
		Offset:    0,
	})
	if err != nil {
		log.Errorf("failed to get products: %v", err)
		return err
	} else if total <= 0 {
		return errors.New("product not found")
	}

	if err := repositories.ProductRepo.Delete(ctx, id); err != nil {
		log.Errorf("failed to delete product: %v", err)
		return err
	}
	return nil
}

func (p *Products) Select(ctx context.Context, query *models.ProductsQuery) (total int, result []models.ProductsResponse, err error) {
	return repositories.ProductRepo.Select(ctx, query)
}

func (p *Products) SelectScroll(ctx context.Context, query *models.ProductsQuery) (result []models.ProductsResponse, err error) {
	return repositories.ProductRepo.SelectScroll(ctx, query)
}
