// Code generated by ent, DO NOT EDIT.

package apppermission

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the apppermission type in the database.
	Label = "app_permission"
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
	// EdgeApp holds the string denoting the app edge name in mutations.
	EdgeApp = "app"
	// EdgeMenus holds the string denoting the menus edge name in mutations.
	EdgeMenus = "menus"
	// Table holds the table name of the apppermission in the database.
	Table = "app_permission"
	// AppTable is the table that holds the app relation/edge.
	AppTable = "app_permission"
	// AppInverseTable is the table name for the App entity.
	// It exists in this package in order to avoid circular dependency with the "app" package.
	AppInverseTable = "app"
	// AppColumn is the table column denoting the app relation/edge.
	AppColumn = "app_id"
	// MenusTable is the table that holds the menus relation/edge.
	MenusTable = "app_menu"
	// MenusInverseTable is the table name for the AppMenu entity.
	// It exists in this package in order to avoid circular dependency with the "appmenu" package.
	MenusInverseTable = "app_menu"
	// MenusColumn is the table column denoting the menus relation/edge.
	MenusColumn = "permission_id"
)

// Columns holds all SQL columns for apppermission fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldAppID,
	FieldName,
	FieldComments,
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
	Hooks [3]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int
)