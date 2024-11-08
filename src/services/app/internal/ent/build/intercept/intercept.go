// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/newsitem"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/organization"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/person"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/predicate"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/rssauthor"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/rssfeed"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next build.Querier) build.Querier {
	return build.QuerierFunc(func(ctx context.Context, q build.Query) (build.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next build.Querier) build.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q build.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The NewsItemFunc type is an adapter to allow the use of ordinary function as a Querier.
type NewsItemFunc func(context.Context, *build.NewsItemQuery) (build.Value, error)

// Query calls f(ctx, q).
func (f NewsItemFunc) Query(ctx context.Context, q build.Query) (build.Value, error) {
	if q, ok := q.(*build.NewsItemQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *build.NewsItemQuery", q)
}

// The TraverseNewsItem type is an adapter to allow the use of ordinary function as Traverser.
type TraverseNewsItem func(context.Context, *build.NewsItemQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseNewsItem) Intercept(next build.Querier) build.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseNewsItem) Traverse(ctx context.Context, q build.Query) error {
	if q, ok := q.(*build.NewsItemQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *build.NewsItemQuery", q)
}

// The OrganizationFunc type is an adapter to allow the use of ordinary function as a Querier.
type OrganizationFunc func(context.Context, *build.OrganizationQuery) (build.Value, error)

// Query calls f(ctx, q).
func (f OrganizationFunc) Query(ctx context.Context, q build.Query) (build.Value, error) {
	if q, ok := q.(*build.OrganizationQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *build.OrganizationQuery", q)
}

// The TraverseOrganization type is an adapter to allow the use of ordinary function as Traverser.
type TraverseOrganization func(context.Context, *build.OrganizationQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseOrganization) Intercept(next build.Querier) build.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseOrganization) Traverse(ctx context.Context, q build.Query) error {
	if q, ok := q.(*build.OrganizationQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *build.OrganizationQuery", q)
}

// The PersonFunc type is an adapter to allow the use of ordinary function as a Querier.
type PersonFunc func(context.Context, *build.PersonQuery) (build.Value, error)

// Query calls f(ctx, q).
func (f PersonFunc) Query(ctx context.Context, q build.Query) (build.Value, error) {
	if q, ok := q.(*build.PersonQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *build.PersonQuery", q)
}

// The TraversePerson type is an adapter to allow the use of ordinary function as Traverser.
type TraversePerson func(context.Context, *build.PersonQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraversePerson) Intercept(next build.Querier) build.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraversePerson) Traverse(ctx context.Context, q build.Query) error {
	if q, ok := q.(*build.PersonQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *build.PersonQuery", q)
}

// The RSSAuthorFunc type is an adapter to allow the use of ordinary function as a Querier.
type RSSAuthorFunc func(context.Context, *build.RSSAuthorQuery) (build.Value, error)

// Query calls f(ctx, q).
func (f RSSAuthorFunc) Query(ctx context.Context, q build.Query) (build.Value, error) {
	if q, ok := q.(*build.RSSAuthorQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *build.RSSAuthorQuery", q)
}

// The TraverseRSSAuthor type is an adapter to allow the use of ordinary function as Traverser.
type TraverseRSSAuthor func(context.Context, *build.RSSAuthorQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseRSSAuthor) Intercept(next build.Querier) build.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseRSSAuthor) Traverse(ctx context.Context, q build.Query) error {
	if q, ok := q.(*build.RSSAuthorQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *build.RSSAuthorQuery", q)
}

// The RSSFeedFunc type is an adapter to allow the use of ordinary function as a Querier.
type RSSFeedFunc func(context.Context, *build.RSSFeedQuery) (build.Value, error)

// Query calls f(ctx, q).
func (f RSSFeedFunc) Query(ctx context.Context, q build.Query) (build.Value, error) {
	if q, ok := q.(*build.RSSFeedQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *build.RSSFeedQuery", q)
}

// The TraverseRSSFeed type is an adapter to allow the use of ordinary function as Traverser.
type TraverseRSSFeed func(context.Context, *build.RSSFeedQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseRSSFeed) Intercept(next build.Querier) build.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseRSSFeed) Traverse(ctx context.Context, q build.Query) error {
	if q, ok := q.(*build.RSSFeedQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *build.RSSFeedQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q build.Query) (Query, error) {
	switch q := q.(type) {
	case *build.NewsItemQuery:
		return &query[*build.NewsItemQuery, predicate.NewsItem, newsitem.OrderOption]{typ: build.TypeNewsItem, tq: q}, nil
	case *build.OrganizationQuery:
		return &query[*build.OrganizationQuery, predicate.Organization, organization.OrderOption]{typ: build.TypeOrganization, tq: q}, nil
	case *build.PersonQuery:
		return &query[*build.PersonQuery, predicate.Person, person.OrderOption]{typ: build.TypePerson, tq: q}, nil
	case *build.RSSAuthorQuery:
		return &query[*build.RSSAuthorQuery, predicate.RSSAuthor, rssauthor.OrderOption]{typ: build.TypeRSSAuthor, tq: q}, nil
	case *build.RSSFeedQuery:
		return &query[*build.RSSFeedQuery, predicate.RSSFeed, rssfeed.OrderOption]{typ: build.TypeRSSFeed, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}
