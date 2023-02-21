// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/adminx/ent/app"
	"github.com/woocoos/adminx/ent/appaction"
	"github.com/woocoos/adminx/ent/appmenu"
	"github.com/woocoos/adminx/ent/appres"
	"github.com/woocoos/adminx/ent/predicate"
)

// AppActionUpdate is the builder for updating AppAction entities.
type AppActionUpdate struct {
	config
	hooks    []Hook
	mutation *AppActionMutation
}

// Where appends a list predicates to the AppActionUpdate builder.
func (aau *AppActionUpdate) Where(ps ...predicate.AppAction) *AppActionUpdate {
	aau.mutation.Where(ps...)
	return aau
}

// SetUpdatedBy sets the "updated_by" field.
func (aau *AppActionUpdate) SetUpdatedBy(i int) *AppActionUpdate {
	aau.mutation.ResetUpdatedBy()
	aau.mutation.SetUpdatedBy(i)
	return aau
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (aau *AppActionUpdate) SetNillableUpdatedBy(i *int) *AppActionUpdate {
	if i != nil {
		aau.SetUpdatedBy(*i)
	}
	return aau
}

// AddUpdatedBy adds i to the "updated_by" field.
func (aau *AppActionUpdate) AddUpdatedBy(i int) *AppActionUpdate {
	aau.mutation.AddUpdatedBy(i)
	return aau
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (aau *AppActionUpdate) ClearUpdatedBy() *AppActionUpdate {
	aau.mutation.ClearUpdatedBy()
	return aau
}

// SetUpdatedAt sets the "updated_at" field.
func (aau *AppActionUpdate) SetUpdatedAt(t time.Time) *AppActionUpdate {
	aau.mutation.SetUpdatedAt(t)
	return aau
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aau *AppActionUpdate) SetNillableUpdatedAt(t *time.Time) *AppActionUpdate {
	if t != nil {
		aau.SetUpdatedAt(*t)
	}
	return aau
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (aau *AppActionUpdate) ClearUpdatedAt() *AppActionUpdate {
	aau.mutation.ClearUpdatedAt()
	return aau
}

// SetAppID sets the "app_id" field.
func (aau *AppActionUpdate) SetAppID(i int) *AppActionUpdate {
	aau.mutation.SetAppID(i)
	return aau
}

// SetName sets the "name" field.
func (aau *AppActionUpdate) SetName(s string) *AppActionUpdate {
	aau.mutation.SetName(s)
	return aau
}

// SetKind sets the "kind" field.
func (aau *AppActionUpdate) SetKind(a appaction.Kind) *AppActionUpdate {
	aau.mutation.SetKind(a)
	return aau
}

// SetMethod sets the "method" field.
func (aau *AppActionUpdate) SetMethod(a appaction.Method) *AppActionUpdate {
	aau.mutation.SetMethod(a)
	return aau
}

// SetComments sets the "comments" field.
func (aau *AppActionUpdate) SetComments(s string) *AppActionUpdate {
	aau.mutation.SetComments(s)
	return aau
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (aau *AppActionUpdate) SetNillableComments(s *string) *AppActionUpdate {
	if s != nil {
		aau.SetComments(*s)
	}
	return aau
}

// ClearComments clears the value of the "comments" field.
func (aau *AppActionUpdate) ClearComments() *AppActionUpdate {
	aau.mutation.ClearComments()
	return aau
}

// SetApp sets the "app" edge to the App entity.
func (aau *AppActionUpdate) SetApp(a *App) *AppActionUpdate {
	return aau.SetAppID(a.ID)
}

// AddMenuIDs adds the "menus" edge to the AppMenu entity by IDs.
func (aau *AppActionUpdate) AddMenuIDs(ids ...int) *AppActionUpdate {
	aau.mutation.AddMenuIDs(ids...)
	return aau
}

// AddMenus adds the "menus" edges to the AppMenu entity.
func (aau *AppActionUpdate) AddMenus(a ...*AppMenu) *AppActionUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aau.AddMenuIDs(ids...)
}

// AddResourceIDs adds the "resources" edge to the AppRes entity by IDs.
func (aau *AppActionUpdate) AddResourceIDs(ids ...int) *AppActionUpdate {
	aau.mutation.AddResourceIDs(ids...)
	return aau
}

// AddResources adds the "resources" edges to the AppRes entity.
func (aau *AppActionUpdate) AddResources(a ...*AppRes) *AppActionUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aau.AddResourceIDs(ids...)
}

// Mutation returns the AppActionMutation object of the builder.
func (aau *AppActionUpdate) Mutation() *AppActionMutation {
	return aau.mutation
}

// ClearApp clears the "app" edge to the App entity.
func (aau *AppActionUpdate) ClearApp() *AppActionUpdate {
	aau.mutation.ClearApp()
	return aau
}

// ClearMenus clears all "menus" edges to the AppMenu entity.
func (aau *AppActionUpdate) ClearMenus() *AppActionUpdate {
	aau.mutation.ClearMenus()
	return aau
}

// RemoveMenuIDs removes the "menus" edge to AppMenu entities by IDs.
func (aau *AppActionUpdate) RemoveMenuIDs(ids ...int) *AppActionUpdate {
	aau.mutation.RemoveMenuIDs(ids...)
	return aau
}

// RemoveMenus removes "menus" edges to AppMenu entities.
func (aau *AppActionUpdate) RemoveMenus(a ...*AppMenu) *AppActionUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aau.RemoveMenuIDs(ids...)
}

// ClearResources clears all "resources" edges to the AppRes entity.
func (aau *AppActionUpdate) ClearResources() *AppActionUpdate {
	aau.mutation.ClearResources()
	return aau
}

// RemoveResourceIDs removes the "resources" edge to AppRes entities by IDs.
func (aau *AppActionUpdate) RemoveResourceIDs(ids ...int) *AppActionUpdate {
	aau.mutation.RemoveResourceIDs(ids...)
	return aau
}

// RemoveResources removes "resources" edges to AppRes entities.
func (aau *AppActionUpdate) RemoveResources(a ...*AppRes) *AppActionUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aau.RemoveResourceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aau *AppActionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AppActionMutation](ctx, aau.sqlSave, aau.mutation, aau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aau *AppActionUpdate) SaveX(ctx context.Context) int {
	affected, err := aau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aau *AppActionUpdate) Exec(ctx context.Context) error {
	_, err := aau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aau *AppActionUpdate) ExecX(ctx context.Context) {
	if err := aau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aau *AppActionUpdate) check() error {
	if v, ok := aau.mutation.Name(); ok {
		if err := appaction.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "AppAction.name": %w`, err)}
		}
	}
	if v, ok := aau.mutation.Kind(); ok {
		if err := appaction.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "AppAction.kind": %w`, err)}
		}
	}
	if v, ok := aau.mutation.Method(); ok {
		if err := appaction.MethodValidator(v); err != nil {
			return &ValidationError{Name: "method", err: fmt.Errorf(`ent: validator failed for field "AppAction.method": %w`, err)}
		}
	}
	if _, ok := aau.mutation.AppID(); aau.mutation.AppCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AppAction.app"`)
	}
	return nil
}

