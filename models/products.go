package models

import "github.com/uptrace/bun"

type (
	Products struct {
		bun.BaseModel     `bun:"products,alias:p"`
		ProductId         string            `json:"product_id" bun:"product_id,pk,type:uuid,notnull"`
		ProductName       string            `json:"product_name" bun:"product_name,type:varchar(255),notnull"`
		Status            string            `json:"status" bun:"status,type:varchar(25),notnull"`
		ProductCategoryId string            `json:"product_category_id" bun:"product_category_id,type:uuid,notnull"`
		ProductCategory   ProductCategories `bun:"rel:belongs-to,join:product_category_id=product_category_id"`
		Price             int               `json:"price" bun:"price,type:int,notnull"`
		StockLocation     string            `json:"stock_location" bun:"stock_location,type:varchar(100),notnull"`
		SupplierId        string            `json:"supplier_id" bun:"supplier_id,type:uuid,notnull"`
		Supplier          Suppliers         `bun:"rel:belongs-to,join:supplier_id=supplier_id"`
		Quantity          int               `json:"quantity" bun:"quantity,type:int,notnull"`
		DateCreated       string            `json:"date_created" bun:"date_created,type:date,notnull"`
	}
)
