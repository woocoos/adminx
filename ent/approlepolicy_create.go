// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/adminx/ent/apppolicy"
	"github.com/woocoos/adminx/ent/approle"
	"github.com/woocoos/adminx/ent/approlepolicy"
)

// AppRolePolicyCreate is the builder for creating a AppRolePolicy entity.
type AppRolePolicyCreate struct {
	config
	mutation *AppRolePolicyMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (arpc *AppRolePolicyCreate) SetCreatedBy(i int) *AppRolePolicyCreate {
	arpc.mutation.SetCreatedBy(i)
	return arpc
}

// SetCreatedAt sets the "created_at" field.
func (arpc *AppRolePolicyCreate) SetCreatedAt(t time.Time) *AppRolePolicyCreate {
	arpc.mutation.SetCreatedAt(t)
	return arpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (arpc *AppRolePolicyCreate) SetNillableCreatedAt(t *time.Time) *AppRolePolicyCreate {
	if t != nil {
		arpc.SetCreatedAt(*t)
	}
	return arpc
}

// SetUpdatedBy sets the "updated_by" field.
func (arpc *AppRolePolicyCreate) SetUpdatedBy(i int) *AppRolePolicyCreate {
	arpc.mutation.SetUpdatedBy(i)
	return arpc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (arpc *AppRolePolicyCreate) SetNillableUpdatedBy(i *int) *AppRolePolicyCreate {
	if i != nil {
		arpc.SetUpdatedBy(*i)
	}
	return arpc
}

// SetUpdatedAt sets the "updated_at" field.
func (arpc *AppRolePolicyCreate) SetUpdatedAt(t time.Time) *AppRolePolicyCreate {
	arpc.mutation.SetUpdatedAt(t)
	return arpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (arpc *AppRolePolicyCreate) SetNillableUpdatedAt(t *time.Time) *AppRolePolicyCreate {
	if t != nil {
		arpc.SetUpdatedAt(*t)
	}
	return arpc
}

// SetRoleID sets the "role_id" field.
func (arpc *AppRolePolicyCreate) SetRoleID(i int) *AppRolePolicyCreate {
	arpc.mutation.SetRoleID(i)
	return arpc
}

// SetPolicyID sets the "policy_id" field.
func (arpc *AppRolePolicyCreate) SetPolicyID(i int) *AppRolePolicyCreate {
	arpc.mutation.SetPolicyID(i)
	return arpc
}

// SetRole sets the "role" edge to the AppRole entity.
func (arpc *AppRolePolicyCreate) SetRole(a *AppRole) *AppRolePolicyCreate {
	return arpc.SetRoleID(a.ID)
}

// SetPolicy sets the "policy" edge to the AppPolicy entity.
func (arpc *AppRolePolicyCreate) SetPolicy(a *AppPolicy) *AppRolePolicyCreate {
	return arpc.SetPolicyID(a.ID)
}

// Mutation returns the AppRolePolicyMutation object of the builder.
func (arpc *AppRolePolicyCreate) Mutation() *AppRolePolicyMutation {
	return arpc.mutation
}

// Save creates the AppRolePolicy in the database.
func (arpc *AppRolePolicyCreate) Save(ctx context.Context) (*AppRolePolicy, error) {
	if err := arpc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*AppRolePolicy, AppRolePolicyMutation](ctx, arpc.sqlSave, arpc.mutation, arpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (arpc *AppRolePolicyCreate) SaveX(ctx context.Context) *AppRolePolicy {
	v, err := arpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arpc *AppRolePolicyCreate) Exec(ctx context.Context) error {
	_, err := arpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arpc *AppRolePolicyCreate) ExecX(ctx context.Context) {
	if err := arpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (arpc *AppRolePolicyCreate) defaults() error {
	if _, ok := arpc.mutation.CreatedAt(); !ok {
		if approlepolicy.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized approlepolicy.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := approlepolicy.DefaultCreatedAt()
		arpc.mutation.SetCreatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (arpc *AppRolePolicyCreate) check() error {
	if _, ok := arpc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "AppRolePolicy.created_by"`)}
	}
	if _, ok := arpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppRolePolicy.created_at"`)}
	}
	if _, ok := arpc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role_id", err: errors.New(`ent: missing required field "AppRolePolicy.role_id"`)}
	}
	if _, ok := arpc.mutation.PolicyID(); !ok {
		return &ValidationError{Name: "policy_id", err: errors.New(`ent: missing required field "AppRolePolicy.policy_id"`)}
	}
	if _, ok := arpc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required edge "AppRolePolicy.role"`)}
	}
	if _, ok := arpc.mutation.PolicyID(); !ok {
		return &ValidationError{Name: "policy", err: errors.New(`ent: missing required edge "AppRolePolicy.policy"`)}
	}
	return nil
}

func (arpc *AppRolePolicyCreate) sqlSave(ctx context.Context) (*AppRolePolicy, error) {
	if err := arpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := arpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, arpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (arpc *AppRolePolicyCreate) createSpec() (*AppRolePolicy, *sqlgraph.CreateSpec) {
	var (
		_node = &AppRolePolicy{config: arpc.config}
		_spec = sqlgraph.NewCreateSpec(approlepolicy.Table, nil)
	)
	if value, ok := arpc.mutation.CreatedBy(); ok {
		_spec.SetField(approlepolicy.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := arpc.mutation.CreatedAt(); ok {
		_spec.SetField(approlepolicy.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := arpc.mutation.UpdatedBy(); ok {
		_spec.SetField(approlepolicy.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := arpc.mutation.UpdatedAt(); ok {
		_spec.SetField(approlepolicy.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := arpc.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   approlepolicy.RoleTable,
			Columns: []string{approlepolicy.RoleColumn},
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
		_node.RoleID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := arpc.mutation.PolicyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   approlepolicy.PolicyTable,
			Columns: []string{approlepolicy.PolicyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: apppolicy.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PolicyID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AppRolePolicyCreateBulk is the builder for creating many AppRolePolicy entities in bulk.
type AppRolePolicyCreateBulk struct {
	config
	builders []*AppRolePolicyCreate
}

// Save creates the AppRolePolicy entities in the database.
func (arpcb *AppRolePolicyCreateBulk) Save(ctx context.Context) ([]*AppRolePolicy, error) {
	specs := make([]*sqlgraph.CreateSpec, len(arpcb.builders))
	nodes := make([]*AppRolePolicy, len(arpcb.builders))
	mutators := make([]Mutator, len(arpcb.builders))
	for i := range arpcb.builders {
		func(i int, root context.Context) {
			builder := arpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppRolePolicyMutation)
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
					_, err = mutators[i+1].Mutate(root, arpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, arpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
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
		if _, err := mutators[0].Mutate(ctx, arpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (arpcb *AppRolePolicyCreateBulk) SaveX(ctx context.Context) []*AppRolePolicy {
	v, err := arpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arpcb *AppRolePolicyCreateBulk) Exec(ctx context.Context) error {
	_, err := arpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arpcb *AppRolePolicyCreateBulk) ExecX(ctx context.Context) {
	if err := arpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