func (aau *AppActionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := aau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(appaction.Table, appaction.Columns, sqlgraph.NewFieldSpec(appaction.FieldID, field.TypeInt))
	if ps := aau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aau.mutation.UpdatedBy(); ok {
		_spec.SetField(appaction.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := aau.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(appaction.FieldUpdatedBy, field.TypeInt, value)
	}
	if aau.mutation.UpdatedByCleared() {
		_spec.ClearField(appaction.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := aau.mutation.UpdatedAt(); ok {
		_spec.SetField(appaction.FieldUpdatedAt, field.TypeTime, value)
	}
	if aau.mutation.UpdatedAtCleared() {
		_spec.ClearField(appaction.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := aau.mutation.Name(); ok {
		_spec.SetField(appaction.FieldName, field.TypeString, value)
	}
	if value, ok := aau.mutation.Kind(); ok {
		_spec.SetField(appaction.FieldKind, field.TypeEnum, value)
	}
	if value, ok := aau.mutation.Method(); ok {
		_spec.SetField(appaction.FieldMethod, field.TypeEnum, value)
	}
	if value, ok := aau.mutation.Comments(); ok {
		_spec.SetField(appaction.FieldComments, field.TypeString, value)
	}
	if aau.mutation.CommentsCleared() {
		_spec.ClearField(appaction.FieldComments, field.TypeString)
	}
	if aau.mutation.AppCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appaction.AppTable,
			Columns: []string{appaction.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aau.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appaction.AppTable,
			Columns: []string{appaction.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aau.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appaction.MenusTable,
			Columns: []string{appaction.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appmenu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aau.mutation.RemovedMenusIDs(); len(nodes) > 0 && !aau.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appaction.MenusTable,
			Columns: []string{appaction.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appmenu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aau.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appaction.MenusTable,
			Columns: []string{appaction.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appmenu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aau.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appaction.ResourcesTable,
			Columns: []string{appaction.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appres.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aau.mutation.RemovedResourcesIDs(); len(nodes) > 0 && !aau.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appaction.ResourcesTable,
			Columns: []string{appaction.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appres.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aau.mutation.ResourcesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appaction.ResourcesTable,
			Columns: []string{appaction.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appres.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aau.mutation.done = true
	return n, nil
}

// AppActionUpdateOne is the builder for updating a single AppAction entity.
type AppActionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppActionMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (aauo *AppActionUpdateOne) SetUpdatedBy(i int) *AppActionUpdateOne {
	aauo.mutation.ResetUpdatedBy()
	aauo.mutation.SetUpdatedBy(i)
	return aauo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (aauo *AppActionUpdateOne) SetNillableUpdatedBy(i *int) *AppActionUpdateOne {
	if i != nil {
		aauo.SetUpdatedBy(*i)
	}
	return aauo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (aauo *AppActionUpdateOne) AddUpdatedBy(i int) *AppActionUpdateOne {
	aauo.mutation.AddUpdatedBy(i)
	return aauo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (aauo *AppActionUpdateOne) ClearUpdatedBy() *AppActionUpdateOne {
	aauo.mutation.ClearUpdatedBy()
	return aauo
}

// SetUpdatedAt sets the "updated_at" field.
func (aauo *AppActionUpdateOne) SetUpdatedAt(t time.Time) *AppActionUpdateOne {
	aauo.mutation.SetUpdatedAt(t)
	return aauo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aauo *AppActionUpdateOne) SetNillableUpdatedAt(t *time.Time) *AppActionUpdateOne {
	if t != nil {
		aauo.SetUpdatedAt(*t)
	}
	return aauo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (aauo *AppActionUpdateOne) ClearUpdatedAt() *AppActionUpdateOne {
	aauo.mutation.ClearUpdatedAt()
	return aauo
}

// SetAppID sets the "app_id" field.
func (aauo *AppActionUpdateOne) SetAppID(i int) *AppActionUpdateOne {
	aauo.mutation.SetAppID(i)
	return aauo
}

// SetName sets the "name" field.
func (aauo *AppActionUpdateOne) SetName(s string) *AppActionUpdateOne {
	aauo.mutation.SetName(s)
	return aauo
}

// SetKind sets the "kind" field.
func (aauo *AppActionUpdateOne) SetKind(a appaction.Kind) *AppActionUpdateOne {
	aauo.mutation.SetKind(a)
	return aauo
}

// SetMethod sets the "method" field.
func (aauo *AppActionUpdateOne) SetMethod(a appaction.Method) *AppActionUpdateOne {
	aauo.mutation.SetMethod(a)
	return aauo
}

// SetComments sets the "comments" field.
func (aauo *AppActionUpdateOne) SetComments(s string) *AppActionUpdateOne {
	aauo.mutation.SetComments(s)
	return aauo
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (aauo *AppActionUpdateOne) SetNillableComments(s *string) *AppActionUpdateOne {
	if s != nil {
		aauo.SetComments(*s)
	}
	return aauo
}

// ClearComments clears the value of the "comments" field.
func (aauo *AppActionUpdateOne) ClearComments() *AppActionUpdateOne {
	aauo.mutation.ClearComments()
	return aauo
}

// SetApp sets the "app" edge to the App entity.
func (aauo *AppActionUpdateOne) SetApp(a *App) *AppActionUpdateOne {
	return aauo.SetAppID(a.ID)
}

// AddMenuIDs adds the "menus" edge to the AppMenu entity by IDs.
func (aauo *AppActionUpdateOne) AddMenuIDs(ids ...int) *AppActionUpdateOne {
	aauo.mutation.AddMenuIDs(ids...)
	return aauo
}

// AddMenus adds the "menus" edges to the AppMenu entity.
func (aauo *AppActionUpdateOne) AddMenus(a ...*AppMenu) *AppActionUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aauo.AddMenuIDs(ids...)
}

// AddResourceIDs adds the "resources" edge to the AppRes entity by IDs.
func (aauo *AppActionUpdateOne) AddResourceIDs(ids ...int) *AppActionUpdateOne {
	aauo.mutation.AddResourceIDs(ids...)
	return aauo
}

// AddResources adds the "resources" edges to the AppRes entity.
func (aauo *AppActionUpdateOne) AddResources(a ...*AppRes) *AppActionUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aauo.AddResourceIDs(ids...)
}

// Mutation returns the AppActionMutation object of the builder.
func (aauo *AppActionUpdateOne) Mutation() *AppActionMutation {
	return aauo.mutation
}

// ClearApp clears the "app" edge to the App entity.
func (aauo *AppActionUpdateOne) ClearApp() *AppActionUpdateOne {
	aauo.mutation.ClearApp()
	return aauo
}

// ClearMenus clears all "menus" edges to the AppMenu entity.
func (aauo *AppActionUpdateOne) ClearMenus() *AppActionUpdateOne {
	aauo.mutation.ClearMenus()
	return aauo
}

// RemoveMenuIDs removes the "menus" edge to AppMenu entities by IDs.
func (aauo *AppActionUpdateOne) RemoveMenuIDs(ids ...int) *AppActionUpdateOne {
	aauo.mutation.RemoveMenuIDs(ids...)
	return aauo
}

// RemoveMenus removes "menus" edges to AppMenu entities.
func (aauo *AppActionUpdateOne) RemoveMenus(a ...*AppMenu) *AppActionUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aauo.RemoveMenuIDs(ids...)
}

// ClearResources clears all "resources" edges to the AppRes entity.
func (aauo *AppActionUpdateOne) ClearResources() *AppActionUpdateOne {
	aauo.mutation.ClearResources()
	return aauo
}

// RemoveResourceIDs removes the "resources" edge to AppRes entities by IDs.
func (aauo *AppActionUpdateOne) RemoveResourceIDs(ids ...int) *AppActionUpdateOne {
	aauo.mutation.RemoveResourceIDs(ids...)
	return aauo
}

// RemoveResources removes "resources" edges to AppRes entities.
func (aauo *AppActionUpdateOne) RemoveResources(a ...*AppRes) *AppActionUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aauo.RemoveResourceIDs(ids...)
}

// Where appends a list predicates to the AppActionUpdate builder.
func (aauo *AppActionUpdateOne) Where(ps ...predicate.AppAction) *AppActionUpdateOne {
	aauo.mutation.Where(ps...)
	return aauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aauo *AppActionUpdateOne) Select(field string, fields ...string) *AppActionUpdateOne {
	aauo.fields = append([]string{field}, fields...)
	return aauo
}

// Save executes the query and returns the updated AppAction entity.
func (aauo *AppActionUpdateOne) Save(ctx context.Context) (*AppAction, error) {
	return withHooks[*AppAction, AppActionMutation](ctx, aauo.sqlSave, aauo.mutation, aauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aauo *AppActionUpdateOne) SaveX(ctx context.Context) *AppAction {
	node, err := aauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aauo *AppActionUpdateOne) Exec(ctx context.Context) error {
	_, err := aauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aauo *AppActionUpdateOne) ExecX(ctx context.Context) {
	if err := aauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aauo *AppActionUpdateOne) check() error {
	if v, ok := aauo.mutation.Name(); ok {
		if err := appaction.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "AppAction.name": %w`, err)}
		}
	}
	if v, ok := aauo.mutation.Kind(); ok {
		if err := appaction.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "AppAction.kind": %w`, err)}
		}
	}
	if v, ok := aauo.mutation.Method(); ok {
		if err := appaction.MethodValidator(v); err != nil {
			return &ValidationError{Name: "method", err: fmt.Errorf(`ent: validator failed for field "AppAction.method": %w`, err)}
		}
	}
	if _, ok := aauo.mutation.AppID(); aauo.mutation.AppCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AppAction.app"`)
	}
	return nil
}

func (aauo *AppActionUpdateOne) sqlSave(ctx context.Context) (_node *AppAction, err error) {
	if err := aauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(appaction.Table, appaction.Columns, sqlgraph.NewFieldSpec(appaction.FieldID, field.TypeInt))
	id, ok := aauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppAction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appaction.FieldID)
		for _, f := range fields {
			if !appaction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aauo.mutation.UpdatedBy(); ok {
		_spec.SetField(appaction.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := aauo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(appaction.FieldUpdatedBy, field.TypeInt, value)
	}
	if aauo.mutation.UpdatedByCleared() {
		_spec.ClearField(appaction.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := aauo.mutation.UpdatedAt(); ok {
		_spec.SetField(appaction.FieldUpdatedAt, field.TypeTime, value)
	}
	if aauo.mutation.UpdatedAtCleared() {
		_spec.ClearField(appaction.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := aauo.mutation.Name(); ok {
		_spec.SetField(appaction.FieldName, field.TypeString, value)
	}
	if value, ok := aauo.mutation.Kind(); ok {
		_spec.SetField(appaction.FieldKind, field.TypeEnum, value)
	}
	if value, ok := aauo.mutation.Method(); ok {
		_spec.SetField(appaction.FieldMethod, field.TypeEnum, value)
	}
	if value, ok := aauo.mutation.Comments(); ok {
		_spec.SetField(appaction.FieldComments, field.TypeString, value)
	}
	if aauo.mutation.CommentsCleared() {
		_spec.ClearField(appaction.FieldComments, field.TypeString)
	}
	if aauo.mutation.AppCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appaction.AppTable,
			Columns: []string{appaction.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aauo.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appaction.AppTable,
			Columns: []string{appaction.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aauo.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appaction.MenusTable,
			Columns: []string{appaction.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appmenu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aauo.mutation.RemovedMenusIDs(); len(nodes) > 0 && !aauo.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appaction.MenusTable,
			Columns: []string{appaction.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appmenu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aauo.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appaction.MenusTable,
			Columns: []string{appaction.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appmenu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aauo.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appaction.ResourcesTable,
			Columns: []string{appaction.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appres.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aauo.mutation.RemovedResourcesIDs(); len(nodes) > 0 && !aauo.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appaction.ResourcesTable,
			Columns: []string{appaction.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appres.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aauo.mutation.ResourcesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appaction.ResourcesTable,
			Columns: []string{appaction.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appres.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AppAction{config: aauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aauo.mutation.done = true
	return _node, nil
}