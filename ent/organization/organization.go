// Code generated by ent, DO NOT EDIT.

package organization

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent"
	"github.com/99designs/gqlgen/graphql"
	"github.com/woocoos/adminx/graph/entgen/types"
)

const (
	// Label holds the string label denoting the organization type in the database.
	Label = "organization"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// FieldKind holds the string denoting the kind field in the database.
	FieldKind = "kind"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldDomain holds the string denoting the domain field in the database.
	FieldDomain = "domain"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldProfile holds the string denoting the profile field in the database.
	FieldProfile = "profile"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldDisplaySort holds the string denoting the display_sort field in the database.
	FieldDisplaySort = "display_sort"
	// FieldCountryCode holds the string denoting the country_code field in the database.
	FieldCountryCode = "country_code"
	// FieldTimezone holds the string denoting the timezone field in the database.
	FieldTimezone = "timezone"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeRolesAndGroups holds the string denoting the rolesandgroups edge name in mutations.
	EdgeRolesAndGroups = "rolesAndGroups"
	// EdgePermissions holds the string denoting the permissions edge name in mutations.
	EdgePermissions = "permissions"
	// EdgePolicies holds the string denoting the policies edge name in mutations.
	EdgePolicies = "policies"
	// EdgeApps holds the string denoting the apps edge name in mutations.
	EdgeApps = "apps"
	// EdgeOrganizationUser holds the string denoting the organization_user edge name in mutations.
	EdgeOrganizationUser = "organization_user"
	// EdgeOrganizationApp holds the string denoting the organization_app edge name in mutations.
	EdgeOrganizationApp = "organization_app"
	// Table holds the table name of the organization in the database.
	Table = "organization"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "organization"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_id"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "organization"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "parent_id"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "organization"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "user"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "owner_id"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "organization_user"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "user"
	// RolesAndGroupsTable is the table that holds the rolesAndGroups relation/edge.
	RolesAndGroupsTable = "organization_role"
	// RolesAndGroupsInverseTable is the table name for the OrganizationRole entity.
	// It exists in this package in order to avoid circular dependency with the "organizationrole" package.
	RolesAndGroupsInverseTable = "organization_role"
	// RolesAndGroupsColumn is the table column denoting the rolesAndGroups relation/edge.
	RolesAndGroupsColumn = "org_id"
	// PermissionsTable is the table that holds the permissions relation/edge.
	PermissionsTable = "permission"
	// PermissionsInverseTable is the table name for the Permission entity.
	// It exists in this package in order to avoid circular dependency with the "permission" package.
	PermissionsInverseTable = "permission"
	// PermissionsColumn is the table column denoting the permissions relation/edge.
	PermissionsColumn = "org_id"
	// PoliciesTable is the table that holds the policies relation/edge.
	PoliciesTable = "organization_policy"
	// PoliciesInverseTable is the table name for the PermissionPolicy entity.
	// It exists in this package in order to avoid circular dependency with the "permissionpolicy" package.
	PoliciesInverseTable = "organization_policy"
	// PoliciesColumn is the table column denoting the policies relation/edge.
	PoliciesColumn = "org_id"
	// AppsTable is the table that holds the apps relation/edge. The primary key declared below.
	AppsTable = "organization_app"
	// AppsInverseTable is the table name for the App entity.
	// It exists in this package in order to avoid circular dependency with the "app" package.
	AppsInverseTable = "app"
	// OrganizationUserTable is the table that holds the organization_user relation/edge.
	OrganizationUserTable = "organization_user"
	// OrganizationUserInverseTable is the table name for the OrganizationUser entity.
	// It exists in this package in order to avoid circular dependency with the "organizationuser" package.
	OrganizationUserInverseTable = "organization_user"
	// OrganizationUserColumn is the table column denoting the organization_user relation/edge.
	OrganizationUserColumn = "org_id"
	// OrganizationAppTable is the table that holds the organization_app relation/edge.
	OrganizationAppTable = "organization_app"
	// OrganizationAppInverseTable is the table name for the OrganizationApp entity.
	// It exists in this package in order to avoid circular dependency with the "organizationapp" package.
	OrganizationAppInverseTable = "organization_app"
	// OrganizationAppColumn is the table column denoting the organization_app relation/edge.
	OrganizationAppColumn = "org_id"
)

// Columns holds all SQL columns for organization fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldOwnerID,
	FieldKind,
	FieldParentID,
	FieldDomain,
	FieldCode,
	FieldName,
	FieldProfile,
	FieldStatus,
	FieldPath,
	FieldDisplaySort,
	FieldCountryCode,
	FieldTimezone,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"org_id", "user_id"}
	// AppsPrimaryKey and AppsColumn2 are the table columns denoting the
	// primary key for the apps relation (M2M).
	AppsPrimaryKey = []string{"org_id", "app_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/woocoos/adminx/ent/runtime"
var (
	Hooks        [5]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultParentID holds the default value on creation for the "parent_id" field.
	DefaultParentID int
	// DomainValidator is a validator for the "domain" field. It is called by the builders before save.
	DomainValidator func(string) error
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// CountryCodeValidator is a validator for the "country_code" field. It is called by the builders before save.
	CountryCodeValidator func(string) error
	// TimezoneValidator is a validator for the "timezone" field. It is called by the builders before save.
	TimezoneValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int
)

// Kind defines the type for the "kind" enum field.
type Kind string

// Kind values.
const (
	KindRoot         Kind = "root"
	KindOrganization Kind = "org"
)

func (k Kind) String() string {
	return string(k)
}

// KindValidator is a validator for the "kind" field enum values. It is called by the builders before save.
func KindValidator(k Kind) error {
	switch k {
	case KindRoot, KindOrganization:
		return nil
	default:
		return fmt.Errorf("organization: invalid enum value for kind field: %q", k)
	}
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s types.SimpleStatus) error {
	switch s.String() {
	case "active", "inactive", "processing":
		return nil
	default:
		return fmt.Errorf("organization: invalid enum value for status field: %q", s)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Kind) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Kind) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Kind(str)
	if err := KindValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Kind", str)
	}
	return nil
}

var (
	// types.SimpleStatus must implement graphql.Marshaler.
	_ graphql.Marshaler = (*types.SimpleStatus)(nil)
	// types.SimpleStatus must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*types.SimpleStatus)(nil)
)
