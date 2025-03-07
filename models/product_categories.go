package models

import (
	"errors"
	"slices"

	"github.com/uptrace/bun"
)

type (
	ProductCategories struct {
		bun.BaseModel       `bun:"product_categories,alias:pc"`
		ProductCategoryId   string `json:"product_category_id" bun:"product_category_id,pk,type:uuid,notnull"`
		ProductCategoryName string `json:"product_category_name" bun:"product_category_name,type:varchar(255),notnull"`
		Status              string `json:"status" bun:"status,type:varchar(25),notnull"`
		CreatedAt           string `json:"created_at" bun:"created_at,notnull"`
		UpdatedAt           string `json:"updated_at" bun:"updated_at,notnull"`
	}

	ProductCategoriesRequest struct {
		ProductCategoryName string `json:"product_category_name" form:"product_category_name" binding:"required"`
		Status              string `json:"status" form:"status" binding:"required"`
	}

	ProductCategoriesResponse struct {
		bun.BaseModel       `bun:"product_categories,alias:pc"`
		ProductCategoryId   string `json:"product_category_id" bun:"product_category_id"`
		ProductCategoryName string `json:"product_category_name" bun:"product_category_name"`
		Status              string `json:"status" bun:"status"`
		CreatedAt           string `json:"created_at" bun:"created_at"`
		UpdatedAt           string `json:"updated_at" bun:"updated_at"`
	}

	ProductCategoriesQuery struct {
		ProductCategoryId   string `json:"product_category_id"`
		ProductCategoryName string `json:"product_category_name"`
		Status              string `json:"status"`
		Limit               int    `json:"limit"`
		Offset              int    `json:"offset"`
	}
)

func (m *ProductCategoriesRequest) Validate() error {
	if len(m.ProductCategoryName) == 0 {
		return errors.New("product category name is required")
	}
	if !slices.Contains([]string{"active", "inactive"}, m.Status) {
		return errors.New("invalid status")
	}
	return nil
}
