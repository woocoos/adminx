// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/woocoos/adminx/ent/approlepolicy"
	"github.com/woocoos/adminx/ent/predicate"
)

// AppRolePolicyDelete is the builder for deleting a AppRolePolicy entity.
type AppRolePolicyDelete struct {
	config
	hooks    []Hook
	mutation *AppRolePolicyMutation
}

// Where appends a list predicates to the AppRolePolicyDelete builder.
func (arpd *AppRolePolicyDelete) Where(ps ...predicate.AppRolePolicy) *AppRolePolicyDelete {
	arpd.mutation.Where(ps...)
	return arpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (arpd *AppRolePolicyDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, AppRolePolicyMutation](ctx, arpd.sqlExec, arpd.mutation, arpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (arpd *AppRolePolicyDelete) ExecX(ctx context.Context) int {
	n, err := arpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (arpd *AppRolePolicyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(approlepolicy.Table, nil)
	if ps := arpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, arpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	arpd.mutation.done = true
	return affected, err
}

// AppRolePolicyDeleteOne is the builder for deleting a single AppRolePolicy entity.
type AppRolePolicyDeleteOne struct {
	arpd *AppRolePolicyDelete
}

// Where appends a list predicates to the AppRolePolicyDelete builder.
func (arpdo *AppRolePolicyDeleteOne) Where(ps ...predicate.AppRolePolicy) *AppRolePolicyDeleteOne {
	arpdo.arpd.mutation.Where(ps...)
	return arpdo
}

// Exec executes the deletion query.
func (arpdo *AppRolePolicyDeleteOne) Exec(ctx context.Context) error {
	n, err := arpdo.arpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{approlepolicy.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (arpdo *AppRolePolicyDeleteOne) ExecX(ctx context.Context) {
	if err := arpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
