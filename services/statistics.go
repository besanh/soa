package services

import (
	"context"

	"github.com/besanh/soa/models"
	"github.com/besanh/soa/repositories"
)

type (
	IStatistics interface {
		GetStatisticsProductsPerCategory(ctx context.Context) ([]models.ProductsPerCategoryStat, error)
		GetStatisticsProductsPerSupplier(ctx context.Context) ([]models.ProductsPerSupplierStat, error)
	}

	Statistics struct{}
)

func NewStatistics() IStatistics {
	return &Statistics{}
}

func (s *Statistics) GetStatisticsProductsPerCategory(ctx context.Context) ([]models.ProductsPerCategoryStat, error) {
	total, _, err := repositories.ProductRepo.Select(ctx, &models.ProductsQuery{
		Limit:  -1,
		Offset: 0,
	})
	if err != nil {
		return nil, err
	}

	stats, err := repositories.StatisticsRepo.GetProductsPerCategory(ctx)
	if err != nil {
		return nil, err
	}

	for i, stat := range stats {
		stats[i].Percent = (float64(stat.Count) / float64(total)) * 100
	}

	return stats, nil
}

func (s *Statistics) GetStatisticsProductsPerSupplier(ctx context.Context) ([]models.ProductsPerSupplierStat, error) {
	total, _, err := repositories.ProductRepo.Select(ctx, &models.ProductsQuery{
		Limit:  -1,
		Offset: 0,
	})
	if err != nil {
		return nil, err
	}

	stats, err := repositories.StatisticsRepo.GetProductsPerSupplier(ctx)
	if err != nil {
		return nil, err
	}

	for i, stat := range stats {
		stats[i].Percent = (float64(stat.Count) / float64(total)) * 100
	}

	return stats, nil
}
