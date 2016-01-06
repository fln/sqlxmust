package sqlxmust

import (
	"github.com/jmoiron/sqlx"
)

// MustExecGetId executes query and returns the value of
// `result.LastInsertId()`. It panics if error is encountered at any stage.
func MustExecGetId(e sqlx.Execer, query string, args ...interface{}) int64 {
	result := sqlx.MustExec(e, query, args...)
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return id
}
