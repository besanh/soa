package v1

import (
	api "github.com/besanh/soa/apis"
	"github.com/besanh/soa/services"
	"github.com/gin-gonic/gin"
)

type StatisticsHandler struct {
	statisticService services.IStatistics
}

func NewStatistics(engine *gin.Engine, statisticService services.IStatistics) {
	handler := &StatisticsHandler{
		statisticService: statisticService,
	}

	// I have already changed path name as follows the problems
	group := engine.Group("api/statistics").Use(api.Validate())
	{
		group.GET("products-per-category", handler.GetStatisticsProductsPerCategory)
		group.GET("products-per-supplier", handler.GetStatisticsProductsPerSupplier)
	}
}

func (h *StatisticsHandler) GetStatisticsProductsPerCategory(ctx *gin.Context) {
	result, err := h.statisticService.GetStatisticsProductsPerCategory(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, result)
}

func (h *StatisticsHandler) GetStatisticsProductsPerSupplier(ctx *gin.Context) {
	result, err := h.statisticService.GetStatisticsProductsPerSupplier(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, result)
}
