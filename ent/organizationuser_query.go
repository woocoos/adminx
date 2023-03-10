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
	"github.com/woocoos/adminx/ent/organizationuser"
	"github.com/woocoos/adminx/ent/predicate"
	"github.com/woocoos/adminx/ent/user"
)

// OrganizationUserQuery is the builder for querying OrganizationUser entities.
type OrganizationUserQuery struct {
	config
	ctx              *QueryContext
	order            []OrderFunc
	inters           []Interceptor
	predicates       []predicate.OrganizationUser
	withOrganization *OrganizationQuery
	withUser         *UserQuery
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*OrganizationUser) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrganizationUserQuery builder.
func (ouq *OrganizationUserQuery) Where(ps ...predicate.OrganizationUser) *OrganizationUserQuery {
	ouq.predicates = append(ouq.predicates, ps...)
	return ouq
}

// Limit the number of records to be returned by this query.
func (ouq *OrganizationUserQuery) Limit(limit int) *OrganizationUserQuery {
	ouq.ctx.Limit = &limit
	return ouq
}

// Offset to start from.
func (ouq *OrganizationUserQuery) Offset(offset int) *OrganizationUserQuery {
	ouq.ctx.Offset = &offset
	return ouq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ouq *OrganizationUserQuery) Unique(unique bool) *OrganizationUserQuery {
	ouq.ctx.Unique = &unique
	return ouq
}

// Order specifies how the records should be ordered.
func (ouq *OrganizationUserQuery) Order(o ...OrderFunc) *OrganizationUserQuery {
	ouq.order = append(ouq.order, o...)
	return ouq
}

