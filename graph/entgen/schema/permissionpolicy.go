package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/adminx/graph/entgen/types"
)

// PermissionPolicy 组织中的策略.基本包括来源于应用初始化的策略和组织自定义的策略.
//
// 在策略制定UI时,原则上策略应该对应一个应用,但是也可以是自定义策略.
type PermissionPolicy struct {
	ent.Schema
}

func (PermissionPolicy) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "organization_policy"},
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (PermissionPolicy) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SnowFlakeID{},
		AuditMixin{},
	}
}

// Fields of the OrganizationPolicy.
func (PermissionPolicy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_id").Comment("组织ID"),
		field.Int("app_id").Optional().Comment("所属应用").Annotations(entgql.Skip(entgql.SkipAll)),
		field.Int("app_policy_id").Optional().Comment("所属应用策略,如果是自定义应用策略,则为空"),
		field.String("name").Comment("策略名称"),
		field.String("comments").Comment("描述"),
		field.JSON("rules", []types.PolicyRule{}).Comment("策略规则,如果是应用策略,则为空"),
	}
}

// Edges of the OrganizationPolicy.
func (PermissionPolicy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).Ref("policies").Unique().Required().Field("org_id"),
	}
}
