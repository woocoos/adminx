// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/adminx/ent/app"
	"github.com/woocoos/adminx/ent/appmenu"
	"github.com/woocoos/adminx/ent/apppermission"
)

// AppCreate is the builder for creating a App entity.
type AppCreate struct {
	config
	mutation *AppMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (ac *AppCreate) SetCreatedBy(i int) *AppCreate {
	ac.mutation.SetCreatedBy(i)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AppCreate) SetCreatedAt(t time.Time) *AppCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AppCreate) SetNillableCreatedAt(t *time.Time) *AppCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedBy sets the "updated_by" field.
func (ac *AppCreate) SetUpdatedBy(i int) *AppCreate {
	ac.mutation.SetUpdatedBy(i)
	return ac
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ac *AppCreate) SetNillableUpdatedBy(i *int) *AppCreate {
	if i != nil {
		ac.SetUpdatedBy(*i)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AppCreate) SetUpdatedAt(t time.Time) *AppCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AppCreate) SetNillableUpdatedAt(t *time.Time) *AppCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetName sets the "name" field.
func (ac *AppCreate) SetName(s string) *AppCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetCode sets the "code" field.
func (ac *AppCreate) SetCode(s string) *AppCreate {
	ac.mutation.SetCode(s)
	return ac
}

// SetKind sets the "kind" field.
func (ac *AppCreate) SetKind(a app.Kind) *AppCreate {
	ac.mutation.SetKind(a)
	return ac
}

// SetRedirectURI sets the "redirect_uri" field.
func (ac *AppCreate) SetRedirectURI(s string) *AppCreate {
	ac.mutation.SetRedirectURI(s)
	return ac
}

// SetNillableRedirectURI sets the "redirect_uri" field if the given value is not nil.
func (ac *AppCreate) SetNillableRedirectURI(s *string) *AppCreate {
	if s != nil {
		ac.SetRedirectURI(*s)
	}
	return ac
}

// SetAppKey sets the "app_key" field.
func (ac *AppCreate) SetAppKey(s string) *AppCreate {
	ac.mutation.SetAppKey(s)
	return ac
}

// SetNillableAppKey sets the "app_key" field if the given value is not nil.
func (ac *AppCreate) SetNillableAppKey(s *string) *AppCreate {
	if s != nil {
		ac.SetAppKey(*s)
	}
	return ac
}

// SetAppSecret sets the "app_secret" field.
func (ac *AppCreate) SetAppSecret(s string) *AppCreate {
	ac.mutation.SetAppSecret(s)
	return ac
}

// SetNillableAppSecret sets the "app_secret" field if the given value is not nil.
func (ac *AppCreate) SetNillableAppSecret(s *string) *AppCreate {
	if s != nil {
		ac.SetAppSecret(*s)
	}
	return ac
}

// SetScopes sets the "scopes" field.
func (ac *AppCreate) SetScopes(s string) *AppCreate {
	ac.mutation.SetScopes(s)
	return ac
}

// SetNillableScopes sets the "scopes" field if the given value is not nil.
func (ac *AppCreate) SetNillableScopes(s *string) *AppCreate {
	if s != nil {
		ac.SetScopes(*s)
	}
	return ac
}

// SetTokenValidity sets the "token_validity" field.
func (ac *AppCreate) SetTokenValidity(i int32) *AppCreate {
	ac.mutation.SetTokenValidity(i)
	return ac
}

// SetNillableTokenValidity sets the "token_validity" field if the given value is not nil.
func (ac *AppCreate) SetNillableTokenValidity(i *int32) *AppCreate {
	if i != nil {
		ac.SetTokenValidity(*i)
	}
	return ac
}

// SetRefreshTokenValidity sets the "refresh_token_validity" field.
func (ac *AppCreate) SetRefreshTokenValidity(i int32) *AppCreate {
	ac.mutation.SetRefreshTokenValidity(i)
	return ac
}

// SetNillableRefreshTokenValidity sets the "refresh_token_validity" field if the given value is not nil.
func (ac *AppCreate) SetNillableRefreshTokenValidity(i *int32) *AppCreate {
	if i != nil {
		ac.SetRefreshTokenValidity(*i)
	}
	return ac
}

