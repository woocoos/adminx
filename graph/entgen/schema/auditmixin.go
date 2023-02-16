package schema

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/tsingsun/woocoo/pkg/security"
	"github.com/woocoos/adminx/ent/hook"
	"strconv"
	"time"
)

type AuditMixin struct {
	mixin.Schema
}

func (e AuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("created_by").Immutable().SchemaType(SnowFlakeID{}.SchemaType()).Immutable().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
		field.Time("created_at").Immutable().Default(time.Now).Immutable().
			Annotations(entgql.OrderField("createdAt"), entgql.Skip(entgql.SkipMutationCreateInput)),
		field.Int("updated_by").Optional().SchemaType(SnowFlakeID{}.SchemaType()).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Time("updated_at").Optional().Default(time.Now).UpdateDefault(time.Now).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}

func (AuditMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		onUpdateHook(),
		onCreateHook(),
	}
}

func onUpdateHook() ent.Hook {
	type setter interface {
		SetUpdatedBy(int)
	}
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			updatedBy, ok := m.Field("updated_by")
			if ok {
				s, ok := m.(setter)
				if ok && (updatedBy == 0 || updatedBy == nil) {
					up := security.GenericIdentityFromContext(ctx)
					uid, _ := strconv.Atoi(up.Name())
					s.SetUpdatedBy(uid)
				}
			}
			return next.Mutate(ctx, m)
		})
	}, ent.OpUpdateOne|ent.OpUpdate)
}

func onCreateHook() ent.Hook {
	type setter interface {
		SetUpdatedBy(int)
		SetCreatedBy(int)
	}
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			s, ok := m.(setter)
			if ok {
				up := security.GenericIdentityFromContext(ctx)
				uid, _ := strconv.Atoi(up.Name())
				s.SetUpdatedBy(uid)
				s.SetCreatedBy(uid)
			}
			return next.Mutate(ctx, m)
		})
	}, ent.OpCreate)
}
