// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/woocoos/adminx/ent/organizationapp"
	"github.com/woocoos/adminx/ent/predicate"
)

// OrganizationAppDelete is the builder for deleting a OrganizationApp entity.
type OrganizationAppDelete struct {
	config
	hooks    []Hook
	mutation *OrganizationAppMutation
}

// Where appends a list predicates to the OrganizationAppDelete builder.
func (oad *OrganizationAppDelete) Where(ps ...predicate.OrganizationApp) *OrganizationAppDelete {
	oad.mutation.Where(ps...)
	return oad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (oad *OrganizationAppDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, OrganizationAppMutation](ctx, oad.sqlExec, oad.mutation, oad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (oad *OrganizationAppDelete) ExecX(ctx context.Context) int {
	n, err := oad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (oad *OrganizationAppDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(organizationapp.Table, nil)
	if ps := oad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, oad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	oad.mutation.done = true
	return affected, err
}

// OrganizationAppDeleteOne is the builder for deleting a single OrganizationApp entity.
type OrganizationAppDeleteOne struct {
	oad *OrganizationAppDelete
}

// Where appends a list predicates to the OrganizationAppDelete builder.
func (oado *OrganizationAppDeleteOne) Where(ps ...predicate.OrganizationApp) *OrganizationAppDeleteOne {
	oado.oad.mutation.Where(ps...)
	return oado
}

// Exec executes the deletion query.
func (oado *OrganizationAppDeleteOne) Exec(ctx context.Context) error {
	n, err := oado.oad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{organizationapp.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oado *OrganizationAppDeleteOne) ExecX(ctx context.Context) {
	if err := oado.Exec(ctx); err != nil {
		panic(err)
	}
}