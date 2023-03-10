package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OrganizationUser holds the schema definition for the OrganizationUser entity.
type OrganizationUser struct {
	ent.Schema
}

func (OrganizationUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "organization_user"},
		entgql.Skip(entgql.SkipType),
	}
}

func (OrganizationUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IntID{},
		AuditMixin{},
	}
}

// Fields of the OrganizationUser.
func (OrganizationUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_id").Comment("组织ID"),
		field.Int("user_id").Comment("用户ID"),
		field.String("display_name").Comment("在组织内的显示名称"),
	}
}

// Edges of the OrganizationUser.
func (OrganizationUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("organization", Organization.Type).Unique().Required().Field("org_id").
			Annotations(entgql.Skip(entgql.SkipAll)),
		edge.To("user", User.Type).Unique().Required().Field("user_id").
			Annotations(entgql.Skip(entgql.SkipAll)),
	}
}
