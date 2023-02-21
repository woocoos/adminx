// Code generated by ent, DO NOT EDIT.

package permission

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/woocoos/adminx/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldID, id))
}

// OrgID applies equality check predicate on the "org_id" field. It's identical to OrgIDEQ.
func OrgID(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldOrgID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldUserID, v))
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldRoleID, v))
}

// OrgPolicyID applies equality check predicate on the "org_policy_id" field. It's identical to OrgPolicyIDEQ.
func OrgPolicyID(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldOrgPolicyID, v))
}

// StartAt applies equality check predicate on the "start_at" field. It's identical to StartAtEQ.
func StartAt(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldStartAt, v))
}

// EndAt applies equality check predicate on the "end_at" field. It's identical to EndAtEQ.
func EndAt(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldEndAt, v))
}

// OrgIDEQ applies the EQ predicate on the "org_id" field.
func OrgIDEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldOrgID, v))
}

// OrgIDNEQ applies the NEQ predicate on the "org_id" field.
func OrgIDNEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldOrgID, v))
}

// OrgIDIn applies the In predicate on the "org_id" field.
func OrgIDIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldOrgID, vs...))
}

// OrgIDNotIn applies the NotIn predicate on the "org_id" field.
func OrgIDNotIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldOrgID, vs...))
}

// PrincipalKindEQ applies the EQ predicate on the "principal_kind" field.
func PrincipalKindEQ(v PrincipalKind) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldPrincipalKind, v))
}

// PrincipalKindNEQ applies the NEQ predicate on the "principal_kind" field.
func PrincipalKindNEQ(v PrincipalKind) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldPrincipalKind, v))
}

// PrincipalKindIn applies the In predicate on the "principal_kind" field.
func PrincipalKindIn(vs ...PrincipalKind) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldPrincipalKind, vs...))
}

// PrincipalKindNotIn applies the NotIn predicate on the "principal_kind" field.
func PrincipalKindNotIn(vs ...PrincipalKind) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldPrincipalKind, vs...))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Permission {
	return predicate.Permission(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Permission {
	return predicate.Permission(sql.FieldNotNull(FieldUserID))
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldRoleID, v))
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldRoleID, v))
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldRoleID, vs...))
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldRoleID, vs...))
}

// RoleIDGT applies the GT predicate on the "role_id" field.
func RoleIDGT(v int) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldRoleID, v))
}

// RoleIDGTE applies the GTE predicate on the "role_id" field.
func RoleIDGTE(v int) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldRoleID, v))
}

// RoleIDLT applies the LT predicate on the "role_id" field.
func RoleIDLT(v int) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldRoleID, v))
}

// RoleIDLTE applies the LTE predicate on the "role_id" field.
func RoleIDLTE(v int) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldRoleID, v))
}

// RoleIDIsNil applies the IsNil predicate on the "role_id" field.
func RoleIDIsNil() predicate.Permission {
	return predicate.Permission(sql.FieldIsNull(FieldRoleID))
}

// RoleIDNotNil applies the NotNil predicate on the "role_id" field.
func RoleIDNotNil() predicate.Permission {
	return predicate.Permission(sql.FieldNotNull(FieldRoleID))
}

// OrgPolicyIDEQ applies the EQ predicate on the "org_policy_id" field.
func OrgPolicyIDEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldOrgPolicyID, v))
}

// OrgPolicyIDNEQ applies the NEQ predicate on the "org_policy_id" field.
func OrgPolicyIDNEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldOrgPolicyID, v))
}

// OrgPolicyIDIn applies the In predicate on the "org_policy_id" field.
func OrgPolicyIDIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldOrgPolicyID, vs...))
}

// OrgPolicyIDNotIn applies the NotIn predicate on the "org_policy_id" field.
func OrgPolicyIDNotIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldOrgPolicyID, vs...))
}

// OrgPolicyIDGT applies the GT predicate on the "org_policy_id" field.
func OrgPolicyIDGT(v int) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldOrgPolicyID, v))
}

// OrgPolicyIDGTE applies the GTE predicate on the "org_policy_id" field.
func OrgPolicyIDGTE(v int) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldOrgPolicyID, v))
}

// OrgPolicyIDLT applies the LT predicate on the "org_policy_id" field.
func OrgPolicyIDLT(v int) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldOrgPolicyID, v))
}

// OrgPolicyIDLTE applies the LTE predicate on the "org_policy_id" field.
func OrgPolicyIDLTE(v int) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldOrgPolicyID, v))
}

// StartAtEQ applies the EQ predicate on the "start_at" field.
func StartAtEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldStartAt, v))
}

// StartAtNEQ applies the NEQ predicate on the "start_at" field.
func StartAtNEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldStartAt, v))
}

// StartAtIn applies the In predicate on the "start_at" field.
func StartAtIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldStartAt, vs...))
}

// StartAtNotIn applies the NotIn predicate on the "start_at" field.
func StartAtNotIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldStartAt, vs...))
}

// StartAtGT applies the GT predicate on the "start_at" field.
func StartAtGT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldStartAt, v))
}

// StartAtGTE applies the GTE predicate on the "start_at" field.
func StartAtGTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldStartAt, v))
}

// StartAtLT applies the LT predicate on the "start_at" field.
func StartAtLT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldStartAt, v))
}

// StartAtLTE applies the LTE predicate on the "start_at" field.
func StartAtLTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldStartAt, v))
}

// StartAtIsNil applies the IsNil predicate on the "start_at" field.
func StartAtIsNil() predicate.Permission {
	return predicate.Permission(sql.FieldIsNull(FieldStartAt))
}

// StartAtNotNil applies the NotNil predicate on the "start_at" field.
func StartAtNotNil() predicate.Permission {
	return predicate.Permission(sql.FieldNotNull(FieldStartAt))
}

// EndAtEQ applies the EQ predicate on the "end_at" field.
func EndAtEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldEndAt, v))
}

// EndAtNEQ applies the NEQ predicate on the "end_at" field.
func EndAtNEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldEndAt, v))
}

// EndAtIn applies the In predicate on the "end_at" field.
func EndAtIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldEndAt, vs...))
}

// EndAtNotIn applies the NotIn predicate on the "end_at" field.
func EndAtNotIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldEndAt, vs...))
}

// EndAtGT applies the GT predicate on the "end_at" field.
func EndAtGT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldEndAt, v))
}

// EndAtGTE applies the GTE predicate on the "end_at" field.
func EndAtGTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldEndAt, v))
}

// EndAtLT applies the LT predicate on the "end_at" field.
func EndAtLT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldEndAt, v))
}

// EndAtLTE applies the LTE predicate on the "end_at" field.
func EndAtLTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldEndAt, v))
}

// EndAtIsNil applies the IsNil predicate on the "end_at" field.
func EndAtIsNil() predicate.Permission {
	return predicate.Permission(sql.FieldIsNull(FieldEndAt))
}

// EndAtNotNil applies the NotNil predicate on the "end_at" field.
func EndAtNotNil() predicate.Permission {
	return predicate.Permission(sql.FieldNotNull(FieldEndAt))
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrganizationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Permission) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		p(s.Not())
	})
}
