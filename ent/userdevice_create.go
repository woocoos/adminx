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
	"github.com/woocoos/adminx/ent/userdevice"
)

// UserDeviceCreate is the builder for creating a UserDevice entity.
type UserDeviceCreate struct {
	config
	mutation *UserDeviceMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (udc *UserDeviceCreate) SetCreatedBy(i int) *UserDeviceCreate {
	udc.mutation.SetCreatedBy(i)
	return udc
}

// SetCreatedAt sets the "created_at" field.
func (udc *UserDeviceCreate) SetCreatedAt(t time.Time) *UserDeviceCreate {
	udc.mutation.SetCreatedAt(t)
	return udc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableCreatedAt(t *time.Time) *UserDeviceCreate {
	if t != nil {
		udc.SetCreatedAt(*t)
	}
	return udc
}

// SetUpdatedBy sets the "updated_by" field.
func (udc *UserDeviceCreate) SetUpdatedBy(i int) *UserDeviceCreate {
	udc.mutation.SetUpdatedBy(i)
	return udc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableUpdatedBy(i *int) *UserDeviceCreate {
	if i != nil {
		udc.SetUpdatedBy(*i)
	}
	return udc
}

// SetUpdatedAt sets the "updated_at" field.
func (udc *UserDeviceCreate) SetUpdatedAt(t time.Time) *UserDeviceCreate {
	udc.mutation.SetUpdatedAt(t)
	return udc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableUpdatedAt(t *time.Time) *UserDeviceCreate {
	if t != nil {
		udc.SetUpdatedAt(*t)
	}
	return udc
}

// SetUserID sets the "user_id" field.
func (udc *UserDeviceCreate) SetUserID(i int) *UserDeviceCreate {
	udc.mutation.SetUserID(i)
	return udc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableUserID(i *int) *UserDeviceCreate {
	if i != nil {
		udc.SetUserID(*i)
	}
	return udc
}

// SetDeviceUID sets the "device_uid" field.
func (udc *UserDeviceCreate) SetDeviceUID(s string) *UserDeviceCreate {
	udc.mutation.SetDeviceUID(s)
	return udc
}

// SetDeviceName sets the "device_name" field.
func (udc *UserDeviceCreate) SetDeviceName(s string) *UserDeviceCreate {
	udc.mutation.SetDeviceName(s)
	return udc
}

// SetNillableDeviceName sets the "device_name" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableDeviceName(s *string) *UserDeviceCreate {
	if s != nil {
		udc.SetDeviceName(*s)
	}
	return udc
}

// SetSystemName sets the "system_name" field.
func (udc *UserDeviceCreate) SetSystemName(s string) *UserDeviceCreate {
	udc.mutation.SetSystemName(s)
	return udc
}

// SetNillableSystemName sets the "system_name" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableSystemName(s *string) *UserDeviceCreate {
	if s != nil {
		udc.SetSystemName(*s)
	}
	return udc
}

// SetSystemVersion sets the "system_version" field.
func (udc *UserDeviceCreate) SetSystemVersion(s string) *UserDeviceCreate {
	udc.mutation.SetSystemVersion(s)
	return udc
}

// SetNillableSystemVersion sets the "system_version" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableSystemVersion(s *string) *UserDeviceCreate {
	if s != nil {
		udc.SetSystemVersion(*s)
	}
	return udc
}

// SetAppVersion sets the "app_version" field.
func (udc *UserDeviceCreate) SetAppVersion(s string) *UserDeviceCreate {
	udc.mutation.SetAppVersion(s)
	return udc
}

// SetNillableAppVersion sets the "app_version" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableAppVersion(s *string) *UserDeviceCreate {
	if s != nil {
		udc.SetAppVersion(*s)
	}
	return udc
}

// SetDeviceModel sets the "device_model" field.
func (udc *UserDeviceCreate) SetDeviceModel(s string) *UserDeviceCreate {
	udc.mutation.SetDeviceModel(s)
	return udc
}

// SetNillableDeviceModel sets the "device_model" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableDeviceModel(s *string) *UserDeviceCreate {
	if s != nil {
		udc.SetDeviceModel(*s)
	}
	return udc
}

