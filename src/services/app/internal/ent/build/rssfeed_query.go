// Code generated by ent, DO NOT EDIT.

package build

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/newsitem"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/organization"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/predicate"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/rssfeed"
)

// RSSFeedQuery is the builder for querying RSSFeed entities.
type RSSFeedQuery struct {
	config
	ctx              *QueryContext
	order            []rssfeed.OrderOption
	inters           []Interceptor
	predicates       []predicate.RSSFeed
	withItems        *NewsItemQuery
	withOrganization *OrganizationQuery
	withFKs          bool
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RSSFeedQuery builder.
func (rfq *RSSFeedQuery) Where(ps ...predicate.RSSFeed) *RSSFeedQuery {
	rfq.predicates = append(rfq.predicates, ps...)
	return rfq
}

// Limit the number of records to be returned by this query.
func (rfq *RSSFeedQuery) Limit(limit int) *RSSFeedQuery {
	rfq.ctx.Limit = &limit
	return rfq
}

// Offset to start from.
func (rfq *RSSFeedQuery) Offset(offset int) *RSSFeedQuery {
	rfq.ctx.Offset = &offset
	return rfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rfq *RSSFeedQuery) Unique(unique bool) *RSSFeedQuery {
	rfq.ctx.Unique = &unique
	return rfq
}

// Order specifies how the records should be ordered.
func (rfq *RSSFeedQuery) Order(o ...rssfeed.OrderOption) *RSSFeedQuery {
	rfq.order = append(rfq.order, o...)
	return rfq
}

// QueryItems chains the current query on the "items" edge.
func (rfq *RSSFeedQuery) QueryItems() *NewsItemQuery {
	query := (&NewsItemClient{config: rfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rssfeed.Table, rssfeed.FieldID, selector),
			sqlgraph.To(newsitem.Table, newsitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, rssfeed.ItemsTable, rssfeed.ItemsColumn),
		)
		fromU = sqlgraph.SetNeighbors(rfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (rfq *RSSFeedQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: rfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rssfeed.Table, rssfeed.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, rssfeed.OrganizationTable, rssfeed.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(rfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RSSFeed entity from the query.
// Returns a *NotFoundError when no RSSFeed was found.
func (rfq *RSSFeedQuery) First(ctx context.Context) (*RSSFeed, error) {
	nodes, err := rfq.Limit(1).All(setContextOp(ctx, rfq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rssfeed.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rfq *RSSFeedQuery) FirstX(ctx context.Context) *RSSFeed {
	node, err := rfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RSSFeed ID from the query.
// Returns a *NotFoundError when no RSSFeed ID was found.
func (rfq *RSSFeedQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rfq.Limit(1).IDs(setContextOp(ctx, rfq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{rssfeed.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rfq *RSSFeedQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := rfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RSSFeed entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RSSFeed entity is found.
// Returns a *NotFoundError when no RSSFeed entities are found.
func (rfq *RSSFeedQuery) Only(ctx context.Context) (*RSSFeed, error) {
	nodes, err := rfq.Limit(2).All(setContextOp(ctx, rfq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rssfeed.Label}
	default:
		return nil, &NotSingularError{rssfeed.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rfq *RSSFeedQuery) OnlyX(ctx context.Context) *RSSFeed {
	node, err := rfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RSSFeed ID in the query.
// Returns a *NotSingularError when more than one RSSFeed ID is found.
// Returns a *NotFoundError when no entities are found.
func (rfq *RSSFeedQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rfq.Limit(2).IDs(setContextOp(ctx, rfq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{rssfeed.Label}
	default:
		err = &NotSingularError{rssfeed.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rfq *RSSFeedQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := rfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RSSFeeds.
func (rfq *RSSFeedQuery) All(ctx context.Context) ([]*RSSFeed, error) {
	ctx = setContextOp(ctx, rfq.ctx, "All")
	if err := rfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RSSFeed, *RSSFeedQuery]()
	return withInterceptors[[]*RSSFeed](ctx, rfq, qr, rfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rfq *RSSFeedQuery) AllX(ctx context.Context) []*RSSFeed {
	nodes, err := rfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RSSFeed IDs.
func (rfq *RSSFeedQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if rfq.ctx.Unique == nil && rfq.path != nil {
		rfq.Unique(true)
	}
	ctx = setContextOp(ctx, rfq.ctx, "IDs")
	if err = rfq.Select(rssfeed.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rfq *RSSFeedQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := rfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rfq *RSSFeedQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rfq.ctx, "Count")
	if err := rfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rfq, querierCount[*RSSFeedQuery](), rfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rfq *RSSFeedQuery) CountX(ctx context.Context) int {
	count, err := rfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rfq *RSSFeedQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rfq.ctx, "Exist")
	switch _, err := rfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("build: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rfq *RSSFeedQuery) ExistX(ctx context.Context) bool {
	exist, err := rfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RSSFeedQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rfq *RSSFeedQuery) Clone() *RSSFeedQuery {
	if rfq == nil {
		return nil
	}
	return &RSSFeedQuery{
		config:           rfq.config,
		ctx:              rfq.ctx.Clone(),
		order:            append([]rssfeed.OrderOption{}, rfq.order...),
		inters:           append([]Interceptor{}, rfq.inters...),
		predicates:       append([]predicate.RSSFeed{}, rfq.predicates...),
		withItems:        rfq.withItems.Clone(),
		withOrganization: rfq.withOrganization.Clone(),
		// clone intermediate query.
		sql:  rfq.sql.Clone(),
		path: rfq.path,
	}
}

// WithItems tells the query-builder to eager-load the nodes that are connected to
// the "items" edge. The optional arguments are used to configure the query builder of the edge.
func (rfq *RSSFeedQuery) WithItems(opts ...func(*NewsItemQuery)) *RSSFeedQuery {
	query := (&NewsItemClient{config: rfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rfq.withItems = query
	return rfq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (rfq *RSSFeedQuery) WithOrganization(opts ...func(*OrganizationQuery)) *RSSFeedQuery {
	query := (&OrganizationClient{config: rfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rfq.withOrganization = query
	return rfq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RSSFeed.Query().
//		GroupBy(rssfeed.FieldCreateTime).
//		Aggregate(build.Count()).
//		Scan(ctx, &v)
func (rfq *RSSFeedQuery) GroupBy(field string, fields ...string) *RSSFeedGroupBy {
	rfq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RSSFeedGroupBy{build: rfq}
	grbuild.flds = &rfq.ctx.Fields
	grbuild.label = rssfeed.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.RSSFeed.Query().
//		Select(rssfeed.FieldCreateTime).
//		Scan(ctx, &v)
func (rfq *RSSFeedQuery) Select(fields ...string) *RSSFeedSelect {
	rfq.ctx.Fields = append(rfq.ctx.Fields, fields...)
	sbuild := &RSSFeedSelect{RSSFeedQuery: rfq}
	sbuild.label = rssfeed.Label
	sbuild.flds, sbuild.scan = &rfq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RSSFeedSelect configured with the given aggregations.
func (rfq *RSSFeedQuery) Aggregate(fns ...AggregateFunc) *RSSFeedSelect {
	return rfq.Select().Aggregate(fns...)
}

func (rfq *RSSFeedQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rfq.inters {
		if inter == nil {
			return fmt.Errorf("build: uninitialized interceptor (forgotten import build/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rfq); err != nil {
				return err
			}
		}
	}
	for _, f := range rfq.ctx.Fields {
		if !rssfeed.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("build: invalid field %q for query", f)}
		}
	}
	if rfq.path != nil {
		prev, err := rfq.path(ctx)
		if err != nil {
			return err
		}
		rfq.sql = prev
	}
	return nil
}

func (rfq *RSSFeedQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RSSFeed, error) {
	var (
		nodes       = []*RSSFeed{}
		withFKs     = rfq.withFKs
		_spec       = rfq.querySpec()
		loadedTypes = [2]bool{
			rfq.withItems != nil,
			rfq.withOrganization != nil,
		}
	)
	if rfq.withOrganization != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, rssfeed.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RSSFeed).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RSSFeed{config: rfq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(rfq.modifiers) > 0 {
		_spec.Modifiers = rfq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rfq.withItems; query != nil {
		if err := rfq.loadItems(ctx, query, nodes,
			func(n *RSSFeed) { n.Edges.Items = []*NewsItem{} },
			func(n *RSSFeed, e *NewsItem) { n.Edges.Items = append(n.Edges.Items, e) }); err != nil {
			return nil, err
		}
	}
	if query := rfq.withOrganization; query != nil {
		if err := rfq.loadOrganization(ctx, query, nodes, nil,
			func(n *RSSFeed, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rfq *RSSFeedQuery) loadItems(ctx context.Context, query *NewsItemQuery, nodes []*RSSFeed, init func(*RSSFeed), assign func(*RSSFeed, *NewsItem)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*RSSFeed)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.NewsItem(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(rssfeed.ItemsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.news_item_feed
		if fk == nil {
			return fmt.Errorf(`foreign-key "news_item_feed" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "news_item_feed" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (rfq *RSSFeedQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*RSSFeed, init func(*RSSFeed), assign func(*RSSFeed, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*RSSFeed)
	for i := range nodes {
		if nodes[i].organization_feeds == nil {
			continue
		}
		fk := *nodes[i].organization_feeds
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(organization.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "organization_feeds" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rfq *RSSFeedQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rfq.querySpec()
	if len(rfq.modifiers) > 0 {
		_spec.Modifiers = rfq.modifiers
	}
	_spec.Node.Columns = rfq.ctx.Fields
	if len(rfq.ctx.Fields) > 0 {
		_spec.Unique = rfq.ctx.Unique != nil && *rfq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rfq.driver, _spec)
}

func (rfq *RSSFeedQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(rssfeed.Table, rssfeed.Columns, sqlgraph.NewFieldSpec(rssfeed.FieldID, field.TypeUUID))
	_spec.From = rfq.sql
	if unique := rfq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rfq.path != nil {
		_spec.Unique = true
	}
	if fields := rfq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rssfeed.FieldID)
		for i := range fields {
			if fields[i] != rssfeed.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rfq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rfq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rfq *RSSFeedQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rfq.driver.Dialect())
	t1 := builder.Table(rssfeed.Table)
	columns := rfq.ctx.Fields
	if len(columns) == 0 {
		columns = rssfeed.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rfq.sql != nil {
		selector = rfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rfq.ctx.Unique != nil && *rfq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range rfq.modifiers {
		m(selector)
	}
	for _, p := range rfq.predicates {
		p(selector)
	}
	for _, p := range rfq.order {
		p(selector)
	}
	if offset := rfq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rfq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (rfq *RSSFeedQuery) ForUpdate(opts ...sql.LockOption) *RSSFeedQuery {
	if rfq.driver.Dialect() == dialect.Postgres {
		rfq.Unique(false)
	}
	rfq.modifiers = append(rfq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return rfq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (rfq *RSSFeedQuery) ForShare(opts ...sql.LockOption) *RSSFeedQuery {
	if rfq.driver.Dialect() == dialect.Postgres {
		rfq.Unique(false)
	}
	rfq.modifiers = append(rfq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return rfq
}

// RSSFeedGroupBy is the group-by builder for RSSFeed entities.
type RSSFeedGroupBy struct {
	selector
	build *RSSFeedQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rfgb *RSSFeedGroupBy) Aggregate(fns ...AggregateFunc) *RSSFeedGroupBy {
	rfgb.fns = append(rfgb.fns, fns...)
	return rfgb
}

// Scan applies the selector query and scans the result into the given value.
func (rfgb *RSSFeedGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rfgb.build.ctx, "GroupBy")
	if err := rfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RSSFeedQuery, *RSSFeedGroupBy](ctx, rfgb.build, rfgb, rfgb.build.inters, v)
}

func (rfgb *RSSFeedGroupBy) sqlScan(ctx context.Context, root *RSSFeedQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rfgb.fns))
	for _, fn := range rfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rfgb.flds)+len(rfgb.fns))
		for _, f := range *rfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RSSFeedSelect is the builder for selecting fields of RSSFeed entities.
type RSSFeedSelect struct {
	*RSSFeedQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rfs *RSSFeedSelect) Aggregate(fns ...AggregateFunc) *RSSFeedSelect {
	rfs.fns = append(rfs.fns, fns...)
	return rfs
}

// Scan applies the selector query and scans the result into the given value.
func (rfs *RSSFeedSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rfs.ctx, "Select")
	if err := rfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RSSFeedQuery, *RSSFeedSelect](ctx, rfs.RSSFeedQuery, rfs, rfs.inters, v)
}

func (rfs *RSSFeedSelect) sqlScan(ctx context.Context, root *RSSFeedQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rfs.fns))
	for _, fn := range rfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
