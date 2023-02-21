// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/adminx/ent/user"
	"github.com/woocoos/adminx/ent/useridentity"
	"github.com/woocoos/adminx/graph/entgen/types"
)

// UserIdentity is the model entity for the UserIdentity schema.
type UserIdentity struct {
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
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// 身份标识类型 手机、邮箱、用户名、微信、qq
	Kind useridentity.Kind `json:"kind,omitempty"`
	// 用户名、邮箱、手机、unionid、qq
	Code string `json:"code,omitempty"`
	// 扩展标识码,比如微信的openID
	CodeExtend string `json:"code_extend,omitempty"`
	// 状态,部分登陆方式需要验证通过才可启用
	Status types.SimpleStatus `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserIdentityQuery when eager-loading is set.
	Edges UserIdentityEdges `json:"edges"`
}

// UserIdentityEdges holds the relations/edges for other nodes in the graph.
type UserIdentityEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserIdentityEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserIdentity) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case useridentity.FieldID, useridentity.FieldCreatedBy, useridentity.FieldUpdatedBy, useridentity.FieldUserID:
			values[i] = new(sql.NullInt64)
		case useridentity.FieldKind, useridentity.FieldCode, useridentity.FieldCodeExtend, useridentity.FieldStatus:
			values[i] = new(sql.NullString)
		case useridentity.FieldCreatedAt, useridentity.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserIdentity", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserIdentity fields.
func (ui *UserIdentity) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case useridentity.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ui.ID = int(value.Int64)
		case useridentity.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ui.CreatedBy = int(value.Int64)
			}
		case useridentity.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ui.CreatedAt = value.Time
			}
		case useridentity.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ui.UpdatedBy = int(value.Int64)
			}
		case useridentity.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ui.UpdatedAt = value.Time
			}
		case useridentity.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ui.UserID = int(value.Int64)
			}
		case useridentity.FieldKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[i])
			} else if value.Valid {
				ui.Kind = useridentity.Kind(value.String)
			}
		case useridentity.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				ui.Code = value.String
			}
		case useridentity.FieldCodeExtend:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code_extend", values[i])
			} else if value.Valid {
				ui.CodeExtend = value.String
			}
		case useridentity.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ui.Status = types.SimpleStatus(value.String)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the UserIdentity entity.
func (ui *UserIdentity) QueryUser() *UserQuery {
	return NewUserIdentityClient(ui.config).QueryUser(ui)
}

// Update returns a builder for updating this UserIdentity.
// Note that you need to call UserIdentity.Unwrap() before calling this method if this UserIdentity
// was returned from a transaction, and the transaction was committed or rolled back.
func (ui *UserIdentity) Update() *UserIdentityUpdateOne {
	return NewUserIdentityClient(ui.config).UpdateOne(ui)
}

// Unwrap unwraps the UserIdentity entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ui *UserIdentity) Unwrap() *UserIdentity {
	_tx, ok := ui.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserIdentity is not a transactional entity")
	}
	ui.config.driver = _tx.drv
	return ui
}

// String implements the fmt.Stringer.
func (ui *UserIdentity) String() string {
	var builder strings.Builder
	builder.WriteString("UserIdentity(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ui.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", ui.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ui.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", ui.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ui.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ui.UserID))
	builder.WriteString(", ")
	builder.WriteString("kind=")
	builder.WriteString(fmt.Sprintf("%v", ui.Kind))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(ui.Code)
	builder.WriteString(", ")
	builder.WriteString("code_extend=")
	builder.WriteString(ui.CodeExtend)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", ui.Status))
	builder.WriteByte(')')
	return builder.String()
}

// UserIdentities is a parsable slice of UserIdentity.
type UserIdentities []*UserIdentity
