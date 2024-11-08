// Code generated by ent, DO NOT EDIT.

package build

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/organization"
)

// Organization is the model entity for the Organization schema.
type Organization struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UniqueName holds the value of the "unique_name" field.
	UniqueName string `json:"unique_name,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// WebsiteURL holds the value of the "website_url" field.
	WebsiteURL string `json:"website_url,omitempty"`
	// LogoImageURL holds the value of the "logo_image_url" field.
	LogoImageURL string `json:"logo_image_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrganizationQuery when eager-loading is set.
	Edges        OrganizationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrganizationEdges holds the relations/edges for other nodes in the graph.
type OrganizationEdges struct {
	// Feeds holds the value of the feeds edge.
	Feeds []*RSSFeed `json:"feeds,omitempty"`
	// Author holds the value of the author edge.
	Author []*RSSAuthor `json:"author,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FeedsOrErr returns the Feeds value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) FeedsOrErr() ([]*RSSFeed, error) {
	if e.loadedTypes[0] {
		return e.Feeds, nil
	}
	return nil, &NotLoadedError{edge: "feeds"}
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) AuthorOrErr() ([]*RSSAuthor, error) {
	if e.loadedTypes[1] {
		return e.Author, nil
	}
	return nil, &NotLoadedError{edge: "author"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Organization) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case organization.FieldUniqueName, organization.FieldName, organization.FieldDescription, organization.FieldWebsiteURL, organization.FieldLogoImageURL:
			values[i] = new(sql.NullString)
		case organization.FieldCreateTime:
			values[i] = new(sql.NullTime)
		case organization.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Organization fields.
func (o *Organization) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case organization.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				o.ID = *value
			}
		case organization.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				o.CreateTime = value.Time
			}
		case organization.FieldUniqueName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field unique_name", values[i])
			} else if value.Valid {
				o.UniqueName = value.String
			}
		case organization.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				o.Name = value.String
			}
		case organization.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				o.Description = value.String
			}
		case organization.FieldWebsiteURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field website_url", values[i])
			} else if value.Valid {
				o.WebsiteURL = value.String
			}
		case organization.FieldLogoImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo_image_url", values[i])
			} else if value.Valid {
				o.LogoImageURL = value.String
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Organization.
// This includes values selected through modifiers, order, etc.
func (o *Organization) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryFeeds queries the "feeds" edge of the Organization entity.
func (o *Organization) QueryFeeds() *RSSFeedQuery {
	return NewOrganizationClient(o.config).QueryFeeds(o)
}

// QueryAuthor queries the "author" edge of the Organization entity.
func (o *Organization) QueryAuthor() *RSSAuthorQuery {
	return NewOrganizationClient(o.config).QueryAuthor(o)
}

// Update returns a builder for updating this Organization.
// Note that you need to call Organization.Unwrap() before calling this method if this Organization
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Organization) Update() *OrganizationUpdateOne {
	return NewOrganizationClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Organization entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Organization) Unwrap() *Organization {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("build: Organization is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Organization) String() string {
	var builder strings.Builder
	builder.WriteString("Organization(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("create_time=")
	builder.WriteString(o.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("unique_name=")
	builder.WriteString(o.UniqueName)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(o.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(o.Description)
	builder.WriteString(", ")
	builder.WriteString("website_url=")
	builder.WriteString(o.WebsiteURL)
	builder.WriteString(", ")
	builder.WriteString("logo_image_url=")
	builder.WriteString(o.LogoImageURL)
	builder.WriteByte(')')
	return builder.String()
}

// Organizations is a parsable slice of Organization.
type Organizations []*Organization
