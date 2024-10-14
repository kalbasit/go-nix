// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package eval_cache_v5

import (
	"context"
	"database/sql"
)

const insertAttribute = `-- name: InsertAttribute :exec
insert or replace into Attributes(parent, name, type, value) values (?, ?, ?, ?)
`

type InsertAttributeParams struct {
	Parent int64
	Name   sql.NullString
	Type   int64
	Value  sql.NullString
}

func (q *Queries) InsertAttribute(ctx context.Context, arg InsertAttributeParams) error {
	_, err := q.db.ExecContext(ctx, insertAttribute,
		arg.Parent,
		arg.Name,
		arg.Type,
		arg.Value,
	)
	return err
}

const insertAttributeWithContext = `-- name: InsertAttributeWithContext :exec
insert or replace into Attributes(parent, name, type, value, context) values (?, ?, ?, ?, ?)
`

type InsertAttributeWithContextParams struct {
	Parent  int64
	Name    sql.NullString
	Type    int64
	Value   sql.NullString
	Context sql.NullString
}

func (q *Queries) InsertAttributeWithContext(ctx context.Context, arg InsertAttributeWithContextParams) error {
	_, err := q.db.ExecContext(ctx, insertAttributeWithContext,
		arg.Parent,
		arg.Name,
		arg.Type,
		arg.Value,
		arg.Context,
	)
	return err
}

const queryAttribute = `-- name: QueryAttribute :one
select type, value, context from Attributes where parent = ? and name = ?
`

type QueryAttributeParams struct {
	Parent int64
	Name   sql.NullString
}

type QueryAttributeRow struct {
	Type    int64
	Value   sql.NullString
	Context sql.NullString
}

// todo sqlc doesn't like the rowid column being included below
func (q *Queries) QueryAttribute(ctx context.Context, arg QueryAttributeParams) (QueryAttributeRow, error) {
	row := q.db.QueryRowContext(ctx, queryAttribute, arg.Parent, arg.Name)
	var i QueryAttributeRow
	err := row.Scan(&i.Type, &i.Value, &i.Context)
	return i, err
}

const queryAttributes = `-- name: QueryAttributes :many
select name from Attributes where parent = ?
`

func (q *Queries) QueryAttributes(ctx context.Context, parent int64) ([]sql.NullString, error) {
	rows, err := q.db.QueryContext(ctx, queryAttributes, parent)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var name sql.NullString
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
