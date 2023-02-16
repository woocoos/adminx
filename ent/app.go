// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/adminx/ent/app"
)

// App is the model entity for the App schema.
type App struct {
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
	// 应用名称
	Name string `json:"name,omitempty"`
	// 应用类型
	Kind app.Kind `json:"kind,omitempty"`
	// 回调地址
	RedirectURI string `json:"redirect_uri,omitempty"`
	// 应用ID
	AppKey string `json:"app_key,omitempty"`
	// 应用密钥
	AppSecret string `json:"app_secret,omitempty"`
	// 权限范围
	Scopes string `json:"scopes,omitempty"`
	// token有效期
	TokenValidity int32 `json:"token_validity,omitempty"`
	// refresh_token有效期
	RefreshTokenValidity int32 `json:"refresh_token_validity,omitempty"`
	// 图标
	Logo string `json:"logo,omitempty"`
	// 备注
	Comments string `json:"comments,omitempty"`
	// 状态
	Status app.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AppQuery when eager-loading is set.
	Edges AppEdges `json:"edges"`
}

// AppEdges holds the relations/edges for other nodes in the graph.
type AppEdges struct {
	// 菜单
	Menus []*AppMenu `json:"menus,omitempty"`
	// 权限
	Permissions []*AppPermission `json:"permissions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedMenus       map[string][]*AppMenu
	namedPermissions map[string][]*AppPermission
}

// MenusOrErr returns the Menus value or an error if the edge
// was not loaded in eager-loading.
func (e AppEdges) MenusOrErr() ([]*AppMenu, error) {
	if e.loadedTypes[0] {
		return e.Menus, nil
	}
	return nil, &NotLoadedError{edge: "menus"}
}

// PermissionsOrErr returns the Permissions value or an error if the edge
// was not loaded in eager-loading.
func (e AppEdges) PermissionsOrErr() ([]*AppPermission, error) {
	if e.loadedTypes[1] {
		return e.Permissions, nil
	}
	return nil, &NotLoadedError{edge: "permissions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*App) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case app.FieldID, app.FieldCreatedBy, app.FieldUpdatedBy, app.FieldTokenValidity, app.FieldRefreshTokenValidity:
			values[i] = new(sql.NullInt64)
		case app.FieldName, app.FieldKind, app.FieldRedirectURI, app.FieldAppKey, app.FieldAppSecret, app.FieldScopes, app.FieldLogo, app.FieldComments, app.FieldStatus:
			values[i] = new(sql.NullString)
		case app.FieldCreatedAt, app.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type App", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the App fields.
func (a *App) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case app.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case app.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				a.CreatedBy = int(value.Int64)
			}
		case app.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case app.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				a.UpdatedBy = int(value.Int64)
			}
		case app.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case app.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case app.FieldKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[i])
			} else if value.Valid {
				a.Kind = app.Kind(value.String)
			}
		case app.FieldRedirectURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field redirect_uri", values[i])
			} else if value.Valid {
				a.RedirectURI = value.String
			}
		case app.FieldAppKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_key", values[i])
			} else if value.Valid {
				a.AppKey = value.String
			}
		case app.FieldAppSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_secret", values[i])
			} else if value.Valid {
				a.AppSecret = value.String
			}
		case app.FieldScopes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scopes", values[i])
			} else if value.Valid {
				a.Scopes = value.String
			}
		case app.FieldTokenValidity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field token_validity", values[i])
			} else if value.Valid {
				a.TokenValidity = int32(value.Int64)
			}
		case app.FieldRefreshTokenValidity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field refresh_token_validity", values[i])
			} else if value.Valid {
				a.RefreshTokenValidity = int32(value.Int64)
			}
		case app.FieldLogo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo", values[i])
			} else if value.Valid {
				a.Logo = value.String
			}
		case app.FieldComments:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comments", values[i])
			} else if value.Valid {
				a.Comments = value.String
			}
		case app.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				a.Status = app.Status(value.String)
			}
		}
	}
	return nil
}

// QueryMenus queries the "menus" edge of the App entity.
func (a *App) QueryMenus() *AppMenuQuery {
	return NewAppClient(a.config).QueryMenus(a)
}

// QueryPermissions queries the "permissions" edge of the App entity.
func (a *App) QueryPermissions() *AppPermissionQuery {
	return NewAppClient(a.config).QueryPermissions(a)
}

// Update returns a builder for updating this App.
// Note that you need to call App.Unwrap() before calling this method if this App
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *App) Update() *AppUpdateOne {
	return NewAppClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the App entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *App) Unwrap() *App {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: App is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *App) String() string {
	var builder strings.Builder
	builder.WriteString("App(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", a.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", a.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("kind=")
	builder.WriteString(fmt.Sprintf("%v", a.Kind))
	builder.WriteString(", ")
	builder.WriteString("redirect_uri=")
	builder.WriteString(a.RedirectURI)
	builder.WriteString(", ")
	builder.WriteString("app_key=")
	builder.WriteString(a.AppKey)
	builder.WriteString(", ")
	builder.WriteString("app_secret=")
	builder.WriteString(a.AppSecret)
	builder.WriteString(", ")
	builder.WriteString("scopes=")
	builder.WriteString(a.Scopes)
	builder.WriteString(", ")
	builder.WriteString("token_validity=")
	builder.WriteString(fmt.Sprintf("%v", a.TokenValidity))
	builder.WriteString(", ")
	builder.WriteString("refresh_token_validity=")
	builder.WriteString(fmt.Sprintf("%v", a.RefreshTokenValidity))
	builder.WriteString(", ")
	builder.WriteString("logo=")
	builder.WriteString(a.Logo)
	builder.WriteString(", ")
	builder.WriteString("comments=")
	builder.WriteString(a.Comments)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", a.Status))
	builder.WriteByte(')')
	return builder.String()
}

// NamedMenus returns the Menus named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *App) NamedMenus(name string) ([]*AppMenu, error) {
	if a.Edges.namedMenus == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedMenus[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *App) appendNamedMenus(name string, edges ...*AppMenu) {
	if a.Edges.namedMenus == nil {
		a.Edges.namedMenus = make(map[string][]*AppMenu)
	}
	if len(edges) == 0 {
		a.Edges.namedMenus[name] = []*AppMenu{}
	} else {
		a.Edges.namedMenus[name] = append(a.Edges.namedMenus[name], edges...)
	}
}

// NamedPermissions returns the Permissions named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *App) NamedPermissions(name string) ([]*AppPermission, error) {
	if a.Edges.namedPermissions == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedPermissions[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *App) appendNamedPermissions(name string, edges ...*AppPermission) {
	if a.Edges.namedPermissions == nil {
		a.Edges.namedPermissions = make(map[string][]*AppPermission)
	}
	if len(edges) == 0 {
		a.Edges.namedPermissions[name] = []*AppPermission{}
	} else {
		a.Edges.namedPermissions[name] = append(a.Edges.namedPermissions[name], edges...)
	}
}

// Apps is a parsable slice of App.
type Apps []*App