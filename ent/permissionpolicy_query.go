// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/adminx/ent/organization"
	"github.com/woocoos/adminx/ent/permissionpolicy"
	"github.com/woocoos/adminx/ent/predicate"
)

// PermissionPolicyQuery is the builder for querying PermissionPolicy entities.
type PermissionPolicyQuery struct {
	config
	ctx              *QueryContext
	order            []OrderFunc
	inters           []Interceptor
	predicates       []predicate.PermissionPolicy
	withOrganization *OrganizationQuery
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*PermissionPolicy) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PermissionPolicyQuery builder.
func (ppq *PermissionPolicyQuery) Where(ps ...predicate.PermissionPolicy) *PermissionPolicyQuery {
	ppq.predicates = append(ppq.predicates, ps...)
	return ppq
}

// Limit the number of records to be returned by this query.
func (ppq *PermissionPolicyQuery) Limit(limit int) *PermissionPolicyQuery {
	ppq.ctx.Limit = &limit
	return ppq
}

// Offset to start from.
func (ppq *PermissionPolicyQuery) Offset(offset int) *PermissionPolicyQuery {
	ppq.ctx.Offset = &offset
	return ppq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ppq *PermissionPolicyQuery) Unique(unique bool) *PermissionPolicyQuery {
	ppq.ctx.Unique = &unique
	return ppq
}

// Order specifies how the records should be ordered.
func (ppq *PermissionPolicyQuery) Order(o ...OrderFunc) *PermissionPolicyQuery {
	ppq.order = append(ppq.order, o...)
	return ppq
}

