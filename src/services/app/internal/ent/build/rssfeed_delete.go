// Code generated by ent, DO NOT EDIT.

package build

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/predicate"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/rssfeed"
)

// RSSFeedDelete is the builder for deleting a RSSFeed entity.
type RSSFeedDelete struct {
	config
	hooks    []Hook
	mutation *RSSFeedMutation
}

// Where appends a list predicates to the RSSFeedDelete builder.
func (rfd *RSSFeedDelete) Where(ps ...predicate.RSSFeed) *RSSFeedDelete {
	rfd.mutation.Where(ps...)
	return rfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rfd *RSSFeedDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rfd.sqlExec, rfd.mutation, rfd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rfd *RSSFeedDelete) ExecX(ctx context.Context) int {
	n, err := rfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rfd *RSSFeedDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(rssfeed.Table, sqlgraph.NewFieldSpec(rssfeed.FieldID, field.TypeUUID))
	if ps := rfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rfd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rfd.mutation.done = true
	return affected, err
}

// RSSFeedDeleteOne is the builder for deleting a single RSSFeed entity.
type RSSFeedDeleteOne struct {
	rfd *RSSFeedDelete
}

// Where appends a list predicates to the RSSFeedDelete builder.
func (rfdo *RSSFeedDeleteOne) Where(ps ...predicate.RSSFeed) *RSSFeedDeleteOne {
	rfdo.rfd.mutation.Where(ps...)
	return rfdo
}

// Exec executes the deletion query.
func (rfdo *RSSFeedDeleteOne) Exec(ctx context.Context) error {
	n, err := rfdo.rfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{rssfeed.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rfdo *RSSFeedDeleteOne) ExecX(ctx context.Context) {
	if err := rfdo.Exec(ctx); err != nil {
		panic(err)
	}
}
