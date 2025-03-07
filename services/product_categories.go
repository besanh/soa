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
	IProductCategories interface {
		Insert(ctx context.Context, data *models.ProductCategoriesRequest) error
		Update(ctx context.Context, id string, data *models.ProductCategoriesRequest) error
		Delete(ctx context.Context, categoryId string) error
		Select(ctx context.Context, filter *models.ProductCategoriesQuery) (int, []models.ProductCategoriesResponse, error)
	}

	ProductCategories struct{}
)

var ProductCategoriesService IProductCategories

func NewProductCategories() IProductCategories {
	return &ProductCategories{}
}

func (s *ProductCategories) Insert(ctx context.Context, data *models.ProductCategoriesRequest) error {
	total, _, err := repositories.ProductCategoryRepo.Select(ctx, &models.ProductCategoriesQuery{
		ProductCategoryName: data.ProductCategoryName,
		Status:              "active",
		Limit:               1,
		Offset:              0,
	})
	if err != nil {
		log.Errorf("failed to get product categories: %v", err)
		return err
	} else if total > 0 {
		return errors.New("product category already exists")
	}

	productCategory := &models.ProductCategories{
		ProductCategoryId:   uuid.NewString(),
		ProductCategoryName: data.ProductCategoryName,
		Status:              data.Status,
		CreatedAt:           time.Now().Format("2006-01-02"),
		UpdatedAt:           time.Now().Format("2006-01-02"),
	}

	if err := repositories.ProductCategoryRepo.Insert(ctx, productCategory); err != nil {
		log.Errorf("failed to insert product category: %v", err)
		return err
	}
	return nil
}

func (s *ProductCategories) Update(ctx context.Context, id string, data *models.ProductCategoriesRequest) error {
	total, productCategoryExist, err := repositories.ProductCategoryRepo.Select(ctx, &models.ProductCategoriesQuery{
		ProductCategoryId: id,
		Limit:             1,
		Offset:            0,
	})
	if err != nil {
		log.Errorf("failed to get product categories: %v", err)
		return err
	} else if total <= 0 {
		return errors.New("product category not found")
	}

	productCategory := &models.ProductCategories{
		ProductCategoryId:   productCategoryExist[0].ProductCategoryId,
		ProductCategoryName: data.ProductCategoryName,
		Status:              data.Status,
		CreatedAt:           productCategoryExist[0].CreatedAt,
		UpdatedAt:           time.Now().Format("2006-01-02"),
	}

	if err := repositories.ProductCategoryRepo.Update(ctx, productCategory); err != nil {
		log.Errorf("failed to update product category: %v", err)
		return err
	}
	return nil
}

func (s *ProductCategories) Delete(ctx context.Context, categoryId string) error {
	total, productCategoryExist, err := repositories.ProductCategoryRepo.Select(ctx, &models.ProductCategoriesQuery{
		ProductCategoryId: categoryId,
		Limit:             1,
		Offset:            0,
	})
	if err != nil {
		log.Errorf("failed to get product categories: %v", err)
		return err
	} else if total <= 0 {
		return errors.New("product category not found")
	}

	if err := repositories.ProductCategoryRepo.Delete(ctx, productCategoryExist[0].ProductCategoryId); err != nil {
		log.Errorf("failed to delete product category: %v", err)
		return err
	}
	return nil
}

func (s *ProductCategories) Select(ctx context.Context, filter *models.ProductCategoriesQuery) (int, []models.ProductCategoriesResponse, error) {
	return repositories.ProductCategoryRepo.Select(ctx, filter)
}
