sqlxmust
========

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/fln/sqlxmust)

This package contains a set of helper functions for `github.com/jmoiron/sqlx`
package.

Functions in this package can be grouped into three categories:
* `...GetAlloc...` - functions similar to `sqlx.Get` except it performs memory
allocation for new data or returns `nil` if no rows were returned. For these
functions `dest` argument must be a double pointer.
* `...Must...` - converts database errors to panics. It is useful if database
communication errors are very rare.
* `...OrRollback` - calls tx.Rollback() if error occurs.