// SetLogo sets the "logo" field.
func (ac *AppCreate) SetLogo(s string) *AppCreate {
	ac.mutation.SetLogo(s)
	return ac
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (ac *AppCreate) SetNillableLogo(s *string) *AppCreate {
	if s != nil {
		ac.SetLogo(*s)
	}
	return ac
}

// SetComments sets the "comments" field.
func (ac *AppCreate) SetComments(s string) *AppCreate {
	ac.mutation.SetComments(s)
	return ac
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (ac *AppCreate) SetNillableComments(s *string) *AppCreate {
	if s != nil {
		ac.SetComments(*s)
	}
	return ac
}

// SetStatus sets the "status" field.
func (ac *AppCreate) SetStatus(a app.Status) *AppCreate {
	ac.mutation.SetStatus(a)
	return ac
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ac *AppCreate) SetNillableStatus(a *app.Status) *AppCreate {
	if a != nil {
		ac.SetStatus(*a)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AppCreate) SetID(i int) *AppCreate {
	ac.mutation.SetID(i)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AppCreate) SetNillableID(i *int) *AppCreate {
	if i != nil {
		ac.SetID(*i)
	}
	return ac
}

// AddMenuIDs adds the "menus" edge to the AppMenu entity by IDs.
func (ac *AppCreate) AddMenuIDs(ids ...int) *AppCreate {
	ac.mutation.AddMenuIDs(ids...)
	return ac
}

// AddMenus adds the "menus" edges to the AppMenu entity.
func (ac *AppCreate) AddMenus(a ...*AppMenu) *AppCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddMenuIDs(ids...)
}

// AddPermissionIDs adds the "permissions" edge to the AppPermission entity by IDs.
func (ac *AppCreate) AddPermissionIDs(ids ...int) *AppCreate {
	ac.mutation.AddPermissionIDs(ids...)
	return ac
}

// AddPermissions adds the "permissions" edges to the AppPermission entity.
func (ac *AppCreate) AddPermissions(a ...*AppPermission) *AppCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddPermissionIDs(ids...)
}

// Mutation returns the AppMutation object of the builder.
func (ac *AppCreate) Mutation() *AppMutation {
	return ac.mutation
}

// Save creates the App in the database.
func (ac *AppCreate) Save(ctx context.Context) (*App, error) {
	if err := ac.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*App, AppMutation](ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AppCreate) SaveX(ctx context.Context) *App {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AppCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AppCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AppCreate) defaults() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		if app.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized app.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := app.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		if app.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized app.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := app.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		if app.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized app.DefaultID (forgotten import ent/runtime?)")
		}
		v := app.DefaultID()
		ac.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ac *AppCreate) check() error {
	if _, ok := ac.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "App.created_by"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "App.created_at"`)}
	}
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "App.name"`)}
	}
	if v, ok := ac.mutation.Name(); ok {
		if err := app.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "App.name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "App.code"`)}
	}
	if v, ok := ac.mutation.Code(); ok {
		if err := app.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "App.code": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New(`ent: missing required field "App.kind"`)}
	}
	if v, ok := ac.mutation.Kind(); ok {
		if err := app.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "App.kind": %w`, err)}
		}
	}
	if v, ok := ac.mutation.RedirectURI(); ok {
		if err := app.RedirectURIValidator(v); err != nil {
			return &ValidationError{Name: "redirect_uri", err: fmt.Errorf(`ent: validator failed for field "App.redirect_uri": %w`, err)}
		}
	}
	if v, ok := ac.mutation.AppSecret(); ok {
		if err := app.AppSecretValidator(v); err != nil {
			return &ValidationError{Name: "app_secret", err: fmt.Errorf(`ent: validator failed for field "App.app_secret": %w`, err)}
		}
	}
	if v, ok := ac.mutation.Scopes(); ok {
		if err := app.ScopesValidator(v); err != nil {
			return &ValidationError{Name: "scopes", err: fmt.Errorf(`ent: validator failed for field "App.scopes": %w`, err)}
		}
	}
	if v, ok := ac.mutation.Status(); ok {
		if err := app.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "App.status": %w`, err)}
		}
	}
	return nil
}

func (ac *AppCreate) sqlSave(ctx context.Context) (*App, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AppCreate) createSpec() (*App, *sqlgraph.CreateSpec) {
	var (
		_node = &App{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(app.Table, sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.CreatedBy(); ok {
		_spec.SetField(app.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(app.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedBy(); ok {
		_spec.SetField(app.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(app.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(app.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.Code(); ok {
		_spec.SetField(app.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := ac.mutation.Kind(); ok {
		_spec.SetField(app.FieldKind, field.TypeEnum, value)
		_node.Kind = value
	}
	if value, ok := ac.mutation.RedirectURI(); ok {
		_spec.SetField(app.FieldRedirectURI, field.TypeString, value)
		_node.RedirectURI = value
	}
	if value, ok := ac.mutation.AppKey(); ok {
		_spec.SetField(app.FieldAppKey, field.TypeString, value)
		_node.AppKey = value
	}
	if value, ok := ac.mutation.AppSecret(); ok {
		_spec.SetField(app.FieldAppSecret, field.TypeString, value)
		_node.AppSecret = value
	}
	if value, ok := ac.mutation.Scopes(); ok {
		_spec.SetField(app.FieldScopes, field.TypeString, value)
		_node.Scopes = value
	}
	if value, ok := ac.mutation.TokenValidity(); ok {
		_spec.SetField(app.FieldTokenValidity, field.TypeInt32, value)
		_node.TokenValidity = value
	}
	if value, ok := ac.mutation.RefreshTokenValidity(); ok {
		_spec.SetField(app.FieldRefreshTokenValidity, field.TypeInt32, value)
		_node.RefreshTokenValidity = value
	}
	if value, ok := ac.mutation.Logo(); ok {
		_spec.SetField(app.FieldLogo, field.TypeString, value)
		_node.Logo = value
	}
	if value, ok := ac.mutation.Comments(); ok {
		_spec.SetField(app.FieldComments, field.TypeString, value)
		_node.Comments = value
	}
	if value, ok := ac.mutation.Status(); ok {
		_spec.SetField(app.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := ac.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   app.MenusTable,
			Columns: []string{app.MenusColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   app.PermissionsTable,
			Columns: []string{app.PermissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: apppermission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AppCreateBulk is the builder for creating many App entities in bulk.
type AppCreateBulk struct {
	config
	builders []*AppCreate
}

// Save creates the App entities in the database.
func (acb *AppCreateBulk) Save(ctx context.Context) ([]*App, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*App, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AppCreateBulk) SaveX(ctx context.Context) []*App {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AppCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AppCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
