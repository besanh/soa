package v1

import (
	api "github.com/besanh/soa/apis"
	"github.com/besanh/soa/common/util"
	"github.com/besanh/soa/models"
	"github.com/besanh/soa/services"
	"github.com/gin-gonic/gin"
)

type ProductCategoiesHandler struct {
	productCategoriesService services.IProductCategories
}

func NewProductCategories(engine *gin.Engine, productCategoriesService services.IProductCategories) {
	handler := &ProductCategoiesHandler{
		productCategoriesService: productCategoriesService,
	}

	group := engine.Group("v1/product-categories").Use(api.Validate())
	{
		group.POST("", handler.Insert)
		group.PUT(":id", handler.Update)
		group.DELETE(":id", handler.Delete)
		group.GET("", handler.Select)
	}
}

func (h *ProductCategoiesHandler) Insert(ctx *gin.Context) {
	var data models.ProductCategoriesRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := data.Validate(); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.productCategoriesService.Insert(ctx, &data); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, data)
}

func (h *ProductCategoiesHandler) Update(ctx *gin.Context) {
	var data models.ProductCategoriesRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := data.Validate(); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.productCategoriesService.Update(ctx, ctx.Param("id"), &data); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, data)
}

func (h *ProductCategoiesHandler) Delete(ctx *gin.Context) {
	if err := h.productCategoriesService.Delete(ctx, ctx.Param("id")); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "delete product category successfully"})
}

func (h *ProductCategoiesHandler) Select(ctx *gin.Context) {
	limit, offset := util.GetLimitOffset(ctx.Query("limit"), ctx.Query("offset"))

	query := &models.ProductCategoriesQuery{
		ProductCategoryName: ctx.Query("product_category_name"),
		Status:              ctx.Query("status"),
		Limit:               limit,
		Offset:              offset,
	}

	total, result, err := h.productCategoriesService.Select(ctx, query)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"total": total,
		"data":  result,
	})
}
