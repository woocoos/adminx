// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/adminx/ent/organization"
	"github.com/woocoos/adminx/ent/user"
	"github.com/woocoos/adminx/graph/entgen/types"
)

// Organization is the model entity for the Organization schema.
type Organization struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy int `json:"created_by,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy int `json:"updated_by,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// 管理账户ID
	OwnerID int `json:"owner_id,omitempty"`
	// 分类: 根节点,组织节点
	Kind organization.Kind `json:"kind,omitempty"`
	// 父级ID,0为根组织.
	ParentID int `json:"parent_id,omitempty"`
	// 默认域名
	Domain string `json:"domain,omitempty"`
	// 系统代码
	Code string `json:"code,omitempty"`
	// 组织名称
	Name string `json:"name,omitempty"`
	// 简介
	Profile string `json:"profile,omitempty"`
	// 状态
	Status types.SimpleStatus `json:"status,omitempty"`
	// 路径编码
	Path string `json:"path,omitempty"`
	// DisplaySort holds the value of the "display_sort" field.
	DisplaySort int32 `json:"display_sort,omitempty"`
	// 国家或地区2字码
	CountryCode string `json:"country_code,omitempty"`
	// 时区
	Timezone string `json:"timezone,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrganizationQuery when eager-loading is set.
	Edges OrganizationEdges `json:"edges"`
}

// OrganizationEdges holds the relations/edges for other nodes in the graph.
type OrganizationEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Organization `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Organization `json:"children,omitempty"`
	// 管理账户
	Owner *User `json:"owner,omitempty"`
	// 组织下用户
	Users []*User `json:"users,omitempty"`
	// 组织下角色及用户组.
	RolesAndGroups []*OrganizationRole `json:"rolesAndGroups,omitempty"`
	// 组织授权信息
	Permissions []*Permission `json:"permissions,omitempty"`
	// 组织下权限策略
	Policies []*PermissionPolicy `json:"policies,omitempty"`
	// 组织下应用
	Apps []*App `json:"apps,omitempty"`
	// OrganizationUser holds the value of the organization_user edge.
	OrganizationUser []*OrganizationUser `json:"organization_user,omitempty"`
	// OrganizationApp holds the value of the organization_app edge.
	OrganizationApp []*OrganizationApp `json:"organization_app,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [10]bool
	// totalCount holds the count of the edges above.
	totalCount [6]map[string]int

	namedChildren         map[string][]*Organization
	namedUsers            map[string][]*User
	namedRolesAndGroups   map[string][]*OrganizationRole
	namedPermissions      map[string][]*Permission
	namedPolicies         map[string][]*PermissionPolicy
	namedApps             map[string][]*App
	namedOrganizationUser map[string][]*OrganizationUser
	namedOrganizationApp  map[string][]*OrganizationApp
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrganizationEdges) ParentOrErr() (*Organization, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organization.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) ChildrenOrErr() ([]*Organization, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrganizationEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[3] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// RolesAndGroupsOrErr returns the RolesAndGroups value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) RolesAndGroupsOrErr() ([]*OrganizationRole, error) {
	if e.loadedTypes[4] {
		return e.RolesAndGroups, nil
	}
	return nil, &NotLoadedError{edge: "rolesAndGroups"}
}

// PermissionsOrErr returns the Permissions value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) PermissionsOrErr() ([]*Permission, error) {
	if e.loadedTypes[5] {
		return e.Permissions, nil
	}
	return nil, &NotLoadedError{edge: "permissions"}
}

// PoliciesOrErr returns the Policies value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) PoliciesOrErr() ([]*PermissionPolicy, error) {
	if e.loadedTypes[6] {
		return e.Policies, nil
	}
	return nil, &NotLoadedError{edge: "policies"}
}

// AppsOrErr returns the Apps value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) AppsOrErr() ([]*App, error) {
	if e.loadedTypes[7] {
		return e.Apps, nil
	}
	return nil, &NotLoadedError{edge: "apps"}
}

// OrganizationUserOrErr returns the OrganizationUser value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) OrganizationUserOrErr() ([]*OrganizationUser, error) {
	if e.loadedTypes[8] {
		return e.OrganizationUser, nil
	}
	return nil, &NotLoadedError{edge: "organization_user"}
}

// OrganizationAppOrErr returns the OrganizationApp value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) OrganizationAppOrErr() ([]*OrganizationApp, error) {
	if e.loadedTypes[9] {
		return e.OrganizationApp, nil
	}
	return nil, &NotLoadedError{edge: "organization_app"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Organization) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case organization.FieldID, organization.FieldCreatedBy, organization.FieldUpdatedBy, organization.FieldOwnerID, organization.FieldParentID, organization.FieldDisplaySort:
			values[i] = new(sql.NullInt64)
		case organization.FieldKind, organization.FieldDomain, organization.FieldCode, organization.FieldName, organization.FieldProfile, organization.FieldStatus, organization.FieldPath, organization.FieldCountryCode, organization.FieldTimezone:
			values[i] = new(sql.NullString)
		case organization.FieldCreatedAt, organization.FieldUpdatedAt, organization.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Organization", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Organization fields.
func (o *Organization) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case organization.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case organization.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				o.CreatedBy = int(value.Int64)
			}
		case organization.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case organization.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				o.UpdatedBy = int(value.Int64)
			}
		case organization.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				o.UpdatedAt = value.Time
			}
		case organization.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				o.DeletedAt = value.Time
			}
		case organization.FieldOwnerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				o.OwnerID = int(value.Int64)
			}
		case organization.FieldKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[i])
			} else if value.Valid {
				o.Kind = organization.Kind(value.String)
			}
		case organization.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				o.ParentID = int(value.Int64)
			}
		case organization.FieldDomain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field domain", values[i])
			} else if value.Valid {
				o.Domain = value.String
			}
		case organization.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				o.Code = value.String
			}
		case organization.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				o.Name = value.String
			}
		case organization.FieldProfile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field profile", values[i])
			} else if value.Valid {
				o.Profile = value.String
			}
		case organization.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				o.Status = types.SimpleStatus(value.String)
			}
		case organization.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				o.Path = value.String
			}
		case organization.FieldDisplaySort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field display_sort", values[i])
			} else if value.Valid {
				o.DisplaySort = int32(value.Int64)
			}
		case organization.FieldCountryCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country_code", values[i])
			} else if value.Valid {
				o.CountryCode = value.String
			}
		case organization.FieldTimezone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field timezone", values[i])
			} else if value.Valid {
				o.Timezone = value.String
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the Organization entity.
func (o *Organization) QueryParent() *OrganizationQuery {
	return NewOrganizationClient(o.config).QueryParent(o)
}

// QueryChildren queries the "children" edge of the Organization entity.
func (o *Organization) QueryChildren() *OrganizationQuery {
	return NewOrganizationClient(o.config).QueryChildren(o)
}

// QueryOwner queries the "owner" edge of the Organization entity.
func (o *Organization) QueryOwner() *UserQuery {
	return NewOrganizationClient(o.config).QueryOwner(o)
}

// QueryUsers queries the "users" edge of the Organization entity.
func (o *Organization) QueryUsers() *UserQuery {
	return NewOrganizationClient(o.config).QueryUsers(o)
}

// QueryRolesAndGroups queries the "rolesAndGroups" edge of the Organization entity.
func (o *Organization) QueryRolesAndGroups() *OrganizationRoleQuery {
	return NewOrganizationClient(o.config).QueryRolesAndGroups(o)
}

// QueryPermissions queries the "permissions" edge of the Organization entity.
func (o *Organization) QueryPermissions() *PermissionQuery {
	return NewOrganizationClient(o.config).QueryPermissions(o)
}

// QueryPolicies queries the "policies" edge of the Organization entity.
func (o *Organization) QueryPolicies() *PermissionPolicyQuery {
	return NewOrganizationClient(o.config).QueryPolicies(o)
}

// QueryApps queries the "apps" edge of the Organization entity.
func (o *Organization) QueryApps() *AppQuery {
	return NewOrganizationClient(o.config).QueryApps(o)
}

// QueryOrganizationUser queries the "organization_user" edge of the Organization entity.
func (o *Organization) QueryOrganizationUser() *OrganizationUserQuery {
	return NewOrganizationClient(o.config).QueryOrganizationUser(o)
}

// QueryOrganizationApp queries the "organization_app" edge of the Organization entity.
func (o *Organization) QueryOrganizationApp() *OrganizationAppQuery {
	return NewOrganizationClient(o.config).QueryOrganizationApp(o)
}

// Update returns a builder for updating this Organization.
// Note that you need to call Organization.Unwrap() before calling this method if this Organization
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Organization) Update() *OrganizationUpdateOne {
	return NewOrganizationClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Organization entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Organization) Unwrap() *Organization {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Organization is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Organization) String() string {
	var builder strings.Builder
	builder.WriteString("Organization(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", o.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", o.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(o.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(o.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(fmt.Sprintf("%v", o.OwnerID))
	builder.WriteString(", ")
	builder.WriteString("kind=")
	builder.WriteString(fmt.Sprintf("%v", o.Kind))
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", o.ParentID))
	builder.WriteString(", ")
	builder.WriteString("domain=")
	builder.WriteString(o.Domain)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(o.Code)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(o.Name)
	builder.WriteString(", ")
	builder.WriteString("profile=")
	builder.WriteString(o.Profile)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", o.Status))
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(o.Path)
	builder.WriteString(", ")
	builder.WriteString("display_sort=")
	builder.WriteString(fmt.Sprintf("%v", o.DisplaySort))
	builder.WriteString(", ")
	builder.WriteString("country_code=")
	builder.WriteString(o.CountryCode)
	builder.WriteString(", ")
	builder.WriteString("timezone=")
	builder.WriteString(o.Timezone)
	builder.WriteByte(')')
	return builder.String()
}

// NamedChildren returns the Children named value or an error if the edge was not
// loaded in eager-loading with this name.
func (o *Organization) NamedChildren(name string) ([]*Organization, error) {
	if o.Edges.namedChildren == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := o.Edges.namedChildren[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (o *Organization) appendNamedChildren(name string, edges ...*Organization) {
	if o.Edges.namedChildren == nil {
		o.Edges.namedChildren = make(map[string][]*Organization)
	}
	if len(edges) == 0 {
		o.Edges.namedChildren[name] = []*Organization{}
	} else {
		o.Edges.namedChildren[name] = append(o.Edges.namedChildren[name], edges...)
	}
}

// NamedUsers returns the Users named value or an error if the edge was not
// loaded in eager-loading with this name.
func (o *Organization) NamedUsers(name string) ([]*User, error) {
	if o.Edges.namedUsers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := o.Edges.namedUsers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (o *Organization) appendNamedUsers(name string, edges ...*User) {
	if o.Edges.namedUsers == nil {
		o.Edges.namedUsers = make(map[string][]*User)
	}
	if len(edges) == 0 {
		o.Edges.namedUsers[name] = []*User{}
	} else {
		o.Edges.namedUsers[name] = append(o.Edges.namedUsers[name], edges...)
	}
}

// NamedRolesAndGroups returns the RolesAndGroups named value or an error if the edge was not
// loaded in eager-loading with this name.
func (o *Organization) NamedRolesAndGroups(name string) ([]*OrganizationRole, error) {
	if o.Edges.namedRolesAndGroups == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := o.Edges.namedRolesAndGroups[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (o *Organization) appendNamedRolesAndGroups(name string, edges ...*OrganizationRole) {
	if o.Edges.namedRolesAndGroups == nil {
		o.Edges.namedRolesAndGroups = make(map[string][]*OrganizationRole)
	}
	if len(edges) == 0 {
		o.Edges.namedRolesAndGroups[name] = []*OrganizationRole{}
	} else {
		o.Edges.namedRolesAndGroups[name] = append(o.Edges.namedRolesAndGroups[name], edges...)
	}
}

// NamedPermissions returns the Permissions named value or an error if the edge was not
// loaded in eager-loading with this name.
func (o *Organization) NamedPermissions(name string) ([]*Permission, error) {
	if o.Edges.namedPermissions == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := o.Edges.namedPermissions[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (o *Organization) appendNamedPermissions(name string, edges ...*Permission) {
	if o.Edges.namedPermissions == nil {
		o.Edges.namedPermissions = make(map[string][]*Permission)
	}
	if len(edges) == 0 {
		o.Edges.namedPermissions[name] = []*Permission{}
	} else {
		o.Edges.namedPermissions[name] = append(o.Edges.namedPermissions[name], edges...)
	}
}

// NamedPolicies returns the Policies named value or an error if the edge was not
// loaded in eager-loading with this name.
func (o *Organization) NamedPolicies(name string) ([]*PermissionPolicy, error) {
	if o.Edges.namedPolicies == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := o.Edges.namedPolicies[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (o *Organization) appendNamedPolicies(name string, edges ...*PermissionPolicy) {
	if o.Edges.namedPolicies == nil {
		o.Edges.namedPolicies = make(map[string][]*PermissionPolicy)
	}
	if len(edges) == 0 {
		o.Edges.namedPolicies[name] = []*PermissionPolicy{}
	} else {
		o.Edges.namedPolicies[name] = append(o.Edges.namedPolicies[name], edges...)
	}
}

// NamedApps returns the Apps named value or an error if the edge was not
// loaded in eager-loading with this name.
func (o *Organization) NamedApps(name string) ([]*App, error) {
	if o.Edges.namedApps == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := o.Edges.namedApps[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (o *Organization) appendNamedApps(name string, edges ...*App) {
	if o.Edges.namedApps == nil {
		o.Edges.namedApps = make(map[string][]*App)
	}
	if len(edges) == 0 {
		o.Edges.namedApps[name] = []*App{}
	} else {
		o.Edges.namedApps[name] = append(o.Edges.namedApps[name], edges...)
	}
}

// NamedOrganizationUser returns the OrganizationUser named value or an error if the edge was not
// loaded in eager-loading with this name.
func (o *Organization) NamedOrganizationUser(name string) ([]*OrganizationUser, error) {
	if o.Edges.namedOrganizationUser == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := o.Edges.namedOrganizationUser[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (o *Organization) appendNamedOrganizationUser(name string, edges ...*OrganizationUser) {
	if o.Edges.namedOrganizationUser == nil {
		o.Edges.namedOrganizationUser = make(map[string][]*OrganizationUser)
	}
	if len(edges) == 0 {
		o.Edges.namedOrganizationUser[name] = []*OrganizationUser{}
	} else {
		o.Edges.namedOrganizationUser[name] = append(o.Edges.namedOrganizationUser[name], edges...)
	}
}

// NamedOrganizationApp returns the OrganizationApp named value or an error if the edge was not
// loaded in eager-loading with this name.
func (o *Organization) NamedOrganizationApp(name string) ([]*OrganizationApp, error) {
	if o.Edges.namedOrganizationApp == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := o.Edges.namedOrganizationApp[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (o *Organization) appendNamedOrganizationApp(name string, edges ...*OrganizationApp) {
	if o.Edges.namedOrganizationApp == nil {
		o.Edges.namedOrganizationApp = make(map[string][]*OrganizationApp)
	}
	if len(edges) == 0 {
		o.Edges.namedOrganizationApp[name] = []*OrganizationApp{}
	} else {
		o.Edges.namedOrganizationApp[name] = append(o.Edges.namedOrganizationApp[name], edges...)
	}
}

// Organizations is a parsable slice of Organization.
type Organizations []*Organization
