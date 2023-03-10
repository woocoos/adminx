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
	"github.com/woocoos/adminx/ent/organization"
	"github.com/woocoos/adminx/ent/organizationrole"
	"github.com/woocoos/adminx/ent/organizationuser"
	"github.com/woocoos/adminx/ent/permission"
	"github.com/woocoos/adminx/ent/permissionpolicy"
	"github.com/woocoos/adminx/ent/user"
	"github.com/woocoos/adminx/graph/entgen/types"
)

// OrganizationCreate is the builder for creating a Organization entity.
type OrganizationCreate struct {
	config
	mutation *OrganizationMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (oc *OrganizationCreate) SetCreatedBy(i int) *OrganizationCreate {
	oc.mutation.SetCreatedBy(i)
	return oc
}

// SetCreatedAt sets the "created_at" field.
func (oc *OrganizationCreate) SetCreatedAt(t time.Time) *OrganizationCreate {
	oc.mutation.SetCreatedAt(t)
	return oc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableCreatedAt(t *time.Time) *OrganizationCreate {
	if t != nil {
		oc.SetCreatedAt(*t)
	}
	return oc
}

// SetUpdatedBy sets the "updated_by" field.
func (oc *OrganizationCreate) SetUpdatedBy(i int) *OrganizationCreate {
	oc.mutation.SetUpdatedBy(i)
	return oc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableUpdatedBy(i *int) *OrganizationCreate {
	if i != nil {
		oc.SetUpdatedBy(*i)
	}
	return oc
}

// SetUpdatedAt sets the "updated_at" field.
func (oc *OrganizationCreate) SetUpdatedAt(t time.Time) *OrganizationCreate {
	oc.mutation.SetUpdatedAt(t)
	return oc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableUpdatedAt(t *time.Time) *OrganizationCreate {
	if t != nil {
		oc.SetUpdatedAt(*t)
	}
	return oc
}

// SetDeletedAt sets the "deleted_at" field.
func (oc *OrganizationCreate) SetDeletedAt(t time.Time) *OrganizationCreate {
	oc.mutation.SetDeletedAt(t)
	return oc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableDeletedAt(t *time.Time) *OrganizationCreate {
	if t != nil {
		oc.SetDeletedAt(*t)
	}
	return oc
}

// SetOwnerID sets the "owner_id" field.
func (oc *OrganizationCreate) SetOwnerID(i int) *OrganizationCreate {
	oc.mutation.SetOwnerID(i)
	return oc
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableOwnerID(i *int) *OrganizationCreate {
	if i != nil {
		oc.SetOwnerID(*i)
	}
	return oc
}

// SetKind sets the "kind" field.
func (oc *OrganizationCreate) SetKind(o organization.Kind) *OrganizationCreate {
	oc.mutation.SetKind(o)
	return oc
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableKind(o *organization.Kind) *OrganizationCreate {
	if o != nil {
		oc.SetKind(*o)
	}
	return oc
}

// SetParentID sets the "parent_id" field.
func (oc *OrganizationCreate) SetParentID(i int) *OrganizationCreate {
	oc.mutation.SetParentID(i)
	return oc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableParentID(i *int) *OrganizationCreate {
	if i != nil {
		oc.SetParentID(*i)
	}
	return oc
}

// SetDomain sets the "domain" field.
func (oc *OrganizationCreate) SetDomain(s string) *OrganizationCreate {
	oc.mutation.SetDomain(s)
	return oc
}

// SetNillableDomain sets the "domain" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableDomain(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetDomain(*s)
	}
	return oc
}

// SetCode sets the "code" field.
func (oc *OrganizationCreate) SetCode(s string) *OrganizationCreate {
	oc.mutation.SetCode(s)
	return oc
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableCode(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetCode(*s)
	}
	return oc
}

// SetName sets the "name" field.
func (oc *OrganizationCreate) SetName(s string) *OrganizationCreate {
	oc.mutation.SetName(s)
	return oc
}

// SetProfile sets the "profile" field.
func (oc *OrganizationCreate) SetProfile(s string) *OrganizationCreate {
	oc.mutation.SetProfile(s)
	return oc
}

// SetNillableProfile sets the "profile" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableProfile(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetProfile(*s)
	}
	return oc
}

// SetStatus sets the "status" field.
func (oc *OrganizationCreate) SetStatus(ts types.SimpleStatus) *OrganizationCreate {
	oc.mutation.SetStatus(ts)
	return oc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableStatus(ts *types.SimpleStatus) *OrganizationCreate {
	if ts != nil {
		oc.SetStatus(*ts)
	}
	return oc
}

// SetPath sets the "path" field.
func (oc *OrganizationCreate) SetPath(s string) *OrganizationCreate {
	oc.mutation.SetPath(s)
	return oc
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillablePath(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetPath(*s)
	}
	return oc
}

// SetDisplaySort sets the "display_sort" field.
func (oc *OrganizationCreate) SetDisplaySort(i int32) *OrganizationCreate {
	oc.mutation.SetDisplaySort(i)
	return oc
}

// SetNillableDisplaySort sets the "display_sort" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableDisplaySort(i *int32) *OrganizationCreate {
	if i != nil {
		oc.SetDisplaySort(*i)
	}
	return oc
}

// SetCountryCode sets the "country_code" field.
func (oc *OrganizationCreate) SetCountryCode(s string) *OrganizationCreate {
	oc.mutation.SetCountryCode(s)
	return oc
}

// SetNillableCountryCode sets the "country_code" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableCountryCode(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetCountryCode(*s)
	}
	return oc
}

// SetTimezone sets the "timezone" field.
func (oc *OrganizationCreate) SetTimezone(s string) *OrganizationCreate {
	oc.mutation.SetTimezone(s)
	return oc
}

// SetNillableTimezone sets the "timezone" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableTimezone(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetTimezone(*s)
	}
	return oc
}

