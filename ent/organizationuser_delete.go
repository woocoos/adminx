// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/adminx/ent/organizationuser"
	"github.com/woocoos/adminx/ent/predicate"
)

// OrganizationUserDelete is the builder for deleting a OrganizationUser entity.
type OrganizationUserDelete struct {
	config
	hooks    []Hook
	mutation *OrganizationUserMutation
}

// Where appends a list predicates to the OrganizationUserDelete builder.
func (oud *OrganizationUserDelete) Where(ps ...predicate.OrganizationUser) *OrganizationUserDelete {
	oud.mutation.Where(ps...)
	return oud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (oud *OrganizationUserDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, OrganizationUserMutation](ctx, oud.sqlExec, oud.mutation, oud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (oud *OrganizationUserDelete) ExecX(ctx context.Context) int {
	n, err := oud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (oud *OrganizationUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(organizationuser.Table, sqlgraph.NewFieldSpec(organizationuser.FieldID, field.TypeInt))
	if ps := oud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, oud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	oud.mutation.done = true
	return affected, err
}

// OrganizationUserDeleteOne is the builder for deleting a single OrganizationUser entity.
type OrganizationUserDeleteOne struct {
	oud *OrganizationUserDelete
}

// Where appends a list predicates to the OrganizationUserDelete builder.
func (oudo *OrganizationUserDeleteOne) Where(ps ...predicate.OrganizationUser) *OrganizationUserDeleteOne {
	oudo.oud.mutation.Where(ps...)
	return oudo
}

// Exec executes the deletion query.
func (oudo *OrganizationUserDeleteOne) Exec(ctx context.Context) error {
	n, err := oudo.oud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{organizationuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oudo *OrganizationUserDeleteOne) ExecX(ctx context.Context) {
	if err := oudo.Exec(ctx); err != nil {
		panic(err)
	}
}
