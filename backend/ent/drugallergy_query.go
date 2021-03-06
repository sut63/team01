// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team01/ent/drugallergy"
	"github.com/sut63/team01/ent/medicine"
	"github.com/sut63/team01/ent/patientinfo"
	"github.com/sut63/team01/ent/pharmacist"
	"github.com/sut63/team01/ent/predicate"
)

// DrugAllergyQuery is the builder for querying DrugAllergy entities.
type DrugAllergyQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.DrugAllergy
	// eager-loading edges.
	withPatient    *PatientInfoQuery
	withMedicine   *MedicineQuery
	withPharmacist *PharmacistQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (daq *DrugAllergyQuery) Where(ps ...predicate.DrugAllergy) *DrugAllergyQuery {
	daq.predicates = append(daq.predicates, ps...)
	return daq
}

// Limit adds a limit step to the query.
func (daq *DrugAllergyQuery) Limit(limit int) *DrugAllergyQuery {
	daq.limit = &limit
	return daq
}

// Offset adds an offset step to the query.
func (daq *DrugAllergyQuery) Offset(offset int) *DrugAllergyQuery {
	daq.offset = &offset
	return daq
}

// Order adds an order step to the query.
func (daq *DrugAllergyQuery) Order(o ...OrderFunc) *DrugAllergyQuery {
	daq.order = append(daq.order, o...)
	return daq
}

