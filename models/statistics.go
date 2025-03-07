package models

type (
	ProductsPerCategoryStat struct {
		Category string  `bun:"category" json:"category"`
		Count    int64   `bun:"count" json:"count"`
		Percent  float64 `json:"percent"`
	}

	ProductsPerSupplierStat struct {
		Supplier string  `bun:"supplier" json:"supplier"`
		Count    int64   `bun:"count" json:"count"`
		Percent  float64 `json:"percent"`
	}
)
