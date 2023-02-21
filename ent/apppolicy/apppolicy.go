// Code generated by ent, DO NOT EDIT.

package apppolicy

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"github.com/99designs/gqlgen/graphql"
	"github.com/woocoos/adminx/graph/entgen/types"
)

const (
	// Label holds the string label denoting the apppolicy type in the database.
	Label = "app_policy"
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
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// FieldRules holds the string denoting the rules field in the database.
	FieldRules = "rules"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldAutoGrant holds the string denoting the auto_grant field in the database.
	FieldAutoGrant = "auto_grant"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeApp holds the string denoting the app edge name in mutations.
	EdgeApp = "app"
	// EdgeRoles holds the string denoting the roles edge name in mutations.
	EdgeRoles = "roles"
	// EdgeAppRolePolicy holds the string denoting the app_role_policy edge name in mutations.
	EdgeAppRolePolicy = "app_role_policy"
	// Table holds the table name of the apppolicy in the database.
	Table = "app_policy"
	// AppTable is the table that holds the app relation/edge.
	AppTable = "app_policy"
	// AppInverseTable is the table name for the App entity.
	// It exists in this package in order to avoid circular dependency with the "app" package.
	AppInverseTable = "app"
	// AppColumn is the table column denoting the app relation/edge.
	AppColumn = "app_id"
	// RolesTable is the table that holds the roles relation/edge. The primary key declared below.
	RolesTable = "app_role_policy"
	// RolesInverseTable is the table name for the AppRole entity.
	// It exists in this package in order to avoid circular dependency with the "approle" package.
	RolesInverseTable = "app_role"
	// AppRolePolicyTable is the table that holds the app_role_policy relation/edge.
	AppRolePolicyTable = "app_role_policy"
	// AppRolePolicyInverseTable is the table name for the AppRolePolicy entity.
	// It exists in this package in order to avoid circular dependency with the "approlepolicy" package.
	AppRolePolicyInverseTable = "app_role_policy"
	// AppRolePolicyColumn is the table column denoting the app_role_policy relation/edge.
	AppRolePolicyColumn = "policy_id"
)

// Columns holds all SQL columns for apppolicy fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldAppID,
	FieldName,
	FieldComments,
	FieldRules,
	FieldVersion,
	FieldAutoGrant,
	FieldStatus,
}

var (
	// RolesPrimaryKey and RolesColumn2 are the table columns denoting the
	// primary key for the roles relation (M2M).
	RolesPrimaryKey = []string{"role_id", "policy_id"}
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
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultAutoGrant holds the default value on creation for the "auto_grant" field.
	DefaultAutoGrant bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int
)

const DefaultStatus types.SimpleStatus = "active"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s types.SimpleStatus) error {
	switch s.String() {
	case "active", "inactive", "processing":
		return nil
	default:
		return fmt.Errorf("apppolicy: invalid enum value for status field: %q", s)
	}
}

var (
	// types.SimpleStatus must implement graphql.Marshaler.
	_ graphql.Marshaler = (*types.SimpleStatus)(nil)
	// types.SimpleStatus must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*types.SimpleStatus)(nil)
)
