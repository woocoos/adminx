// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/adminx/ent/user"
	"github.com/woocoos/adminx/ent/userloginprofile"
)

// User is the model entity for the User schema.
type User struct {
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
	// 登陆名称
	PrincipalName string `json:"principal_name,omitempty"`
	// 显示名
	DisplayName string `json:"display_name,omitempty"`
	// 用户类型
	UserType user.UserType `json:"user_type,omitempty"`
	// 创建类型,邀请，注册,手工创建
	CreationType user.CreationType `json:"creation_type,omitempty"`
	// 注册时IP
	RegisterIP *string `json:"register_ip,omitempty"`
	// 状态
	Status user.Status `json:"status,omitempty"`
	// 备注
	Comments string `json:"comments,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// 用户身份标识
	Identities []*UserIdentity `json:"identities,omitempty"`
	// 登陆设置
	LoginProfile *UserLoginProfile `json:"login_profile,omitempty"`
	// 用户密码
	Passwords []*UserPassword `json:"passwords,omitempty"`
	// 用户设备
	Devices []*UserDevice `json:"devices,omitempty"`
	// 用户所属组织
	Organizations []*Organization `json:"organizations,omitempty"`
	// OrganizationUser holds the value of the organization_user edge.
	OrganizationUser []*OrganizationUser `json:"organization_user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
	// totalCount holds the count of the edges above.
	totalCount [4]map[string]int

	namedIdentities       map[string][]*UserIdentity
	namedPasswords        map[string][]*UserPassword
	namedDevices          map[string][]*UserDevice
	namedOrganizations    map[string][]*Organization
	namedOrganizationUser map[string][]*OrganizationUser
}

// IdentitiesOrErr returns the Identities value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) IdentitiesOrErr() ([]*UserIdentity, error) {
	if e.loadedTypes[0] {
		return e.Identities, nil
	}
	return nil, &NotLoadedError{edge: "identities"}
}

// LoginProfileOrErr returns the LoginProfile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) LoginProfileOrErr() (*UserLoginProfile, error) {
	if e.loadedTypes[1] {
		if e.LoginProfile == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: userloginprofile.Label}
		}
		return e.LoginProfile, nil
	}
	return nil, &NotLoadedError{edge: "login_profile"}
}

// PasswordsOrErr returns the Passwords value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PasswordsOrErr() ([]*UserPassword, error) {
	if e.loadedTypes[2] {
		return e.Passwords, nil
	}
	return nil, &NotLoadedError{edge: "passwords"}
}

// DevicesOrErr returns the Devices value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) DevicesOrErr() ([]*UserDevice, error) {
	if e.loadedTypes[3] {
		return e.Devices, nil
	}
	return nil, &NotLoadedError{edge: "devices"}
}

// OrganizationsOrErr returns the Organizations value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OrganizationsOrErr() ([]*Organization, error) {
	if e.loadedTypes[4] {
		return e.Organizations, nil
	}
	return nil, &NotLoadedError{edge: "organizations"}
}

// OrganizationUserOrErr returns the OrganizationUser value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OrganizationUserOrErr() ([]*OrganizationUser, error) {
	if e.loadedTypes[5] {
		return e.OrganizationUser, nil
	}
	return nil, &NotLoadedError{edge: "organization_user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldCreatedBy, user.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case user.FieldPrincipalName, user.FieldDisplayName, user.FieldUserType, user.FieldCreationType, user.FieldRegisterIP, user.FieldStatus, user.FieldComments:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				u.CreatedBy = int(value.Int64)
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				u.UpdatedBy = int(value.Int64)
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				u.DeletedAt = value.Time
			}
		case user.FieldPrincipalName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field principal_name", values[i])
			} else if value.Valid {
				u.PrincipalName = value.String
			}
		case user.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				u.DisplayName = value.String
			}
		case user.FieldUserType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_type", values[i])
			} else if value.Valid {
				u.UserType = user.UserType(value.String)
			}
		case user.FieldCreationType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creation_type", values[i])
			} else if value.Valid {
				u.CreationType = user.CreationType(value.String)
			}
		case user.FieldRegisterIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field register_ip", values[i])
			} else if value.Valid {
				u.RegisterIP = new(string)
				*u.RegisterIP = value.String
			}
		case user.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = user.Status(value.String)
			}
		case user.FieldComments:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comments", values[i])
			} else if value.Valid {
				u.Comments = value.String
			}
		}
	}
	return nil
}

// QueryIdentities queries the "identities" edge of the User entity.
func (u *User) QueryIdentities() *UserIdentityQuery {
	return NewUserClient(u.config).QueryIdentities(u)
}

