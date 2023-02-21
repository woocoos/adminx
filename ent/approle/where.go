// Code generated by ent, DO NOT EDIT.

package approle

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/woocoos/adminx/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AppRole {
	return predicate.AppRole(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AppRole {
	return predicate.AppRole(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AppRole {
	return predicate.AppRole(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AppRole {
	return predicate.AppRole(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AppRole {
	return predicate.AppRole(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AppRole {
	return predicate.AppRole(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AppRole {
	return predicate.AppRole(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldUpdatedAt, v))
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v int) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldAppID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldName, v))
}

// Comments applies equality check predicate on the "comments" field. It's identical to CommentsEQ.
func Comments(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldComments, v))
}

// AutoGrant applies equality check predicate on the "auto_grant" field. It's identical to AutoGrantEQ.
func AutoGrant(v bool) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldAutoGrant, v))
}

// Editable applies equality check predicate on the "editable" field. It's identical to EditableEQ.
func Editable(v bool) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldEditable, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int) predicate.AppRole {
	return predicate.AppRole(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int) predicate.AppRole {
	return predicate.AppRole(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int) predicate.AppRole {
	return predicate.AppRole(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int) predicate.AppRole {
	return predicate.AppRole(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int) predicate.AppRole {
	return predicate.AppRole(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int) predicate.AppRole {
	return predicate.AppRole(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int) predicate.AppRole {
	return predicate.AppRole(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AppRole {
	return predicate.AppRole(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AppRole {
	return predicate.AppRole(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AppRole {
	return predicate.AppRole(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AppRole {
	return predicate.AppRole(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AppRole {
	return predicate.AppRole(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AppRole {
	return predicate.AppRole(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AppRole {
	return predicate.AppRole(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int) predicate.AppRole {
	return predicate.AppRole(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int) predicate.AppRole {
	return predicate.AppRole(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int) predicate.AppRole {
	return predicate.AppRole(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int) predicate.AppRole {
	return predicate.AppRole(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int) predicate.AppRole {
	return predicate.AppRole(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int) predicate.AppRole {
	return predicate.AppRole(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int) predicate.AppRole {
	return predicate.AppRole(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.AppRole {
	return predicate.AppRole(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.AppRole {
	return predicate.AppRole(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AppRole {
	return predicate.AppRole(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AppRole {
	return predicate.AppRole(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AppRole {
	return predicate.AppRole(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AppRole {
	return predicate.AppRole(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AppRole {
	return predicate.AppRole(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AppRole {
	return predicate.AppRole(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AppRole {
	return predicate.AppRole(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.AppRole {
	return predicate.AppRole(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.AppRole {
	return predicate.AppRole(sql.FieldNotNull(FieldUpdatedAt))
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v int) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldAppID, v))
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v int) predicate.AppRole {
	return predicate.AppRole(sql.FieldNEQ(FieldAppID, v))
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...int) predicate.AppRole {
	return predicate.AppRole(sql.FieldIn(FieldAppID, vs...))
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...int) predicate.AppRole {
	return predicate.AppRole(sql.FieldNotIn(FieldAppID, vs...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.AppRole {
	return predicate.AppRole(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.AppRole {
	return predicate.AppRole(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldContainsFold(FieldName, v))
}

// CommentsEQ applies the EQ predicate on the "comments" field.
func CommentsEQ(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldComments, v))
}

// CommentsNEQ applies the NEQ predicate on the "comments" field.
func CommentsNEQ(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldNEQ(FieldComments, v))
}

// CommentsIn applies the In predicate on the "comments" field.
func CommentsIn(vs ...string) predicate.AppRole {
	return predicate.AppRole(sql.FieldIn(FieldComments, vs...))
}

// CommentsNotIn applies the NotIn predicate on the "comments" field.
func CommentsNotIn(vs ...string) predicate.AppRole {
	return predicate.AppRole(sql.FieldNotIn(FieldComments, vs...))
}

// CommentsGT applies the GT predicate on the "comments" field.
func CommentsGT(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldGT(FieldComments, v))
}

// CommentsGTE applies the GTE predicate on the "comments" field.
func CommentsGTE(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldGTE(FieldComments, v))
}

// CommentsLT applies the LT predicate on the "comments" field.
func CommentsLT(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldLT(FieldComments, v))
}

// CommentsLTE applies the LTE predicate on the "comments" field.
func CommentsLTE(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldLTE(FieldComments, v))
}

// CommentsContains applies the Contains predicate on the "comments" field.
func CommentsContains(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldContains(FieldComments, v))
}

// CommentsHasPrefix applies the HasPrefix predicate on the "comments" field.
func CommentsHasPrefix(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldHasPrefix(FieldComments, v))
}

// CommentsHasSuffix applies the HasSuffix predicate on the "comments" field.
func CommentsHasSuffix(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldHasSuffix(FieldComments, v))
}

// CommentsIsNil applies the IsNil predicate on the "comments" field.
func CommentsIsNil() predicate.AppRole {
	return predicate.AppRole(sql.FieldIsNull(FieldComments))
}

// CommentsNotNil applies the NotNil predicate on the "comments" field.
func CommentsNotNil() predicate.AppRole {
	return predicate.AppRole(sql.FieldNotNull(FieldComments))
}

// CommentsEqualFold applies the EqualFold predicate on the "comments" field.
func CommentsEqualFold(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldEqualFold(FieldComments, v))
}

// CommentsContainsFold applies the ContainsFold predicate on the "comments" field.
func CommentsContainsFold(v string) predicate.AppRole {
	return predicate.AppRole(sql.FieldContainsFold(FieldComments, v))
}

// AutoGrantEQ applies the EQ predicate on the "auto_grant" field.
func AutoGrantEQ(v bool) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldAutoGrant, v))
}

// AutoGrantNEQ applies the NEQ predicate on the "auto_grant" field.
func AutoGrantNEQ(v bool) predicate.AppRole {
	return predicate.AppRole(sql.FieldNEQ(FieldAutoGrant, v))
}

// EditableEQ applies the EQ predicate on the "editable" field.
func EditableEQ(v bool) predicate.AppRole {
	return predicate.AppRole(sql.FieldEQ(FieldEditable, v))
}

// EditableNEQ applies the NEQ predicate on the "editable" field.
func EditableNEQ(v bool) predicate.AppRole {
	return predicate.AppRole(sql.FieldNEQ(FieldEditable, v))
}

// HasApp applies the HasEdge predicate on the "app" edge.
func HasApp() predicate.AppRole {
	return predicate.AppRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AppTable, AppColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAppWith applies the HasEdge predicate on the "app" edge with a given conditions (other predicates).
func HasAppWith(preds ...predicate.App) predicate.AppRole {
	return predicate.AppRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AppInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AppTable, AppColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPolicies applies the HasEdge predicate on the "policies" edge.
func HasPolicies() predicate.AppRole {
	return predicate.AppRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PoliciesTable, PoliciesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPoliciesWith applies the HasEdge predicate on the "policies" edge with a given conditions (other predicates).
func HasPoliciesWith(preds ...predicate.AppPolicy) predicate.AppRole {
	return predicate.AppRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PoliciesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PoliciesTable, PoliciesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAppRolePolicy applies the HasEdge predicate on the "app_role_policy" edge.
func HasAppRolePolicy() predicate.AppRole {
	return predicate.AppRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, AppRolePolicyTable, AppRolePolicyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAppRolePolicyWith applies the HasEdge predicate on the "app_role_policy" edge with a given conditions (other predicates).
func HasAppRolePolicyWith(preds ...predicate.AppRolePolicy) predicate.AppRole {
	return predicate.AppRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AppRolePolicyInverseTable, AppRolePolicyColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, AppRolePolicyTable, AppRolePolicyColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppRole) predicate.AppRole {
	return predicate.AppRole(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppRole) predicate.AppRole {
	return predicate.AppRole(func(s *sql.Selector) {
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
func Not(p predicate.AppRole) predicate.AppRole {
	return predicate.AppRole(func(s *sql.Selector) {
		p(s.Not())
	})
}