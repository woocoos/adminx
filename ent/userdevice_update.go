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
	"github.com/woocoos/adminx/ent/predicate"
	"github.com/woocoos/adminx/ent/userdevice"
)

// UserDeviceUpdate is the builder for updating UserDevice entities.
type UserDeviceUpdate struct {
	config
	hooks    []Hook
	mutation *UserDeviceMutation
}

// Where appends a list predicates to the UserDeviceUpdate builder.
func (udu *UserDeviceUpdate) Where(ps ...predicate.UserDevice) *UserDeviceUpdate {
	udu.mutation.Where(ps...)
	return udu
}

// SetUpdatedBy sets the "updated_by" field.
func (udu *UserDeviceUpdate) SetUpdatedBy(i int) *UserDeviceUpdate {
	udu.mutation.ResetUpdatedBy()
	udu.mutation.SetUpdatedBy(i)
	return udu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (udu *UserDeviceUpdate) SetNillableUpdatedBy(i *int) *UserDeviceUpdate {
	if i != nil {
		udu.SetUpdatedBy(*i)
	}
	return udu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (udu *UserDeviceUpdate) AddUpdatedBy(i int) *UserDeviceUpdate {
	udu.mutation.AddUpdatedBy(i)
	return udu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (udu *UserDeviceUpdate) ClearUpdatedBy() *UserDeviceUpdate {
	udu.mutation.ClearUpdatedBy()
	return udu
}

// SetUpdatedAt sets the "updated_at" field.
func (udu *UserDeviceUpdate) SetUpdatedAt(t time.Time) *UserDeviceUpdate {
	udu.mutation.SetUpdatedAt(t)
	return udu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (udu *UserDeviceUpdate) ClearUpdatedAt() *UserDeviceUpdate {
	udu.mutation.ClearUpdatedAt()
	return udu
}

// SetDeviceUID sets the "device_uid" field.
func (udu *UserDeviceUpdate) SetDeviceUID(s string) *UserDeviceUpdate {
	udu.mutation.SetDeviceUID(s)
	return udu
}

// SetDeviceName sets the "device_name" field.
func (udu *UserDeviceUpdate) SetDeviceName(s string) *UserDeviceUpdate {
	udu.mutation.SetDeviceName(s)
	return udu
}

// SetNillableDeviceName sets the "device_name" field if the given value is not nil.
func (udu *UserDeviceUpdate) SetNillableDeviceName(s *string) *UserDeviceUpdate {
	if s != nil {
		udu.SetDeviceName(*s)
	}
	return udu
}

// ClearDeviceName clears the value of the "device_name" field.
func (udu *UserDeviceUpdate) ClearDeviceName() *UserDeviceUpdate {
	udu.mutation.ClearDeviceName()
	return udu
}

// SetSystemName sets the "system_name" field.
func (udu *UserDeviceUpdate) SetSystemName(s string) *UserDeviceUpdate {
	udu.mutation.SetSystemName(s)
	return udu
}

// SetNillableSystemName sets the "system_name" field if the given value is not nil.
func (udu *UserDeviceUpdate) SetNillableSystemName(s *string) *UserDeviceUpdate {
	if s != nil {
		udu.SetSystemName(*s)
	}
	return udu
}

// ClearSystemName clears the value of the "system_name" field.
func (udu *UserDeviceUpdate) ClearSystemName() *UserDeviceUpdate {
	udu.mutation.ClearSystemName()
	return udu
}

// SetSystemVersion sets the "system_version" field.
func (udu *UserDeviceUpdate) SetSystemVersion(s string) *UserDeviceUpdate {
	udu.mutation.SetSystemVersion(s)
	return udu
}

// SetNillableSystemVersion sets the "system_version" field if the given value is not nil.
func (udu *UserDeviceUpdate) SetNillableSystemVersion(s *string) *UserDeviceUpdate {
	if s != nil {
		udu.SetSystemVersion(*s)
	}
	return udu
}

// ClearSystemVersion clears the value of the "system_version" field.
func (udu *UserDeviceUpdate) ClearSystemVersion() *UserDeviceUpdate {
	udu.mutation.ClearSystemVersion()
	return udu
}

// SetAppVersion sets the "app_version" field.
func (udu *UserDeviceUpdate) SetAppVersion(s string) *UserDeviceUpdate {
	udu.mutation.SetAppVersion(s)
	return udu
}

// SetNillableAppVersion sets the "app_version" field if the given value is not nil.
func (udu *UserDeviceUpdate) SetNillableAppVersion(s *string) *UserDeviceUpdate {
	if s != nil {
		udu.SetAppVersion(*s)
	}
	return udu
}

// ClearAppVersion clears the value of the "app_version" field.
func (udu *UserDeviceUpdate) ClearAppVersion() *UserDeviceUpdate {
	udu.mutation.ClearAppVersion()
	return udu
}

// SetDeviceModel sets the "device_model" field.
func (udu *UserDeviceUpdate) SetDeviceModel(s string) *UserDeviceUpdate {
	udu.mutation.SetDeviceModel(s)
	return udu
}

// SetNillableDeviceModel sets the "device_model" field if the given value is not nil.
func (udu *UserDeviceUpdate) SetNillableDeviceModel(s *string) *UserDeviceUpdate {
	if s != nil {
		udu.SetDeviceModel(*s)
	}
	return udu
}

// ClearDeviceModel clears the value of the "device_model" field.
func (udu *UserDeviceUpdate) ClearDeviceModel() *UserDeviceUpdate {
	udu.mutation.ClearDeviceModel()
	return udu
}

// SetStatus sets the "status" field.
func (udu *UserDeviceUpdate) SetStatus(u userdevice.Status) *UserDeviceUpdate {
	udu.mutation.SetStatus(u)
	return udu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (udu *UserDeviceUpdate) SetNillableStatus(u *userdevice.Status) *UserDeviceUpdate {
	if u != nil {
		udu.SetStatus(*u)
	}
	return udu
}

// ClearStatus clears the value of the "status" field.
func (udu *UserDeviceUpdate) ClearStatus() *UserDeviceUpdate {
	udu.mutation.ClearStatus()
	return udu
}

// SetComments sets the "comments" field.
func (udu *UserDeviceUpdate) SetComments(s string) *UserDeviceUpdate {
	udu.mutation.SetComments(s)
	return udu
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (udu *UserDeviceUpdate) SetNillableComments(s *string) *UserDeviceUpdate {
	if s != nil {
		udu.SetComments(*s)
	}
	return udu
}

// ClearComments clears the value of the "comments" field.
func (udu *UserDeviceUpdate) ClearComments() *UserDeviceUpdate {
	udu.mutation.ClearComments()
	return udu
}

// Mutation returns the UserDeviceMutation object of the builder.
func (udu *UserDeviceUpdate) Mutation() *UserDeviceMutation {
	return udu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (udu *UserDeviceUpdate) Save(ctx context.Context) (int, error) {
	if err := udu.defaults(); err != nil {
		return 0, err
	}
	return withHooks[int, UserDeviceMutation](ctx, udu.sqlSave, udu.mutation, udu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (udu *UserDeviceUpdate) SaveX(ctx context.Context) int {
	affected, err := udu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (udu *UserDeviceUpdate) Exec(ctx context.Context) error {
	_, err := udu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (udu *UserDeviceUpdate) ExecX(ctx context.Context) {
	if err := udu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (udu *UserDeviceUpdate) defaults() error {
	if _, ok := udu.mutation.UpdatedAt(); !ok && !udu.mutation.UpdatedAtCleared() {
		if userdevice.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized userdevice.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := userdevice.UpdateDefaultUpdatedAt()
		udu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (udu *UserDeviceUpdate) check() error {
	if v, ok := udu.mutation.DeviceUID(); ok {
		if err := userdevice.DeviceUIDValidator(v); err != nil {
			return &ValidationError{Name: "device_uid", err: fmt.Errorf(`ent: validator failed for field "UserDevice.device_uid": %w`, err)}
		}
	}
	if v, ok := udu.mutation.DeviceName(); ok {
		if err := userdevice.DeviceNameValidator(v); err != nil {
			return &ValidationError{Name: "device_name", err: fmt.Errorf(`ent: validator failed for field "UserDevice.device_name": %w`, err)}
		}
	}
	if v, ok := udu.mutation.SystemName(); ok {
		if err := userdevice.SystemNameValidator(v); err != nil {
			return &ValidationError{Name: "system_name", err: fmt.Errorf(`ent: validator failed for field "UserDevice.system_name": %w`, err)}
		}
	}
	if v, ok := udu.mutation.SystemVersion(); ok {
		if err := userdevice.SystemVersionValidator(v); err != nil {
			return &ValidationError{Name: "system_version", err: fmt.Errorf(`ent: validator failed for field "UserDevice.system_version": %w`, err)}
		}
	}
	if v, ok := udu.mutation.AppVersion(); ok {
		if err := userdevice.AppVersionValidator(v); err != nil {
			return &ValidationError{Name: "app_version", err: fmt.Errorf(`ent: validator failed for field "UserDevice.app_version": %w`, err)}
		}
	}
	if v, ok := udu.mutation.DeviceModel(); ok {
		if err := userdevice.DeviceModelValidator(v); err != nil {
			return &ValidationError{Name: "device_model", err: fmt.Errorf(`ent: validator failed for field "UserDevice.device_model": %w`, err)}
		}
	}
	if v, ok := udu.mutation.Status(); ok {
		if err := userdevice.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "UserDevice.status": %w`, err)}
		}
	}
	return nil
}

func (udu *UserDeviceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := udu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userdevice.Table, userdevice.Columns, sqlgraph.NewFieldSpec(userdevice.FieldID, field.TypeInt))
	if ps := udu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := udu.mutation.UpdatedBy(); ok {
		_spec.SetField(userdevice.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := udu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(userdevice.FieldUpdatedBy, field.TypeInt, value)
	}
	if udu.mutation.UpdatedByCleared() {
		_spec.ClearField(userdevice.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := udu.mutation.UpdatedAt(); ok {
		_spec.SetField(userdevice.FieldUpdatedAt, field.TypeTime, value)
	}
	if udu.mutation.UpdatedAtCleared() {
		_spec.ClearField(userdevice.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := udu.mutation.DeviceUID(); ok {
		_spec.SetField(userdevice.FieldDeviceUID, field.TypeString, value)
	}
	if value, ok := udu.mutation.DeviceName(); ok {
		_spec.SetField(userdevice.FieldDeviceName, field.TypeString, value)
	}
	if udu.mutation.DeviceNameCleared() {
		_spec.ClearField(userdevice.FieldDeviceName, field.TypeString)
	}
	if value, ok := udu.mutation.SystemName(); ok {
		_spec.SetField(userdevice.FieldSystemName, field.TypeString, value)
	}
	if udu.mutation.SystemNameCleared() {
		_spec.ClearField(userdevice.FieldSystemName, field.TypeString)
	}
	if value, ok := udu.mutation.SystemVersion(); ok {
		_spec.SetField(userdevice.FieldSystemVersion, field.TypeString, value)
	}
	if udu.mutation.SystemVersionCleared() {
		_spec.ClearField(userdevice.FieldSystemVersion, field.TypeString)
	}
	if value, ok := udu.mutation.AppVersion(); ok {
		_spec.SetField(userdevice.FieldAppVersion, field.TypeString, value)
	}
	if udu.mutation.AppVersionCleared() {
		_spec.ClearField(userdevice.FieldAppVersion, field.TypeString)
	}
	if value, ok := udu.mutation.DeviceModel(); ok {
		_spec.SetField(userdevice.FieldDeviceModel, field.TypeString, value)
	}
	if udu.mutation.DeviceModelCleared() {
		_spec.ClearField(userdevice.FieldDeviceModel, field.TypeString)
	}
	if value, ok := udu.mutation.Status(); ok {
		_spec.SetField(userdevice.FieldStatus, field.TypeEnum, value)
	}
	if udu.mutation.StatusCleared() {
		_spec.ClearField(userdevice.FieldStatus, field.TypeEnum)
	}
	if value, ok := udu.mutation.Comments(); ok {
		_spec.SetField(userdevice.FieldComments, field.TypeString, value)
	}
	if udu.mutation.CommentsCleared() {
		_spec.ClearField(userdevice.FieldComments, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, udu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userdevice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	udu.mutation.done = true
	return n, nil
}

// UserDeviceUpdateOne is the builder for updating a single UserDevice entity.
type UserDeviceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserDeviceMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (uduo *UserDeviceUpdateOne) SetUpdatedBy(i int) *UserDeviceUpdateOne {
	uduo.mutation.ResetUpdatedBy()
	uduo.mutation.SetUpdatedBy(i)
	return uduo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (uduo *UserDeviceUpdateOne) SetNillableUpdatedBy(i *int) *UserDeviceUpdateOne {
	if i != nil {
		uduo.SetUpdatedBy(*i)
	}
	return uduo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (uduo *UserDeviceUpdateOne) AddUpdatedBy(i int) *UserDeviceUpdateOne {
	uduo.mutation.AddUpdatedBy(i)
	return uduo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (uduo *UserDeviceUpdateOne) ClearUpdatedBy() *UserDeviceUpdateOne {
	uduo.mutation.ClearUpdatedBy()
	return uduo
}

// SetUpdatedAt sets the "updated_at" field.
func (uduo *UserDeviceUpdateOne) SetUpdatedAt(t time.Time) *UserDeviceUpdateOne {
	uduo.mutation.SetUpdatedAt(t)
	return uduo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (uduo *UserDeviceUpdateOne) ClearUpdatedAt() *UserDeviceUpdateOne {
	uduo.mutation.ClearUpdatedAt()
	return uduo
}

// SetDeviceUID sets the "device_uid" field.
func (uduo *UserDeviceUpdateOne) SetDeviceUID(s string) *UserDeviceUpdateOne {
	uduo.mutation.SetDeviceUID(s)
	return uduo
}

// SetDeviceName sets the "device_name" field.
func (uduo *UserDeviceUpdateOne) SetDeviceName(s string) *UserDeviceUpdateOne {
	uduo.mutation.SetDeviceName(s)
	return uduo
}

// SetNillableDeviceName sets the "device_name" field if the given value is not nil.
func (uduo *UserDeviceUpdateOne) SetNillableDeviceName(s *string) *UserDeviceUpdateOne {
	if s != nil {
		uduo.SetDeviceName(*s)
	}
	return uduo
}

// ClearDeviceName clears the value of the "device_name" field.
func (uduo *UserDeviceUpdateOne) ClearDeviceName() *UserDeviceUpdateOne {
	uduo.mutation.ClearDeviceName()
	return uduo
}

// SetSystemName sets the "system_name" field.
func (uduo *UserDeviceUpdateOne) SetSystemName(s string) *UserDeviceUpdateOne {
	uduo.mutation.SetSystemName(s)
	return uduo
}

// SetNillableSystemName sets the "system_name" field if the given value is not nil.
func (uduo *UserDeviceUpdateOne) SetNillableSystemName(s *string) *UserDeviceUpdateOne {
	if s != nil {
		uduo.SetSystemName(*s)
	}
	return uduo
}

// ClearSystemName clears the value of the "system_name" field.
func (uduo *UserDeviceUpdateOne) ClearSystemName() *UserDeviceUpdateOne {
	uduo.mutation.ClearSystemName()
	return uduo
}

// SetSystemVersion sets the "system_version" field.
func (uduo *UserDeviceUpdateOne) SetSystemVersion(s string) *UserDeviceUpdateOne {
	uduo.mutation.SetSystemVersion(s)
	return uduo
}

// SetNillableSystemVersion sets the "system_version" field if the given value is not nil.
func (uduo *UserDeviceUpdateOne) SetNillableSystemVersion(s *string) *UserDeviceUpdateOne {
	if s != nil {
		uduo.SetSystemVersion(*s)
	}
	return uduo
}

// ClearSystemVersion clears the value of the "system_version" field.
func (uduo *UserDeviceUpdateOne) ClearSystemVersion() *UserDeviceUpdateOne {
	uduo.mutation.ClearSystemVersion()
	return uduo
}

// SetAppVersion sets the "app_version" field.
func (uduo *UserDeviceUpdateOne) SetAppVersion(s string) *UserDeviceUpdateOne {
	uduo.mutation.SetAppVersion(s)
	return uduo
}

// SetNillableAppVersion sets the "app_version" field if the given value is not nil.
func (uduo *UserDeviceUpdateOne) SetNillableAppVersion(s *string) *UserDeviceUpdateOne {
	if s != nil {
		uduo.SetAppVersion(*s)
	}
	return uduo
}

// ClearAppVersion clears the value of the "app_version" field.
func (uduo *UserDeviceUpdateOne) ClearAppVersion() *UserDeviceUpdateOne {
	uduo.mutation.ClearAppVersion()
	return uduo
}

// SetDeviceModel sets the "device_model" field.
func (uduo *UserDeviceUpdateOne) SetDeviceModel(s string) *UserDeviceUpdateOne {
	uduo.mutation.SetDeviceModel(s)
	return uduo
}

// SetNillableDeviceModel sets the "device_model" field if the given value is not nil.
func (uduo *UserDeviceUpdateOne) SetNillableDeviceModel(s *string) *UserDeviceUpdateOne {
	if s != nil {
		uduo.SetDeviceModel(*s)
	}
	return uduo
}

// ClearDeviceModel clears the value of the "device_model" field.
func (uduo *UserDeviceUpdateOne) ClearDeviceModel() *UserDeviceUpdateOne {
	uduo.mutation.ClearDeviceModel()
	return uduo
}

// SetStatus sets the "status" field.
func (uduo *UserDeviceUpdateOne) SetStatus(u userdevice.Status) *UserDeviceUpdateOne {
	uduo.mutation.SetStatus(u)
	return uduo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uduo *UserDeviceUpdateOne) SetNillableStatus(u *userdevice.Status) *UserDeviceUpdateOne {
	if u != nil {
		uduo.SetStatus(*u)
	}
	return uduo
}

// ClearStatus clears the value of the "status" field.
func (uduo *UserDeviceUpdateOne) ClearStatus() *UserDeviceUpdateOne {
	uduo.mutation.ClearStatus()
	return uduo
}

// SetComments sets the "comments" field.
func (uduo *UserDeviceUpdateOne) SetComments(s string) *UserDeviceUpdateOne {
	uduo.mutation.SetComments(s)
	return uduo
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (uduo *UserDeviceUpdateOne) SetNillableComments(s *string) *UserDeviceUpdateOne {
	if s != nil {
		uduo.SetComments(*s)
	}
	return uduo
}

// ClearComments clears the value of the "comments" field.
func (uduo *UserDeviceUpdateOne) ClearComments() *UserDeviceUpdateOne {
	uduo.mutation.ClearComments()
	return uduo
}

// Mutation returns the UserDeviceMutation object of the builder.
func (uduo *UserDeviceUpdateOne) Mutation() *UserDeviceMutation {
	return uduo.mutation
}

// Where appends a list predicates to the UserDeviceUpdate builder.
func (uduo *UserDeviceUpdateOne) Where(ps ...predicate.UserDevice) *UserDeviceUpdateOne {
	uduo.mutation.Where(ps...)
	return uduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uduo *UserDeviceUpdateOne) Select(field string, fields ...string) *UserDeviceUpdateOne {
	uduo.fields = append([]string{field}, fields...)
	return uduo
}

// Save executes the query and returns the updated UserDevice entity.
func (uduo *UserDeviceUpdateOne) Save(ctx context.Context) (*UserDevice, error) {
	if err := uduo.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*UserDevice, UserDeviceMutation](ctx, uduo.sqlSave, uduo.mutation, uduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uduo *UserDeviceUpdateOne) SaveX(ctx context.Context) *UserDevice {
	node, err := uduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uduo *UserDeviceUpdateOne) Exec(ctx context.Context) error {
	_, err := uduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uduo *UserDeviceUpdateOne) ExecX(ctx context.Context) {
	if err := uduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uduo *UserDeviceUpdateOne) defaults() error {
	if _, ok := uduo.mutation.UpdatedAt(); !ok && !uduo.mutation.UpdatedAtCleared() {
		if userdevice.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized userdevice.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := userdevice.UpdateDefaultUpdatedAt()
		uduo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uduo *UserDeviceUpdateOne) check() error {
	if v, ok := uduo.mutation.DeviceUID(); ok {
		if err := userdevice.DeviceUIDValidator(v); err != nil {
			return &ValidationError{Name: "device_uid", err: fmt.Errorf(`ent: validator failed for field "UserDevice.device_uid": %w`, err)}
		}
	}
	if v, ok := uduo.mutation.DeviceName(); ok {
		if err := userdevice.DeviceNameValidator(v); err != nil {
			return &ValidationError{Name: "device_name", err: fmt.Errorf(`ent: validator failed for field "UserDevice.device_name": %w`, err)}
		}
	}
	if v, ok := uduo.mutation.SystemName(); ok {
		if err := userdevice.SystemNameValidator(v); err != nil {
			return &ValidationError{Name: "system_name", err: fmt.Errorf(`ent: validator failed for field "UserDevice.system_name": %w`, err)}
		}
	}
	if v, ok := uduo.mutation.SystemVersion(); ok {
		if err := userdevice.SystemVersionValidator(v); err != nil {
			return &ValidationError{Name: "system_version", err: fmt.Errorf(`ent: validator failed for field "UserDevice.system_version": %w`, err)}
		}
	}
	if v, ok := uduo.mutation.AppVersion(); ok {
		if err := userdevice.AppVersionValidator(v); err != nil {
			return &ValidationError{Name: "app_version", err: fmt.Errorf(`ent: validator failed for field "UserDevice.app_version": %w`, err)}
		}
	}
	if v, ok := uduo.mutation.DeviceModel(); ok {
		if err := userdevice.DeviceModelValidator(v); err != nil {
			return &ValidationError{Name: "device_model", err: fmt.Errorf(`ent: validator failed for field "UserDevice.device_model": %w`, err)}
		}
	}
	if v, ok := uduo.mutation.Status(); ok {
		if err := userdevice.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "UserDevice.status": %w`, err)}
		}
	}
	return nil
}

func (uduo *UserDeviceUpdateOne) sqlSave(ctx context.Context) (_node *UserDevice, err error) {
	if err := uduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userdevice.Table, userdevice.Columns, sqlgraph.NewFieldSpec(userdevice.FieldID, field.TypeInt))
	id, ok := uduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserDevice.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userdevice.FieldID)
		for _, f := range fields {
			if !userdevice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userdevice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uduo.mutation.UpdatedBy(); ok {
		_spec.SetField(userdevice.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := uduo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(userdevice.FieldUpdatedBy, field.TypeInt, value)
	}
	if uduo.mutation.UpdatedByCleared() {
		_spec.ClearField(userdevice.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := uduo.mutation.UpdatedAt(); ok {
		_spec.SetField(userdevice.FieldUpdatedAt, field.TypeTime, value)
	}
	if uduo.mutation.UpdatedAtCleared() {
		_spec.ClearField(userdevice.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := uduo.mutation.DeviceUID(); ok {
		_spec.SetField(userdevice.FieldDeviceUID, field.TypeString, value)
	}
	if value, ok := uduo.mutation.DeviceName(); ok {
		_spec.SetField(userdevice.FieldDeviceName, field.TypeString, value)
	}
	if uduo.mutation.DeviceNameCleared() {
		_spec.ClearField(userdevice.FieldDeviceName, field.TypeString)
	}
	if value, ok := uduo.mutation.SystemName(); ok {
		_spec.SetField(userdevice.FieldSystemName, field.TypeString, value)
	}
	if uduo.mutation.SystemNameCleared() {
		_spec.ClearField(userdevice.FieldSystemName, field.TypeString)
	}
	if value, ok := uduo.mutation.SystemVersion(); ok {
		_spec.SetField(userdevice.FieldSystemVersion, field.TypeString, value)
	}
	if uduo.mutation.SystemVersionCleared() {
		_spec.ClearField(userdevice.FieldSystemVersion, field.TypeString)
	}
	if value, ok := uduo.mutation.AppVersion(); ok {
		_spec.SetField(userdevice.FieldAppVersion, field.TypeString, value)
	}
	if uduo.mutation.AppVersionCleared() {
		_spec.ClearField(userdevice.FieldAppVersion, field.TypeString)
	}
	if value, ok := uduo.mutation.DeviceModel(); ok {
		_spec.SetField(userdevice.FieldDeviceModel, field.TypeString, value)
	}
	if uduo.mutation.DeviceModelCleared() {
		_spec.ClearField(userdevice.FieldDeviceModel, field.TypeString)
	}
	if value, ok := uduo.mutation.Status(); ok {
		_spec.SetField(userdevice.FieldStatus, field.TypeEnum, value)
	}
	if uduo.mutation.StatusCleared() {
		_spec.ClearField(userdevice.FieldStatus, field.TypeEnum)
	}
	if value, ok := uduo.mutation.Comments(); ok {
		_spec.SetField(userdevice.FieldComments, field.TypeString, value)
	}
	if uduo.mutation.CommentsCleared() {
		_spec.ClearField(userdevice.FieldComments, field.TypeString)
	}
	_node = &UserDevice{config: uduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userdevice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uduo.mutation.done = true
	return _node, nil
}