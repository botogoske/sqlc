// Code generated by sqlc. DO NOT EDIT.

package querytest

import (
	"database/sql"
)

type Foo struct {
	Bar sql.NullString
	Bat string
	Baz sql.NullInt64
	Qux int64
}