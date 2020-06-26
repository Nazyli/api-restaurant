package dbdialect

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

// SeQuery
// MySQL struct
type MySQL struct {
	db *sqlx.DB
}

// New init mysql
func New(db *sqlx.DB) *MySQL {
	return &MySQL{db}
}

// SetQuery . . .
func (m *MySQL) SetQuery(query string) (newQuery string) {
	switch m.db.DriverName() {
	case "postgres":
		count := strings.Count(query, "?")
		for i := 1; i <= count; i++ {
			query = strings.Replace(query, "?", fmt.Sprintf("$%d", i), 1)
		}
		return query
	default:
		return query
	}
}
