package ogx

import (
	"context"
	"database/sql"
	"strings"
	"sync/atomic"
	"time"
)

type QueryEvent struct {
	DB *DB

	IQuery        Query
	Query         string
	QueryTemplate string
	QueryArgs     []interface{}
	Model         Model

	StartTime time.Time
	Result    sql.Result
	Err       error

	Stash map[interface{}]interface{}
}

func (e *QueryEvent) Operation() string {
	if e.IQuery != nil {
		return e.IQuery.Operation()
	}
	return queryOperation(e.Query)
}

func queryOperation(query string) string {
	if idx := strings.IndexByte(query, ' '); idx > 0 {
		query = query[:idx]
	}
	if len(query) > 16 {
		query = query[:16]
	}
	return query
}

type QueryHook interface {
	BeforeQuery(context.Context, *QueryEvent) context.Context
	AfterQuery(context.Context, *QueryEvent)
}

func (db *DB) beforeQuery(
	ctx context.Context,
	iquery Query,
	queryTemplate string,
	queryArgs []interface{},
	query string,
	model Model,
) (context.Context, *QueryEvent) {
	atomic.AddUint32(&db.stats.Queries, 1)

	if len(db.queryHooks) == 0 {
		return ctx, nil
	}

	event := &QueryEvent{
		DB: db,

		Model:         model,
		IQuery:        iquery,
		Query:         query,
		QueryTemplate: queryTemplate,
		QueryArgs:     queryArgs,

		StartTime: time.Now(),
	}

	for _, hook := range db.queryHooks {
		ctx = hook.BeforeQuery(ctx, event)
	}

	return ctx, event
}

func (db *DB) afterQuery(
	ctx context.Context,
	event *QueryEvent,
	res sql.Result,
	err error,
) {
	switch err {
	case nil, sql.ErrNoRows:
		// nothing
	default:
		atomic.AddUint32(&db.stats.Errors, 1)
	}

	if event == nil {
		return
	}

	event.Result = res
	event.Err = err

	db.afterQueryFromIndex(ctx, event, len(db.queryHooks)-1)
}

func (db *DB) afterQueryFromIndex(ctx context.Context, event *QueryEvent, hookIndex int) {
	for ; hookIndex >= 0; hookIndex-- {
		db.queryHooks[hookIndex].AfterQuery(ctx, event)
	}
}
