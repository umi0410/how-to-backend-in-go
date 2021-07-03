// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"db/ent/predicate"
	"db/ent/tourproduct"
	"db/ent/user"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TourProductQuery is the builder for querying TourProduct entities.
type TourProductQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TourProduct
	// eager-loading edges.
	withManager *UserQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TourProductQuery builder.
func (tpq *TourProductQuery) Where(ps ...predicate.TourProduct) *TourProductQuery {
	tpq.predicates = append(tpq.predicates, ps...)
	return tpq
}

// Limit adds a limit step to the query.
func (tpq *TourProductQuery) Limit(limit int) *TourProductQuery {
	tpq.limit = &limit
	return tpq
}

// Offset adds an offset step to the query.
func (tpq *TourProductQuery) Offset(offset int) *TourProductQuery {
	tpq.offset = &offset
	return tpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tpq *TourProductQuery) Unique(unique bool) *TourProductQuery {
	tpq.unique = &unique
	return tpq
}

// Order adds an order step to the query.
func (tpq *TourProductQuery) Order(o ...OrderFunc) *TourProductQuery {
	tpq.order = append(tpq.order, o...)
	return tpq
}

// QueryManager chains the current query on the "manager" edge.
func (tpq *TourProductQuery) QueryManager() *UserQuery {
	query := &UserQuery{config: tpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tourproduct.Table, tourproduct.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tourproduct.ManagerTable, tourproduct.ManagerColumn),
		)
		fromU = sqlgraph.SetNeighbors(tpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TourProduct entity from the query.
// Returns a *NotFoundError when no TourProduct was found.
func (tpq *TourProductQuery) First(ctx context.Context) (*TourProduct, error) {
	nodes, err := tpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tourproduct.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tpq *TourProductQuery) FirstX(ctx context.Context) *TourProduct {
	node, err := tpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TourProduct ID from the query.
// Returns a *NotFoundError when no TourProduct ID was found.
func (tpq *TourProductQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tourproduct.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tpq *TourProductQuery) FirstIDX(ctx context.Context) int {
	id, err := tpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TourProduct entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one TourProduct entity is not found.
// Returns a *NotFoundError when no TourProduct entities are found.
func (tpq *TourProductQuery) Only(ctx context.Context) (*TourProduct, error) {
	nodes, err := tpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tourproduct.Label}
	default:
		return nil, &NotSingularError{tourproduct.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tpq *TourProductQuery) OnlyX(ctx context.Context) *TourProduct {
	node, err := tpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TourProduct ID in the query.
// Returns a *NotSingularError when exactly one TourProduct ID is not found.
// Returns a *NotFoundError when no entities are found.
func (tpq *TourProductQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tourproduct.Label}
	default:
		err = &NotSingularError{tourproduct.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tpq *TourProductQuery) OnlyIDX(ctx context.Context) int {
	id, err := tpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TourProducts.
func (tpq *TourProductQuery) All(ctx context.Context) ([]*TourProduct, error) {
	if err := tpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tpq *TourProductQuery) AllX(ctx context.Context) []*TourProduct {
	nodes, err := tpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TourProduct IDs.
func (tpq *TourProductQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tpq.Select(tourproduct.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tpq *TourProductQuery) IDsX(ctx context.Context) []int {
	ids, err := tpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tpq *TourProductQuery) Count(ctx context.Context) (int, error) {
	if err := tpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tpq *TourProductQuery) CountX(ctx context.Context) int {
	count, err := tpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tpq *TourProductQuery) Exist(ctx context.Context) (bool, error) {
	if err := tpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tpq *TourProductQuery) ExistX(ctx context.Context) bool {
	exist, err := tpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TourProductQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tpq *TourProductQuery) Clone() *TourProductQuery {
	if tpq == nil {
		return nil
	}
	return &TourProductQuery{
		config:      tpq.config,
		limit:       tpq.limit,
		offset:      tpq.offset,
		order:       append([]OrderFunc{}, tpq.order...),
		predicates:  append([]predicate.TourProduct{}, tpq.predicates...),
		withManager: tpq.withManager.Clone(),
		// clone intermediate query.
		sql:  tpq.sql.Clone(),
		path: tpq.path,
	}
}

// WithManager tells the query-builder to eager-load the nodes that are connected to
// the "manager" edge. The optional arguments are used to configure the query builder of the edge.
func (tpq *TourProductQuery) WithManager(opts ...func(*UserQuery)) *TourProductQuery {
	query := &UserQuery{config: tpq.config}
	for _, opt := range opts {
		opt(query)
	}
	tpq.withManager = query
	return tpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TourProduct.Query().
//		GroupBy(tourproduct.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tpq *TourProductQuery) GroupBy(field string, fields ...string) *TourProductGroupBy {
	group := &TourProductGroupBy{config: tpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tpq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.TourProduct.Query().
//		Select(tourproduct.FieldName).
//		Scan(ctx, &v)
//
func (tpq *TourProductQuery) Select(field string, fields ...string) *TourProductSelect {
	tpq.fields = append([]string{field}, fields...)
	return &TourProductSelect{TourProductQuery: tpq}
}

func (tpq *TourProductQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tpq.fields {
		if !tourproduct.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tpq.path != nil {
		prev, err := tpq.path(ctx)
		if err != nil {
			return err
		}
		tpq.sql = prev
	}
	return nil
}

func (tpq *TourProductQuery) sqlAll(ctx context.Context) ([]*TourProduct, error) {
	var (
		nodes       = []*TourProduct{}
		withFKs     = tpq.withFKs
		_spec       = tpq.querySpec()
		loadedTypes = [1]bool{
			tpq.withManager != nil,
		}
	)
	if tpq.withManager != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, tourproduct.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TourProduct{config: tpq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, tpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tpq.withManager; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*TourProduct)
		for i := range nodes {
			if nodes[i].user_products == nil {
				continue
			}
			fk := *nodes[i].user_products
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_products" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Manager = n
			}
		}
	}

	return nodes, nil
}

func (tpq *TourProductQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tpq.querySpec()
	return sqlgraph.CountNodes(ctx, tpq.driver, _spec)
}

func (tpq *TourProductQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tpq *TourProductQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tourproduct.Table,
			Columns: tourproduct.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tourproduct.FieldID,
			},
		},
		From:   tpq.sql,
		Unique: true,
	}
	if unique := tpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tourproduct.FieldID)
		for i := range fields {
			if fields[i] != tourproduct.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tpq *TourProductQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tpq.driver.Dialect())
	t1 := builder.Table(tourproduct.Table)
	selector := builder.Select(t1.Columns(tourproduct.Columns...)...).From(t1)
	if tpq.sql != nil {
		selector = tpq.sql
		selector.Select(selector.Columns(tourproduct.Columns...)...)
	}
	for _, p := range tpq.predicates {
		p(selector)
	}
	for _, p := range tpq.order {
		p(selector)
	}
	if offset := tpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TourProductGroupBy is the group-by builder for TourProduct entities.
type TourProductGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tpgb *TourProductGroupBy) Aggregate(fns ...AggregateFunc) *TourProductGroupBy {
	tpgb.fns = append(tpgb.fns, fns...)
	return tpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tpgb *TourProductGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tpgb.path(ctx)
	if err != nil {
		return err
	}
	tpgb.sql = query
	return tpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tpgb *TourProductGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tpgb *TourProductGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tpgb.fields) > 1 {
		return nil, errors.New("ent: TourProductGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tpgb *TourProductGroupBy) StringsX(ctx context.Context) []string {
	v, err := tpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tpgb *TourProductGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tourproduct.Label}
	default:
		err = fmt.Errorf("ent: TourProductGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tpgb *TourProductGroupBy) StringX(ctx context.Context) string {
	v, err := tpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tpgb *TourProductGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tpgb.fields) > 1 {
		return nil, errors.New("ent: TourProductGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tpgb *TourProductGroupBy) IntsX(ctx context.Context) []int {
	v, err := tpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tpgb *TourProductGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tourproduct.Label}
	default:
		err = fmt.Errorf("ent: TourProductGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tpgb *TourProductGroupBy) IntX(ctx context.Context) int {
	v, err := tpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tpgb *TourProductGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tpgb.fields) > 1 {
		return nil, errors.New("ent: TourProductGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tpgb *TourProductGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tpgb *TourProductGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tourproduct.Label}
	default:
		err = fmt.Errorf("ent: TourProductGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tpgb *TourProductGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tpgb *TourProductGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tpgb.fields) > 1 {
		return nil, errors.New("ent: TourProductGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tpgb *TourProductGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tpgb *TourProductGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tourproduct.Label}
	default:
		err = fmt.Errorf("ent: TourProductGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tpgb *TourProductGroupBy) BoolX(ctx context.Context) bool {
	v, err := tpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tpgb *TourProductGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tpgb.fields {
		if !tourproduct.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tpgb *TourProductGroupBy) sqlQuery() *sql.Selector {
	selector := tpgb.sql
	columns := make([]string, 0, len(tpgb.fields)+len(tpgb.fns))
	columns = append(columns, tpgb.fields...)
	for _, fn := range tpgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(tpgb.fields...)
}

// TourProductSelect is the builder for selecting fields of TourProduct entities.
type TourProductSelect struct {
	*TourProductQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tps *TourProductSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tps.prepareQuery(ctx); err != nil {
		return err
	}
	tps.sql = tps.TourProductQuery.sqlQuery(ctx)
	return tps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tps *TourProductSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tps *TourProductSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tps.fields) > 1 {
		return nil, errors.New("ent: TourProductSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tps *TourProductSelect) StringsX(ctx context.Context) []string {
	v, err := tps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tps *TourProductSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tourproduct.Label}
	default:
		err = fmt.Errorf("ent: TourProductSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tps *TourProductSelect) StringX(ctx context.Context) string {
	v, err := tps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tps *TourProductSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tps.fields) > 1 {
		return nil, errors.New("ent: TourProductSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tps *TourProductSelect) IntsX(ctx context.Context) []int {
	v, err := tps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tps *TourProductSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tourproduct.Label}
	default:
		err = fmt.Errorf("ent: TourProductSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tps *TourProductSelect) IntX(ctx context.Context) int {
	v, err := tps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tps *TourProductSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tps.fields) > 1 {
		return nil, errors.New("ent: TourProductSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tps *TourProductSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tps *TourProductSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tourproduct.Label}
	default:
		err = fmt.Errorf("ent: TourProductSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tps *TourProductSelect) Float64X(ctx context.Context) float64 {
	v, err := tps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tps *TourProductSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tps.fields) > 1 {
		return nil, errors.New("ent: TourProductSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tps *TourProductSelect) BoolsX(ctx context.Context) []bool {
	v, err := tps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tps *TourProductSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tourproduct.Label}
	default:
		err = fmt.Errorf("ent: TourProductSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tps *TourProductSelect) BoolX(ctx context.Context) bool {
	v, err := tps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tps *TourProductSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tps.sqlQuery().Query()
	if err := tps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tps *TourProductSelect) sqlQuery() sql.Querier {
	selector := tps.sql
	selector.Select(selector.Columns(tps.fields...)...)
	return selector
}
