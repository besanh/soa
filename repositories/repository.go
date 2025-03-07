package repositories

import (
	"context"
	"strings"

	"github.com/besanh/soa/pkgs/sqlclient"
	"github.com/uptrace/bun/schema"
)

var PgSqlClient sqlclient.ISqlClientConn

func CreateTable(client sqlclient.ISqlClientConn, ctx context.Context, table any) error {
	query := client.GetDB().NewCreateTable().Model(table).IfNotExists()
	value, _ := query.AppendQuery(schema.NewFormatter(query.Dialect()), nil)
	queryStr := string(value)
	queryStr = strings.ReplaceAll(queryStr, " char(36)", " uuid")
	queryStr = strings.ReplaceAll(queryStr, " timestamp", " timestamptz")
	queryStr = strings.ReplaceAll(queryStr, " timestamptz_only", " timestamp")
	_, err := client.GetDB().QueryContext(ctx, queryStr)
	return err
}
