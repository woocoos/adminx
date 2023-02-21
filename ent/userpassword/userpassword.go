// Code generated by ent, DO NOT EDIT.

package userpassword

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
	// Label holds the string label denoting the userpassword type in the database.
	Label = "user_password"
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
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldScene holds the string denoting the scene field in the database.
	FieldScene = "scene"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldSalt holds the string denoting the salt field in the database.
	FieldSalt = "salt"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldMemo holds the string denoting the memo field in the database.
	FieldMemo = "memo"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the userpassword in the database.
	Table = "user_password"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_password"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "user"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for userpassword fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldUserID,
	FieldScene,
	FieldPassword,
	FieldSalt,
	FieldStatus,
	FieldMemo,
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
	// SaltValidator is a validator for the "salt" field. It is called by the builders before save.
	SaltValidator func(string) error
)

// Scene defines the type for the "scene" enum field.
type Scene string

// Scene values.
const (
	SceneLogin Scene = "login"
)

func (s Scene) String() string {
	return string(s)
}

// SceneValidator is a validator for the "scene" field enum values. It is called by the builders before save.
func SceneValidator(s Scene) error {
	switch s {
	case SceneLogin:
		return nil
	default:
		return fmt.Errorf("userpassword: invalid enum value for scene field: %q", s)
	}
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s types.SimpleStatus) error {
	switch s.String() {
	case "active", "inactive", "processing":
		return nil
	default:
		return fmt.Errorf("userpassword: invalid enum value for status field: %q", s)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Scene) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Scene) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Scene(str)
	if err := SceneValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Scene", str)
	}
	return nil
}

var (
	// types.SimpleStatus must implement graphql.Marshaler.
	_ graphql.Marshaler = (*types.SimpleStatus)(nil)
	// types.SimpleStatus must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*types.SimpleStatus)(nil)
)
