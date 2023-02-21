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
	"github.com/woocoos/adminx/ent/apppolicy"
	"github.com/woocoos/adminx/ent/approle"
	"github.com/woocoos/adminx/graph/entgen/types"
)

// AppPolicyCreate is the builder for creating a AppPolicy entity.
type AppPolicyCreate struct {
	config
	mutation *AppPolicyMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (apc *AppPolicyCreate) SetCreatedBy(i int) *AppPolicyCreate {
	apc.mutation.SetCreatedBy(i)
	return apc
}

// SetCreatedAt sets the "created_at" field.
func (apc *AppPolicyCreate) SetCreatedAt(t time.Time) *AppPolicyCreate {
	apc.mutation.SetCreatedAt(t)
	return apc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (apc *AppPolicyCreate) SetNillableCreatedAt(t *time.Time) *AppPolicyCreate {
	if t != nil {
		apc.SetCreatedAt(*t)
	}
	return apc
}

// SetUpdatedBy sets the "updated_by" field.
func (apc *AppPolicyCreate) SetUpdatedBy(i int) *AppPolicyCreate {
	apc.mutation.SetUpdatedBy(i)
	return apc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (apc *AppPolicyCreate) SetNillableUpdatedBy(i *int) *AppPolicyCreate {
	if i != nil {
		apc.SetUpdatedBy(*i)
	}
	return apc
}

// SetUpdatedAt sets the "updated_at" field.
func (apc *AppPolicyCreate) SetUpdatedAt(t time.Time) *AppPolicyCreate {
	apc.mutation.SetUpdatedAt(t)
	return apc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (apc *AppPolicyCreate) SetNillableUpdatedAt(t *time.Time) *AppPolicyCreate {
	if t != nil {
		apc.SetUpdatedAt(*t)
	}
	return apc
}

// SetAppID sets the "app_id" field.
func (apc *AppPolicyCreate) SetAppID(i int) *AppPolicyCreate {
	apc.mutation.SetAppID(i)
	return apc
}

// SetName sets the "name" field.
func (apc *AppPolicyCreate) SetName(s string) *AppPolicyCreate {
	apc.mutation.SetName(s)
	return apc
}

// SetComments sets the "comments" field.
func (apc *AppPolicyCreate) SetComments(s string) *AppPolicyCreate {
	apc.mutation.SetComments(s)
	return apc
}

// SetRules sets the "rules" field.
func (apc *AppPolicyCreate) SetRules(tr []types.PolicyRule) *AppPolicyCreate {
	apc.mutation.SetRules(tr)
	return apc
}

// SetVersion sets the "version" field.
func (apc *AppPolicyCreate) SetVersion(s string) *AppPolicyCreate {
	apc.mutation.SetVersion(s)
	return apc
}

// SetAutoGrant sets the "auto_grant" field.
func (apc *AppPolicyCreate) SetAutoGrant(b bool) *AppPolicyCreate {
	apc.mutation.SetAutoGrant(b)
	return apc
}

// SetNillableAutoGrant sets the "auto_grant" field if the given value is not nil.
func (apc *AppPolicyCreate) SetNillableAutoGrant(b *bool) *AppPolicyCreate {
	if b != nil {
		apc.SetAutoGrant(*b)
	}
	return apc
}

// SetStatus sets the "status" field.
func (apc *AppPolicyCreate) SetStatus(ts types.SimpleStatus) *AppPolicyCreate {
	apc.mutation.SetStatus(ts)
	return apc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (apc *AppPolicyCreate) SetNillableStatus(ts *types.SimpleStatus) *AppPolicyCreate {
	if ts != nil {
		apc.SetStatus(*ts)
	}
	return apc
}

// SetID sets the "id" field.
func (apc *AppPolicyCreate) SetID(i int) *AppPolicyCreate {
	apc.mutation.SetID(i)
	return apc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (apc *AppPolicyCreate) SetNillableID(i *int) *AppPolicyCreate {
	if i != nil {
		apc.SetID(*i)
	}
	return apc
}

// SetApp sets the "app" edge to the App entity.
func (apc *AppPolicyCreate) SetApp(a *App) *AppPolicyCreate {
	return apc.SetAppID(a.ID)
}

// AddRoleIDs adds the "roles" edge to the AppRole entity by IDs.
func (apc *AppPolicyCreate) AddRoleIDs(ids ...int) *AppPolicyCreate {
	apc.mutation.AddRoleIDs(ids...)
	return apc
}

// AddRoles adds the "roles" edges to the AppRole entity.
func (apc *AppPolicyCreate) AddRoles(a ...*AppRole) *AppPolicyCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apc.AddRoleIDs(ids...)
}

// Mutation returns the AppPolicyMutation object of the builder.
func (apc *AppPolicyCreate) Mutation() *AppPolicyMutation {
	return apc.mutation
}

// Save creates the AppPolicy in the database.
func (apc *AppPolicyCreate) Save(ctx context.Context) (*AppPolicy, error) {
	if err := apc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*AppPolicy, AppPolicyMutation](ctx, apc.sqlSave, apc.mutation, apc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (apc *AppPolicyCreate) SaveX(ctx context.Context) *AppPolicy {
	v, err := apc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (apc *AppPolicyCreate) Exec(ctx context.Context) error {
	_, err := apc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apc *AppPolicyCreate) ExecX(ctx context.Context) {
	if err := apc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (apc *AppPolicyCreate) defaults() error {
	if _, ok := apc.mutation.CreatedAt(); !ok {
		if apppolicy.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized apppolicy.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := apppolicy.DefaultCreatedAt()
		apc.mutation.SetCreatedAt(v)
	}
	if _, ok := apc.mutation.AutoGrant(); !ok {
		v := apppolicy.DefaultAutoGrant
		apc.mutation.SetAutoGrant(v)
	}
	if _, ok := apc.mutation.Status(); !ok {
		v := apppolicy.DefaultStatus
		apc.mutation.SetStatus(v)
	}
	if _, ok := apc.mutation.ID(); !ok {
		if apppolicy.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized apppolicy.DefaultID (forgotten import ent/runtime?)")
		}
		v := apppolicy.DefaultID()
		apc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (apc *AppPolicyCreate) check() error {
	if _, ok := apc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "AppPolicy.created_by"`)}
	}
	if _, ok := apc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppPolicy.created_at"`)}
	}
	if _, ok := apc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "AppPolicy.app_id"`)}
	}
	if _, ok := apc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "AppPolicy.name"`)}
	}
	if _, ok := apc.mutation.Comments(); !ok {
		return &ValidationError{Name: "comments", err: errors.New(`ent: missing required field "AppPolicy.comments"`)}
	}
	if _, ok := apc.mutation.Rules(); !ok {
		return &ValidationError{Name: "rules", err: errors.New(`ent: missing required field "AppPolicy.rules"`)}
	}
	if _, ok := apc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "AppPolicy.version"`)}
	}
	if _, ok := apc.mutation.AutoGrant(); !ok {
		return &ValidationError{Name: "auto_grant", err: errors.New(`ent: missing required field "AppPolicy.auto_grant"`)}
	}
	if v, ok := apc.mutation.Status(); ok {
		if err := apppolicy.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "AppPolicy.status": %w`, err)}
		}
	}
	if _, ok := apc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app", err: errors.New(`ent: missing required edge "AppPolicy.app"`)}
	}
	return nil
}

