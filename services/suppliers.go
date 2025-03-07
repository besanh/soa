package services

import (
	"context"
	"errors"

	"github.com/besanh/soa/common/log"
	"github.com/besanh/soa/models"
	"github.com/besanh/soa/repositories"
	"github.com/google/uuid"
)

type (
	ISuppliers interface {
		Insert(ctx context.Context, data *models.SuppliersRequest) error
		Update(ctx context.Context, id string, data *models.SuppliersRequest) error
		Delete(ctx context.Context, supplierId string) error
		Select(ctx context.Context, filter *models.SuppliersQuery) (int, []models.SuppliersResponse, error)
	}

	Suppliers struct{}
)

var SuppliersService ISuppliers

func NewSuppliers() ISuppliers {
	service := &Suppliers{}
	return service
}

func (s *Suppliers) Insert(ctx context.Context, data *models.SuppliersRequest) error {
	// Get supplier
	total, _, err := repositories.SupplierRepo.Select(ctx, &models.SuppliersQuery{
		SupplierName: data.SupplierName,
		Status:       "active",
		Limit:        1,
		Offset:       0,
	})
	if err != nil {
		return err
	} else if total > 0 {
		return errors.New("supplier already exists")
	}

	// Insert
	supplier := &models.Suppliers{
		SupplierId:   uuid.NewString(),
		SupplierName: data.SupplierName,
		Status:       data.Status,
	}

	if err := repositories.SupplierRepo.Insert(ctx, supplier); err != nil {
		log.Errorf("failed to insert supplier: %v", err)
		return err
	}

	return nil
}

func (s *Suppliers) Update(ctx context.Context, id string, data *models.SuppliersRequest) error {
	// Get supplier
	total, supplierExist, err := repositories.SupplierRepo.Select(ctx, &models.SuppliersQuery{
		SupplierId: id,
		Limit:      1,
		Offset:     0,
	})
	if err != nil {
		return err
	} else if total <= 0 {
		return errors.New("supplier not found")
	}

	// Update
	supplier := &models.Suppliers{
		SupplierId:   supplierExist[0].SupplierId,
		SupplierName: data.SupplierName,
		Status:       data.Status,
	}

	if err := repositories.SupplierRepo.Update(ctx, supplier); err != nil {
		log.Errorf("failed to update supplier: %v", err)
		return err
	}
	return nil
}

func (s *Suppliers) Delete(ctx context.Context, supplierId string) error {
	// Get supplier
	total, _, err := repositories.SupplierRepo.Select(ctx, &models.SuppliersQuery{
		SupplierId: supplierId,
		Limit:      1,
		Offset:     0,
	})
	if err != nil {
		log.Errorf("failed to get supplier: %v", err)
		return err
	} else if total <= 0 {
		return errors.New("supplier not found")
	}

	// Delete
	if err := repositories.SupplierRepo.Delete(ctx, supplierId); err != nil {
		log.Errorf("failed to delete supplier: %v", err)
		return err
	}
	return nil
}

func (s *Suppliers) Select(ctx context.Context, filter *models.SuppliersQuery) (int, []models.SuppliersResponse, error) {
	return repositories.SupplierRepo.Select(ctx, filter)
}