// SetID sets the "id" field.
func (oc *OrganizationCreate) SetID(i int) *OrganizationCreate {
	oc.mutation.SetID(i)
	return oc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableID(i *int) *OrganizationCreate {
	if i != nil {
		oc.SetID(*i)
	}
	return oc
}

// SetParent sets the "parent" edge to the Organization entity.
func (oc *OrganizationCreate) SetParent(o *Organization) *OrganizationCreate {
	return oc.SetParentID(o.ID)
}

// AddChildIDs adds the "children" edge to the Organization entity by IDs.
func (oc *OrganizationCreate) AddChildIDs(ids ...int) *OrganizationCreate {
	oc.mutation.AddChildIDs(ids...)
	return oc
}

// AddChildren adds the "children" edges to the Organization entity.
func (oc *OrganizationCreate) AddChildren(o ...*Organization) *OrganizationCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddChildIDs(ids...)
}

// SetOwner sets the "owner" edge to the User entity.
func (oc *OrganizationCreate) SetOwner(u *User) *OrganizationCreate {
	return oc.SetOwnerID(u.ID)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (oc *OrganizationCreate) AddUserIDs(ids ...int) *OrganizationCreate {
	oc.mutation.AddUserIDs(ids...)
	return oc
}

// AddUsers adds the "users" edges to the User entity.
func (oc *OrganizationCreate) AddUsers(u ...*User) *OrganizationCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return oc.AddUserIDs(ids...)
}

// AddRolesAndGroupIDs adds the "rolesAndGroups" edge to the OrganizationRole entity by IDs.
func (oc *OrganizationCreate) AddRolesAndGroupIDs(ids ...int) *OrganizationCreate {
	oc.mutation.AddRolesAndGroupIDs(ids...)
	return oc
}

// AddRolesAndGroups adds the "rolesAndGroups" edges to the OrganizationRole entity.
func (oc *OrganizationCreate) AddRolesAndGroups(o ...*OrganizationRole) *OrganizationCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddRolesAndGroupIDs(ids...)
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (oc *OrganizationCreate) AddPermissionIDs(ids ...int) *OrganizationCreate {
	oc.mutation.AddPermissionIDs(ids...)
	return oc
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (oc *OrganizationCreate) AddPermissions(p ...*Permission) *OrganizationCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return oc.AddPermissionIDs(ids...)
}

