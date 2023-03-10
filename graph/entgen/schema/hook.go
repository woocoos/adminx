package schema

import (
	"context"
	"entgo.io/ent"
	gen "github.com/woocoos/adminx/ent"
	"github.com/woocoos/adminx/ent/hook"
	"github.com/woocoos/adminx/ent/organization"
)

// InitDisplaySortHook 初始化displaySort字段,表需要有parent_id字段.
func InitDisplaySortHook(table string) ent.Hook {
	parentField := organization.FieldParentID
	displayField := organization.FieldDisplaySort
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				pid, _ := m.Field(parentField)
				if mx, ok := m.(interface {
					SetDisplaySort(int32)
					Client() *gen.Client
				}); ok {
					var old int
					switch table {
					case organization.Table:
						old, _ = mx.Client().Organization.Query().Where(organization.ParentID(pid.(int))).
							Aggregate(gen.Max(displayField)).Int(ctx)
					}
					mx.SetDisplaySort(int32(old + 1))
				}
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate)
}
