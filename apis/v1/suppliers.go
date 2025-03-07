package v1

import (
	api "github.com/besanh/soa/apis"
	"github.com/besanh/soa/common/util"
	"github.com/besanh/soa/models"
	"github.com/besanh/soa/services"
	"github.com/gin-gonic/gin"
)

type SuppliersHandler struct {
	suppliersService services.ISuppliers
}

func NewSuppliers(engine *gin.Engine, suppliersService services.ISuppliers) {
	handler := &SuppliersHandler{
		suppliersService: suppliersService,
	}

	group := engine.Group("v1/suppliers").Use(api.Validate())
	{
		group.POST("", handler.Insert)
		group.PUT(":id", handler.Update)
		group.DELETE(":id", handler.Delete)
		group.GET("", handler.Select)
	}
}

func (h *SuppliersHandler) Insert(ctx *gin.Context) {
	var data models.SuppliersRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := data.Validate(); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.suppliersService.Insert(ctx, &data); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, data)
}

func (h *SuppliersHandler) Update(ctx *gin.Context) {
	var data models.SuppliersRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := data.Validate(); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.suppliersService.Update(ctx, ctx.Param("id"), &data); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, data)
}

func (h *SuppliersHandler) Delete(ctx *gin.Context) {
	if err := h.suppliersService.Delete(ctx, ctx.Param("id")); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "delete supplier successfully"})
}

func (h *SuppliersHandler) Select(ctx *gin.Context) {
	limit, offset := util.GetLimitOffset(ctx.Query("limit"), ctx.Query("offset"))

	query := &models.SuppliersQuery{
		SupplierName: ctx.Query("supplier_name"),
		Status:       ctx.Query("status"),
		Limit:        limit,
		Offset:       offset,
	}

	total, result, err := h.suppliersService.Select(ctx, query)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"total": total,
		"data":  result,
	})
}