// AddPolicyIDs adds the "policies" edge to the PermissionPolicy entity by IDs.
func (oc *OrganizationCreate) AddPolicyIDs(ids ...int) *OrganizationCreate {
	oc.mutation.AddPolicyIDs(ids...)
	return oc
}

// AddPolicies adds the "policies" edges to the PermissionPolicy entity.
func (oc *OrganizationCreate) AddPolicies(p ...*PermissionPolicy) *OrganizationCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return oc.AddPolicyIDs(ids...)
}

// AddAppIDs adds the "apps" edge to the App entity by IDs.
func (oc *OrganizationCreate) AddAppIDs(ids ...int) *OrganizationCreate {
	oc.mutation.AddAppIDs(ids...)
	return oc
}

// AddApps adds the "apps" edges to the App entity.
func (oc *OrganizationCreate) AddApps(a ...*App) *OrganizationCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return oc.AddAppIDs(ids...)
}

// AddOrganizationUserIDs adds the "organization_user" edge to the OrganizationUser entity by IDs.
func (oc *OrganizationCreate) AddOrganizationUserIDs(ids ...int) *OrganizationCreate {
	oc.mutation.AddOrganizationUserIDs(ids...)
	return oc
}

// AddOrganizationUser adds the "organization_user" edges to the OrganizationUser entity.
func (oc *OrganizationCreate) AddOrganizationUser(o ...*OrganizationUser) *OrganizationCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddOrganizationUserIDs(ids...)
}

// Mutation returns the OrganizationMutation object of the builder.
func (oc *OrganizationCreate) Mutation() *OrganizationMutation {
	return oc.mutation
}

// Save creates the Organization in the database.
func (oc *OrganizationCreate) Save(ctx context.Context) (*Organization, error) {
	if err := oc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*Organization, OrganizationMutation](ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrganizationCreate) SaveX(ctx context.Context) *Organization {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrganizationCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrganizationCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrganizationCreate) defaults() error {
	if _, ok := oc.mutation.CreatedAt(); !ok {
		if organization.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized organization.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := organization.DefaultCreatedAt()
		oc.mutation.SetCreatedAt(v)
	}
	if _, ok := oc.mutation.ParentID(); !ok {
		v := organization.DefaultParentID
		oc.mutation.SetParentID(v)
	}
	if _, ok := oc.mutation.ID(); !ok {
		if organization.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized organization.DefaultID (forgotten import ent/runtime?)")
		}
		v := organization.DefaultID()
		oc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrganizationCreate) check() error {
	if _, ok := oc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Organization.created_by"`)}
	}
	if _, ok := oc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Organization.created_at"`)}
	}
	if v, ok := oc.mutation.Kind(); ok {
		if err := organization.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "Organization.kind": %w`, err)}
		}
	}
	if _, ok := oc.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent_id", err: errors.New(`ent: missing required field "Organization.parent_id"`)}
	}
	if v, ok := oc.mutation.Domain(); ok {
		if err := organization.DomainValidator(v); err != nil {
			return &ValidationError{Name: "domain", err: fmt.Errorf(`ent: validator failed for field "Organization.domain": %w`, err)}
		}
	}
	if v, ok := oc.mutation.Code(); ok {
		if err := organization.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Organization.code": %w`, err)}
		}
	}
	if _, ok := oc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Organization.name"`)}
	}
	if v, ok := oc.mutation.Name(); ok {
		if err := organization.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Organization.name": %w`, err)}
		}
	}
	if v, ok := oc.mutation.Status(); ok {
		if err := organization.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Organization.status": %w`, err)}
		}
	}
	if v, ok := oc.mutation.CountryCode(); ok {
		if err := organization.CountryCodeValidator(v); err != nil {
			return &ValidationError{Name: "country_code", err: fmt.Errorf(`ent: validator failed for field "Organization.country_code": %w`, err)}
		}
	}
	if v, ok := oc.mutation.Timezone(); ok {
		if err := organization.TimezoneValidator(v); err != nil {
			return &ValidationError{Name: "timezone", err: fmt.Errorf(`ent: validator failed for field "Organization.timezone": %w`, err)}
		}
	}
	if _, ok := oc.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent", err: errors.New(`ent: missing required edge "Organization.parent"`)}
	}
	return nil
}

