package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/besanh/soa/models"
)

type (
	IProductCategories interface {
		Insert(ctx context.Context, data *models.ProductCategories) error
		Update(ctx context.Context, data *models.ProductCategories) error
		Delete(ctx context.Context, categoryId string) error
		Select(ctx context.Context, filter *models.ProductCategoriesQuery) (int, []models.ProductCategoriesResponse, error)
	}

	ProductCategoies struct{}
)

var ProductCategoryRepo IProductCategories

func NewProductCategories() IProductCategories {
	repo := &ProductCategoies{}
	repo.initTable()
	repo.initColumns()
	repo.initIndexes()
	return repo
}

func (repo *ProductCategoies) initTable() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	PgSqlClient.GetDB().RegisterModel((*models.ProductCategories)(nil))
	if err := CreateTable(PgSqlClient, ctx, (*models.ProductCategories)(nil)); err != nil {
		panic(err)
	}
}

func (repo *ProductCategoies) initColumns() {

}

func (repo *ProductCategoies) initIndexes() {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// defer cancel()

	// if _, err := PgSqlClient.GetDB().NewCreateIndex().Model((*models.ProductCategories)(nil)).IfNotExists().Index("idx_product_categories_combination").Column("product_category_name", "status").Exec(ctx); err != nil {
	// 	panic(err)
	// }
}

func (repo *ProductCategoies) Insert(ctx context.Context, data *models.ProductCategories) error {
	resp, err := PgSqlClient.GetDB().NewInsert().Model(data).Exec(ctx)
	if err != nil {
		return err
	} else if affected, _ := resp.RowsAffected(); affected < 1 {
		return errors.New("insert product category failed")
	}
	return nil
}

func (repo *ProductCategoies) Update(ctx context.Context, data *models.ProductCategories) error {
	query := PgSqlClient.GetDB().NewUpdate().Model(data).
		Where("product_category_id = ?", data.ProductCategoryId)
	res, err := query.Exec(ctx)
	if err != nil {
		return err
	} else if affected, _ := res.RowsAffected(); affected < 1 {
		return errors.New("update product category failed")
	}
	return nil
}

func (repo *ProductCategoies) Delete(ctx context.Context, categoryId string) error {
	query := PgSqlClient.GetDB().NewDelete().Model(&models.ProductCategories{}).
		Where("product_category_id = ?", categoryId)
	res, err := query.Exec(ctx)
	if err != nil {
		return err
	} else if affected, _ := res.RowsAffected(); affected < 1 {
		return errors.New("delete product category failed")
	}
	return nil
}

func (repo *ProductCategoies) Select(ctx context.Context, filter *models.ProductCategoriesQuery) (int, []models.ProductCategoriesResponse, error) {
	result := new([]models.ProductCategoriesResponse)
	query := PgSqlClient.GetDB().NewSelect().Model(result)

	if len(filter.ProductCategoryName) > 0 {
		query.Where("product_category_name = ?", filter.ProductCategoryName)
	}
	if len(filter.Status) > 0 {
		query.Where("status = ?", filter.Status)
	}
	if filter.Limit > 0 {
		query.Limit(filter.Limit).Offset(filter.Offset)
	}
	query.Order("created_at desc")

	total, err := query.ScanAndCount(ctx)
	if err != nil {
		return 0, nil, err
	} else if err == sql.ErrNoRows {
		return 0, nil, nil
	}

	return total, *result, nil
}
