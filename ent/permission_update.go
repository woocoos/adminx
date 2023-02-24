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
	"github.com/woocoos/adminx/ent/organization"
	"github.com/woocoos/adminx/ent/permission"
	"github.com/woocoos/adminx/ent/predicate"
	"github.com/woocoos/adminx/ent/user"
	"github.com/woocoos/adminx/graph/entgen/types"
)

// PermissionUpdate is the builder for updating Permission entities.
type PermissionUpdate struct {
	config
	hooks    []Hook
	mutation *PermissionMutation
}

// Where appends a list predicates to the PermissionUpdate builder.
func (pu *PermissionUpdate) Where(ps ...predicate.Permission) *PermissionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedBy sets the "updated_by" field.
func (pu *PermissionUpdate) SetUpdatedBy(i int) *PermissionUpdate {
	pu.mutation.ResetUpdatedBy()
	pu.mutation.SetUpdatedBy(i)
	return pu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableUpdatedBy(i *int) *PermissionUpdate {
	if i != nil {
		pu.SetUpdatedBy(*i)
	}
	return pu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (pu *PermissionUpdate) AddUpdatedBy(i int) *PermissionUpdate {
	pu.mutation.AddUpdatedBy(i)
	return pu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (pu *PermissionUpdate) ClearUpdatedBy() *PermissionUpdate {
	pu.mutation.ClearUpdatedBy()
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PermissionUpdate) SetUpdatedAt(t time.Time) *PermissionUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableUpdatedAt(t *time.Time) *PermissionUpdate {
	if t != nil {
		pu.SetUpdatedAt(*t)
	}
	return pu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (pu *PermissionUpdate) ClearUpdatedAt() *PermissionUpdate {
	pu.mutation.ClearUpdatedAt()
	return pu
}

// SetOrgID sets the "org_id" field.
func (pu *PermissionUpdate) SetOrgID(i int) *PermissionUpdate {
	pu.mutation.SetOrgID(i)
	return pu
}

// SetPrincipalKind sets the "principal_kind" field.
func (pu *PermissionUpdate) SetPrincipalKind(pk permission.PrincipalKind) *PermissionUpdate {
	pu.mutation.SetPrincipalKind(pk)
	return pu
}

// SetUserID sets the "user_id" field.
func (pu *PermissionUpdate) SetUserID(i int) *PermissionUpdate {
	pu.mutation.SetUserID(i)
	return pu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableUserID(i *int) *PermissionUpdate {
	if i != nil {
		pu.SetUserID(*i)
	}
	return pu
}

// ClearUserID clears the value of the "user_id" field.
func (pu *PermissionUpdate) ClearUserID() *PermissionUpdate {
	pu.mutation.ClearUserID()
	return pu
}

// SetRoleID sets the "role_id" field.
func (pu *PermissionUpdate) SetRoleID(i int) *PermissionUpdate {
	pu.mutation.ResetRoleID()
	pu.mutation.SetRoleID(i)
	return pu
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableRoleID(i *int) *PermissionUpdate {
	if i != nil {
		pu.SetRoleID(*i)
	}
	return pu
}

// AddRoleID adds i to the "role_id" field.
func (pu *PermissionUpdate) AddRoleID(i int) *PermissionUpdate {
	pu.mutation.AddRoleID(i)
	return pu
}

// ClearRoleID clears the value of the "role_id" field.
func (pu *PermissionUpdate) ClearRoleID() *PermissionUpdate {
	pu.mutation.ClearRoleID()
	return pu
}

// SetOrgPolicyID sets the "org_policy_id" field.
func (pu *PermissionUpdate) SetOrgPolicyID(i int) *PermissionUpdate {
	pu.mutation.ResetOrgPolicyID()
	pu.mutation.SetOrgPolicyID(i)
	return pu
}

// AddOrgPolicyID adds i to the "org_policy_id" field.
func (pu *PermissionUpdate) AddOrgPolicyID(i int) *PermissionUpdate {
	pu.mutation.AddOrgPolicyID(i)
	return pu
}

// SetStartAt sets the "start_at" field.
func (pu *PermissionUpdate) SetStartAt(t time.Time) *PermissionUpdate {
	pu.mutation.SetStartAt(t)
	return pu
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableStartAt(t *time.Time) *PermissionUpdate {
	if t != nil {
		pu.SetStartAt(*t)
	}
	return pu
}

// ClearStartAt clears the value of the "start_at" field.
func (pu *PermissionUpdate) ClearStartAt() *PermissionUpdate {
	pu.mutation.ClearStartAt()
	return pu
}

// SetEndAt sets the "end_at" field.
func (pu *PermissionUpdate) SetEndAt(t time.Time) *PermissionUpdate {
	pu.mutation.SetEndAt(t)
	return pu
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableEndAt(t *time.Time) *PermissionUpdate {
	if t != nil {
		pu.SetEndAt(*t)
	}
	return pu
}

// ClearEndAt clears the value of the "end_at" field.
func (pu *PermissionUpdate) ClearEndAt() *PermissionUpdate {
	pu.mutation.ClearEndAt()
	return pu
}

// SetStatus sets the "status" field.
func (pu *PermissionUpdate) SetStatus(ts types.SimpleStatus) *PermissionUpdate {
	pu.mutation.SetStatus(ts)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableStatus(ts *types.SimpleStatus) *PermissionUpdate {
	if ts != nil {
		pu.SetStatus(*ts)
	}
	return pu
}

// ClearStatus clears the value of the "status" field.
func (pu *PermissionUpdate) ClearStatus() *PermissionUpdate {
	pu.mutation.ClearStatus()
	return pu
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (pu *PermissionUpdate) SetOrganizationID(id int) *PermissionUpdate {
	pu.mutation.SetOrganizationID(id)
	return pu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (pu *PermissionUpdate) SetOrganization(o *Organization) *PermissionUpdate {
	return pu.SetOrganizationID(o.ID)
}

// SetUser sets the "user" edge to the User entity.
func (pu *PermissionUpdate) SetUser(u *User) *PermissionUpdate {
	return pu.SetUserID(u.ID)
}

// Mutation returns the PermissionMutation object of the builder.
func (pu *PermissionUpdate) Mutation() *PermissionMutation {
	return pu.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (pu *PermissionUpdate) ClearOrganization() *PermissionUpdate {
	pu.mutation.ClearOrganization()
	return pu
}

// ClearUser clears the "user" edge to the User entity.
func (pu *PermissionUpdate) ClearUser() *PermissionUpdate {
	pu.mutation.ClearUser()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PermissionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PermissionMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PermissionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PermissionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PermissionUpdate) check() error {
	if v, ok := pu.mutation.PrincipalKind(); ok {
		if err := permission.PrincipalKindValidator(v); err != nil {
			return &ValidationError{Name: "principal_kind", err: fmt.Errorf(`ent: validator failed for field "Permission.principal_kind": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Status(); ok {
		if err := permission.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Permission.status": %w`, err)}
		}
	}
	if _, ok := pu.mutation.OrganizationID(); pu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Permission.organization"`)
	}
	return nil
}

func (pu *PermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedBy(); ok {
		_spec.SetField(permission.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(permission.FieldUpdatedBy, field.TypeInt, value)
	}
	if pu.mutation.UpdatedByCleared() {
		_spec.ClearField(permission.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.UpdatedAtCleared() {
		_spec.ClearField(permission.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := pu.mutation.PrincipalKind(); ok {
		_spec.SetField(permission.FieldPrincipalKind, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.RoleID(); ok {
		_spec.SetField(permission.FieldRoleID, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedRoleID(); ok {
		_spec.AddField(permission.FieldRoleID, field.TypeInt, value)
	}
	if pu.mutation.RoleIDCleared() {
		_spec.ClearField(permission.FieldRoleID, field.TypeInt)
	}
	if value, ok := pu.mutation.OrgPolicyID(); ok {
		_spec.SetField(permission.FieldOrgPolicyID, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedOrgPolicyID(); ok {
		_spec.AddField(permission.FieldOrgPolicyID, field.TypeInt, value)
	}
	if value, ok := pu.mutation.StartAt(); ok {
		_spec.SetField(permission.FieldStartAt, field.TypeTime, value)
	}
	if pu.mutation.StartAtCleared() {
		_spec.ClearField(permission.FieldStartAt, field.TypeTime)
	}
	if value, ok := pu.mutation.EndAt(); ok {
		_spec.SetField(permission.FieldEndAt, field.TypeTime, value)
	}
	if pu.mutation.EndAtCleared() {
		_spec.ClearField(permission.FieldEndAt, field.TypeTime)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(permission.FieldStatus, field.TypeEnum, value)
	}
	if pu.mutation.StatusCleared() {
		_spec.ClearField(permission.FieldStatus, field.TypeEnum)
	}
	if pu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permission.OrganizationTable,
			Columns: []string{permission.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permission.OrganizationTable,
			Columns: []string{permission.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   permission.UserTable,
			Columns: []string{permission.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   permission.UserTable,
			Columns: []string{permission.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PermissionUpdateOne is the builder for updating a single Permission entity.
type PermissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PermissionMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (puo *PermissionUpdateOne) SetUpdatedBy(i int) *PermissionUpdateOne {
	puo.mutation.ResetUpdatedBy()
	puo.mutation.SetUpdatedBy(i)
	return puo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableUpdatedBy(i *int) *PermissionUpdateOne {
	if i != nil {
		puo.SetUpdatedBy(*i)
	}
	return puo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (puo *PermissionUpdateOne) AddUpdatedBy(i int) *PermissionUpdateOne {
	puo.mutation.AddUpdatedBy(i)
	return puo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (puo *PermissionUpdateOne) ClearUpdatedBy() *PermissionUpdateOne {
	puo.mutation.ClearUpdatedBy()
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PermissionUpdateOne) SetUpdatedAt(t time.Time) *PermissionUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableUpdatedAt(t *time.Time) *PermissionUpdateOne {
	if t != nil {
		puo.SetUpdatedAt(*t)
	}
	return puo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (puo *PermissionUpdateOne) ClearUpdatedAt() *PermissionUpdateOne {
	puo.mutation.ClearUpdatedAt()
	return puo
}

// SetOrgID sets the "org_id" field.
func (puo *PermissionUpdateOne) SetOrgID(i int) *PermissionUpdateOne {
	puo.mutation.SetOrgID(i)
	return puo
}

// SetPrincipalKind sets the "principal_kind" field.
func (puo *PermissionUpdateOne) SetPrincipalKind(pk permission.PrincipalKind) *PermissionUpdateOne {
	puo.mutation.SetPrincipalKind(pk)
	return puo
}

// SetUserID sets the "user_id" field.
func (puo *PermissionUpdateOne) SetUserID(i int) *PermissionUpdateOne {
	puo.mutation.SetUserID(i)
	return puo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableUserID(i *int) *PermissionUpdateOne {
	if i != nil {
		puo.SetUserID(*i)
	}
	return puo
}

// ClearUserID clears the value of the "user_id" field.
func (puo *PermissionUpdateOne) ClearUserID() *PermissionUpdateOne {
	puo.mutation.ClearUserID()
	return puo
}

// SetRoleID sets the "role_id" field.
func (puo *PermissionUpdateOne) SetRoleID(i int) *PermissionUpdateOne {
	puo.mutation.ResetRoleID()
	puo.mutation.SetRoleID(i)
	return puo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableRoleID(i *int) *PermissionUpdateOne {
	if i != nil {
		puo.SetRoleID(*i)
	}
	return puo
}

// AddRoleID adds i to the "role_id" field.
func (puo *PermissionUpdateOne) AddRoleID(i int) *PermissionUpdateOne {
	puo.mutation.AddRoleID(i)
	return puo
}

// ClearRoleID clears the value of the "role_id" field.
func (puo *PermissionUpdateOne) ClearRoleID() *PermissionUpdateOne {
	puo.mutation.ClearRoleID()
	return puo
}

// SetOrgPolicyID sets the "org_policy_id" field.
func (puo *PermissionUpdateOne) SetOrgPolicyID(i int) *PermissionUpdateOne {
	puo.mutation.ResetOrgPolicyID()
	puo.mutation.SetOrgPolicyID(i)
	return puo
}

// AddOrgPolicyID adds i to the "org_policy_id" field.
func (puo *PermissionUpdateOne) AddOrgPolicyID(i int) *PermissionUpdateOne {
	puo.mutation.AddOrgPolicyID(i)
	return puo
}

// SetStartAt sets the "start_at" field.
func (puo *PermissionUpdateOne) SetStartAt(t time.Time) *PermissionUpdateOne {
	puo.mutation.SetStartAt(t)
	return puo
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableStartAt(t *time.Time) *PermissionUpdateOne {
	if t != nil {
		puo.SetStartAt(*t)
	}
	return puo
}

// ClearStartAt clears the value of the "start_at" field.
func (puo *PermissionUpdateOne) ClearStartAt() *PermissionUpdateOne {
	puo.mutation.ClearStartAt()
	return puo
}

// SetEndAt sets the "end_at" field.
func (puo *PermissionUpdateOne) SetEndAt(t time.Time) *PermissionUpdateOne {
	puo.mutation.SetEndAt(t)
	return puo
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableEndAt(t *time.Time) *PermissionUpdateOne {
	if t != nil {
		puo.SetEndAt(*t)
	}
	return puo
}

// ClearEndAt clears the value of the "end_at" field.
func (puo *PermissionUpdateOne) ClearEndAt() *PermissionUpdateOne {
	puo.mutation.ClearEndAt()
	return puo
}

// SetStatus sets the "status" field.
func (puo *PermissionUpdateOne) SetStatus(ts types.SimpleStatus) *PermissionUpdateOne {
	puo.mutation.SetStatus(ts)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableStatus(ts *types.SimpleStatus) *PermissionUpdateOne {
	if ts != nil {
		puo.SetStatus(*ts)
	}
	return puo
}

// ClearStatus clears the value of the "status" field.
func (puo *PermissionUpdateOne) ClearStatus() *PermissionUpdateOne {
	puo.mutation.ClearStatus()
	return puo
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (puo *PermissionUpdateOne) SetOrganizationID(id int) *PermissionUpdateOne {
	puo.mutation.SetOrganizationID(id)
	return puo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (puo *PermissionUpdateOne) SetOrganization(o *Organization) *PermissionUpdateOne {
	return puo.SetOrganizationID(o.ID)
}

// SetUser sets the "user" edge to the User entity.
func (puo *PermissionUpdateOne) SetUser(u *User) *PermissionUpdateOne {
	return puo.SetUserID(u.ID)
}

// Mutation returns the PermissionMutation object of the builder.
func (puo *PermissionUpdateOne) Mutation() *PermissionMutation {
	return puo.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (puo *PermissionUpdateOne) ClearOrganization() *PermissionUpdateOne {
	puo.mutation.ClearOrganization()
	return puo
}

// ClearUser clears the "user" edge to the User entity.
func (puo *PermissionUpdateOne) ClearUser() *PermissionUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// Where appends a list predicates to the PermissionUpdate builder.
func (puo *PermissionUpdateOne) Where(ps ...predicate.Permission) *PermissionUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PermissionUpdateOne) Select(field string, fields ...string) *PermissionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Permission entity.
func (puo *PermissionUpdateOne) Save(ctx context.Context) (*Permission, error) {
	return withHooks[*Permission, PermissionMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PermissionUpdateOne) SaveX(ctx context.Context) *Permission {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PermissionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PermissionUpdateOne) check() error {
	if v, ok := puo.mutation.PrincipalKind(); ok {
		if err := permission.PrincipalKindValidator(v); err != nil {
			return &ValidationError{Name: "principal_kind", err: fmt.Errorf(`ent: validator failed for field "Permission.principal_kind": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Status(); ok {
		if err := permission.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Permission.status": %w`, err)}
		}
	}
	if _, ok := puo.mutation.OrganizationID(); puo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Permission.organization"`)
	}
	return nil
}

func (puo *PermissionUpdateOne) sqlSave(ctx context.Context) (_node *Permission, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Permission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permission.FieldID)
		for _, f := range fields {
			if !permission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != permission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedBy(); ok {
		_spec.SetField(permission.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(permission.FieldUpdatedBy, field.TypeInt, value)
	}
	if puo.mutation.UpdatedByCleared() {
		_spec.ClearField(permission.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.UpdatedAtCleared() {
		_spec.ClearField(permission.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := puo.mutation.PrincipalKind(); ok {
		_spec.SetField(permission.FieldPrincipalKind, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.RoleID(); ok {
		_spec.SetField(permission.FieldRoleID, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedRoleID(); ok {
		_spec.AddField(permission.FieldRoleID, field.TypeInt, value)
	}
	if puo.mutation.RoleIDCleared() {
		_spec.ClearField(permission.FieldRoleID, field.TypeInt)
	}
	if value, ok := puo.mutation.OrgPolicyID(); ok {
		_spec.SetField(permission.FieldOrgPolicyID, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedOrgPolicyID(); ok {
		_spec.AddField(permission.FieldOrgPolicyID, field.TypeInt, value)
	}
	if value, ok := puo.mutation.StartAt(); ok {
		_spec.SetField(permission.FieldStartAt, field.TypeTime, value)
	}
	if puo.mutation.StartAtCleared() {
		_spec.ClearField(permission.FieldStartAt, field.TypeTime)
	}
	if value, ok := puo.mutation.EndAt(); ok {
		_spec.SetField(permission.FieldEndAt, field.TypeTime, value)
	}
	if puo.mutation.EndAtCleared() {
		_spec.ClearField(permission.FieldEndAt, field.TypeTime)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(permission.FieldStatus, field.TypeEnum, value)
	}
	if puo.mutation.StatusCleared() {
		_spec.ClearField(permission.FieldStatus, field.TypeEnum)
	}
	if puo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permission.OrganizationTable,
			Columns: []string{permission.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permission.OrganizationTable,
			Columns: []string{permission.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   permission.UserTable,
			Columns: []string{permission.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   permission.UserTable,
			Columns: []string{permission.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Permission{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