// QueryOrganization chains the current query on the "organization" edge.
func (ppq *PermissionPolicyQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: ppq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ppq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ppq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(permissionpolicy.Table, permissionpolicy.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, permissionpolicy.OrganizationTable, permissionpolicy.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(ppq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PermissionPolicy entity from the query.
// Returns a *NotFoundError when no PermissionPolicy was found.
func (ppq *PermissionPolicyQuery) First(ctx context.Context) (*PermissionPolicy, error) {
	nodes, err := ppq.Limit(1).All(setContextOp(ctx, ppq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{permissionpolicy.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ppq *PermissionPolicyQuery) FirstX(ctx context.Context) *PermissionPolicy {
	node, err := ppq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PermissionPolicy ID from the query.
// Returns a *NotFoundError when no PermissionPolicy ID was found.
func (ppq *PermissionPolicyQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ppq.Limit(1).IDs(setContextOp(ctx, ppq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{permissionpolicy.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ppq *PermissionPolicyQuery) FirstIDX(ctx context.Context) int {
	id, err := ppq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PermissionPolicy entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PermissionPolicy entity is found.
// Returns a *NotFoundError when no PermissionPolicy entities are found.
func (ppq *PermissionPolicyQuery) Only(ctx context.Context) (*PermissionPolicy, error) {
	nodes, err := ppq.Limit(2).All(setContextOp(ctx, ppq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{permissionpolicy.Label}
	default:
		return nil, &NotSingularError{permissionpolicy.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ppq *PermissionPolicyQuery) OnlyX(ctx context.Context) *PermissionPolicy {
	node, err := ppq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PermissionPolicy ID in the query.
// Returns a *NotSingularError when more than one PermissionPolicy ID is found.
// Returns a *NotFoundError when no entities are found.
func (ppq *PermissionPolicyQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ppq.Limit(2).IDs(setContextOp(ctx, ppq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{permissionpolicy.Label}
	default:
		err = &NotSingularError{permissionpolicy.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ppq *PermissionPolicyQuery) OnlyIDX(ctx context.Context) int {
	id, err := ppq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PermissionPolicies.
func (ppq *PermissionPolicyQuery) All(ctx context.Context) ([]*PermissionPolicy, error) {
	ctx = setContextOp(ctx, ppq.ctx, "All")
	if err := ppq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PermissionPolicy, *PermissionPolicyQuery]()
	return withInterceptors[[]*PermissionPolicy](ctx, ppq, qr, ppq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ppq *PermissionPolicyQuery) AllX(ctx context.Context) []*PermissionPolicy {
	nodes, err := ppq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PermissionPolicy IDs.
func (ppq *PermissionPolicyQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ppq.ctx.Unique == nil && ppq.path != nil {
		ppq.Unique(true)
	}
	ctx = setContextOp(ctx, ppq.ctx, "IDs")
	if err = ppq.Select(permissionpolicy.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ppq *PermissionPolicyQuery) IDsX(ctx context.Context) []int {
	ids, err := ppq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ppq *PermissionPolicyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ppq.ctx, "Count")
	if err := ppq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ppq, querierCount[*PermissionPolicyQuery](), ppq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ppq *PermissionPolicyQuery) CountX(ctx context.Context) int {
	count, err := ppq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ppq *PermissionPolicyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ppq.ctx, "Exist")
	switch _, err := ppq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ppq *PermissionPolicyQuery) ExistX(ctx context.Context) bool {
	exist, err := ppq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PermissionPolicyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ppq *PermissionPolicyQuery) Clone() *PermissionPolicyQuery {
	if ppq == nil {
		return nil
	}
	return &PermissionPolicyQuery{
		config:           ppq.config,
		ctx:              ppq.ctx.Clone(),
		order:            append([]OrderFunc{}, ppq.order...),
		inters:           append([]Interceptor{}, ppq.inters...),
		predicates:       append([]predicate.PermissionPolicy{}, ppq.predicates...),
		withOrganization: ppq.withOrganization.Clone(),
		// clone intermediate query.
		sql:  ppq.sql.Clone(),
		path: ppq.path,
	}
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (ppq *PermissionPolicyQuery) WithOrganization(opts ...func(*OrganizationQuery)) *PermissionPolicyQuery {
	query := (&OrganizationClient{config: ppq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ppq.withOrganization = query
	return ppq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy int `json:"created_by,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PermissionPolicy.Query().
//		GroupBy(permissionpolicy.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ppq *PermissionPolicyQuery) GroupBy(field string, fields ...string) *PermissionPolicyGroupBy {
	ppq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PermissionPolicyGroupBy{build: ppq}
	grbuild.flds = &ppq.ctx.Fields
	grbuild.label = permissionpolicy.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy int `json:"created_by,omitempty"`
//	}
//
//	client.PermissionPolicy.Query().
//		Select(permissionpolicy.FieldCreatedBy).
//		Scan(ctx, &v)
func (ppq *PermissionPolicyQuery) Select(fields ...string) *PermissionPolicySelect {
	ppq.ctx.Fields = append(ppq.ctx.Fields, fields...)
	sbuild := &PermissionPolicySelect{PermissionPolicyQuery: ppq}
	sbuild.label = permissionpolicy.Label
	sbuild.flds, sbuild.scan = &ppq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PermissionPolicySelect configured with the given aggregations.
func (ppq *PermissionPolicyQuery) Aggregate(fns ...AggregateFunc) *PermissionPolicySelect {
	return ppq.Select().Aggregate(fns...)
}

func (ppq *PermissionPolicyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ppq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ppq); err != nil {
				return err
			}
		}
	}
	for _, f := range ppq.ctx.Fields {
		if !permissionpolicy.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ppq.path != nil {
		prev, err := ppq.path(ctx)
		if err != nil {
			return err
		}
		ppq.sql = prev
	}
	return nil
}

func (ppq *PermissionPolicyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PermissionPolicy, error) {
	var (
		nodes       = []*PermissionPolicy{}
		_spec       = ppq.querySpec()
		loadedTypes = [1]bool{
			ppq.withOrganization != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PermissionPolicy).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PermissionPolicy{config: ppq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ppq.modifiers) > 0 {
		_spec.Modifiers = ppq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ppq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ppq.withOrganization; query != nil {
		if err := ppq.loadOrganization(ctx, query, nodes, nil,
			func(n *PermissionPolicy, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	for i := range ppq.loadTotal {
		if err := ppq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ppq *PermissionPolicyQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*PermissionPolicy, init func(*PermissionPolicy), assign func(*PermissionPolicy, *Organization)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PermissionPolicy)
	for i := range nodes {
		fk := nodes[i].OrgID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(organization.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "org_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ppq *PermissionPolicyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ppq.querySpec()
	if len(ppq.modifiers) > 0 {
		_spec.Modifiers = ppq.modifiers
	}
	_spec.Node.Columns = ppq.ctx.Fields
	if len(ppq.ctx.Fields) > 0 {
		_spec.Unique = ppq.ctx.Unique != nil && *ppq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ppq.driver, _spec)
}

func (ppq *PermissionPolicyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(permissionpolicy.Table, permissionpolicy.Columns, sqlgraph.NewFieldSpec(permissionpolicy.FieldID, field.TypeInt))
	_spec.From = ppq.sql
	if unique := ppq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ppq.path != nil {
		_spec.Unique = true
	}
	if fields := ppq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permissionpolicy.FieldID)
		for i := range fields {
			if fields[i] != permissionpolicy.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ppq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ppq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ppq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ppq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ppq *PermissionPolicyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ppq.driver.Dialect())
	t1 := builder.Table(permissionpolicy.Table)
	columns := ppq.ctx.Fields
	if len(columns) == 0 {
		columns = permissionpolicy.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ppq.sql != nil {
		selector = ppq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ppq.ctx.Unique != nil && *ppq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ppq.predicates {
		p(selector)
	}
	for _, p := range ppq.order {
		p(selector)
	}
	if offset := ppq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ppq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PermissionPolicyGroupBy is the group-by builder for PermissionPolicy entities.
type PermissionPolicyGroupBy struct {
	selector
	build *PermissionPolicyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ppgb *PermissionPolicyGroupBy) Aggregate(fns ...AggregateFunc) *PermissionPolicyGroupBy {
	ppgb.fns = append(ppgb.fns, fns...)
	return ppgb
}

// Scan applies the selector query and scans the result into the given value.
func (ppgb *PermissionPolicyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ppgb.build.ctx, "GroupBy")
	if err := ppgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PermissionPolicyQuery, *PermissionPolicyGroupBy](ctx, ppgb.build, ppgb, ppgb.build.inters, v)
}

func (ppgb *PermissionPolicyGroupBy) sqlScan(ctx context.Context, root *PermissionPolicyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ppgb.fns))
	for _, fn := range ppgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ppgb.flds)+len(ppgb.fns))
		for _, f := range *ppgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ppgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ppgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PermissionPolicySelect is the builder for selecting fields of PermissionPolicy entities.
type PermissionPolicySelect struct {
	*PermissionPolicyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pps *PermissionPolicySelect) Aggregate(fns ...AggregateFunc) *PermissionPolicySelect {
	pps.fns = append(pps.fns, fns...)
	return pps
}

// Scan applies the selector query and scans the result into the given value.
func (pps *PermissionPolicySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pps.ctx, "Select")
	if err := pps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PermissionPolicyQuery, *PermissionPolicySelect](ctx, pps.PermissionPolicyQuery, pps, pps.inters, v)
}

func (pps *PermissionPolicySelect) sqlScan(ctx context.Context, root *PermissionPolicyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pps.fns))
	for _, fn := range pps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