// SetStatus sets the "status" field.
func (udc *UserDeviceCreate) SetStatus(u userdevice.Status) *UserDeviceCreate {
	udc.mutation.SetStatus(u)
	return udc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableStatus(u *userdevice.Status) *UserDeviceCreate {
	if u != nil {
		udc.SetStatus(*u)
	}
	return udc
}

// SetComments sets the "comments" field.
func (udc *UserDeviceCreate) SetComments(s string) *UserDeviceCreate {
	udc.mutation.SetComments(s)
	return udc
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableComments(s *string) *UserDeviceCreate {
	if s != nil {
		udc.SetComments(*s)
	}
	return udc
}

// SetID sets the "id" field.
func (udc *UserDeviceCreate) SetID(i int) *UserDeviceCreate {
	udc.mutation.SetID(i)
	return udc
}

// SetUser sets the "user" edge to the User entity.
func (udc *UserDeviceCreate) SetUser(u *User) *UserDeviceCreate {
	return udc.SetUserID(u.ID)
}

// Mutation returns the UserDeviceMutation object of the builder.
func (udc *UserDeviceCreate) Mutation() *UserDeviceMutation {
	return udc.mutation
}

// Save creates the UserDevice in the database.
func (udc *UserDeviceCreate) Save(ctx context.Context) (*UserDevice, error) {
	if err := udc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*UserDevice, UserDeviceMutation](ctx, udc.sqlSave, udc.mutation, udc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (udc *UserDeviceCreate) SaveX(ctx context.Context) *UserDevice {
	v, err := udc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (udc *UserDeviceCreate) Exec(ctx context.Context) error {
	_, err := udc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (udc *UserDeviceCreate) ExecX(ctx context.Context) {
	if err := udc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (udc *UserDeviceCreate) defaults() error {
	if _, ok := udc.mutation.CreatedAt(); !ok {
		if userdevice.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized userdevice.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := userdevice.DefaultCreatedAt()
		udc.mutation.SetCreatedAt(v)
	}
	if _, ok := udc.mutation.UpdatedAt(); !ok {
		if userdevice.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized userdevice.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := userdevice.DefaultUpdatedAt()
		udc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (udc *UserDeviceCreate) check() error {
	if _, ok := udc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "UserDevice.created_by"`)}
	}
	if _, ok := udc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserDevice.created_at"`)}
	}
	if _, ok := udc.mutation.DeviceUID(); !ok {
		return &ValidationError{Name: "device_uid", err: errors.New(`ent: missing required field "UserDevice.device_uid"`)}
	}
	if v, ok := udc.mutation.DeviceUID(); ok {
		if err := userdevice.DeviceUIDValidator(v); err != nil {
			return &ValidationError{Name: "device_uid", err: fmt.Errorf(`ent: validator failed for field "UserDevice.device_uid": %w`, err)}
		}
	}
	if v, ok := udc.mutation.DeviceName(); ok {
		if err := userdevice.DeviceNameValidator(v); err != nil {
			return &ValidationError{Name: "device_name", err: fmt.Errorf(`ent: validator failed for field "UserDevice.device_name": %w`, err)}
		}
	}
	if v, ok := udc.mutation.SystemName(); ok {
		if err := userdevice.SystemNameValidator(v); err != nil {
			return &ValidationError{Name: "system_name", err: fmt.Errorf(`ent: validator failed for field "UserDevice.system_name": %w`, err)}
		}
	}
	if v, ok := udc.mutation.SystemVersion(); ok {
		if err := userdevice.SystemVersionValidator(v); err != nil {
			return &ValidationError{Name: "system_version", err: fmt.Errorf(`ent: validator failed for field "UserDevice.system_version": %w`, err)}
		}
	}
	if v, ok := udc.mutation.AppVersion(); ok {
		if err := userdevice.AppVersionValidator(v); err != nil {
			return &ValidationError{Name: "app_version", err: fmt.Errorf(`ent: validator failed for field "UserDevice.app_version": %w`, err)}
		}
	}
	if v, ok := udc.mutation.DeviceModel(); ok {
		if err := userdevice.DeviceModelValidator(v); err != nil {
			return &ValidationError{Name: "device_model", err: fmt.Errorf(`ent: validator failed for field "UserDevice.device_model": %w`, err)}
		}
	}
	if v, ok := udc.mutation.Status(); ok {
		if err := userdevice.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "UserDevice.status": %w`, err)}
		}
	}
	return nil
}

func (udc *UserDeviceCreate) sqlSave(ctx context.Context) (*UserDevice, error) {
	if err := udc.check(); err != nil {
		return nil, err
	}
	_node, _spec := udc.createSpec()
	if err := sqlgraph.CreateNode(ctx, udc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	udc.mutation.id = &_node.ID
	udc.mutation.done = true
	return _node, nil
}

func (udc *UserDeviceCreate) createSpec() (*UserDevice, *sqlgraph.CreateSpec) {
	var (
		_node = &UserDevice{config: udc.config}
		_spec = sqlgraph.NewCreateSpec(userdevice.Table, sqlgraph.NewFieldSpec(userdevice.FieldID, field.TypeInt))
	)
	if id, ok := udc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := udc.mutation.CreatedBy(); ok {
		_spec.SetField(userdevice.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := udc.mutation.CreatedAt(); ok {
		_spec.SetField(userdevice.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := udc.mutation.UpdatedBy(); ok {
		_spec.SetField(userdevice.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := udc.mutation.UpdatedAt(); ok {
		_spec.SetField(userdevice.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := udc.mutation.DeviceUID(); ok {
		_spec.SetField(userdevice.FieldDeviceUID, field.TypeString, value)
		_node.DeviceUID = value
	}
	if value, ok := udc.mutation.DeviceName(); ok {
		_spec.SetField(userdevice.FieldDeviceName, field.TypeString, value)
		_node.DeviceName = value
	}
	if value, ok := udc.mutation.SystemName(); ok {
		_spec.SetField(userdevice.FieldSystemName, field.TypeString, value)
		_node.SystemName = value
	}
	if value, ok := udc.mutation.SystemVersion(); ok {
		_spec.SetField(userdevice.FieldSystemVersion, field.TypeString, value)
		_node.SystemVersion = value
	}
	if value, ok := udc.mutation.AppVersion(); ok {
		_spec.SetField(userdevice.FieldAppVersion, field.TypeString, value)
		_node.AppVersion = value
	}
	if value, ok := udc.mutation.DeviceModel(); ok {
		_spec.SetField(userdevice.FieldDeviceModel, field.TypeString, value)
		_node.DeviceModel = value
	}
	if value, ok := udc.mutation.Status(); ok {
		_spec.SetField(userdevice.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := udc.mutation.Comments(); ok {
		_spec.SetField(userdevice.FieldComments, field.TypeString, value)
		_node.Comments = value
	}
	if nodes := udc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userdevice.UserTable,
			Columns: []string{userdevice.UserColumn},
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

// UserDeviceCreateBulk is the builder for creating many UserDevice entities in bulk.
type UserDeviceCreateBulk struct {
	config
	builders []*UserDeviceCreate
}

// Save creates the UserDevice entities in the database.
func (udcb *UserDeviceCreateBulk) Save(ctx context.Context) ([]*UserDevice, error) {
	specs := make([]*sqlgraph.CreateSpec, len(udcb.builders))
	nodes := make([]*UserDevice, len(udcb.builders))
	mutators := make([]Mutator, len(udcb.builders))
	for i := range udcb.builders {
		func(i int, root context.Context) {
			builder := udcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserDeviceMutation)
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
					_, err = mutators[i+1].Mutate(root, udcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, udcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, udcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (udcb *UserDeviceCreateBulk) SaveX(ctx context.Context) []*UserDevice {
	v, err := udcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (udcb *UserDeviceCreateBulk) Exec(ctx context.Context) error {
	_, err := udcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (udcb *UserDeviceCreateBulk) ExecX(ctx context.Context) {
	if err := udcb.Exec(ctx); err != nil {
		panic(err)
	}
}