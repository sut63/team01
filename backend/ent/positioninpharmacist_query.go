// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team01/ent/pharmacist"
	"github.com/sut63/team01/ent/positioninpharmacist"
	"github.com/sut63/team01/ent/predicate"
)

// PositionInPharmacistQuery is the builder for querying PositionInPharmacist entities.
type PositionInPharmacistQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.PositionInPharmacist
	// eager-loading edges.
	withPharmacist *PharmacistQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (pipq *PositionInPharmacistQuery) Where(ps ...predicate.PositionInPharmacist) *PositionInPharmacistQuery {
	pipq.predicates = append(pipq.predicates, ps...)
	return pipq
}

// Limit adds a limit step to the query.
func (pipq *PositionInPharmacistQuery) Limit(limit int) *PositionInPharmacistQuery {
	pipq.limit = &limit
	return pipq
}

// Offset adds an offset step to the query.
func (pipq *PositionInPharmacistQuery) Offset(offset int) *PositionInPharmacistQuery {
	pipq.offset = &offset
	return pipq
}

// Order adds an order step to the query.
func (pipq *PositionInPharmacistQuery) Order(o ...OrderFunc) *PositionInPharmacistQuery {
	pipq.order = append(pipq.order, o...)
	return pipq
}

// QueryPharmacist chains the current query on the pharmacist edge.
func (pipq *PositionInPharmacistQuery) QueryPharmacist() *PharmacistQuery {
	query := &PharmacistQuery{config: pipq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pipq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(positioninpharmacist.Table, positioninpharmacist.FieldID, pipq.sqlQuery()),
			sqlgraph.To(pharmacist.Table, pharmacist.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, positioninpharmacist.PharmacistTable, positioninpharmacist.PharmacistColumn),
		)
		fromU = sqlgraph.SetNeighbors(pipq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PositionInPharmacist entity in the query. Returns *NotFoundError when no positioninpharmacist was found.
func (pipq *PositionInPharmacistQuery) First(ctx context.Context) (*PositionInPharmacist, error) {
	pips, err := pipq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(pips) == 0 {
		return nil, &NotFoundError{positioninpharmacist.Label}
	}
	return pips[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pipq *PositionInPharmacistQuery) FirstX(ctx context.Context) *PositionInPharmacist {
	pip, err := pipq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return pip
}

// FirstID returns the first PositionInPharmacist id in the query. Returns *NotFoundError when no id was found.
func (pipq *PositionInPharmacistQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pipq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{positioninpharmacist.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (pipq *PositionInPharmacistQuery) FirstXID(ctx context.Context) int {
	id, err := pipq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only PositionInPharmacist entity in the query, returns an error if not exactly one entity was returned.
func (pipq *PositionInPharmacistQuery) Only(ctx context.Context) (*PositionInPharmacist, error) {
	pips, err := pipq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(pips) {
	case 1:
		return pips[0], nil
	case 0:
		return nil, &NotFoundError{positioninpharmacist.Label}
	default:
		return nil, &NotSingularError{positioninpharmacist.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pipq *PositionInPharmacistQuery) OnlyX(ctx context.Context) *PositionInPharmacist {
	pip, err := pipq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return pip
}

// OnlyID returns the only PositionInPharmacist id in the query, returns an error if not exactly one id was returned.
func (pipq *PositionInPharmacistQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pipq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{positioninpharmacist.Label}
	default:
		err = &NotSingularError{positioninpharmacist.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pipq *PositionInPharmacistQuery) OnlyIDX(ctx context.Context) int {
	id, err := pipq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PositionInPharmacists.
func (pipq *PositionInPharmacistQuery) All(ctx context.Context) ([]*PositionInPharmacist, error) {
	if err := pipq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pipq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pipq *PositionInPharmacistQuery) AllX(ctx context.Context) []*PositionInPharmacist {
	pips, err := pipq.All(ctx)
	if err != nil {
		panic(err)
	}
	return pips
}

// IDs executes the query and returns a list of PositionInPharmacist ids.
func (pipq *PositionInPharmacistQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pipq.Select(positioninpharmacist.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pipq *PositionInPharmacistQuery) IDsX(ctx context.Context) []int {
	ids, err := pipq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pipq *PositionInPharmacistQuery) Count(ctx context.Context) (int, error) {
	if err := pipq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pipq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pipq *PositionInPharmacistQuery) CountX(ctx context.Context) int {
	count, err := pipq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pipq *PositionInPharmacistQuery) Exist(ctx context.Context) (bool, error) {
	if err := pipq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pipq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pipq *PositionInPharmacistQuery) ExistX(ctx context.Context) bool {
	exist, err := pipq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pipq *PositionInPharmacistQuery) Clone() *PositionInPharmacistQuery {
	return &PositionInPharmacistQuery{
		config:     pipq.config,
		limit:      pipq.limit,
		offset:     pipq.offset,
		order:      append([]OrderFunc{}, pipq.order...),
		unique:     append([]string{}, pipq.unique...),
		predicates: append([]predicate.PositionInPharmacist{}, pipq.predicates...),
		// clone intermediate query.
		sql:  pipq.sql.Clone(),
		path: pipq.path,
	}
}

//  WithPharmacist tells the query-builder to eager-loads the nodes that are connected to
// the "pharmacist" edge. The optional arguments used to configure the query builder of the edge.
func (pipq *PositionInPharmacistQuery) WithPharmacist(opts ...func(*PharmacistQuery)) *PositionInPharmacistQuery {
	query := &PharmacistQuery{config: pipq.config}
	for _, opt := range opts {
		opt(query)
	}
	pipq.withPharmacist = query
	return pipq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Position string `json:"position,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PositionInPharmacist.Query().
//		GroupBy(positioninpharmacist.FieldPosition).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pipq *PositionInPharmacistQuery) GroupBy(field string, fields ...string) *PositionInPharmacistGroupBy {
	group := &PositionInPharmacistGroupBy{config: pipq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pipq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pipq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Position string `json:"position,omitempty"`
//	}
//
//	client.PositionInPharmacist.Query().
//		Select(positioninpharmacist.FieldPosition).
//		Scan(ctx, &v)
//
func (pipq *PositionInPharmacistQuery) Select(field string, fields ...string) *PositionInPharmacistSelect {
	selector := &PositionInPharmacistSelect{config: pipq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pipq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pipq.sqlQuery(), nil
	}
	return selector
}

func (pipq *PositionInPharmacistQuery) prepareQuery(ctx context.Context) error {
	if pipq.path != nil {
		prev, err := pipq.path(ctx)
		if err != nil {
			return err
		}
		pipq.sql = prev
	}
	return nil
}

func (pipq *PositionInPharmacistQuery) sqlAll(ctx context.Context) ([]*PositionInPharmacist, error) {
	var (
		nodes       = []*PositionInPharmacist{}
		_spec       = pipq.querySpec()
		loadedTypes = [1]bool{
			pipq.withPharmacist != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &PositionInPharmacist{config: pipq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, pipq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pipq.withPharmacist; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*PositionInPharmacist)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Pharmacist(func(s *sql.Selector) {
			s.Where(sql.InValues(positioninpharmacist.PharmacistColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.positioninpharmacist_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "positioninpharmacist_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "positioninpharmacist_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Pharmacist = append(node.Edges.Pharmacist, n)
		}
	}

	return nodes, nil
}

func (pipq *PositionInPharmacistQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pipq.querySpec()
	return sqlgraph.CountNodes(ctx, pipq.driver, _spec)
}

func (pipq *PositionInPharmacistQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pipq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (pipq *PositionInPharmacistQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   positioninpharmacist.Table,
			Columns: positioninpharmacist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: positioninpharmacist.FieldID,
			},
		},
		From:   pipq.sql,
		Unique: true,
	}
	if ps := pipq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pipq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pipq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pipq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pipq *PositionInPharmacistQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(pipq.driver.Dialect())
	t1 := builder.Table(positioninpharmacist.Table)
	selector := builder.Select(t1.Columns(positioninpharmacist.Columns...)...).From(t1)
	if pipq.sql != nil {
		selector = pipq.sql
		selector.Select(selector.Columns(positioninpharmacist.Columns...)...)
	}
	for _, p := range pipq.predicates {
		p(selector)
	}
	for _, p := range pipq.order {
		p(selector)
	}
	if offset := pipq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pipq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PositionInPharmacistGroupBy is the builder for group-by PositionInPharmacist entities.
type PositionInPharmacistGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pipgb *PositionInPharmacistGroupBy) Aggregate(fns ...AggregateFunc) *PositionInPharmacistGroupBy {
	pipgb.fns = append(pipgb.fns, fns...)
	return pipgb
}

// Scan applies the group-by query and scan the result into the given value.
func (pipgb *PositionInPharmacistGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pipgb.path(ctx)
	if err != nil {
		return err
	}
	pipgb.sql = query
	return pipgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pipgb *PositionInPharmacistGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pipgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (pipgb *PositionInPharmacistGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pipgb.fields) > 1 {
		return nil, errors.New("ent: PositionInPharmacistGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pipgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pipgb *PositionInPharmacistGroupBy) StringsX(ctx context.Context) []string {
	v, err := pipgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (pipgb *PositionInPharmacistGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pipgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{positioninpharmacist.Label}
	default:
		err = fmt.Errorf("ent: PositionInPharmacistGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pipgb *PositionInPharmacistGroupBy) StringX(ctx context.Context) string {
	v, err := pipgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (pipgb *PositionInPharmacistGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pipgb.fields) > 1 {
		return nil, errors.New("ent: PositionInPharmacistGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pipgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pipgb *PositionInPharmacistGroupBy) IntsX(ctx context.Context) []int {
	v, err := pipgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (pipgb *PositionInPharmacistGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pipgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{positioninpharmacist.Label}
	default:
		err = fmt.Errorf("ent: PositionInPharmacistGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pipgb *PositionInPharmacistGroupBy) IntX(ctx context.Context) int {
	v, err := pipgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (pipgb *PositionInPharmacistGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pipgb.fields) > 1 {
		return nil, errors.New("ent: PositionInPharmacistGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pipgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pipgb *PositionInPharmacistGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pipgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (pipgb *PositionInPharmacistGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pipgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{positioninpharmacist.Label}
	default:
		err = fmt.Errorf("ent: PositionInPharmacistGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pipgb *PositionInPharmacistGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pipgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (pipgb *PositionInPharmacistGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pipgb.fields) > 1 {
		return nil, errors.New("ent: PositionInPharmacistGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pipgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pipgb *PositionInPharmacistGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pipgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (pipgb *PositionInPharmacistGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pipgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{positioninpharmacist.Label}
	default:
		err = fmt.Errorf("ent: PositionInPharmacistGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pipgb *PositionInPharmacistGroupBy) BoolX(ctx context.Context) bool {
	v, err := pipgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pipgb *PositionInPharmacistGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pipgb.sqlQuery().Query()
	if err := pipgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pipgb *PositionInPharmacistGroupBy) sqlQuery() *sql.Selector {
	selector := pipgb.sql
	columns := make([]string, 0, len(pipgb.fields)+len(pipgb.fns))
	columns = append(columns, pipgb.fields...)
	for _, fn := range pipgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(pipgb.fields...)
}

// PositionInPharmacistSelect is the builder for select fields of PositionInPharmacist entities.
type PositionInPharmacistSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (pips *PositionInPharmacistSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := pips.path(ctx)
	if err != nil {
		return err
	}
	pips.sql = query
	return pips.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pips *PositionInPharmacistSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pips.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (pips *PositionInPharmacistSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pips.fields) > 1 {
		return nil, errors.New("ent: PositionInPharmacistSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pips.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pips *PositionInPharmacistSelect) StringsX(ctx context.Context) []string {
	v, err := pips.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (pips *PositionInPharmacistSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pips.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{positioninpharmacist.Label}
	default:
		err = fmt.Errorf("ent: PositionInPharmacistSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pips *PositionInPharmacistSelect) StringX(ctx context.Context) string {
	v, err := pips.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (pips *PositionInPharmacistSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pips.fields) > 1 {
		return nil, errors.New("ent: PositionInPharmacistSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pips.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pips *PositionInPharmacistSelect) IntsX(ctx context.Context) []int {
	v, err := pips.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (pips *PositionInPharmacistSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pips.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{positioninpharmacist.Label}
	default:
		err = fmt.Errorf("ent: PositionInPharmacistSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pips *PositionInPharmacistSelect) IntX(ctx context.Context) int {
	v, err := pips.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (pips *PositionInPharmacistSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pips.fields) > 1 {
		return nil, errors.New("ent: PositionInPharmacistSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pips.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pips *PositionInPharmacistSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pips.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (pips *PositionInPharmacistSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pips.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{positioninpharmacist.Label}
	default:
		err = fmt.Errorf("ent: PositionInPharmacistSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pips *PositionInPharmacistSelect) Float64X(ctx context.Context) float64 {
	v, err := pips.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (pips *PositionInPharmacistSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pips.fields) > 1 {
		return nil, errors.New("ent: PositionInPharmacistSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pips.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pips *PositionInPharmacistSelect) BoolsX(ctx context.Context) []bool {
	v, err := pips.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (pips *PositionInPharmacistSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pips.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{positioninpharmacist.Label}
	default:
		err = fmt.Errorf("ent: PositionInPharmacistSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pips *PositionInPharmacistSelect) BoolX(ctx context.Context) bool {
	v, err := pips.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pips *PositionInPharmacistSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pips.sqlQuery().Query()
	if err := pips.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pips *PositionInPharmacistSelect) sqlQuery() sql.Querier {
	selector := pips.sql
	selector.Select(selector.Columns(pips.fields...)...)
	return selector
}
