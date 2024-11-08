// Code generated by ent, DO NOT EDIT.

package build

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/newsitem"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/organization"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/predicate"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/rssfeed"
)

// RSSFeedUpdate is the builder for updating RSSFeed entities.
type RSSFeedUpdate struct {
	config
	hooks    []Hook
	mutation *RSSFeedMutation
}

// Where appends a list predicates to the RSSFeedUpdate builder.
func (rfu *RSSFeedUpdate) Where(ps ...predicate.RSSFeed) *RSSFeedUpdate {
	rfu.mutation.Where(ps...)
	return rfu
}

// SetUpdateTime sets the "update_time" field.
func (rfu *RSSFeedUpdate) SetUpdateTime(t time.Time) *RSSFeedUpdate {
	rfu.mutation.SetUpdateTime(t)
	return rfu
}

// SetRssFeedURL sets the "rss_feed_url" field.
func (rfu *RSSFeedUpdate) SetRssFeedURL(s string) *RSSFeedUpdate {
	rfu.mutation.SetRssFeedURL(s)
	return rfu
}

// SetNillableRssFeedURL sets the "rss_feed_url" field if the given value is not nil.
func (rfu *RSSFeedUpdate) SetNillableRssFeedURL(s *string) *RSSFeedUpdate {
	if s != nil {
		rfu.SetRssFeedURL(*s)
	}
	return rfu
}

// SetContentWhitelistRegex sets the "content_whitelist_regex" field.
func (rfu *RSSFeedUpdate) SetContentWhitelistRegex(s string) *RSSFeedUpdate {
	rfu.mutation.SetContentWhitelistRegex(s)
	return rfu
}

// SetNillableContentWhitelistRegex sets the "content_whitelist_regex" field if the given value is not nil.
func (rfu *RSSFeedUpdate) SetNillableContentWhitelistRegex(s *string) *RSSFeedUpdate {
	if s != nil {
		rfu.SetContentWhitelistRegex(*s)
	}
	return rfu
}

// ClearContentWhitelistRegex clears the value of the "content_whitelist_regex" field.
func (rfu *RSSFeedUpdate) ClearContentWhitelistRegex() *RSSFeedUpdate {
	rfu.mutation.ClearContentWhitelistRegex()
	return rfu
}

// SetHTMLPaywallRegex sets the "html_paywall_regex" field.
func (rfu *RSSFeedUpdate) SetHTMLPaywallRegex(s string) *RSSFeedUpdate {
	rfu.mutation.SetHTMLPaywallRegex(s)
	return rfu
}

// SetNillableHTMLPaywallRegex sets the "html_paywall_regex" field if the given value is not nil.
func (rfu *RSSFeedUpdate) SetNillableHTMLPaywallRegex(s *string) *RSSFeedUpdate {
	if s != nil {
		rfu.SetHTMLPaywallRegex(*s)
	}
	return rfu
}

// ClearHTMLPaywallRegex clears the value of the "html_paywall_regex" field.
func (rfu *RSSFeedUpdate) ClearHTMLPaywallRegex() *RSSFeedUpdate {
	rfu.mutation.ClearHTMLPaywallRegex()
	return rfu
}

// SetTitleTrimRight sets the "title_trim_right" field.
func (rfu *RSSFeedUpdate) SetTitleTrimRight(s string) *RSSFeedUpdate {
	rfu.mutation.SetTitleTrimRight(s)
	return rfu
}

// SetNillableTitleTrimRight sets the "title_trim_right" field if the given value is not nil.
func (rfu *RSSFeedUpdate) SetNillableTitleTrimRight(s *string) *RSSFeedUpdate {
	if s != nil {
		rfu.SetTitleTrimRight(*s)
	}
	return rfu
}

// ClearTitleTrimRight clears the value of the "title_trim_right" field.
func (rfu *RSSFeedUpdate) ClearTitleTrimRight() *RSSFeedUpdate {
	rfu.mutation.ClearTitleTrimRight()
	return rfu
}

