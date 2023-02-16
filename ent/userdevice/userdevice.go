// Code generated by ent, DO NOT EDIT.

package userdevice

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the userdevice type in the database.
	Label = "user_device"
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
	// FieldDeviceUID holds the string denoting the device_uid field in the database.
	FieldDeviceUID = "device_uid"
	// FieldDeviceName holds the string denoting the device_name field in the database.
	FieldDeviceName = "device_name"
	// FieldSystemName holds the string denoting the system_name field in the database.
	FieldSystemName = "system_name"
	// FieldSystemVersion holds the string denoting the system_version field in the database.
	FieldSystemVersion = "system_version"
	// FieldAppVersion holds the string denoting the app_version field in the database.
	FieldAppVersion = "app_version"
	// FieldDeviceModel holds the string denoting the device_model field in the database.
	FieldDeviceModel = "device_model"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the userdevice in the database.
	Table = "user_device"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_device"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "user"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for userdevice fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldUserID,
	FieldDeviceUID,
	FieldDeviceName,
	FieldSystemName,
	FieldSystemVersion,
	FieldAppVersion,
	FieldDeviceModel,
	FieldStatus,
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
	Hooks [2]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DeviceUIDValidator is a validator for the "device_uid" field. It is called by the builders before save.
	DeviceUIDValidator func(string) error
	// DeviceNameValidator is a validator for the "device_name" field. It is called by the builders before save.
	DeviceNameValidator func(string) error
	// SystemNameValidator is a validator for the "system_name" field. It is called by the builders before save.
	SystemNameValidator func(string) error
	// SystemVersionValidator is a validator for the "system_version" field. It is called by the builders before save.
	SystemVersionValidator func(string) error
	// AppVersionValidator is a validator for the "app_version" field. It is called by the builders before save.
	AppVersionValidator func(string) error
	// DeviceModelValidator is a validator for the "device_model" field. It is called by the builders before save.
	DeviceModelValidator func(string) error
)

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
		return fmt.Errorf("userdevice: invalid enum value for status field: %q", s)
	}
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