// QueryOrganization chains the current query on the "organization" edge.
func (ouq *OrganizationUserQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: ouq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ouq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ouq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organizationuser.Table, organizationuser.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, organizationuser.OrganizationTable, organizationuser.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(ouq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (ouq *OrganizationUserQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ouq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ouq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ouq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organizationuser.Table, organizationuser.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, organizationuser.UserTable, organizationuser.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ouq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrganizationUser entity from the query.
// Returns a *NotFoundError when no OrganizationUser was found.
func (ouq *OrganizationUserQuery) First(ctx context.Context) (*OrganizationUser, error) {
	nodes, err := ouq.Limit(1).All(setContextOp(ctx, ouq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{organizationuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ouq *OrganizationUserQuery) FirstX(ctx context.Context) *OrganizationUser {
	node, err := ouq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrganizationUser ID from the query.
// Returns a *NotFoundError when no OrganizationUser ID was found.
func (ouq *OrganizationUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ouq.Limit(1).IDs(setContextOp(ctx, ouq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{organizationuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ouq *OrganizationUserQuery) FirstIDX(ctx context.Context) int {
	id, err := ouq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrganizationUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrganizationUser entity is found.
// Returns a *NotFoundError when no OrganizationUser entities are found.
func (ouq *OrganizationUserQuery) Only(ctx context.Context) (*OrganizationUser, error) {
	nodes, err := ouq.Limit(2).All(setContextOp(ctx, ouq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{organizationuser.Label}
	default:
		return nil, &NotSingularError{organizationuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ouq *OrganizationUserQuery) OnlyX(ctx context.Context) *OrganizationUser {
	node, err := ouq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrganizationUser ID in the query.
// Returns a *NotSingularError when more than one OrganizationUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (ouq *OrganizationUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ouq.Limit(2).IDs(setContextOp(ctx, ouq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{organizationuser.Label}
	default:
		err = &NotSingularError{organizationuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ouq *OrganizationUserQuery) OnlyIDX(ctx context.Context) int {
	id, err := ouq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrganizationUsers.
func (ouq *OrganizationUserQuery) All(ctx context.Context) ([]*OrganizationUser, error) {
	ctx = setContextOp(ctx, ouq.ctx, "All")
	if err := ouq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrganizationUser, *OrganizationUserQuery]()
	return withInterceptors[[]*OrganizationUser](ctx, ouq, qr, ouq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ouq *OrganizationUserQuery) AllX(ctx context.Context) []*OrganizationUser {
	nodes, err := ouq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrganizationUser IDs.
func (ouq *OrganizationUserQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ouq.ctx.Unique == nil && ouq.path != nil {
		ouq.Unique(true)
	}
	ctx = setContextOp(ctx, ouq.ctx, "IDs")
	if err = ouq.Select(organizationuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ouq *OrganizationUserQuery) IDsX(ctx context.Context) []int {
	ids, err := ouq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ouq *OrganizationUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ouq.ctx, "Count")
	if err := ouq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ouq, querierCount[*OrganizationUserQuery](), ouq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ouq *OrganizationUserQuery) CountX(ctx context.Context) int {
	count, err := ouq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ouq *OrganizationUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ouq.ctx, "Exist")
	switch _, err := ouq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ouq *OrganizationUserQuery) ExistX(ctx context.Context) bool {
	exist, err := ouq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrganizationUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ouq *OrganizationUserQuery) Clone() *OrganizationUserQuery {
	if ouq == nil {
		return nil
	}
	return &OrganizationUserQuery{
		config:           ouq.config,
		ctx:              ouq.ctx.Clone(),
		order:            append([]OrderFunc{}, ouq.order...),
		inters:           append([]Interceptor{}, ouq.inters...),
		predicates:       append([]predicate.OrganizationUser{}, ouq.predicates...),
		withOrganization: ouq.withOrganization.Clone(),
		withUser:         ouq.withUser.Clone(),
		// clone intermediate query.
		sql:  ouq.sql.Clone(),
		path: ouq.path,
	}
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (ouq *OrganizationUserQuery) WithOrganization(opts ...func(*OrganizationQuery)) *OrganizationUserQuery {
	query := (&OrganizationClient{config: ouq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ouq.withOrganization = query
	return ouq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ouq *OrganizationUserQuery) WithUser(opts ...func(*UserQuery)) *OrganizationUserQuery {
	query := (&UserClient{config: ouq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ouq.withUser = query
	return ouq
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
//	client.OrganizationUser.Query().
//		GroupBy(organizationuser.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ouq *OrganizationUserQuery) GroupBy(field string, fields ...string) *OrganizationUserGroupBy {
	ouq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrganizationUserGroupBy{build: ouq}
	grbuild.flds = &ouq.ctx.Fields
	grbuild.label = organizationuser.Label
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
//	client.OrganizationUser.Query().
//		Select(organizationuser.FieldCreatedBy).
//		Scan(ctx, &v)
func (ouq *OrganizationUserQuery) Select(fields ...string) *OrganizationUserSelect {
	ouq.ctx.Fields = append(ouq.ctx.Fields, fields...)
	sbuild := &OrganizationUserSelect{OrganizationUserQuery: ouq}
	sbuild.label = organizationuser.Label
	sbuild.flds, sbuild.scan = &ouq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrganizationUserSelect configured with the given aggregations.
func (ouq *OrganizationUserQuery) Aggregate(fns ...AggregateFunc) *OrganizationUserSelect {
	return ouq.Select().Aggregate(fns...)
}

func (ouq *OrganizationUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ouq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ouq); err != nil {
				return err
			}
		}
	}
	for _, f := range ouq.ctx.Fields {
		if !organizationuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ouq.path != nil {
		prev, err := ouq.path(ctx)
		if err != nil {
			return err
		}
		ouq.sql = prev
	}
	return nil
}

func (ouq *OrganizationUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrganizationUser, error) {
	var (
		nodes       = []*OrganizationUser{}
		_spec       = ouq.querySpec()
		loadedTypes = [2]bool{
			ouq.withOrganization != nil,
			ouq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrganizationUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrganizationUser{config: ouq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ouq.modifiers) > 0 {
		_spec.Modifiers = ouq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ouq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ouq.withOrganization; query != nil {
		if err := ouq.loadOrganization(ctx, query, nodes, nil,
			func(n *OrganizationUser, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := ouq.withUser; query != nil {
		if err := ouq.loadUser(ctx, query, nodes, nil,
			func(n *OrganizationUser, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	for i := range ouq.loadTotal {
		if err := ouq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ouq *OrganizationUserQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*OrganizationUser, init func(*OrganizationUser), assign func(*OrganizationUser, *Organization)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*OrganizationUser)
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
func (ouq *OrganizationUserQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*OrganizationUser, init func(*OrganizationUser), assign func(*OrganizationUser, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*OrganizationUser)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ouq *OrganizationUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ouq.querySpec()
	if len(ouq.modifiers) > 0 {
		_spec.Modifiers = ouq.modifiers
	}
	_spec.Node.Columns = ouq.ctx.Fields
	if len(ouq.ctx.Fields) > 0 {
		_spec.Unique = ouq.ctx.Unique != nil && *ouq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ouq.driver, _spec)
}

func (ouq *OrganizationUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(organizationuser.Table, organizationuser.Columns, sqlgraph.NewFieldSpec(organizationuser.FieldID, field.TypeInt))
	_spec.From = ouq.sql
	if unique := ouq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ouq.path != nil {
		_spec.Unique = true
	}
	if fields := ouq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationuser.FieldID)
		for i := range fields {
			if fields[i] != organizationuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ouq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ouq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ouq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ouq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ouq *OrganizationUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ouq.driver.Dialect())
	t1 := builder.Table(organizationuser.Table)
	columns := ouq.ctx.Fields
	if len(columns) == 0 {
		columns = organizationuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ouq.sql != nil {
		selector = ouq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ouq.ctx.Unique != nil && *ouq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ouq.predicates {
		p(selector)
	}
	for _, p := range ouq.order {
		p(selector)
	}
	if offset := ouq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ouq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrganizationUserGroupBy is the group-by builder for OrganizationUser entities.
type OrganizationUserGroupBy struct {
	selector
	build *OrganizationUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ougb *OrganizationUserGroupBy) Aggregate(fns ...AggregateFunc) *OrganizationUserGroupBy {
	ougb.fns = append(ougb.fns, fns...)
	return ougb
}

// Scan applies the selector query and scans the result into the given value.
func (ougb *OrganizationUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ougb.build.ctx, "GroupBy")
	if err := ougb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrganizationUserQuery, *OrganizationUserGroupBy](ctx, ougb.build, ougb, ougb.build.inters, v)
}

func (ougb *OrganizationUserGroupBy) sqlScan(ctx context.Context, root *OrganizationUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ougb.fns))
	for _, fn := range ougb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ougb.flds)+len(ougb.fns))
		for _, f := range *ougb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ougb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ougb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrganizationUserSelect is the builder for selecting fields of OrganizationUser entities.
type OrganizationUserSelect struct {
	*OrganizationUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ous *OrganizationUserSelect) Aggregate(fns ...AggregateFunc) *OrganizationUserSelect {
	ous.fns = append(ous.fns, fns...)
	return ous
}

// Scan applies the selector query and scans the result into the given value.
func (ous *OrganizationUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ous.ctx, "Select")
	if err := ous.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrganizationUserQuery, *OrganizationUserSelect](ctx, ous.OrganizationUserQuery, ous, ous.inters, v)
}

func (ous *OrganizationUserSelect) sqlScan(ctx context.Context, root *OrganizationUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ous.fns))
	for _, fn := range ous.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ous.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ous.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
