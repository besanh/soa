package v1

import (
	api "github.com/besanh/soa/apis"
	"github.com/besanh/soa/common/util"
	"github.com/besanh/soa/models"
	"github.com/besanh/soa/services"
	"github.com/gin-gonic/gin"
)

type ProductsHandler struct {
	productsService services.IProducts
}

func NewProduct(engine *gin.Engine, productsService services.IProducts) {
	handler := &ProductsHandler{
		productsService: productsService,
	}

	group := engine.Group("v1/products").Use(api.Validate())
	{
		group.POST("", handler.Insert)
		group.PUT(":id", handler.Update)
		group.DELETE(":id", handler.Delete)
		group.GET("", handler.Select)
		group.GET("scroll", handler.SelectScroll)
	}
}

func (h *ProductsHandler) Insert(ctx *gin.Context) {
	var data models.ProductsRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := data.Validate(); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.productsService.Insert(ctx, &data); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, data)
}

func (h *ProductsHandler) Update(ctx *gin.Context) {
	var data models.ProductsRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := data.Validate(); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.productsService.Update(ctx, ctx.Param("id"), &data); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, data)
}

func (h *ProductsHandler) Delete(ctx *gin.Context) {
	if err := h.productsService.Delete(ctx, ctx.Param("id")); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "delete product successfully"})
}

func (h *ProductsHandler) Select(ctx *gin.Context) {
	limit, offset := util.GetLimitOffset(ctx.Query("limit"), ctx.Query("offset"))

	query := &models.ProductsQuery{
		ProductName:       ctx.Query("product_name"),
		ProductReference:  ctx.Query("product_reference"),
		Status:            util.ParseQueryArray(ctx.QueryArray("status")),
		ProductCategoryId: util.ParseQueryArray(ctx.QueryArray("product_category_id")),
		SupplierId:        util.ParseQueryArray(ctx.QueryArray("supplier_id")),
		FromDateCreated:   ctx.Query("from_date_created"),
		ToDateCreated:     ctx.Query("to_date_created"),
		FromPrice:         ctx.Query("from_price"),
		ToPrice:           ctx.Query("to_price"),
		FromQuantity:      ctx.Query("from_quantity"),
		ToQuantity:        ctx.Query("to_quantity"),
		Limit:             limit,
		Offset:            offset,
	}

	total, result, err := h.productsService.Select(ctx, query)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"total": total,
		"data":  result,
	})
}

func (h *ProductsHandler) SelectScroll(ctx *gin.Context) {
	limit, offset := util.GetLimitOffset(ctx.Query("limit"), ctx.Query("offset"))

	query := &models.ProductsQuery{
		LastSeenId:        ctx.Query("last_seen_id"),
		CreatedAt:         ctx.Query("created_at"),
		ProductName:       ctx.Query("product_name"),
		ProductReference:  ctx.Query("product_reference"),
		Status:            util.ParseQueryArray(ctx.QueryArray("status")),
		ProductCategoryId: util.ParseQueryArray(ctx.QueryArray("product_category_id")),
		SupplierId:        util.ParseQueryArray(ctx.QueryArray("supplier_id")),
		FromDateCreated:   ctx.Query("from_date_created"),
		ToDateCreated:     ctx.Query("to_date_created"),
		FromPrice:         ctx.Query("from_price"),
		ToPrice:           ctx.Query("to_price"),
		FromQuantity:      ctx.Query("from_quantity"),
		ToQuantity:        ctx.Query("to_quantity"),
		Limit:             limit,
		Offset:            offset,
	}

	total, result, err := h.productsService.Select(ctx, query)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"total": total,
		"data":  result,
	})
}
