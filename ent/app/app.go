// Code generated by ent, DO NOT EDIT.

package app

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the app type in the database.
	Label = "app"
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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldKind holds the string denoting the kind field in the database.
	FieldKind = "kind"
	// FieldRedirectURI holds the string denoting the redirect_uri field in the database.
	FieldRedirectURI = "redirect_uri"
	// FieldAppKey holds the string denoting the app_key field in the database.
	FieldAppKey = "app_key"
	// FieldAppSecret holds the string denoting the app_secret field in the database.
	FieldAppSecret = "app_secret"
	// FieldScopes holds the string denoting the scopes field in the database.
	FieldScopes = "scopes"
	// FieldTokenValidity holds the string denoting the token_validity field in the database.
	FieldTokenValidity = "token_validity"
	// FieldRefreshTokenValidity holds the string denoting the refresh_token_validity field in the database.
	FieldRefreshTokenValidity = "refresh_token_validity"
	// FieldLogo holds the string denoting the logo field in the database.
	FieldLogo = "logo"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeMenus holds the string denoting the menus edge name in mutations.
	EdgeMenus = "menus"
	// EdgePermissions holds the string denoting the permissions edge name in mutations.
	EdgePermissions = "permissions"
	// Table holds the table name of the app in the database.
	Table = "app"
	// MenusTable is the table that holds the menus relation/edge.
	MenusTable = "app_menu"
	// MenusInverseTable is the table name for the AppMenu entity.
	// It exists in this package in order to avoid circular dependency with the "appmenu" package.
	MenusInverseTable = "app_menu"
	// MenusColumn is the table column denoting the menus relation/edge.
	MenusColumn = "app_id"
	// PermissionsTable is the table that holds the permissions relation/edge.
	PermissionsTable = "app_permission"
	// PermissionsInverseTable is the table name for the AppPermission entity.
	// It exists in this package in order to avoid circular dependency with the "apppermission" package.
	PermissionsInverseTable = "app_permission"
	// PermissionsColumn is the table column denoting the permissions relation/edge.
	PermissionsColumn = "app_id"
)

// Columns holds all SQL columns for app fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldName,
	FieldKind,
	FieldRedirectURI,
	FieldAppKey,
	FieldAppSecret,
	FieldScopes,
	FieldTokenValidity,
	FieldRefreshTokenValidity,
	FieldLogo,
	FieldComments,
	FieldStatus,
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
	Hooks [2]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// RedirectURIValidator is a validator for the "redirect_uri" field. It is called by the builders before save.
	RedirectURIValidator func(string) error
	// AppSecretValidator is a validator for the "app_secret" field. It is called by the builders before save.
	AppSecretValidator func(string) error
	// ScopesValidator is a validator for the "scopes" field. It is called by the builders before save.
	ScopesValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int
)

// Kind defines the type for the "kind" enum field.
type Kind string

// Kind values.
const (
	KindWeb    Kind = "web"
	KindNative Kind = "native"
	KindServer Kind = "server"
)

func (k Kind) String() string {
	return string(k)
}

// KindValidator is a validator for the "kind" field enum values. It is called by the builders before save.
func KindValidator(k Kind) error {
	switch k {
	case KindWeb, KindNative, KindServer:
		return nil
	default:
		return fmt.Errorf("app: invalid enum value for kind field: %q", k)
	}
}

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusActive, StatusInactive:
		return nil
	default:
		return fmt.Errorf("app: invalid enum value for status field: %q", s)
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

// MarshalGQL implements graphql.Marshaler interface.
func (e Status) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Status) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Status(str)
	if err := StatusValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}
