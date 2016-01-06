package sqlxmust

import (
	"database/sql"
	"errors"
	"reflect"

	"github.com/jmoiron/sqlx"
)

// GetAlloc is the same as `slqx.Get` except it must be given pointer to a
// pointer to destination data structure. If query returns no rows it sets
// pointer to nil. Otherwise it allocates new value and performs normal
// `sqlx.Get` on it.
func GetAlloc(q sqlx.Queryer, destpp interface{}, query string, args ...interface{}) error {
	t := reflect.TypeOf(destpp)

	if t.Kind() != reflect.Ptr {
		return errors.New("must pass a double pointer to destination value, not a value")
	}

	if t.Elem().Kind() != reflect.Ptr {
		return errors.New("must pass a double pointer to destination value, not a pointer to value")
	}

	destp := reflect.New(t.Elem().Elem())
	err := sqlx.Get(q, destp.Interface(), query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			reflect.ValueOf(destpp).Elem().Set(reflect.Zero(t.Elem()))
			return nil
		}
		return err
	}
	reflect.ValueOf(destpp).Elem().Set(destp)
	return nil
}

// MustGetAlloc is same as `GetAlloc` but panics on error.
func MustGetAlloc(q sqlx.Queryer, destpp interface{}, query string, args ...interface{}) {
	err := GetAlloc(q, destpp, query, args...)
	if err != nil {
		panic(err)
	}
}

// MustGet is the same as `sqlx.Get` but panics on error.
func MustGet(q sqlx.Queryer, dest interface{}, query string, args ...interface{}) {
	err := sqlx.Get(q, dest, query, args...)
	if err != nil {
		panic(err)
	}
}

// MustSelect is the same as `sqlx.Select` but panics on error.
func MustSelect(q sqlx.Queryer, dest interface{}, query string, args ...interface{}) {
	err := sqlx.Select(q, dest, query, args...)
	if err != nil {
		panic(err)
	}
}
