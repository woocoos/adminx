// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/adminx/ent/appres"
	"github.com/woocoos/adminx/ent/predicate"
)

// AppResDelete is the builder for deleting a AppRes entity.
type AppResDelete struct {
	config
	hooks    []Hook
	mutation *AppResMutation
}

// Where appends a list predicates to the AppResDelete builder.
func (ard *AppResDelete) Where(ps ...predicate.AppRes) *AppResDelete {
	ard.mutation.Where(ps...)
	return ard
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ard *AppResDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, AppResMutation](ctx, ard.sqlExec, ard.mutation, ard.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ard *AppResDelete) ExecX(ctx context.Context) int {
	n, err := ard.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ard *AppResDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(appres.Table, sqlgraph.NewFieldSpec(appres.FieldID, field.TypeInt))
	if ps := ard.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ard.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ard.mutation.done = true
	return affected, err
}

// AppResDeleteOne is the builder for deleting a single AppRes entity.
type AppResDeleteOne struct {
	ard *AppResDelete
}

// Where appends a list predicates to the AppResDelete builder.
func (ardo *AppResDeleteOne) Where(ps ...predicate.AppRes) *AppResDeleteOne {
	ardo.ard.mutation.Where(ps...)
	return ardo
}

// Exec executes the deletion query.
func (ardo *AppResDeleteOne) Exec(ctx context.Context) error {
	n, err := ardo.ard.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appres.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ardo *AppResDeleteOne) ExecX(ctx context.Context) {
	if err := ardo.Exec(ctx); err != nil {
		panic(err)
	}
}
