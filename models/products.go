package models

import (
	"errors"
	"slices"

	"github.com/uptrace/bun"
)

type (
	Products struct {
		bun.BaseModel     `bun:"product,alias:p"`
		ProductId         string            `json:"product_id" bun:"product_id,pk,type:uuid,notnull"`
		ProductName       string            `json:"product_name" bun:"product_name,type:varchar(255),notnull"`
		ProductReference  string            `json:"product_reference" bun:"product_reference,type:varchar(255),notnull"`
		Status            string            `json:"status" bun:"status,type:varchar(25),notnull"`
		ProductCategoryId string            `json:"product_category_id" bun:"product_category_id,type:uuid,notnull"`
		ProductCategory   ProductCategories `json:"-" bun:"rel:belongs-to,join:product_category_id=product_category_id"`
		Price             int64             `json:"price" bun:"price,type:int,notnull"`
		StockLocation     string            `json:"stock_location" bun:"stock_location,type:varchar(100),notnull"`
		SupplierId        string            `json:"supplier_id" bun:"supplier_id,type:uuid,notnull"`
		Supplier          Suppliers         `json:"-" bun:"rel:belongs-to,join:supplier_id=supplier_id"`
		Quantity          int               `json:"quantity" bun:"quantity,type:int,notnull"`
		DateCreated       string            `json:"date_created" bun:"date_created,type:date,notnull"`
	}

	ProductsRequest struct {
		ProductName       string `json:"product_name" form:"product_name" binding:"required"`
		ProductReference  string `json:"product_reference" form:"product_reference" binding:"required"`
		Status            string `json:"status" form:"status" binding:"required"`
		ProductCategoryId string `json:"product_category_id" form:"product_category_id" binding:"required"`
		Price             int64  `json:"price" form:"price" binding:"required"`
		StockLocation     string `json:"stock_location" form:"stock_location" binding:"required"`
		SupplierId        string `json:"supplier_id" form:"supplier_id" binding:"required"`
		Quantity          int    `json:"quantity" form:"quantity" binding:"required"`
	}

	ProductsResponse struct {
		bun.BaseModel     `bun:"product,alias:p"`
		ProductId         string                     `json:"product_id" bun:"product_id,pk"`
		ProductName       string                     `json:"product_name" bun:"product_name"`
		ProductReference  string                     `json:"product_reference" bun:"product_reference"`
		Status            string                     `json:"status" bun:"status"`
		ProductCategoryId string                     `json:"product_category_id" bun:"product_category_id"`
		ProductCategory   *ProductCategoriesResponse `bun:"rel:belongs-to,join:product_category_id=product_category_id"`
		Price             int64                      `json:"price" bun:"price"`
		StockLocation     string                     `json:"stock_location" bun:"stock_location"`
		SupplierId        string                     `json:"supplier_id" bun:"supplier_id"`
		Supplier          *SuppliersResponse         `bun:"rel:belongs-to,join:supplier_id=supplier_id"`
		Quantity          int                        `json:"quantity" bun:"quantity"`
		DateCreated       string                     `json:"date_created" bun:"date_created"`
	}

	ProductsQuery struct {
		LastSeenId        string   `json:"last_seen_id"`
		CreatedAt         string   `json:"created_at"`
		ProductId         string   `json:"product_id"`
		ProductName       string   `json:"product_name"`
		ProductReference  string   `json:"product_reference"`
		Status            []string `json:"status"`
		ProductCategoryId []string `json:"product_category_id"`
		SupplierId        []string `json:"supplier_id"`
		FromDateCreated   string   `json:"from_date_created"`
		ToDateCreated     string   `json:"to_date_created"`
		FromPrice         string   `json:"from_price"`
		ToPrice           string   `json:"to_price"`
		FromQuantity      string   `json:"from_quantity"`
		ToQuantity        string   `json:"to_quantity"`
		Limit             int      `json:"limit"`
		Offset            int      `json:"offset"`
		Order             string   `json:"order"`
		Sort              string   `json:"sort"`
	}

	ProductsPage struct {
		Products   []ProductsResponse `json:"products"`
		NextCursor *Cursor            `json:"next_cursor,omitempty"`
	}

	// Cursor represents the composite cursor for pagination.
	Cursor struct {
		DateCreated string `json:"date_created"`
		ProductId   string `json:"product_id"`
	}
)

func (m *ProductsRequest) Validate() error {
	if !slices.Contains([]string{"on_order", "available", "out_of_stock"}, m.Status) {
		return errors.New("invalid status")
	}

	return nil
}
