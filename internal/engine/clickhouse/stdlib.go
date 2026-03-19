package clickhouse

import (
	"github.com/thedemons/sqlc/internal/sql/catalog"
)

func defaultSchema(name string) *catalog.Schema {
	return &catalog.Schema{Name: name}
}
