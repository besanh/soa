package services

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/besanh/soa/common/log"
	"github.com/besanh/soa/models"
	"github.com/besanh/soa/repositories"
	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
)

type (
	IProducts interface {
		Insert(ctx context.Context, data *models.ProductsRequest) error
		Update(ctx context.Context, id string, data *models.ProductsRequest) error
		Delete(ctx context.Context, id string) error
		Select(ctx context.Context, query *models.ProductsQuery) (total int, result []models.ProductsResponse, err error)
		SelectById(ctx context.Context, id string) (result models.ProductsResponse, err error)
		SelectScroll(ctx context.Context, query *models.ProductsQuery) (result *models.ProductsPage, err error)
		ExportPdf(ctx context.Context, query *models.ProductsQuery) (*gofpdf.Fpdf, error)
	}

	Products struct {
	}
)

var ProductsService IProducts

func NewProducts() IProducts {
	return &Products{}
}

func (p *Products) Insert(ctx context.Context, data *models.ProductsRequest) error {
	total, _, err := repositories.ProductRepo.Select(ctx, &models.ProductsQuery{
		ProductName: data.ProductName,
		Limit:       1,
		Offset:      0,
	})
	if err != nil {
		log.Errorf("failed to get products: %v", err)
		return err
	} else if total > 0 {
		return errors.New("product already exists")
	}

	// Get category
	total, _, err = repositories.ProductCategoryRepo.Select(ctx, &models.ProductCategoriesQuery{
		ProductCategoryId: data.ProductCategoryId,
		Limit:             1,
		Offset:            0,
	})
	if err != nil {
		log.Errorf("failed to get product categories: %v", err)
		return err
	} else if total <= 0 {
		return errors.New("product category not found")
	}

	// Get supplier
	total, _, err = repositories.SupplierRepo.Select(ctx, &models.SuppliersQuery{
		SupplierId: data.SupplierId,
		Limit:      1,
		Offset:     0,
	})
	if err != nil {
		log.Errorf("failed to get suppliers: %v", err)
		return err
	} else if total <= 0 {
		return errors.New("supplier not found")
	}

	product := &models.Products{
		ProductId:         uuid.NewString(),
		ProductName:       data.ProductName,
		Price:             data.Price,
		Status:            data.Status,
		DateCreated:       time.Now().Format("2006-01-02"),
		Quantity:          1,
		StockLocation:     data.StockLocation,
		SupplierId:        data.SupplierId,
		ProductCategoryId: data.ProductCategoryId,
	}

	if err := repositories.ProductRepo.Insert(ctx, product); err != nil {
		log.Errorf("failed to insert product: %v", err)
		return err
	}
	return nil
}

func (p *Products) Update(ctx context.Context, id string, data *models.ProductsRequest) error {
	total, _, err := repositories.ProductRepo.Select(ctx, &models.ProductsQuery{
		ProductId: id,
		Limit:     1,
		Offset:    0,
	})
	if err != nil {
		log.Errorf("failed to get products: %v", err)
		return err
	} else if total <= 0 {
		return errors.New("product not found")
	}

	product := &models.Products{
		ProductId:         id,
		ProductName:       data.ProductName,
		Price:             data.Price,
		Status:            data.Status,
		DateCreated:       time.Now().Format("2006-01-02"),
		Quantity:          1,
		StockLocation:     data.StockLocation,
		SupplierId:        data.SupplierId,
		ProductCategoryId: data.ProductCategoryId,
	}

	if err := repositories.ProductRepo.Update(ctx, id, product); err != nil {
		log.Errorf("failed to update product: %v", err)
		return err
	}
	return nil
}

func (p *Products) Delete(ctx context.Context, id string) error {
	total, _, err := repositories.ProductRepo.Select(ctx, &models.ProductsQuery{
		ProductId: id,
		Limit:     1,
		Offset:    0,
	})
	if err != nil {
		log.Errorf("failed to get products: %v", err)
		return err
	} else if total <= 0 {
		return errors.New("product not found")
	}

	if err := repositories.ProductRepo.Delete(ctx, id); err != nil {
		log.Errorf("failed to delete product: %v", err)
		return err
	}
	return nil
}

