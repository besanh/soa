package repositories

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/besanh/soa/models"
	"github.com/uptrace/bun"
)

type (
	IProducts interface {
		Insert(ctx context.Context, product *models.Products) error
		Update(ctx context.Context, id string, product *models.Products) error
		Delete(ctx context.Context, id string) error
		Select(ctx context.Context, filter *models.ProductsQuery) (int, []models.ProductsResponse, error)
		SelectScroll(ctx context.Context, query *models.ProductsQuery) (result []models.ProductsResponse, err error)
	}

	Products struct{}
)

var ProductRepo IProducts

func NewProducts() IProducts {
	repo := &Products{}
	go func() {
		repo.initTable()
		repo.initColumns()
		repo.initIndexes()
	}()
	return repo
}

func (repo *Products) initTable() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	PgSqlClient.GetDB().RegisterModel((*models.Products)(nil))
	if err := CreateTable(PgSqlClient, ctx, (*models.Products)(nil)); err != nil {
		panic(err)
	}
}

func (repo *Products) initColumns() {

}

func (repo *Products) initIndexes() {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// defer cancel()

	// indexes := []string{
	// 	// Composite index on name and code.
	// 	"CREATE INDEX IF NOT EXISTS idx_products_combination ON products(product_name, code)",
	// }

	// if _, err := PgSqlClient.GetDB().NewCreateIndex().Model((*models.Products)(nil)).IfNotExists().Index("idx_products_combination").Column("domain_uuid").Exec(ctx); err != nil {
	// 	panic(err)
	// }
}

func (repo *Products) Insert(ctx context.Context, data *models.Products) error {
	resp, err := PgSqlClient.GetDB().NewInsert().Model(data).Exec(ctx)
	if err != nil {
		return err
	} else if affected, _ := resp.RowsAffected(); affected < 1 {
		return errors.New("insert product failed")
	}
	return nil
}

func (repo *Products) Update(ctx context.Context, id string, data *models.Products) error {
	query := PgSqlClient.GetDB().NewUpdate().Model(data).
		Where("product_id = ?", id)
	res, err := query.Exec(ctx)
	if err != nil {
		return err
	} else if affected, _ := res.RowsAffected(); affected < 1 {
		return errors.New("update product failed")
	}
	return nil
}

func (repo *Products) Delete(ctx context.Context, id string) error {
	query := PgSqlClient.GetDB().NewDelete().Model(&models.Products{}).
		Where("product_id = ?", id)
	res, err := query.Exec(ctx)
	if err != nil {
		return err
	} else if affected, _ := res.RowsAffected(); affected < 1 {
		return errors.New("delete product failed")
	}
	return nil
}

func (repo *Products) Select(ctx context.Context, filter *models.ProductsQuery) (int, []models.ProductsResponse, error) {
	result := new([]models.ProductsResponse)
	query := PgSqlClient.GetDB().NewSelect().Model(result).
		Relation("ProductCategory", func(sq *bun.SelectQuery) *bun.SelectQuery {
			return sq.Where("status = ?", "active").
				Where("product_categories_uuid = IN (?)", bun.In(filter.ProductCategoryId))
		}).
		Relation("Supplier", func(sq *bun.SelectQuery) *bun.SelectQuery {
			return sq.Where("status = ?", "active").
				Where("suppliers_uuid = IN (?)", bun.In(filter.SupplierId))
		})

	if len(filter.ProductName) > 0 {
		query.Where("? = ?", bun.Ident("product_name"), "%"+filter.ProductName+"%")
	}
	if len(filter.ProductReference) > 0 {
		query.Where("? = ?", bun.Ident("product_reference"), filter.ProductReference)
	}
	if len(filter.FromDateCreated) > 0 {
		query.Where("date_created >= ?", filter.FromDateCreated)
	}
	if len(filter.ToDateCreated) > 0 {
		query.Where("date_created <= ?", filter.ToDateCreated)
	}
	if len(filter.Status) > 0 {
		query.Where("status = IN (?)", bun.In(filter.Status))
	}
	if len(filter.FromPrice) > 0 {
		price, _ := strconv.Atoi(filter.FromPrice)
		query.Where("price >= ?", price)
	}
	if len(filter.ToPrice) > 0 {
		price, _ := strconv.Atoi(filter.ToPrice)
		query.Where("price <= ?", price)
	}

	if len(filter.FromQuantity) > 0 {
		quantity, _ := strconv.Atoi(filter.FromQuantity)
		query.Where("quantity >= ?", quantity)
	}
	if len(filter.ToQuantity) > 0 {
		quantity, _ := strconv.Atoi(filter.ToQuantity)
		query.Where("quantity <= ?", quantity)
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

func (repo *Products) SelectScroll(ctx context.Context, filter *models.ProductsQuery) ([]models.ProductsResponse, error) {
	result := new([]models.ProductsResponse)
	query := PgSqlClient.GetDB().NewSelect().Model(result).
		Relation("ProductCategory", func(sq *bun.SelectQuery) *bun.SelectQuery {
			return sq.Where("status = ?", "active").
				Where("product_categories_uuid = IN (?)", bun.In(filter.ProductCategoryId))
		}).
		Relation("Supplier", func(sq *bun.SelectQuery) *bun.SelectQuery {
			return sq.Where("status = ?", "active").
				Where("suppliers_uuid = IN (?)", bun.In(filter.SupplierId))
		})

	if len(filter.ProductName) > 0 {
		query.Where("? = ?", bun.Ident("product_name"), "%"+filter.ProductName+"%")
	}
	if len(filter.ProductReference) > 0 {
		query.Where("? = ?", bun.Ident("product_reference"), filter.ProductReference)
	}
	if len(filter.FromDateCreated) > 0 {
		query.Where("date_created >= ?", filter.FromDateCreated)
	}
	if len(filter.ToDateCreated) > 0 {
		query.Where("date_created <= ?", filter.ToDateCreated)
	}
	if len(filter.Status) > 0 {
		query.Where("status = IN (?)", bun.In(filter.Status))
	}
	if len(filter.FromPrice) > 0 {
		price, _ := strconv.Atoi(filter.FromPrice)
		query.Where("price >= ?", price)
	}
	if len(filter.ToPrice) > 0 {
		price, _ := strconv.Atoi(filter.ToPrice)
		query.Where("price <= ?", price)
	}

	if len(filter.FromQuantity) > 0 {
		quantity, _ := strconv.Atoi(filter.FromQuantity)
		query.Where("quantity >= ?", quantity)
	}
	if len(filter.ToQuantity) > 0 {
		quantity, _ := strconv.Atoi(filter.ToQuantity)
		query.Where("quantity <= ?", quantity)
	}

	if filter.Limit > 0 {
		query.Limit(filter.Limit).Offset(filter.Offset)
	}

	query.Where("(created_at > ?) OR (created_at = ? AND id > ?)", filter.CreatedAt, filter.CreatedAt, filter.LastSeenId).
		Order("created_at ASC, id ASC")

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	} else if err == sql.ErrNoRows {
		return nil, nil
	}

	return *result, nil
}
