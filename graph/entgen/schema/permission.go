package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Permission 授权产生的权限信息.
type Permission struct {
	ent.Schema
}

func (Permission) Permission() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "permission"},
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (Permission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SnowFlakeID{},
		AuditMixin{},
	}
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_id").Comment("授权组织"),
		field.Enum("principal_kind").Values("user", "role").Comment("授权类型:角色,用户"),
		field.Int("user_id").Optional().Comment("授权类型为用户的ID"),
		field.Int("role_id").Optional().Comment("授权类型为角色或用户组的ID"),
		field.Int("org_policy_id").Comment("策略"),
		field.Time("start_at").Optional().Comment("生效开始时间"),
		field.Time("end_at").Optional().Comment("生效结束时间"),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).Ref("permissions").
			Unique().Required().Field("org_id"),
		edge.To("user", User.Type).Unique().Field("user_id"),
	}
}