/*
 * Use basic case query
 */
func (p *Products) Select(ctx context.Context, query *models.ProductsQuery) (total int, result []models.ProductsResponse, err error) {
	// TODO: Apply caching from Redis to optimize performance

	return repositories.ProductRepo.Select(ctx, query)
}

/*
 * Use scroll, optimizing performance
 */
func (p *Products) SelectScroll(ctx context.Context, query *models.ProductsQuery) (result *models.ProductsPage, err error) {
	// TODO: Apply caching from Redis to optimize performance
	products, err := repositories.ProductRepo.SelectScroll(ctx, query)
	if err != nil {
		return
	}
	var nextCursor *models.Cursor
	if len(products) > 0 {
		// Set the next cursor to the last record in the current page.
		lastProduct := products[len(products)-1]
		nextCursor = &models.Cursor{
			DateCreated: lastProduct.DateCreated,
			ProductId:   lastProduct.ProductId,
		}
	}
	return &models.ProductsPage{
		Products:   products,
		NextCursor: nextCursor,
	}, nil
}

func (p *Products) ExportPdf(ctx context.Context, query *models.ProductsQuery) (*gofpdf.Fpdf, error) {
	products, err := repositories.ProductRepo.SelectScroll(ctx, query)
	if err != nil {
		return nil, err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Title.
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Product Information")
	pdf.Ln(12)

	// Table header.
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(10, 10, "#", "1", 0, "C", false, 0, "")
	pdf.CellFormat(80, 10, "Product Reference", "1", 0, "C", false, 0, "")
	pdf.CellFormat(80, 10, "Product Name", "1", 0, "C", false, 0, "")
	pdf.CellFormat(80, 10, "Date Added", "1", 0, "C", false, 0, "")
	pdf.CellFormat(80, 10, "Status", "1", 0, "C", false, 0, "")
	pdf.CellFormat(80, 10, "Product Category", "1", 0, "C", false, 0, "")
	pdf.CellFormat(80, 10, "Price", "1", 0, "C", false, 0, "")
	pdf.CellFormat(80, 10, "Stock Location(city)", "1", 0, "C", false, 0, "")
	pdf.CellFormat(80, 10, "Supplier", "1", 0, "C", false, 0, "")
	pdf.CellFormat(80, 10, "Available Quantity", "1", 0, "C", false, 0, "")
	pdf.Ln(-1)

	// Table rows.
	pdf.SetFont("Arial", "", 12)
	for i, p := range products {
		pdf.CellFormat(10, 10, strconv.Itoa(i+1), "1", 0, "C", false, 0, "")
		pdf.CellFormat(80, 10, p.ProductReference, "1", 0, "L", false, 0, "")
		pdf.CellFormat(80, 10, p.ProductName, "1", 0, "L", false, 0, "")
		pdf.CellFormat(80, 10, p.DateCreated, "1", 0, "C", false, 0, "")
		pdf.CellFormat(80, 10, p.Status, "1", 0, "C", false, 0, "")
		pdf.CellFormat(80, 10, p.ProductCategory.ProductCategoryName, "1", 0, "L", false, 0, "")
		pdf.CellFormat(80, 10, strconv.FormatInt(p.Price, 10), "1", 0, "R", false, 0, "")
		pdf.CellFormat(80, 10, p.StockLocation, "1", 0, "L", false, 0, "")
		pdf.CellFormat(80, 10, p.Supplier.SupplierName, "1", 0, "L", false, 0, "")
		pdf.CellFormat(80, 10, strconv.Itoa(p.Quantity), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}

	return pdf, nil
}

func (p *Products) SelectById(ctx context.Context, id string) (result models.ProductsResponse, err error) {
	return repositories.ProductRepo.SelectById(ctx, id)
}
