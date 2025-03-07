package models

import (
	"errors"
	"slices"

	"github.com/uptrace/bun"
)

type (
	Suppliers struct {
		bun.BaseModel `bun:"supplier,alias:sp"`
		SupplierId    string `json:"supplier_id" bun:"supplier_id,pk,type:uuid,notnull"`
		SupplierName  string `json:"supplier_name" bun:"supplier_name,type:varchar(100),notnull"`
		Status        string `json:"status" bun:"status,type:varchar(25),notnull"`
	}

	SuppliersRequest struct {
		SupplierName string `json:"supplier_name" form:"supplier_name" binding:"required"`
		Status       string `json:"status" form:"status" binding:"required"`
	}

	SuppliersResponse struct {
		bun.BaseModel `bun:"supplier,alias:sp"`
		SupplierId    string `json:"supplier_id" bun:"supplier_id"`
		SupplierName  string `json:"supplier_name" bun:"supplier_name"`
		Status        string `json:"status" bun:"status"`
	}

	SuppliersQuery struct {
		SupplierId   string `json:"supplier_id"`
		SupplierName string `json:"supplier_name"`
		Status       string `json:"status"`
		Limit        int    `json:"limit"`
		Offset       int    `json:"offset"`
	}
)

func (m *SuppliersRequest) Validate() error {
	if len(m.SupplierName) == 0 {
		return errors.New("supplier name is required")
	}
	if !slices.Contains([]string{"active", "inactive"}, m.Status) {
		return errors.New("invalid status")
	}
	return nil
}
