// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/adminx/ent/organization"
	"github.com/woocoos/adminx/ent/organizationuser"
	"github.com/woocoos/adminx/ent/user"
)

// OrganizationUser is the model entity for the OrganizationUser schema.
type OrganizationUser struct {
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
	// 组织ID
	OrgID int `json:"org_id,omitempty"`
	// 用户ID
	UserID int `json:"user_id,omitempty"`
	// 在组织内的显示名称
	DisplayName string `json:"display_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrganizationUserQuery when eager-loading is set.
	Edges OrganizationUserEdges `json:"edges"`
}

// OrganizationUserEdges holds the relations/edges for other nodes in the graph.
type OrganizationUserEdges struct {
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrganizationUserEdges) OrganizationOrErr() (*Organization, error) {
	if e.loadedTypes[0] {
		if e.Organization == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organization.Label}
		}
		return e.Organization, nil
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrganizationUserEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrganizationUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case organizationuser.FieldID, organizationuser.FieldCreatedBy, organizationuser.FieldUpdatedBy, organizationuser.FieldOrgID, organizationuser.FieldUserID:
			values[i] = new(sql.NullInt64)
		case organizationuser.FieldDisplayName:
			values[i] = new(sql.NullString)
		case organizationuser.FieldCreatedAt, organizationuser.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrganizationUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrganizationUser fields.
func (ou *OrganizationUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case organizationuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ou.ID = int(value.Int64)
		case organizationuser.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ou.CreatedBy = int(value.Int64)
			}
		case organizationuser.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ou.CreatedAt = value.Time
			}
		case organizationuser.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ou.UpdatedBy = int(value.Int64)
			}
		case organizationuser.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ou.UpdatedAt = value.Time
			}
		case organizationuser.FieldOrgID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				ou.OrgID = int(value.Int64)
			}
		case organizationuser.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ou.UserID = int(value.Int64)
			}
		case organizationuser.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				ou.DisplayName = value.String
			}
		}
	}
	return nil
}

// QueryOrganization queries the "organization" edge of the OrganizationUser entity.
func (ou *OrganizationUser) QueryOrganization() *OrganizationQuery {
	return NewOrganizationUserClient(ou.config).QueryOrganization(ou)
}

// QueryUser queries the "user" edge of the OrganizationUser entity.
func (ou *OrganizationUser) QueryUser() *UserQuery {
	return NewOrganizationUserClient(ou.config).QueryUser(ou)
}

// Update returns a builder for updating this OrganizationUser.
// Note that you need to call OrganizationUser.Unwrap() before calling this method if this OrganizationUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (ou *OrganizationUser) Update() *OrganizationUserUpdateOne {
	return NewOrganizationUserClient(ou.config).UpdateOne(ou)
}

// Unwrap unwraps the OrganizationUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ou *OrganizationUser) Unwrap() *OrganizationUser {
	_tx, ok := ou.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrganizationUser is not a transactional entity")
	}
	ou.config.driver = _tx.drv
	return ou
}

// String implements the fmt.Stringer.
func (ou *OrganizationUser) String() string {
	var builder strings.Builder
	builder.WriteString("OrganizationUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ou.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", ou.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ou.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", ou.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ou.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("org_id=")
	builder.WriteString(fmt.Sprintf("%v", ou.OrgID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ou.UserID))
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(ou.DisplayName)
	builder.WriteByte(')')
	return builder.String()
}

// OrganizationUsers is a parsable slice of OrganizationUser.
type OrganizationUsers []*OrganizationUser
