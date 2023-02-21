// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/adminx/ent/organizationpolicy"
	"github.com/woocoos/adminx/graph/entgen/types"
)

// OrganizationPolicy is the model entity for the OrganizationPolicy schema.
type OrganizationPolicy struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 组织ID
	OrgID int `json:"org_id,omitempty"`
	// 所属应用
	AppID int `json:"app_id,omitempty"`
	// 所属应用策略,如果是自定义应用策略,则为空
	AppPolicyID int `json:"app_policy_id,omitempty"`
	// 策略名称
	Name string `json:"name,omitempty"`
	// 描述
	Comments string `json:"comments,omitempty"`
	// 策略规则,如果是应用策略,则为空
	Rules []types.PolicyRule `json:"rules,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrganizationPolicy) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case organizationpolicy.FieldRules:
			values[i] = new([]byte)
		case organizationpolicy.FieldID, organizationpolicy.FieldOrgID, organizationpolicy.FieldAppID, organizationpolicy.FieldAppPolicyID:
			values[i] = new(sql.NullInt64)
		case organizationpolicy.FieldName, organizationpolicy.FieldComments:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrganizationPolicy", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrganizationPolicy fields.
func (op *OrganizationPolicy) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case organizationpolicy.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			op.ID = int(value.Int64)
		case organizationpolicy.FieldOrgID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				op.OrgID = int(value.Int64)
			}
		case organizationpolicy.FieldAppID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				op.AppID = int(value.Int64)
			}
		case organizationpolicy.FieldAppPolicyID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_policy_id", values[i])
			} else if value.Valid {
				op.AppPolicyID = int(value.Int64)
			}
		case organizationpolicy.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				op.Name = value.String
			}
		case organizationpolicy.FieldComments:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comments", values[i])
			} else if value.Valid {
				op.Comments = value.String
			}
		case organizationpolicy.FieldRules:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field rules", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &op.Rules); err != nil {
					return fmt.Errorf("unmarshal field rules: %w", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this OrganizationPolicy.
// Note that you need to call OrganizationPolicy.Unwrap() before calling this method if this OrganizationPolicy
// was returned from a transaction, and the transaction was committed or rolled back.
func (op *OrganizationPolicy) Update() *OrganizationPolicyUpdateOne {
	return NewOrganizationPolicyClient(op.config).UpdateOne(op)
}

// Unwrap unwraps the OrganizationPolicy entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (op *OrganizationPolicy) Unwrap() *OrganizationPolicy {
	_tx, ok := op.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrganizationPolicy is not a transactional entity")
	}
	op.config.driver = _tx.drv
	return op
}

// String implements the fmt.Stringer.
func (op *OrganizationPolicy) String() string {
	var builder strings.Builder
	builder.WriteString("OrganizationPolicy(")
	builder.WriteString(fmt.Sprintf("id=%v, ", op.ID))
	builder.WriteString("org_id=")
	builder.WriteString(fmt.Sprintf("%v", op.OrgID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", op.AppID))
	builder.WriteString(", ")
	builder.WriteString("app_policy_id=")
	builder.WriteString(fmt.Sprintf("%v", op.AppPolicyID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(op.Name)
	builder.WriteString(", ")
	builder.WriteString("comments=")
	builder.WriteString(op.Comments)
	builder.WriteString(", ")
	builder.WriteString("rules=")
	builder.WriteString(fmt.Sprintf("%v", op.Rules))
	builder.WriteByte(')')
	return builder.String()
}

// OrganizationPolicies is a parsable slice of OrganizationPolicy.
type OrganizationPolicies []*OrganizationPolicy