// SetRssFeedRank sets the "rss_feed_rank" field.
func (rfu *RSSFeedUpdate) SetRssFeedRank(f float32) *RSSFeedUpdate {
	rfu.mutation.ResetRssFeedRank()
	rfu.mutation.SetRssFeedRank(f)
	return rfu
}

// SetNillableRssFeedRank sets the "rss_feed_rank" field if the given value is not nil.
func (rfu *RSSFeedUpdate) SetNillableRssFeedRank(f *float32) *RSSFeedUpdate {
	if f != nil {
		rfu.SetRssFeedRank(*f)
	}
	return rfu
}

// AddRssFeedRank adds f to the "rss_feed_rank" field.
func (rfu *RSSFeedUpdate) AddRssFeedRank(f float32) *RSSFeedUpdate {
	rfu.mutation.AddRssFeedRank(f)
	return rfu
}

// ClearRssFeedRank clears the value of the "rss_feed_rank" field.
func (rfu *RSSFeedUpdate) ClearRssFeedRank() *RSSFeedUpdate {
	rfu.mutation.ClearRssFeedRank()
	return rfu
}

// SetMaxFetchIntervalMin sets the "max_fetch_interval_min" field.
func (rfu *RSSFeedUpdate) SetMaxFetchIntervalMin(i int64) *RSSFeedUpdate {
	rfu.mutation.ResetMaxFetchIntervalMin()
	rfu.mutation.SetMaxFetchIntervalMin(i)
	return rfu
}

// SetNillableMaxFetchIntervalMin sets the "max_fetch_interval_min" field if the given value is not nil.
func (rfu *RSSFeedUpdate) SetNillableMaxFetchIntervalMin(i *int64) *RSSFeedUpdate {
	if i != nil {
		rfu.SetMaxFetchIntervalMin(*i)
	}
	return rfu
}

// AddMaxFetchIntervalMin adds i to the "max_fetch_interval_min" field.
func (rfu *RSSFeedUpdate) AddMaxFetchIntervalMin(i int64) *RSSFeedUpdate {
	rfu.mutation.AddMaxFetchIntervalMin(i)
	return rfu
}

// ClearMaxFetchIntervalMin clears the value of the "max_fetch_interval_min" field.
func (rfu *RSSFeedUpdate) ClearMaxFetchIntervalMin() *RSSFeedUpdate {
	rfu.mutation.ClearMaxFetchIntervalMin()
	return rfu
}

// SetDiscardOgImage sets the "discard_og_image" field.
func (rfu *RSSFeedUpdate) SetDiscardOgImage(b bool) *RSSFeedUpdate {
	rfu.mutation.SetDiscardOgImage(b)
	return rfu
}

// SetNillableDiscardOgImage sets the "discard_og_image" field if the given value is not nil.
func (rfu *RSSFeedUpdate) SetNillableDiscardOgImage(b *bool) *RSSFeedUpdate {
	if b != nil {
		rfu.SetDiscardOgImage(*b)
	}
	return rfu
}

// AddItemIDs adds the "items" edge to the NewsItem entity by IDs.
func (rfu *RSSFeedUpdate) AddItemIDs(ids ...uuid.UUID) *RSSFeedUpdate {
	rfu.mutation.AddItemIDs(ids...)
	return rfu
}