func (oc *OrganizationCreate) sqlSave(ctx context.Context) (*Organization, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrganizationCreate) createSpec() (*Organization, *sqlgraph.CreateSpec) {
	var (
		_node = &Organization{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(organization.Table, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt))
	)
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oc.mutation.CreatedBy(); ok {
		_spec.SetField(organization.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := oc.mutation.CreatedAt(); ok {
		_spec.SetField(organization.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oc.mutation.UpdatedBy(); ok {
		_spec.SetField(organization.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := oc.mutation.UpdatedAt(); ok {
		_spec.SetField(organization.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := oc.mutation.DeletedAt(); ok {
		_spec.SetField(organization.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := oc.mutation.Kind(); ok {
		_spec.SetField(organization.FieldKind, field.TypeEnum, value)
		_node.Kind = value
	}
	if value, ok := oc.mutation.Domain(); ok {
		_spec.SetField(organization.FieldDomain, field.TypeString, value)
		_node.Domain = value
	}
	if value, ok := oc.mutation.Code(); ok {
		_spec.SetField(organization.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := oc.mutation.Name(); ok {
		_spec.SetField(organization.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := oc.mutation.Profile(); ok {
		_spec.SetField(organization.FieldProfile, field.TypeString, value)
		_node.Profile = value
	}
	if value, ok := oc.mutation.Status(); ok {
		_spec.SetField(organization.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := oc.mutation.Path(); ok {
		_spec.SetField(organization.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := oc.mutation.DisplaySort(); ok {
		_spec.SetField(organization.FieldDisplaySort, field.TypeInt32, value)
		_node.DisplaySort = value
	}
	if value, ok := oc.mutation.CountryCode(); ok {
		_spec.SetField(organization.FieldCountryCode, field.TypeString, value)
		_node.CountryCode = value
	}
	if value, ok := oc.mutation.Timezone(); ok {
		_spec.SetField(organization.FieldTimezone, field.TypeString, value)
		_node.Timezone = value
	}
	if nodes := oc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organization.ParentTable,
			Columns: []string{organization.ParentColumn},
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
		_node.ParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.ChildrenTable,
			Columns: []string{organization.ChildrenColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   organization.OwnerTable,
			Columns: []string{organization.OwnerColumn},
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
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   organization.UsersTable,
			Columns: organization.UsersPrimaryKey,
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
		createE := &OrganizationUserCreate{config: oc.config, mutation: newOrganizationUserMutation(oc.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.RolesAndGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.RolesAndGroupsTable,
			Columns: []string{organization.RolesAndGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organizationrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.PermissionsTable,
			Columns: []string{organization.PermissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: permission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.PoliciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.PoliciesTable,
			Columns: []string{organization.PoliciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: permissionpolicy.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.AppsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   organization.AppsTable,
			Columns: organization.AppsPrimaryKey,
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
		createE := &OrganizationAppCreate{config: oc.config, mutation: newOrganizationAppMutation(oc.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.OrganizationUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   organization.OrganizationUserTable,
			Columns: []string{organization.OrganizationUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organizationuser.FieldID,
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

// OrganizationCreateBulk is the builder for creating many Organization entities in bulk.
type OrganizationCreateBulk struct {
	config
	builders []*OrganizationCreate
}

// Save creates the Organization entities in the database.
func (ocb *OrganizationCreateBulk) Save(ctx context.Context) ([]*Organization, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Organization, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrganizationCreateBulk) SaveX(ctx context.Context) []*Organization {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrganizationCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrganizationCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
