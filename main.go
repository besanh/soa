package main

import (
	v1 "github.com/besanh/soa/apis/v1"
	"github.com/besanh/soa/common/env"
	"github.com/besanh/soa/pkgs/sqlclient"
	"github.com/besanh/soa/repositories"
	"github.com/besanh/soa/servers"
	"github.com/besanh/soa/services"
	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	LogLevel string
	LogFile  string
}

var config Config

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	config = Config{
		Port:     env.GetStringENV("PORT", "8000"),
		LogLevel: env.GetStringENV("LOG_LEVEL", "debug"),
		LogFile:  env.GetStringENV("LOG_FILE", "tmp/console.log"),
	}

	sqlClientConfig := sqlclient.SqlConfig{
		Host:         env.GetStringENV("PG_HOST", "localhost"),
		Database:     env.GetStringENV("PG_DATABASE", "soa"),
		Username:     env.GetStringENV("PG_USERNAME", "soa"),
		Password:     env.GetStringENV("PG_PASSWORD", "soa"),
		Port:         env.GetIntENV("PG_PORT", 5432),
		DialTimeout:  20,
		ReadTimeout:  30,
		WriteTimeout: 30,
		Timeout:      30,
		PoolSize:     10,
		MaxIdleConns: 10,
		MaxOpenConns: 10,
	}

	initRepository(sqlClientConfig)
}

/*
 * author: AnhLe
 */
func main() {
	server := servers.NewServer()

	productsCategoriesService := services.NewProductCategories()
	v1.NewProductCategories(server.Engine, productsCategoriesService)
	services.SECRET_KEY = env.GetStringENV("SECRET_KEY", "")

	server.Start(config.Port)
}

func initRepository(sqlClientConfig sqlclient.SqlConfig) {
	repositories.PgSqlClient = sqlclient.NewSqlClient(sqlClientConfig)

	repositories.ProductRepo = repositories.NewProducts()
	repositories.ProductCategoryRepo = repositories.NewProductCategories()
	repositories.SupplierRepo = repositories.NewSuppliers()
}
