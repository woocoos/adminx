// Code generated by ent, DO NOT EDIT.

package organizationapp

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the organizationapp type in the database.
	Label = "organization_app"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldOrgID holds the string denoting the org_id field in the database.
	FieldOrgID = "org_id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// EdgeApp holds the string denoting the app edge name in mutations.
	EdgeApp = "app"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// AppFieldID holds the string denoting the ID field of the App.
	AppFieldID = "id"
	// OrganizationFieldID holds the string denoting the ID field of the Organization.
	OrganizationFieldID = "id"
	// Table holds the table name of the organizationapp in the database.
	Table = "organization_app"
	// AppTable is the table that holds the app relation/edge.
	AppTable = "organization_app"
	// AppInverseTable is the table name for the App entity.
	// It exists in this package in order to avoid circular dependency with the "app" package.
	AppInverseTable = "app"
	// AppColumn is the table column denoting the app relation/edge.
	AppColumn = "app_id"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "organization_app"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organization"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "org_id"
)

// Columns holds all SQL columns for organizationapp fields.
var Columns = []string{
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldOrgID,
	FieldAppID,
}

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
)