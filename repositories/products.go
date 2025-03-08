package repositories

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/besanh/soa/common/log"
	"github.com/besanh/soa/models"
	"github.com/uptrace/bun"
)

type (
	IProducts interface {
		Insert(ctx context.Context, product *models.Products) error
		Update(ctx context.Context, id string, product *models.Products) error
		Delete(ctx context.Context, id string) error
		Select(ctx context.Context, filter *models.ProductsQuery) (int, []models.ProductsResponse, error)
		SelectById(ctx context.Context, id string) (result models.ProductsResponse, err error)
		SelectScroll(ctx context.Context, query *models.ProductsQuery) (result []models.ProductsResponse, err error)
	}

	Products struct{}
)

var ProductRepo IProducts

func NewProducts() IProducts {
	repo := &Products{}
	repo.initTable()
	repo.initColumns()
	repo.initIndexes()
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if _, err := PgSqlClient.GetDB().NewCreateIndex().Model((*models.Products)(nil)).IfNotExists().Index("idx_products_combination").Column("product_name", "product_reference", "status", "date_created", "price", "quantity").Exec(ctx); err != nil {
		log.Errorf("failed to create index: %v", err)
	}
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
			// sq = sq.Where("pc.status = ?", "active")
			if len(filter.ProductCategoryId) > 0 {
				sq.Where("product_categories_uuid = IN (?)", bun.In(filter.ProductCategoryId))
			}
			return sq
		}).
		Relation("Supplier", func(sq *bun.SelectQuery) *bun.SelectQuery {
			// sq = sq.Where("sp.status = ?", "active")
			if len(filter.SupplierId) > 0 {
				sq.Where("suppliers_uuid = IN (?)", bun.In(filter.SupplierId))
			}
			return sq
		})

	if len(filter.ProductName) > 0 {
		query.Where("? ILIKE ?", bun.Ident("product_name"), "%"+filter.ProductName+"%")
	}
	if len(filter.ProductReference) > 0 {
		query.Where("? ILIKE ?", bun.Ident("product_reference"), filter.ProductReference)
	}
	if len(filter.FromDateCreated) > 0 {
		query.Where("date_created >= ?", filter.FromDateCreated)
	}
	if len(filter.ToDateCreated) > 0 {
		query.Where("date_created <= ?", filter.ToDateCreated)
	}
	if len(filter.Status) > 0 {
		query.Where("status IN (?)", bun.In(filter.Status))
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

	orderValue := "date_created"
	sortValue := "DESC"
	switch filter.Order {
	case "product_name":
		orderValue = "product_name"
	case "product_reference":
		orderValue = "product_reference"
	case "date_created":
		orderValue = "date_created"
	case "status":
		orderValue = "u.username"
	case "price":
		orderValue = "p.id"
	case "stock_location":
		orderValue = "stock_location"
	case "quantity":
		orderValue = "quantity"
	}
	if len(filter.Sort) > 0 && filter.Sort == "asc" {
		sortValue = "ASC"
	}
	query.OrderExpr(orderValue + " " + sortValue)
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
			if len(filter.ProductCategoryId) > 0 {
				sq.Where("product_categories_uuid = IN (?)", bun.In(filter.ProductCategoryId))
			}
			return sq
		}).
		Relation("Supplier", func(sq *bun.SelectQuery) *bun.SelectQuery {
			if len(filter.SupplierId) > 0 {
				sq.Where("suppliers_uuid = IN (?)", bun.In(filter.SupplierId))
			}
			return sq
		})

	if len(filter.ProductName) > 0 {
		query.Where("? ILIKE ?", bun.Ident("product_name"), "%"+filter.ProductName+"%")
	}
	if len(filter.ProductReference) > 0 {
		query.Where("? ILIKE ?", bun.Ident("product_reference"), filter.ProductReference)
	}
	if len(filter.FromDateCreated) > 0 {
		query.Where("date_created >= ?", filter.FromDateCreated)
	}
	if len(filter.ToDateCreated) > 0 {
		query.Where("date_created <= ?", filter.ToDateCreated)
	}
	if len(filter.Status) > 0 {
		query.Where("status IN (?)", bun.In(filter.Status))
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
		query.Limit(filter.Limit)
	}

	if len(filter.LastSeenId) > 0 && len(filter.CreatedAt) > 0 {
		query.Where("(date_created > ?) OR (date_created = ? AND product_id > ?)", filter.CreatedAt, filter.CreatedAt, filter.LastSeenId).
			Order("date_created ASC, id ASC")
	}

	// I forgot order columns

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	} else if err == sql.ErrNoRows {
		return nil, nil
	}

	return *result, nil
}

func (repo *Products) SelectById(ctx context.Context, id string) (result models.ProductsResponse, err error) {
	query := PgSqlClient.GetDB().NewSelect().Model(&result).
		Relation("ProductCategory").
		Relation("Supplier")

	err = query.Where("id = ?", id).Scan(ctx)
	if err != nil {
		return models.ProductsResponse{}, err
	} else if err == sql.ErrNoRows {
		return models.ProductsResponse{}, nil
	}

	return result, nil
}
