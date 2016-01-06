package sqlxmust

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// TxMustGetAllocOrRollback is the same as `sqlxmust.MustGetAlloc` exept it
// calls tx.Rollback() on error before panicking.
func TxMustGetAllocOrRollback(tx *sqlx.Tx, destpp interface{}, query string, args ...interface{}) {
	err := GetAlloc(tx, destpp, query, args...)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
}

// TxMustGetOrRollback is the same as `sqlxmust.MustGet` except it rolls-back
// the transaction on error before panicking.
func TxMustGetOrRollback(tx *sqlx.Tx, dest interface{}, query string, args ...interface{}) {
	err := tx.Get(dest, query, args...)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
}

// TxMustExecGetIdOrRollback is the same as `sqlxmust.MustExecGetId` except it
// rolls-back the transaction before panicking.
func TxMustExecGetIdOrRollback(tx *sqlx.Tx, query string, args ...interface{}) int64 {
	result := TxMustExecOrRollback(tx, query, args...)
	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	return id
}

// TxMustExecOrRollback it the same as `sqlx.Exec` except it rolls-back the
// transaction before panicking.
func TxMustExecOrRollback(tx *sqlx.Tx, query string, args ...interface{}) sql.Result {
	res, err := tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	return res
}

// TxMustCommit is the same as `sql.Commit` except it panics on error.
func TxMustCommit(tx *sqlx.Tx) {
	err := tx.Commit()
	if err != nil {
		panic(err)
	}
}
