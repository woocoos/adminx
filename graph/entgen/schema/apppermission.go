package schema

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"fmt"
	gen "github.com/woocoos/adminx/ent"
	"github.com/woocoos/adminx/ent/apppermission"
	"github.com/woocoos/adminx/ent/hook"
	"regexp"
)

// AppPermission holds the schema definition for the AppPermission entity.
type AppPermission struct {
	ent.Schema
}

func (AppPermission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_permission"},
		entgql.RelayConnection(),
	}
}

func (AppPermission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SnowFlakeID{},
		AuditMixin{},
	}
}

// Fields of the AppPermission.
func (AppPermission) Fields() []ent.Field {
	return []ent.Field{
		field.Int("app_id").Comment("所属应用"),
		field.String("name").Optional().Comment("名称").
			Match(regexp.MustCompile("[a-z/]+$")),
		field.String("comments").Optional().Comment("备注").
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
	}
}

// Edges of the AppPermission.
func (AppPermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app", App.Type).Ref("permissions").Unique().Required().Field("app_id"),
		edge.To("menus", AppMenu.Type).Comment("被引用的菜单项"),
	}
}

func (AppPermission) Hooks() []ent.Hook {
	// app name unique
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.AppPermissionFunc(func(ctx context.Context, m *gen.AppPermissionMutation) (gen.Value, error) {
				if m.Op() == ent.OpCreate {
					appid, _ := m.AppID()
					n, _ := m.Name()
					exists, err := m.Client().AppPermission.Query().
						Where(apppermission.AppID(appid), apppermission.Name(n)).Exist(ctx)
					if err != nil {
						return nil, err
					}
					if exists {
						return nil, fmt.Errorf("app permission name %s already exists", n)
					}
				} else if m.Op() == ent.OpUpdateOne || m.Op() == ent.OpUpdate {
					appid, _ := m.AppID()
					n, _ := m.Name()
					c, err := m.Client().AppPermission.Query().
						Where(apppermission.AppID(appid), apppermission.Name(n)).Count(ctx)
					if err != nil {
						return nil, err
					}
					if c > 1 {
						return nil, fmt.Errorf("app permission name %s already exists", n)
					}
				}
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne),
	}
}
