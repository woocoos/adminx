// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/adminx/ent/app"
	"github.com/woocoos/adminx/ent/appres"
)

// AppRes is the model entity for the AppRes schema.
type AppRes struct {
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
	// 所属应用
	AppID int `json:"app_id,omitempty"`
	// 资源名称
	Name string `json:"name,omitempty"`
	// 资源类型名称,如数据库表名
	TypeName string `json:"type_name,omitempty"`
	// 应用资源表达式
	ArnPattern string `json:"arn_pattern,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AppResQuery when eager-loading is set.
	Edges                AppResEdges `json:"edges"`
	app_action_resources *int
}

// AppResEdges holds the relations/edges for other nodes in the graph.
type AppResEdges struct {
	// App holds the value of the app edge.
	App *App `json:"app,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// AppOrErr returns the App value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppResEdges) AppOrErr() (*App, error) {
	if e.loadedTypes[0] {
		if e.App == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: app.Label}
		}
		return e.App, nil
	}
	return nil, &NotLoadedError{edge: "app"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppRes) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case appres.FieldID, appres.FieldCreatedBy, appres.FieldUpdatedBy, appres.FieldAppID:
			values[i] = new(sql.NullInt64)
		case appres.FieldName, appres.FieldTypeName, appres.FieldArnPattern:
			values[i] = new(sql.NullString)
		case appres.FieldCreatedAt, appres.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case appres.ForeignKeys[0]: // app_action_resources
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppRes", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppRes fields.
func (ar *AppRes) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appres.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ar.ID = int(value.Int64)
		case appres.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ar.CreatedBy = int(value.Int64)
			}
		case appres.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ar.CreatedAt = value.Time
			}
		case appres.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ar.UpdatedBy = int(value.Int64)
			}
		case appres.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ar.UpdatedAt = value.Time
			}
		case appres.FieldAppID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				ar.AppID = int(value.Int64)
			}
		case appres.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ar.Name = value.String
			}
		case appres.FieldTypeName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type_name", values[i])
			} else if value.Valid {
				ar.TypeName = value.String
			}
		case appres.FieldArnPattern:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field arn_pattern", values[i])
			} else if value.Valid {
				ar.ArnPattern = value.String
			}
		case appres.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field app_action_resources", value)
			} else if value.Valid {
				ar.app_action_resources = new(int)
				*ar.app_action_resources = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryApp queries the "app" edge of the AppRes entity.
func (ar *AppRes) QueryApp() *AppQuery {
	return NewAppResClient(ar.config).QueryApp(ar)
}

// Update returns a builder for updating this AppRes.
// Note that you need to call AppRes.Unwrap() before calling this method if this AppRes
// was returned from a transaction, and the transaction was committed or rolled back.
func (ar *AppRes) Update() *AppResUpdateOne {
	return NewAppResClient(ar.config).UpdateOne(ar)
}

// Unwrap unwraps the AppRes entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ar *AppRes) Unwrap() *AppRes {
	_tx, ok := ar.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppRes is not a transactional entity")
	}
	ar.config.driver = _tx.drv
	return ar
}

// String implements the fmt.Stringer.
func (ar *AppRes) String() string {
	var builder strings.Builder
	builder.WriteString("AppRes(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ar.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", ar.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ar.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", ar.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ar.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", ar.AppID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ar.Name)
	builder.WriteString(", ")
	builder.WriteString("type_name=")
	builder.WriteString(ar.TypeName)
	builder.WriteString(", ")
	builder.WriteString("arn_pattern=")
	builder.WriteString(ar.ArnPattern)
	builder.WriteByte(')')
	return builder.String()
}

// AppResSlice is a parsable slice of AppRes.
type AppResSlice []*AppRes