// QueryLoginProfile queries the "login_profile" edge of the User entity.
func (u *User) QueryLoginProfile() *UserLoginProfileQuery {
	return NewUserClient(u.config).QueryLoginProfile(u)
}

// QueryPasswords queries the "passwords" edge of the User entity.
func (u *User) QueryPasswords() *UserPasswordQuery {
	return NewUserClient(u.config).QueryPasswords(u)
}

// QueryDevices queries the "devices" edge of the User entity.
func (u *User) QueryDevices() *UserDeviceQuery {
	return NewUserClient(u.config).QueryDevices(u)
}

// QueryOrganizations queries the "organizations" edge of the User entity.
func (u *User) QueryOrganizations() *OrganizationQuery {
	return NewUserClient(u.config).QueryOrganizations(u)
}

// QueryOrganizationUser queries the "organization_user" edge of the User entity.
func (u *User) QueryOrganizationUser() *OrganizationUserQuery {
	return NewUserClient(u.config).QueryOrganizationUser(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", u.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", u.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(u.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("principal_name=")
	builder.WriteString(u.PrincipalName)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(u.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("user_type=")
	builder.WriteString(fmt.Sprintf("%v", u.UserType))
	builder.WriteString(", ")
	builder.WriteString("creation_type=")
	builder.WriteString(fmt.Sprintf("%v", u.CreationType))
	builder.WriteString(", ")
	if v := u.RegisterIP; v != nil {
		builder.WriteString("register_ip=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteString(", ")
	builder.WriteString("comments=")
	builder.WriteString(u.Comments)
	builder.WriteByte(')')
	return builder.String()
}

// NamedIdentities returns the Identities named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedIdentities(name string) ([]*UserIdentity, error) {
	if u.Edges.namedIdentities == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedIdentities[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedIdentities(name string, edges ...*UserIdentity) {
	if u.Edges.namedIdentities == nil {
		u.Edges.namedIdentities = make(map[string][]*UserIdentity)
	}
	if len(edges) == 0 {
		u.Edges.namedIdentities[name] = []*UserIdentity{}
	} else {
		u.Edges.namedIdentities[name] = append(u.Edges.namedIdentities[name], edges...)
	}
}

// NamedPasswords returns the Passwords named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedPasswords(name string) ([]*UserPassword, error) {
	if u.Edges.namedPasswords == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedPasswords[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedPasswords(name string, edges ...*UserPassword) {
	if u.Edges.namedPasswords == nil {
		u.Edges.namedPasswords = make(map[string][]*UserPassword)
	}
	if len(edges) == 0 {
		u.Edges.namedPasswords[name] = []*UserPassword{}
	} else {
		u.Edges.namedPasswords[name] = append(u.Edges.namedPasswords[name], edges...)
	}
}

// NamedDevices returns the Devices named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedDevices(name string) ([]*UserDevice, error) {
	if u.Edges.namedDevices == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedDevices[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedDevices(name string, edges ...*UserDevice) {
	if u.Edges.namedDevices == nil {
		u.Edges.namedDevices = make(map[string][]*UserDevice)
	}
	if len(edges) == 0 {
		u.Edges.namedDevices[name] = []*UserDevice{}
	} else {
		u.Edges.namedDevices[name] = append(u.Edges.namedDevices[name], edges...)
	}
}

// NamedOrganizations returns the Organizations named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedOrganizations(name string) ([]*Organization, error) {
	if u.Edges.namedOrganizations == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedOrganizations[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedOrganizations(name string, edges ...*Organization) {
	if u.Edges.namedOrganizations == nil {
		u.Edges.namedOrganizations = make(map[string][]*Organization)
	}
	if len(edges) == 0 {
		u.Edges.namedOrganizations[name] = []*Organization{}
	} else {
		u.Edges.namedOrganizations[name] = append(u.Edges.namedOrganizations[name], edges...)
	}
}

// NamedOrganizationUser returns the OrganizationUser named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedOrganizationUser(name string) ([]*OrganizationUser, error) {
	if u.Edges.namedOrganizationUser == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedOrganizationUser[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedOrganizationUser(name string, edges ...*OrganizationUser) {
	if u.Edges.namedOrganizationUser == nil {
		u.Edges.namedOrganizationUser = make(map[string][]*OrganizationUser)
	}
	if len(edges) == 0 {
		u.Edges.namedOrganizationUser[name] = []*OrganizationUser{}
	} else {
		u.Edges.namedOrganizationUser[name] = append(u.Edges.namedOrganizationUser[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User