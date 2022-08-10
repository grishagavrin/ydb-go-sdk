package xsql

import (
	"context"
	"database/sql/driver"
	"fmt"

	"github.com/ydb-platform/ydb-go-genproto/protos/Ydb"
)

type stmt struct {
	nopResult
	namedValueChecker

	conn   *conn
	params map[string]*Ydb.Type
	query  string
}

var (
	_ driver.Stmt             = &stmt{}
	_ driver.StmtQueryContext = &stmt{}
	_ driver.StmtExecContext  = &stmt{}

	_ driver.NamedValueChecker = &stmt{}
)

func (s *stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	if s.conn.isClosed() {
		return nil, errClosedConn
	}
	switch m := queryModeFromContext(withKeepInCache(ctx), s.conn.defaultQueryMode); m {
	case DataQueryMode:
		return s.conn.QueryContext(ctx, s.query, args)
	default:
		return nil, fmt.Errorf("unsupported query mode '%s' for execute query on prepared statement", m)
	}
}

func (s *stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	if s.conn.isClosed() {
		return nil, errClosedConn
	}
	switch m := queryModeFromContext(withKeepInCache(ctx), s.conn.defaultQueryMode); m {
	case DataQueryMode:
		return s.conn.ExecContext(ctx, s.query, args)
	default:
		return nil, fmt.Errorf("unsupported query mode '%s' for execute query on prepared statement", m)
	}
}

func (s *stmt) NumInput() int {
	return len(s.params)
}

func (s *stmt) Close() error {
	return nil
}

func (s stmt) Exec(args []driver.Value) (driver.Result, error) {
	return nil, errDeprecated
}

func (s stmt) Query(args []driver.Value) (driver.Rows, error) {
	return nil, errDeprecated
}