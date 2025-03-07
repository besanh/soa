package repositories

import (
	"context"
	"time"

	"github.com/besanh/soa/models"
)

type (
	IProducts interface{}

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

}