// QueryPatient chains the current query on the patient edge.
func (daq *DrugAllergyQuery) QueryPatient() *PatientInfoQuery {
	query := &PatientInfoQuery{config: daq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := daq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(drugallergy.Table, drugallergy.FieldID, daq.sqlQuery()),
			sqlgraph.To(patientinfo.Table, patientinfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, drugallergy.PatientTable, drugallergy.PatientColumn),
		)
		fromU = sqlgraph.SetNeighbors(daq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMedicine chains the current query on the medicine edge.
func (daq *DrugAllergyQuery) QueryMedicine() *MedicineQuery {
	query := &MedicineQuery{config: daq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := daq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(drugallergy.Table, drugallergy.FieldID, daq.sqlQuery()),
			sqlgraph.To(medicine.Table, medicine.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, drugallergy.MedicineTable, drugallergy.MedicineColumn),
		)
		fromU = sqlgraph.SetNeighbors(daq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPharmacist chains the current query on the pharmacist edge.
func (daq *DrugAllergyQuery) QueryPharmacist() *PharmacistQuery {
	query := &PharmacistQuery{config: daq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := daq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(drugallergy.Table, drugallergy.FieldID, daq.sqlQuery()),
			sqlgraph.To(pharmacist.Table, pharmacist.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, drugallergy.PharmacistTable, drugallergy.PharmacistColumn),
		)
		fromU = sqlgraph.SetNeighbors(daq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DrugAllergy entity in the query. Returns *NotFoundError when no drugallergy was found.
func (daq *DrugAllergyQuery) First(ctx context.Context) (*DrugAllergy, error) {
	das, err := daq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(das) == 0 {
		return nil, &NotFoundError{drugallergy.Label}
	}
	return das[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (daq *DrugAllergyQuery) FirstX(ctx context.Context) *DrugAllergy {
	da, err := daq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return da
}

// FirstID returns the first DrugAllergy id in the query. Returns *NotFoundError when no id was found.
func (daq *DrugAllergyQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = daq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{drugallergy.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (daq *DrugAllergyQuery) FirstXID(ctx context.Context) int {
	id, err := daq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only DrugAllergy entity in the query, returns an error if not exactly one entity was returned.
func (daq *DrugAllergyQuery) Only(ctx context.Context) (*DrugAllergy, error) {
	das, err := daq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(das) {
	case 1:
		return das[0], nil
	case 0:
		return nil, &NotFoundError{drugallergy.Label}
	default:
		return nil, &NotSingularError{drugallergy.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (daq *DrugAllergyQuery) OnlyX(ctx context.Context) *DrugAllergy {
	da, err := daq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return da
}

// OnlyID returns the only DrugAllergy id in the query, returns an error if not exactly one id was returned.
func (daq *DrugAllergyQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = daq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{drugallergy.Label}
	default:
		err = &NotSingularError{drugallergy.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (daq *DrugAllergyQuery) OnlyIDX(ctx context.Context) int {
	id, err := daq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DrugAllergies.
func (daq *DrugAllergyQuery) All(ctx context.Context) ([]*DrugAllergy, error) {
	if err := daq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return daq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (daq *DrugAllergyQuery) AllX(ctx context.Context) []*DrugAllergy {
	das, err := daq.All(ctx)
	if err != nil {
		panic(err)
	}
	return das
}

// IDs executes the query and returns a list of DrugAllergy ids.
func (daq *DrugAllergyQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := daq.Select(drugallergy.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (daq *DrugAllergyQuery) IDsX(ctx context.Context) []int {
	ids, err := daq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (daq *DrugAllergyQuery) Count(ctx context.Context) (int, error) {
	if err := daq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return daq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (daq *DrugAllergyQuery) CountX(ctx context.Context) int {
	count, err := daq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (daq *DrugAllergyQuery) Exist(ctx context.Context) (bool, error) {
	if err := daq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return daq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (daq *DrugAllergyQuery) ExistX(ctx context.Context) bool {
	exist, err := daq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (daq *DrugAllergyQuery) Clone() *DrugAllergyQuery {
	return &DrugAllergyQuery{
		config:     daq.config,
		limit:      daq.limit,
		offset:     daq.offset,
		order:      append([]OrderFunc{}, daq.order...),
		unique:     append([]string{}, daq.unique...),
		predicates: append([]predicate.DrugAllergy{}, daq.predicates...),
		// clone intermediate query.
		sql:  daq.sql.Clone(),
		path: daq.path,
	}
}

//  WithPatient tells the query-builder to eager-loads the nodes that are connected to
// the "patient" edge. The optional arguments used to configure the query builder of the edge.
func (daq *DrugAllergyQuery) WithPatient(opts ...func(*PatientInfoQuery)) *DrugAllergyQuery {
	query := &PatientInfoQuery{config: daq.config}
	for _, opt := range opts {
		opt(query)
	}
	daq.withPatient = query
	return daq
}

//  WithMedicine tells the query-builder to eager-loads the nodes that are connected to
// the "medicine" edge. The optional arguments used to configure the query builder of the edge.
func (daq *DrugAllergyQuery) WithMedicine(opts ...func(*MedicineQuery)) *DrugAllergyQuery {
	query := &MedicineQuery{config: daq.config}
	for _, opt := range opts {
		opt(query)
	}
	daq.withMedicine = query
	return daq
}

//  WithPharmacist tells the query-builder to eager-loads the nodes that are connected to
// the "pharmacist" edge. The optional arguments used to configure the query builder of the edge.
func (daq *DrugAllergyQuery) WithPharmacist(opts ...func(*PharmacistQuery)) *DrugAllergyQuery {
	query := &PharmacistQuery{config: daq.config}
	for _, opt := range opts {
		opt(query)
	}
	daq.withPharmacist = query
	return daq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Symptom string `json:"symptom,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DrugAllergy.Query().
//		GroupBy(drugallergy.FieldSymptom).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (daq *DrugAllergyQuery) GroupBy(field string, fields ...string) *DrugAllergyGroupBy {
	group := &DrugAllergyGroupBy{config: daq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := daq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return daq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Symptom string `json:"symptom,omitempty"`
//	}
//
//	client.DrugAllergy.Query().
//		Select(drugallergy.FieldSymptom).
//		Scan(ctx, &v)
//
func (daq *DrugAllergyQuery) Select(field string, fields ...string) *DrugAllergySelect {
	selector := &DrugAllergySelect{config: daq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := daq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return daq.sqlQuery(), nil
	}
	return selector
}

func (daq *DrugAllergyQuery) prepareQuery(ctx context.Context) error {
	if daq.path != nil {
		prev, err := daq.path(ctx)
		if err != nil {
			return err
		}
		daq.sql = prev
	}
	return nil
}

func (daq *DrugAllergyQuery) sqlAll(ctx context.Context) ([]*DrugAllergy, error) {
	var (
		nodes       = []*DrugAllergy{}
		withFKs     = daq.withFKs
		_spec       = daq.querySpec()
		loadedTypes = [3]bool{
			daq.withPatient != nil,
			daq.withMedicine != nil,
			daq.withPharmacist != nil,
		}
	)
	if daq.withPatient != nil || daq.withMedicine != nil || daq.withPharmacist != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, drugallergy.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &DrugAllergy{config: daq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
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
	if err := sqlgraph.QueryNodes(ctx, daq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := daq.withPatient; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*DrugAllergy)
		for i := range nodes {
			if fk := nodes[i].patient_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(patientinfo.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "patient_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Patient = n
			}
		}
	}

	if query := daq.withMedicine; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*DrugAllergy)
		for i := range nodes {
			if fk := nodes[i].medicine_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(medicine.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "medicine_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Medicine = n
			}
		}
	}

	if query := daq.withPharmacist; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*DrugAllergy)
		for i := range nodes {
			if fk := nodes[i].pharmacist_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(pharmacist.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "pharmacist_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Pharmacist = n
			}
		}
	}

	return nodes, nil
}

func (daq *DrugAllergyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := daq.querySpec()
	return sqlgraph.CountNodes(ctx, daq.driver, _spec)
}

func (daq *DrugAllergyQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := daq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (daq *DrugAllergyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   drugallergy.Table,
			Columns: drugallergy.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: drugallergy.FieldID,
			},
		},
		From:   daq.sql,
		Unique: true,
	}
	if ps := daq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := daq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := daq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := daq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (daq *DrugAllergyQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(daq.driver.Dialect())
	t1 := builder.Table(drugallergy.Table)
	selector := builder.Select(t1.Columns(drugallergy.Columns...)...).From(t1)
	if daq.sql != nil {
		selector = daq.sql
		selector.Select(selector.Columns(drugallergy.Columns...)...)
	}
	for _, p := range daq.predicates {
		p(selector)
	}
	for _, p := range daq.order {
		p(selector)
	}
	if offset := daq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := daq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DrugAllergyGroupBy is the builder for group-by DrugAllergy entities.
type DrugAllergyGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dagb *DrugAllergyGroupBy) Aggregate(fns ...AggregateFunc) *DrugAllergyGroupBy {
	dagb.fns = append(dagb.fns, fns...)
	return dagb
}

// Scan applies the group-by query and scan the result into the given value.
func (dagb *DrugAllergyGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dagb.path(ctx)
	if err != nil {
		return err
	}
	dagb.sql = query
	return dagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dagb *DrugAllergyGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (dagb *DrugAllergyGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dagb.fields) > 1 {
		return nil, errors.New("ent: DrugAllergyGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dagb *DrugAllergyGroupBy) StringsX(ctx context.Context) []string {
	v, err := dagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (dagb *DrugAllergyGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{drugallergy.Label}
	default:
		err = fmt.Errorf("ent: DrugAllergyGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dagb *DrugAllergyGroupBy) StringX(ctx context.Context) string {
	v, err := dagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (dagb *DrugAllergyGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dagb.fields) > 1 {
		return nil, errors.New("ent: DrugAllergyGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dagb *DrugAllergyGroupBy) IntsX(ctx context.Context) []int {
	v, err := dagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (dagb *DrugAllergyGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{drugallergy.Label}
	default:
		err = fmt.Errorf("ent: DrugAllergyGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dagb *DrugAllergyGroupBy) IntX(ctx context.Context) int {
	v, err := dagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (dagb *DrugAllergyGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dagb.fields) > 1 {
		return nil, errors.New("ent: DrugAllergyGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dagb *DrugAllergyGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (dagb *DrugAllergyGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{drugallergy.Label}
	default:
		err = fmt.Errorf("ent: DrugAllergyGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dagb *DrugAllergyGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (dagb *DrugAllergyGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dagb.fields) > 1 {
		return nil, errors.New("ent: DrugAllergyGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dagb *DrugAllergyGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (dagb *DrugAllergyGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{drugallergy.Label}
	default:
		err = fmt.Errorf("ent: DrugAllergyGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dagb *DrugAllergyGroupBy) BoolX(ctx context.Context) bool {
	v, err := dagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dagb *DrugAllergyGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dagb.sqlQuery().Query()
	if err := dagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dagb *DrugAllergyGroupBy) sqlQuery() *sql.Selector {
	selector := dagb.sql
	columns := make([]string, 0, len(dagb.fields)+len(dagb.fns))
	columns = append(columns, dagb.fields...)
	for _, fn := range dagb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(dagb.fields...)
}

// DrugAllergySelect is the builder for select fields of DrugAllergy entities.
type DrugAllergySelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (das *DrugAllergySelect) Scan(ctx context.Context, v interface{}) error {
	query, err := das.path(ctx)
	if err != nil {
		return err
	}
	das.sql = query
	return das.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (das *DrugAllergySelect) ScanX(ctx context.Context, v interface{}) {
	if err := das.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (das *DrugAllergySelect) Strings(ctx context.Context) ([]string, error) {
	if len(das.fields) > 1 {
		return nil, errors.New("ent: DrugAllergySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := das.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (das *DrugAllergySelect) StringsX(ctx context.Context) []string {
	v, err := das.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (das *DrugAllergySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = das.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{drugallergy.Label}
	default:
		err = fmt.Errorf("ent: DrugAllergySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (das *DrugAllergySelect) StringX(ctx context.Context) string {
	v, err := das.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (das *DrugAllergySelect) Ints(ctx context.Context) ([]int, error) {
	if len(das.fields) > 1 {
		return nil, errors.New("ent: DrugAllergySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := das.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (das *DrugAllergySelect) IntsX(ctx context.Context) []int {
	v, err := das.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (das *DrugAllergySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = das.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{drugallergy.Label}
	default:
		err = fmt.Errorf("ent: DrugAllergySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (das *DrugAllergySelect) IntX(ctx context.Context) int {
	v, err := das.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (das *DrugAllergySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(das.fields) > 1 {
		return nil, errors.New("ent: DrugAllergySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := das.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (das *DrugAllergySelect) Float64sX(ctx context.Context) []float64 {
	v, err := das.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (das *DrugAllergySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = das.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{drugallergy.Label}
	default:
		err = fmt.Errorf("ent: DrugAllergySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (das *DrugAllergySelect) Float64X(ctx context.Context) float64 {
	v, err := das.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (das *DrugAllergySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(das.fields) > 1 {
		return nil, errors.New("ent: DrugAllergySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := das.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (das *DrugAllergySelect) BoolsX(ctx context.Context) []bool {
	v, err := das.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (das *DrugAllergySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = das.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{drugallergy.Label}
	default:
		err = fmt.Errorf("ent: DrugAllergySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (das *DrugAllergySelect) BoolX(ctx context.Context) bool {
	v, err := das.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (das *DrugAllergySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := das.sqlQuery().Query()
	if err := das.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (das *DrugAllergySelect) sqlQuery() sql.Querier {
	selector := das.sql
	selector.Select(selector.Columns(das.fields...)...)
	return selector
}
