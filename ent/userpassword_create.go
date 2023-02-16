// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/adminx/ent/user"
	"github.com/woocoos/adminx/ent/userpassword"
)

// UserPasswordCreate is the builder for creating a UserPassword entity.
type UserPasswordCreate struct {
	config
	mutation *UserPasswordMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (upc *UserPasswordCreate) SetCreatedBy(i int) *UserPasswordCreate {
	upc.mutation.SetCreatedBy(i)
	return upc
}

// SetCreatedAt sets the "created_at" field.
func (upc *UserPasswordCreate) SetCreatedAt(t time.Time) *UserPasswordCreate {
	upc.mutation.SetCreatedAt(t)
	return upc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (upc *UserPasswordCreate) SetNillableCreatedAt(t *time.Time) *UserPasswordCreate {
	if t != nil {
		upc.SetCreatedAt(*t)
	}
	return upc
}

// SetUpdatedBy sets the "updated_by" field.
func (upc *UserPasswordCreate) SetUpdatedBy(i int) *UserPasswordCreate {
	upc.mutation.SetUpdatedBy(i)
	return upc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (upc *UserPasswordCreate) SetNillableUpdatedBy(i *int) *UserPasswordCreate {
	if i != nil {
		upc.SetUpdatedBy(*i)
	}
	return upc
}

// SetUpdatedAt sets the "updated_at" field.
func (upc *UserPasswordCreate) SetUpdatedAt(t time.Time) *UserPasswordCreate {
	upc.mutation.SetUpdatedAt(t)
	return upc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (upc *UserPasswordCreate) SetNillableUpdatedAt(t *time.Time) *UserPasswordCreate {
	if t != nil {
		upc.SetUpdatedAt(*t)
	}
	return upc
}

// SetUserID sets the "user_id" field.
func (upc *UserPasswordCreate) SetUserID(i int) *UserPasswordCreate {
	upc.mutation.SetUserID(i)
	return upc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (upc *UserPasswordCreate) SetNillableUserID(i *int) *UserPasswordCreate {
	if i != nil {
		upc.SetUserID(*i)
	}
	return upc
}

// SetScene sets the "scene" field.
func (upc *UserPasswordCreate) SetScene(u userpassword.Scene) *UserPasswordCreate {
	upc.mutation.SetScene(u)
	return upc
}

// SetNillableScene sets the "scene" field if the given value is not nil.
func (upc *UserPasswordCreate) SetNillableScene(u *userpassword.Scene) *UserPasswordCreate {
	if u != nil {
		upc.SetScene(*u)
	}
	return upc
}

// SetPassword sets the "password" field.
func (upc *UserPasswordCreate) SetPassword(s string) *UserPasswordCreate {
	upc.mutation.SetPassword(s)
	return upc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (upc *UserPasswordCreate) SetNillablePassword(s *string) *UserPasswordCreate {
	if s != nil {
		upc.SetPassword(*s)
	}
	return upc
}

// SetSalt sets the "salt" field.
func (upc *UserPasswordCreate) SetSalt(s string) *UserPasswordCreate {
	upc.mutation.SetSalt(s)
	return upc
}

// SetStatus sets the "status" field.
func (upc *UserPasswordCreate) SetStatus(u userpassword.Status) *UserPasswordCreate {
	upc.mutation.SetStatus(u)
	return upc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (upc *UserPasswordCreate) SetNillableStatus(u *userpassword.Status) *UserPasswordCreate {
	if u != nil {
		upc.SetStatus(*u)
	}
	return upc
}

// SetMemo sets the "memo" field.
func (upc *UserPasswordCreate) SetMemo(s string) *UserPasswordCreate {
	upc.mutation.SetMemo(s)
	return upc
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (upc *UserPasswordCreate) SetNillableMemo(s *string) *UserPasswordCreate {
	if s != nil {
		upc.SetMemo(*s)
	}
	return upc
}

// SetID sets the "id" field.
func (upc *UserPasswordCreate) SetID(i int) *UserPasswordCreate {
	upc.mutation.SetID(i)
	return upc
}

// SetUser sets the "user" edge to the User entity.
func (upc *UserPasswordCreate) SetUser(u *User) *UserPasswordCreate {
	return upc.SetUserID(u.ID)
}

// Mutation returns the UserPasswordMutation object of the builder.
func (upc *UserPasswordCreate) Mutation() *UserPasswordMutation {
	return upc.mutation
}

// Save creates the UserPassword in the database.
func (upc *UserPasswordCreate) Save(ctx context.Context) (*UserPassword, error) {
	if err := upc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*UserPassword, UserPasswordMutation](ctx, upc.sqlSave, upc.mutation, upc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (upc *UserPasswordCreate) SaveX(ctx context.Context) *UserPassword {
	v, err := upc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upc *UserPasswordCreate) Exec(ctx context.Context) error {
	_, err := upc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upc *UserPasswordCreate) ExecX(ctx context.Context) {
	if err := upc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (upc *UserPasswordCreate) defaults() error {
	if _, ok := upc.mutation.CreatedAt(); !ok {
		if userpassword.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized userpassword.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := userpassword.DefaultCreatedAt()
		upc.mutation.SetCreatedAt(v)
	}
	if _, ok := upc.mutation.UpdatedAt(); !ok {
		if userpassword.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized userpassword.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := userpassword.DefaultUpdatedAt()
		upc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (upc *UserPasswordCreate) check() error {
	if _, ok := upc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "UserPassword.created_by"`)}
	}
	if _, ok := upc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserPassword.created_at"`)}
	}
	if v, ok := upc.mutation.Scene(); ok {
		if err := userpassword.SceneValidator(v); err != nil {
			return &ValidationError{Name: "scene", err: fmt.Errorf(`ent: validator failed for field "UserPassword.scene": %w`, err)}
		}
	}
	if _, ok := upc.mutation.Salt(); !ok {
		return &ValidationError{Name: "salt", err: errors.New(`ent: missing required field "UserPassword.salt"`)}
	}
	if v, ok := upc.mutation.Salt(); ok {
		if err := userpassword.SaltValidator(v); err != nil {
			return &ValidationError{Name: "salt", err: fmt.Errorf(`ent: validator failed for field "UserPassword.salt": %w`, err)}
		}
	}
	if v, ok := upc.mutation.Status(); ok {
		if err := userpassword.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "UserPassword.status": %w`, err)}
		}
	}
	return nil
}

func (upc *UserPasswordCreate) sqlSave(ctx context.Context) (*UserPassword, error) {
	if err := upc.check(); err != nil {
		return nil, err
	}
	_node, _spec := upc.createSpec()
	if err := sqlgraph.CreateNode(ctx, upc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	upc.mutation.id = &_node.ID
	upc.mutation.done = true
	return _node, nil
}

func (upc *UserPasswordCreate) createSpec() (*UserPassword, *sqlgraph.CreateSpec) {
	var (
		_node = &UserPassword{config: upc.config}
		_spec = sqlgraph.NewCreateSpec(userpassword.Table, sqlgraph.NewFieldSpec(userpassword.FieldID, field.TypeInt))
	)
	if id, ok := upc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := upc.mutation.CreatedBy(); ok {
		_spec.SetField(userpassword.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := upc.mutation.CreatedAt(); ok {
		_spec.SetField(userpassword.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := upc.mutation.UpdatedBy(); ok {
		_spec.SetField(userpassword.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := upc.mutation.UpdatedAt(); ok {
		_spec.SetField(userpassword.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := upc.mutation.Scene(); ok {
		_spec.SetField(userpassword.FieldScene, field.TypeEnum, value)
		_node.Scene = value
	}
	if value, ok := upc.mutation.Password(); ok {
		_spec.SetField(userpassword.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := upc.mutation.Salt(); ok {
		_spec.SetField(userpassword.FieldSalt, field.TypeString, value)
		_node.Salt = value
	}
	if value, ok := upc.mutation.Status(); ok {
		_spec.SetField(userpassword.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := upc.mutation.Memo(); ok {
		_spec.SetField(userpassword.FieldMemo, field.TypeString, value)
		_node.Memo = value
	}
	if nodes := upc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userpassword.UserTable,
			Columns: []string{userpassword.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserPasswordCreateBulk is the builder for creating many UserPassword entities in bulk.
type UserPasswordCreateBulk struct {
	config
	builders []*UserPasswordCreate
}

// Save creates the UserPassword entities in the database.
func (upcb *UserPasswordCreateBulk) Save(ctx context.Context) ([]*UserPassword, error) {
	specs := make([]*sqlgraph.CreateSpec, len(upcb.builders))
	nodes := make([]*UserPassword, len(upcb.builders))
	mutators := make([]Mutator, len(upcb.builders))
	for i := range upcb.builders {
		func(i int, root context.Context) {
			builder := upcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserPasswordMutation)
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
					_, err = mutators[i+1].Mutate(root, upcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, upcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, upcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (upcb *UserPasswordCreateBulk) SaveX(ctx context.Context) []*UserPassword {
	v, err := upcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upcb *UserPasswordCreateBulk) Exec(ctx context.Context) error {
	_, err := upcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upcb *UserPasswordCreateBulk) ExecX(ctx context.Context) {
	if err := upcb.Exec(ctx); err != nil {
		panic(err)
	}
}
