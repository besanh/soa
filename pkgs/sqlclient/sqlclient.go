package sqlclient

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type ISqlClientConn interface {
	GetDB() *bun.DB
}

type SqlConfig struct {
	Host         string
	Port         int
	Database     string
	Username     string
	Password     string
	Timeout      int
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
	PoolSize     int
	MaxIdleConns int
	MaxOpenConns int
}

type SqlClientConn struct {
	SqlConfig
	DB *bun.DB
}

func NewSqlClient(config SqlConfig) ISqlClientConn {
	client := &SqlClientConn{}
	client.SqlConfig = config
	if err := client.Connect(); err != nil {
		log.Fatal(err)
		return nil
	}
	if err := client.DB.Ping(); err != nil {
		log.Fatal(err)
		return nil
	}
	return client
}

func (c *SqlClientConn) Connect() (err error) {
	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(fmt.Sprintf("%s:%d", c.Host, c.Port)),
		pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		pgdriver.WithUser(c.Username),
		pgdriver.WithPassword(c.Password),
		pgdriver.WithDatabase(c.Database),
		pgdriver.WithTimeout(time.Duration(c.Timeout)*time.Second),
		pgdriver.WithDialTimeout(time.Duration(c.DialTimeout)*time.Second),
		pgdriver.WithReadTimeout(time.Duration(c.ReadTimeout)*time.Second),
		pgdriver.WithWriteTimeout(time.Duration(c.WriteTimeout)*time.Second),
		pgdriver.WithInsecure(true),
	)
	sqldb := sql.OpenDB(pgconn)
	sqldb.SetMaxIdleConns(c.MaxIdleConns)
	sqldb.SetMaxOpenConns(c.MaxOpenConns)
	db := bun.NewDB(sqldb, pgdialect.New(), bun.WithDiscardUnknownColumns())
	c.DB = db
	return nil
}

func (c *SqlClientConn) GetDB() *bun.DB {
	return c.DB
}
