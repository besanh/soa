package repositories

import (
	"context"

	"github.com/besanh/soa/models"
)

type (
	IStatistics interface {
		GetProductsPerCategory(ctx context.Context) ([]models.ProductsPerCategoryStat, error)
		GetProductsPerSupplier(ctx context.Context) ([]models.ProductsPerSupplierStat, error)
	}

	Statistics struct{}
)

var StatisticsRepo IStatistics

func NewStatisticsRepo() IStatistics {
	return &Statistics{}
}

func (repo *Statistics) GetProductsPerCategory(ctx context.Context) ([]models.ProductsPerCategoryStat, error) {
	result := new([]models.ProductsPerCategoryStat)
	err := PgSqlClient.GetDB().NewSelect().
		Model(result).
		ColumnExpr("pc.product_category_name as category, COUNT(*) as count").
		TableExpr("products p JOIN product_categories pc ON p.product_category_id = pc.product_category_id").
		Group("pc.product_category_name").
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return *result, nil
}

func (repo *Statistics) GetProductsPerSupplier(ctx context.Context) ([]models.ProductsPerSupplierStat, error) {
	result := new([]models.ProductsPerSupplierStat)
	err := PgSqlClient.GetDB().NewSelect().
		Model(result).
		ColumnExpr("s.supplier_name as supplier, COUNT(*) as count").
		TableExpr("products p JOIN suppliers s ON p.supplier_id = s.supplier_id").
		Group("s.supplier_name").
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return *result, nil
}
