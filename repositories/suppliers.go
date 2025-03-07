package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/besanh/soa/models"
)

type (
	ISuppliers interface {
		Insert(ctx context.Context, data *models.Suppliers) error
		Update(ctx context.Context, data *models.Suppliers) error
		Delete(ctx context.Context, supplierId string) error
		Select(ctx context.Context, filter *models.SuppliersQuery) (int, []models.SuppliersResponse, error)
	}

	Suppliers struct{}
)

var SupplierRepo ISuppliers

func NewSuppliers() ISuppliers {
	repo := &Suppliers{}
	repo.initTable()
	repo.initColumns()
	repo.initIndexes()
	return repo
}

func (repo *Suppliers) initTable() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	PgSqlClient.GetDB().RegisterModel((*models.Suppliers)(nil))
	if err := CreateTable(PgSqlClient, ctx, (*models.Suppliers)(nil)); err != nil {
		panic(err)
	}
}

func (repo *Suppliers) initColumns() {

}

func (repo *Suppliers) initIndexes() {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// defer cancel()

	// if _, err := PgSqlClient.GetDB().NewCreateIndex().Model((*models.Suppliers)(nil)).IfNotExists().Index("idx_suppliers_combination").Column("supplier_name", "status").Exec(ctx); err != nil {
	// 	panic(err)
	// }
}

func (repo *Suppliers) Insert(ctx context.Context, data *models.Suppliers) error {
	resp, err := PgSqlClient.GetDB().NewInsert().Model(data).Exec(ctx)
	if err != nil {
		return err
	} else if affected, _ := resp.RowsAffected(); affected < 1 {
		return errors.New("insert supplier failed")
	}
	return nil
}

func (repo *Suppliers) Update(ctx context.Context, data *models.Suppliers) error {
	query := PgSqlClient.GetDB().NewUpdate().Model(data).
		Where("supplier_id = ?", data.SupplierId)
	res, err := query.Exec(ctx)
	if err != nil {
		return err
	} else if affected, _ := res.RowsAffected(); affected < 1 {
		return errors.New("update supplier failed")
	}
	return nil
}

func (repo *Suppliers) Delete(ctx context.Context, supplierId string) error {
	query := PgSqlClient.GetDB().NewDelete().Model(&models.Suppliers{}).
		Where("supplier_id = ?", supplierId)
	res, err := query.Exec(ctx)
	if err != nil {
		return err
	} else if affected, _ := res.RowsAffected(); affected < 1 {
		return errors.New("delete supplier failed")
	}
	return nil
}

func (repo *Suppliers) Select(ctx context.Context, filter *models.SuppliersQuery) (int, []models.SuppliersResponse, error) {
	result := new([]models.SuppliersResponse)
	query := PgSqlClient.GetDB().NewSelect().Model(result)
	if len(filter.SupplierName) > 0 {
		query.Where("supplier_name = ?", filter.SupplierName)
	}
	if len(filter.Status) > 0 {
		query.Where("status = ?", filter.Status)
	}
	if filter.Limit > 0 {
		query.Limit(filter.Limit).Offset(filter.Offset)
	}

	total, err := query.ScanAndCount(ctx)
	if err != nil {
		return 0, nil, err
	} else if err == sql.ErrNoRows {
		return 0, nil, nil
	}

	return total, *result, nil
}
