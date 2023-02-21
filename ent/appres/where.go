// Code generated by ent, DO NOT EDIT.

package appres

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/woocoos/adminx/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AppRes {
	return predicate.AppRes(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AppRes {
	return predicate.AppRes(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AppRes {
	return predicate.AppRes(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AppRes {
	return predicate.AppRes(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AppRes {
	return predicate.AppRes(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AppRes {
	return predicate.AppRes(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AppRes {
	return predicate.AppRes(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AppRes {
	return predicate.AppRes(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AppRes {
	return predicate.AppRes(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int) predicate.AppRes {
	return predicate.AppRes(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AppRes {
	return predicate.AppRes(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int) predicate.AppRes {
	return predicate.AppRes(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AppRes {
	return predicate.AppRes(sql.FieldEQ(FieldUpdatedAt, v))
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v int) predicate.AppRes {
	return predicate.AppRes(sql.FieldEQ(FieldAppID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldEQ(FieldName, v))
}

// TypeName applies equality check predicate on the "type_name" field. It's identical to TypeNameEQ.
func TypeName(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldEQ(FieldTypeName, v))
}

// ArnPattern applies equality check predicate on the "arn_pattern" field. It's identical to ArnPatternEQ.
func ArnPattern(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldEQ(FieldArnPattern, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int) predicate.AppRes {
	return predicate.AppRes(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int) predicate.AppRes {
	return predicate.AppRes(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int) predicate.AppRes {
	return predicate.AppRes(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int) predicate.AppRes {
	return predicate.AppRes(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int) predicate.AppRes {
	return predicate.AppRes(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int) predicate.AppRes {
	return predicate.AppRes(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int) predicate.AppRes {
	return predicate.AppRes(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int) predicate.AppRes {
	return predicate.AppRes(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AppRes {
	return predicate.AppRes(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AppRes {
	return predicate.AppRes(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AppRes {
	return predicate.AppRes(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AppRes {
	return predicate.AppRes(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AppRes {
	return predicate.AppRes(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AppRes {
	return predicate.AppRes(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AppRes {
	return predicate.AppRes(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AppRes {
	return predicate.AppRes(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int) predicate.AppRes {
	return predicate.AppRes(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int) predicate.AppRes {
	return predicate.AppRes(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int) predicate.AppRes {
	return predicate.AppRes(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int) predicate.AppRes {
	return predicate.AppRes(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int) predicate.AppRes {
	return predicate.AppRes(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int) predicate.AppRes {
	return predicate.AppRes(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int) predicate.AppRes {
	return predicate.AppRes(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int) predicate.AppRes {
	return predicate.AppRes(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.AppRes {
	return predicate.AppRes(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.AppRes {
	return predicate.AppRes(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AppRes {
	return predicate.AppRes(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AppRes {
	return predicate.AppRes(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AppRes {
	return predicate.AppRes(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AppRes {
	return predicate.AppRes(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AppRes {
	return predicate.AppRes(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AppRes {
	return predicate.AppRes(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AppRes {
	return predicate.AppRes(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AppRes {
	return predicate.AppRes(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.AppRes {
	return predicate.AppRes(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.AppRes {
	return predicate.AppRes(sql.FieldNotNull(FieldUpdatedAt))
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v int) predicate.AppRes {
	return predicate.AppRes(sql.FieldEQ(FieldAppID, v))
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v int) predicate.AppRes {
	return predicate.AppRes(sql.FieldNEQ(FieldAppID, v))
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...int) predicate.AppRes {
	return predicate.AppRes(sql.FieldIn(FieldAppID, vs...))
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...int) predicate.AppRes {
	return predicate.AppRes(sql.FieldNotIn(FieldAppID, vs...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.AppRes {
	return predicate.AppRes(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.AppRes {
	return predicate.AppRes(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldContainsFold(FieldName, v))
}

// TypeNameEQ applies the EQ predicate on the "type_name" field.
func TypeNameEQ(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldEQ(FieldTypeName, v))
}

// TypeNameNEQ applies the NEQ predicate on the "type_name" field.
func TypeNameNEQ(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldNEQ(FieldTypeName, v))
}

// TypeNameIn applies the In predicate on the "type_name" field.
func TypeNameIn(vs ...string) predicate.AppRes {
	return predicate.AppRes(sql.FieldIn(FieldTypeName, vs...))
}

// TypeNameNotIn applies the NotIn predicate on the "type_name" field.
func TypeNameNotIn(vs ...string) predicate.AppRes {
	return predicate.AppRes(sql.FieldNotIn(FieldTypeName, vs...))
}

// TypeNameGT applies the GT predicate on the "type_name" field.
func TypeNameGT(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldGT(FieldTypeName, v))
}

// TypeNameGTE applies the GTE predicate on the "type_name" field.
func TypeNameGTE(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldGTE(FieldTypeName, v))
}

// TypeNameLT applies the LT predicate on the "type_name" field.
func TypeNameLT(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldLT(FieldTypeName, v))
}

// TypeNameLTE applies the LTE predicate on the "type_name" field.
func TypeNameLTE(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldLTE(FieldTypeName, v))
}

// TypeNameContains applies the Contains predicate on the "type_name" field.
func TypeNameContains(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldContains(FieldTypeName, v))
}

// TypeNameHasPrefix applies the HasPrefix predicate on the "type_name" field.
func TypeNameHasPrefix(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldHasPrefix(FieldTypeName, v))
}

// TypeNameHasSuffix applies the HasSuffix predicate on the "type_name" field.
func TypeNameHasSuffix(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldHasSuffix(FieldTypeName, v))
}

// TypeNameEqualFold applies the EqualFold predicate on the "type_name" field.
func TypeNameEqualFold(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldEqualFold(FieldTypeName, v))
}

// TypeNameContainsFold applies the ContainsFold predicate on the "type_name" field.
func TypeNameContainsFold(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldContainsFold(FieldTypeName, v))
}

// ArnPatternEQ applies the EQ predicate on the "arn_pattern" field.
func ArnPatternEQ(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldEQ(FieldArnPattern, v))
}

// ArnPatternNEQ applies the NEQ predicate on the "arn_pattern" field.
func ArnPatternNEQ(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldNEQ(FieldArnPattern, v))
}

// ArnPatternIn applies the In predicate on the "arn_pattern" field.
func ArnPatternIn(vs ...string) predicate.AppRes {
	return predicate.AppRes(sql.FieldIn(FieldArnPattern, vs...))
}

// ArnPatternNotIn applies the NotIn predicate on the "arn_pattern" field.
func ArnPatternNotIn(vs ...string) predicate.AppRes {
	return predicate.AppRes(sql.FieldNotIn(FieldArnPattern, vs...))
}

// ArnPatternGT applies the GT predicate on the "arn_pattern" field.
func ArnPatternGT(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldGT(FieldArnPattern, v))
}

// ArnPatternGTE applies the GTE predicate on the "arn_pattern" field.
func ArnPatternGTE(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldGTE(FieldArnPattern, v))
}

// ArnPatternLT applies the LT predicate on the "arn_pattern" field.
func ArnPatternLT(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldLT(FieldArnPattern, v))
}

// ArnPatternLTE applies the LTE predicate on the "arn_pattern" field.
func ArnPatternLTE(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldLTE(FieldArnPattern, v))
}

// ArnPatternContains applies the Contains predicate on the "arn_pattern" field.
func ArnPatternContains(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldContains(FieldArnPattern, v))
}

// ArnPatternHasPrefix applies the HasPrefix predicate on the "arn_pattern" field.
func ArnPatternHasPrefix(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldHasPrefix(FieldArnPattern, v))
}

// ArnPatternHasSuffix applies the HasSuffix predicate on the "arn_pattern" field.
func ArnPatternHasSuffix(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldHasSuffix(FieldArnPattern, v))
}

// ArnPatternEqualFold applies the EqualFold predicate on the "arn_pattern" field.
func ArnPatternEqualFold(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldEqualFold(FieldArnPattern, v))
}

// ArnPatternContainsFold applies the ContainsFold predicate on the "arn_pattern" field.
func ArnPatternContainsFold(v string) predicate.AppRes {
	return predicate.AppRes(sql.FieldContainsFold(FieldArnPattern, v))
}

// HasApp applies the HasEdge predicate on the "app" edge.
func HasApp() predicate.AppRes {
	return predicate.AppRes(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AppTable, AppColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAppWith applies the HasEdge predicate on the "app" edge with a given conditions (other predicates).
func HasAppWith(preds ...predicate.App) predicate.AppRes {
	return predicate.AppRes(func(s *sql.Selector) {
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

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppRes) predicate.AppRes {
	return predicate.AppRes(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppRes) predicate.AppRes {
	return predicate.AppRes(func(s *sql.Selector) {
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
func Not(p predicate.AppRes) predicate.AppRes {
	return predicate.AppRes(func(s *sql.Selector) {
		p(s.Not())
	})
}