func (apc *AppPolicyCreate) sqlSave(ctx context.Context) (*AppPolicy, error) {
	if err := apc.check(); err != nil {
		return nil, err
	}
	_node, _spec := apc.createSpec()
	if err := sqlgraph.CreateNode(ctx, apc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	apc.mutation.id = &_node.ID
	apc.mutation.done = true
	return _node, nil
}

func (apc *AppPolicyCreate) createSpec() (*AppPolicy, *sqlgraph.CreateSpec) {
	var (
		_node = &AppPolicy{config: apc.config}
		_spec = sqlgraph.NewCreateSpec(apppolicy.Table, sqlgraph.NewFieldSpec(apppolicy.FieldID, field.TypeInt))
	)
	if id, ok := apc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := apc.mutation.CreatedBy(); ok {
		_spec.SetField(apppolicy.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := apc.mutation.CreatedAt(); ok {
		_spec.SetField(apppolicy.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := apc.mutation.UpdatedBy(); ok {
		_spec.SetField(apppolicy.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := apc.mutation.UpdatedAt(); ok {
		_spec.SetField(apppolicy.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := apc.mutation.Name(); ok {
		_spec.SetField(apppolicy.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := apc.mutation.Comments(); ok {
		_spec.SetField(apppolicy.FieldComments, field.TypeString, value)
		_node.Comments = value
	}
	if value, ok := apc.mutation.Rules(); ok {
		_spec.SetField(apppolicy.FieldRules, field.TypeJSON, value)
		_node.Rules = value
	}
	if value, ok := apc.mutation.Version(); ok {
		_spec.SetField(apppolicy.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := apc.mutation.AutoGrant(); ok {
		_spec.SetField(apppolicy.FieldAutoGrant, field.TypeBool, value)
		_node.AutoGrant = value
	}
	if value, ok := apc.mutation.Status(); ok {
		_spec.SetField(apppolicy.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := apc.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apppolicy.AppTable,
			Columns: []string{apppolicy.AppColumn},
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
		_node.AppID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := apc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   apppolicy.RolesTable,
			Columns: apppolicy.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: approle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AppRolePolicyCreate{config: apc.config, mutation: newAppRolePolicyMutation(apc.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AppPolicyCreateBulk is the builder for creating many AppPolicy entities in bulk.
type AppPolicyCreateBulk struct {
	config
	builders []*AppPolicyCreate
}

// Save creates the AppPolicy entities in the database.
func (apcb *AppPolicyCreateBulk) Save(ctx context.Context) ([]*AppPolicy, error) {
	specs := make([]*sqlgraph.CreateSpec, len(apcb.builders))
	nodes := make([]*AppPolicy, len(apcb.builders))
	mutators := make([]Mutator, len(apcb.builders))
	for i := range apcb.builders {
		func(i int, root context.Context) {
			builder := apcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppPolicyMutation)
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
					_, err = mutators[i+1].Mutate(root, apcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, apcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, apcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (apcb *AppPolicyCreateBulk) SaveX(ctx context.Context) []*AppPolicy {
	v, err := apcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (apcb *AppPolicyCreateBulk) Exec(ctx context.Context) error {
	_, err := apcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apcb *AppPolicyCreateBulk) ExecX(ctx context.Context) {
	if err := apcb.Exec(ctx); err != nil {
		panic(err)
	}
}