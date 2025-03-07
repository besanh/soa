package v1

import (
	"net/http"

	api "github.com/besanh/soa/apis"
	"github.com/besanh/soa/services"
	"github.com/gin-gonic/gin"
)

type DistanceHandler struct {
	distanceService services.IDistance
}

func NewDistance(engine *gin.Engine, distanceService services.IDistance) {
	handler := &DistanceHandler{
		distanceService: distanceService,
	}

	group := engine.Group("v1/distance").Use(api.Validate())
	{
		group.GET("", handler.CalculateDistance)
	}
}

func (h *DistanceHandler) CalculateDistance(ctx *gin.Context) {
	ip := ctx.Query("ip")
	city := ctx.Query("city")
	if ip == "" || city == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing ip or city parameter"})
		return
	}

	result, err := h.distanceService.CalculateDistance(ip, city)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