// AddItems adds the "items" edges to the NewsItem entity.
func (rfu *RSSFeedUpdate) AddItems(n ...*NewsItem) *RSSFeedUpdate {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return rfu.AddItemIDs(ids...)
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (rfu *RSSFeedUpdate) SetOrganizationID(id uuid.UUID) *RSSFeedUpdate {
	rfu.mutation.SetOrganizationID(id)
	return rfu
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (rfu *RSSFeedUpdate) SetNillableOrganizationID(id *uuid.UUID) *RSSFeedUpdate {
	if id != nil {
		rfu = rfu.SetOrganizationID(*id)
	}
	return rfu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (rfu *RSSFeedUpdate) SetOrganization(o *Organization) *RSSFeedUpdate {
	return rfu.SetOrganizationID(o.ID)
}

// Mutation returns the RSSFeedMutation object of the builder.
func (rfu *RSSFeedUpdate) Mutation() *RSSFeedMutation {
	return rfu.mutation
}

// ClearItems clears all "items" edges to the NewsItem entity.
func (rfu *RSSFeedUpdate) ClearItems() *RSSFeedUpdate {
	rfu.mutation.ClearItems()
	return rfu
}

// RemoveItemIDs removes the "items" edge to NewsItem entities by IDs.
func (rfu *RSSFeedUpdate) RemoveItemIDs(ids ...uuid.UUID) *RSSFeedUpdate {
	rfu.mutation.RemoveItemIDs(ids...)
	return rfu
}

// RemoveItems removes "items" edges to NewsItem entities.
func (rfu *RSSFeedUpdate) RemoveItems(n ...*NewsItem) *RSSFeedUpdate {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return rfu.RemoveItemIDs(ids...)
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (rfu *RSSFeedUpdate) ClearOrganization() *RSSFeedUpdate {
	rfu.mutation.ClearOrganization()
	return rfu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rfu *RSSFeedUpdate) Save(ctx context.Context) (int, error) {
	rfu.defaults()
	return withHooks(ctx, rfu.sqlSave, rfu.mutation, rfu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rfu *RSSFeedUpdate) SaveX(ctx context.Context) int {
	affected, err := rfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rfu *RSSFeedUpdate) Exec(ctx context.Context) error {
	_, err := rfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rfu *RSSFeedUpdate) ExecX(ctx context.Context) {
	if err := rfu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rfu *RSSFeedUpdate) defaults() {
	if _, ok := rfu.mutation.UpdateTime(); !ok {
		v := rssfeed.UpdateDefaultUpdateTime()
		rfu.mutation.SetUpdateTime(v)
	}
}

func (rfu *RSSFeedUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(rssfeed.Table, rssfeed.Columns, sqlgraph.NewFieldSpec(rssfeed.FieldID, field.TypeUUID))
	if ps := rfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rfu.mutation.UpdateTime(); ok {
		_spec.SetField(rssfeed.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := rfu.mutation.RssFeedURL(); ok {
		_spec.SetField(rssfeed.FieldRssFeedURL, field.TypeString, value)
	}
	if value, ok := rfu.mutation.ContentWhitelistRegex(); ok {
		_spec.SetField(rssfeed.FieldContentWhitelistRegex, field.TypeString, value)
	}
	if rfu.mutation.ContentWhitelistRegexCleared() {
		_spec.ClearField(rssfeed.FieldContentWhitelistRegex, field.TypeString)
	}
	if value, ok := rfu.mutation.HTMLPaywallRegex(); ok {
		_spec.SetField(rssfeed.FieldHTMLPaywallRegex, field.TypeString, value)
	}
	if rfu.mutation.HTMLPaywallRegexCleared() {
		_spec.ClearField(rssfeed.FieldHTMLPaywallRegex, field.TypeString)
	}
	if value, ok := rfu.mutation.TitleTrimRight(); ok {
		_spec.SetField(rssfeed.FieldTitleTrimRight, field.TypeString, value)
	}
	if rfu.mutation.TitleTrimRightCleared() {
		_spec.ClearField(rssfeed.FieldTitleTrimRight, field.TypeString)
	}
	if value, ok := rfu.mutation.RssFeedRank(); ok {
		_spec.SetField(rssfeed.FieldRssFeedRank, field.TypeFloat32, value)
	}
	if value, ok := rfu.mutation.AddedRssFeedRank(); ok {
		_spec.AddField(rssfeed.FieldRssFeedRank, field.TypeFloat32, value)
	}
	if rfu.mutation.RssFeedRankCleared() {
		_spec.ClearField(rssfeed.FieldRssFeedRank, field.TypeFloat32)
	}
	if value, ok := rfu.mutation.MaxFetchIntervalMin(); ok {
		_spec.SetField(rssfeed.FieldMaxFetchIntervalMin, field.TypeInt64, value)
	}
	if value, ok := rfu.mutation.AddedMaxFetchIntervalMin(); ok {
		_spec.AddField(rssfeed.FieldMaxFetchIntervalMin, field.TypeInt64, value)
	}
	if rfu.mutation.MaxFetchIntervalMinCleared() {
		_spec.ClearField(rssfeed.FieldMaxFetchIntervalMin, field.TypeInt64)
	}
	if value, ok := rfu.mutation.DiscardOgImage(); ok {
		_spec.SetField(rssfeed.FieldDiscardOgImage, field.TypeBool, value)
	}
	if rfu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   rssfeed.ItemsTable,
			Columns: []string{rssfeed.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(newsitem.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rfu.mutation.RemovedItemsIDs(); len(nodes) > 0 && !rfu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   rssfeed.ItemsTable,
			Columns: []string{rssfeed.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(newsitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rfu.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   rssfeed.ItemsTable,
			Columns: []string{rssfeed.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(newsitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rfu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rssfeed.OrganizationTable,
			Columns: []string{rssfeed.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rfu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rssfeed.OrganizationTable,
			Columns: []string{rssfeed.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rssfeed.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rfu.mutation.done = true
	return n, nil
}

// RSSFeedUpdateOne is the builder for updating a single RSSFeed entity.
type RSSFeedUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RSSFeedMutation
}

// SetUpdateTime sets the "update_time" field.
func (rfuo *RSSFeedUpdateOne) SetUpdateTime(t time.Time) *RSSFeedUpdateOne {
	rfuo.mutation.SetUpdateTime(t)
	return rfuo
}

// SetRssFeedURL sets the "rss_feed_url" field.
func (rfuo *RSSFeedUpdateOne) SetRssFeedURL(s string) *RSSFeedUpdateOne {
	rfuo.mutation.SetRssFeedURL(s)
	return rfuo
}

// SetNillableRssFeedURL sets the "rss_feed_url" field if the given value is not nil.
func (rfuo *RSSFeedUpdateOne) SetNillableRssFeedURL(s *string) *RSSFeedUpdateOne {
	if s != nil {
		rfuo.SetRssFeedURL(*s)
	}
	return rfuo
}

// SetContentWhitelistRegex sets the "content_whitelist_regex" field.
func (rfuo *RSSFeedUpdateOne) SetContentWhitelistRegex(s string) *RSSFeedUpdateOne {
	rfuo.mutation.SetContentWhitelistRegex(s)
	return rfuo
}

// SetNillableContentWhitelistRegex sets the "content_whitelist_regex" field if the given value is not nil.
func (rfuo *RSSFeedUpdateOne) SetNillableContentWhitelistRegex(s *string) *RSSFeedUpdateOne {
	if s != nil {
		rfuo.SetContentWhitelistRegex(*s)
	}
	return rfuo
}

// ClearContentWhitelistRegex clears the value of the "content_whitelist_regex" field.
func (rfuo *RSSFeedUpdateOne) ClearContentWhitelistRegex() *RSSFeedUpdateOne {
	rfuo.mutation.ClearContentWhitelistRegex()
	return rfuo
}

// SetHTMLPaywallRegex sets the "html_paywall_regex" field.
func (rfuo *RSSFeedUpdateOne) SetHTMLPaywallRegex(s string) *RSSFeedUpdateOne {
	rfuo.mutation.SetHTMLPaywallRegex(s)
	return rfuo
}

// SetNillableHTMLPaywallRegex sets the "html_paywall_regex" field if the given value is not nil.
func (rfuo *RSSFeedUpdateOne) SetNillableHTMLPaywallRegex(s *string) *RSSFeedUpdateOne {
	if s != nil {
		rfuo.SetHTMLPaywallRegex(*s)
	}
	return rfuo
}

// ClearHTMLPaywallRegex clears the value of the "html_paywall_regex" field.
func (rfuo *RSSFeedUpdateOne) ClearHTMLPaywallRegex() *RSSFeedUpdateOne {
	rfuo.mutation.ClearHTMLPaywallRegex()
	return rfuo
}

// SetTitleTrimRight sets the "title_trim_right" field.
func (rfuo *RSSFeedUpdateOne) SetTitleTrimRight(s string) *RSSFeedUpdateOne {
	rfuo.mutation.SetTitleTrimRight(s)
	return rfuo
}

// SetNillableTitleTrimRight sets the "title_trim_right" field if the given value is not nil.
func (rfuo *RSSFeedUpdateOne) SetNillableTitleTrimRight(s *string) *RSSFeedUpdateOne {
	if s != nil {
		rfuo.SetTitleTrimRight(*s)
	}
	return rfuo
}

// ClearTitleTrimRight clears the value of the "title_trim_right" field.
func (rfuo *RSSFeedUpdateOne) ClearTitleTrimRight() *RSSFeedUpdateOne {
	rfuo.mutation.ClearTitleTrimRight()
	return rfuo
}

// SetRssFeedRank sets the "rss_feed_rank" field.
func (rfuo *RSSFeedUpdateOne) SetRssFeedRank(f float32) *RSSFeedUpdateOne {
	rfuo.mutation.ResetRssFeedRank()
	rfuo.mutation.SetRssFeedRank(f)
	return rfuo
}

// SetNillableRssFeedRank sets the "rss_feed_rank" field if the given value is not nil.
func (rfuo *RSSFeedUpdateOne) SetNillableRssFeedRank(f *float32) *RSSFeedUpdateOne {
	if f != nil {
		rfuo.SetRssFeedRank(*f)
	}
	return rfuo
}

// AddRssFeedRank adds f to the "rss_feed_rank" field.
func (rfuo *RSSFeedUpdateOne) AddRssFeedRank(f float32) *RSSFeedUpdateOne {
	rfuo.mutation.AddRssFeedRank(f)
	return rfuo
}

// ClearRssFeedRank clears the value of the "rss_feed_rank" field.
func (rfuo *RSSFeedUpdateOne) ClearRssFeedRank() *RSSFeedUpdateOne {
	rfuo.mutation.ClearRssFeedRank()
	return rfuo
}

// SetMaxFetchIntervalMin sets the "max_fetch_interval_min" field.
func (rfuo *RSSFeedUpdateOne) SetMaxFetchIntervalMin(i int64) *RSSFeedUpdateOne {
	rfuo.mutation.ResetMaxFetchIntervalMin()
	rfuo.mutation.SetMaxFetchIntervalMin(i)
	return rfuo
}

// SetNillableMaxFetchIntervalMin sets the "max_fetch_interval_min" field if the given value is not nil.
func (rfuo *RSSFeedUpdateOne) SetNillableMaxFetchIntervalMin(i *int64) *RSSFeedUpdateOne {
	if i != nil {
		rfuo.SetMaxFetchIntervalMin(*i)
	}
	return rfuo
}

// AddMaxFetchIntervalMin adds i to the "max_fetch_interval_min" field.
func (rfuo *RSSFeedUpdateOne) AddMaxFetchIntervalMin(i int64) *RSSFeedUpdateOne {
	rfuo.mutation.AddMaxFetchIntervalMin(i)
	return rfuo
}

// ClearMaxFetchIntervalMin clears the value of the "max_fetch_interval_min" field.
func (rfuo *RSSFeedUpdateOne) ClearMaxFetchIntervalMin() *RSSFeedUpdateOne {
	rfuo.mutation.ClearMaxFetchIntervalMin()
	return rfuo
}

// SetDiscardOgImage sets the "discard_og_image" field.
func (rfuo *RSSFeedUpdateOne) SetDiscardOgImage(b bool) *RSSFeedUpdateOne {
	rfuo.mutation.SetDiscardOgImage(b)
	return rfuo
}

// SetNillableDiscardOgImage sets the "discard_og_image" field if the given value is not nil.
func (rfuo *RSSFeedUpdateOne) SetNillableDiscardOgImage(b *bool) *RSSFeedUpdateOne {
	if b != nil {
		rfuo.SetDiscardOgImage(*b)
	}
	return rfuo
}

// AddItemIDs adds the "items" edge to the NewsItem entity by IDs.
func (rfuo *RSSFeedUpdateOne) AddItemIDs(ids ...uuid.UUID) *RSSFeedUpdateOne {
	rfuo.mutation.AddItemIDs(ids...)
	return rfuo
}

// AddItems adds the "items" edges to the NewsItem entity.
func (rfuo *RSSFeedUpdateOne) AddItems(n ...*NewsItem) *RSSFeedUpdateOne {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return rfuo.AddItemIDs(ids...)
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (rfuo *RSSFeedUpdateOne) SetOrganizationID(id uuid.UUID) *RSSFeedUpdateOne {
	rfuo.mutation.SetOrganizationID(id)
	return rfuo
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (rfuo *RSSFeedUpdateOne) SetNillableOrganizationID(id *uuid.UUID) *RSSFeedUpdateOne {
	if id != nil {
		rfuo = rfuo.SetOrganizationID(*id)
	}
	return rfuo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (rfuo *RSSFeedUpdateOne) SetOrganization(o *Organization) *RSSFeedUpdateOne {
	return rfuo.SetOrganizationID(o.ID)
}

// Mutation returns the RSSFeedMutation object of the builder.
func (rfuo *RSSFeedUpdateOne) Mutation() *RSSFeedMutation {
	return rfuo.mutation
}

// ClearItems clears all "items" edges to the NewsItem entity.
func (rfuo *RSSFeedUpdateOne) ClearItems() *RSSFeedUpdateOne {
	rfuo.mutation.ClearItems()
	return rfuo
}

// RemoveItemIDs removes the "items" edge to NewsItem entities by IDs.
func (rfuo *RSSFeedUpdateOne) RemoveItemIDs(ids ...uuid.UUID) *RSSFeedUpdateOne {
	rfuo.mutation.RemoveItemIDs(ids...)
	return rfuo
}

// RemoveItems removes "items" edges to NewsItem entities.
func (rfuo *RSSFeedUpdateOne) RemoveItems(n ...*NewsItem) *RSSFeedUpdateOne {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return rfuo.RemoveItemIDs(ids...)
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (rfuo *RSSFeedUpdateOne) ClearOrganization() *RSSFeedUpdateOne {
	rfuo.mutation.ClearOrganization()
	return rfuo
}

// Where appends a list predicates to the RSSFeedUpdate builder.
func (rfuo *RSSFeedUpdateOne) Where(ps ...predicate.RSSFeed) *RSSFeedUpdateOne {
	rfuo.mutation.Where(ps...)
	return rfuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rfuo *RSSFeedUpdateOne) Select(field string, fields ...string) *RSSFeedUpdateOne {
	rfuo.fields = append([]string{field}, fields...)
	return rfuo
}

// Save executes the query and returns the updated RSSFeed entity.
func (rfuo *RSSFeedUpdateOne) Save(ctx context.Context) (*RSSFeed, error) {
	rfuo.defaults()
	return withHooks(ctx, rfuo.sqlSave, rfuo.mutation, rfuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rfuo *RSSFeedUpdateOne) SaveX(ctx context.Context) *RSSFeed {
	node, err := rfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rfuo *RSSFeedUpdateOne) Exec(ctx context.Context) error {
	_, err := rfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rfuo *RSSFeedUpdateOne) ExecX(ctx context.Context) {
	if err := rfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rfuo *RSSFeedUpdateOne) defaults() {
	if _, ok := rfuo.mutation.UpdateTime(); !ok {
		v := rssfeed.UpdateDefaultUpdateTime()
		rfuo.mutation.SetUpdateTime(v)
	}
}

func (rfuo *RSSFeedUpdateOne) sqlSave(ctx context.Context) (_node *RSSFeed, err error) {
	_spec := sqlgraph.NewUpdateSpec(rssfeed.Table, rssfeed.Columns, sqlgraph.NewFieldSpec(rssfeed.FieldID, field.TypeUUID))
	id, ok := rfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`build: missing "RSSFeed.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rssfeed.FieldID)
		for _, f := range fields {
			if !rssfeed.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("build: invalid field %q for query", f)}
			}
			if f != rssfeed.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rfuo.mutation.UpdateTime(); ok {
		_spec.SetField(rssfeed.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := rfuo.mutation.RssFeedURL(); ok {
		_spec.SetField(rssfeed.FieldRssFeedURL, field.TypeString, value)
	}
	if value, ok := rfuo.mutation.ContentWhitelistRegex(); ok {
		_spec.SetField(rssfeed.FieldContentWhitelistRegex, field.TypeString, value)
	}
	if rfuo.mutation.ContentWhitelistRegexCleared() {
		_spec.ClearField(rssfeed.FieldContentWhitelistRegex, field.TypeString)
	}
	if value, ok := rfuo.mutation.HTMLPaywallRegex(); ok {
		_spec.SetField(rssfeed.FieldHTMLPaywallRegex, field.TypeString, value)
	}
	if rfuo.mutation.HTMLPaywallRegexCleared() {
		_spec.ClearField(rssfeed.FieldHTMLPaywallRegex, field.TypeString)
	}
	if value, ok := rfuo.mutation.TitleTrimRight(); ok {
		_spec.SetField(rssfeed.FieldTitleTrimRight, field.TypeString, value)
	}
	if rfuo.mutation.TitleTrimRightCleared() {
		_spec.ClearField(rssfeed.FieldTitleTrimRight, field.TypeString)
	}
	if value, ok := rfuo.mutation.RssFeedRank(); ok {
		_spec.SetField(rssfeed.FieldRssFeedRank, field.TypeFloat32, value)
	}
	if value, ok := rfuo.mutation.AddedRssFeedRank(); ok {
		_spec.AddField(rssfeed.FieldRssFeedRank, field.TypeFloat32, value)
	}
	if rfuo.mutation.RssFeedRankCleared() {
		_spec.ClearField(rssfeed.FieldRssFeedRank, field.TypeFloat32)
	}
	if value, ok := rfuo.mutation.MaxFetchIntervalMin(); ok {
		_spec.SetField(rssfeed.FieldMaxFetchIntervalMin, field.TypeInt64, value)
	}
	if value, ok := rfuo.mutation.AddedMaxFetchIntervalMin(); ok {
		_spec.AddField(rssfeed.FieldMaxFetchIntervalMin, field.TypeInt64, value)
	}
	if rfuo.mutation.MaxFetchIntervalMinCleared() {
		_spec.ClearField(rssfeed.FieldMaxFetchIntervalMin, field.TypeInt64)
	}
	if value, ok := rfuo.mutation.DiscardOgImage(); ok {
		_spec.SetField(rssfeed.FieldDiscardOgImage, field.TypeBool, value)
	}
	if rfuo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   rssfeed.ItemsTable,
			Columns: []string{rssfeed.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(newsitem.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rfuo.mutation.RemovedItemsIDs(); len(nodes) > 0 && !rfuo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   rssfeed.ItemsTable,
			Columns: []string{rssfeed.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(newsitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rfuo.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   rssfeed.ItemsTable,
			Columns: []string{rssfeed.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(newsitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rfuo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rssfeed.OrganizationTable,
			Columns: []string{rssfeed.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rfuo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rssfeed.OrganizationTable,
			Columns: []string{rssfeed.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RSSFeed{config: rfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rssfeed.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rfuo.mutation.done = true
	return _node, nil
}
