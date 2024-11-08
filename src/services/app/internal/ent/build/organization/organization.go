// Code generated by ent, DO NOT EDIT.

package organization

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the organization type in the database.
	Label = "organization"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "uuid"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUniqueName holds the string denoting the unique_name field in the database.
	FieldUniqueName = "unique_name"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldWebsiteURL holds the string denoting the website_url field in the database.
	FieldWebsiteURL = "website_url"
	// FieldLogoImageURL holds the string denoting the logo_image_url field in the database.
	FieldLogoImageURL = "logo_image_url"
	// EdgeFeeds holds the string denoting the feeds edge name in mutations.
	EdgeFeeds = "feeds"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// Table holds the table name of the organization in the database.
	Table = "organizations"
	// FeedsTable is the table that holds the feeds relation/edge.
	FeedsTable = "rss_feeds"
	// FeedsInverseTable is the table name for the RSSFeed entity.
	// It exists in this package in order to avoid circular dependency with the "rssfeed" package.
	FeedsInverseTable = "rss_feeds"
	// FeedsColumn is the table column denoting the feeds relation/edge.
	FeedsColumn = "organization_feeds"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "rss_authors"
	// AuthorInverseTable is the table name for the RSSAuthor entity.
	// It exists in this package in order to avoid circular dependency with the "rssauthor" package.
	AuthorInverseTable = "rss_authors"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "organization_author"
)

// Columns holds all SQL columns for organization fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUniqueName,
	FieldName,
	FieldDescription,
	FieldWebsiteURL,
	FieldLogoImageURL,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// UniqueNameValidator is a validator for the "unique_name" field. It is called by the builders before save.
	UniqueNameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Organization queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUniqueName orders the results by the unique_name field.
func ByUniqueName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUniqueName, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByWebsiteURL orders the results by the website_url field.
func ByWebsiteURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWebsiteURL, opts...).ToFunc()
}

// ByLogoImageURL orders the results by the logo_image_url field.
func ByLogoImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogoImageURL, opts...).ToFunc()
}

// ByFeedsCount orders the results by feeds count.
func ByFeedsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFeedsStep(), opts...)
	}
}

// ByFeeds orders the results by feeds terms.
func ByFeeds(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFeedsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAuthorCount orders the results by author count.
func ByAuthorCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAuthorStep(), opts...)
	}
}

// ByAuthor orders the results by author terms.
func ByAuthor(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newFeedsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FeedsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FeedsTable, FeedsColumn),
	)
}
func newAuthorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AuthorTable, AuthorColumn),
	)
}
