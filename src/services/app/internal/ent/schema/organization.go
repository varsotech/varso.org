package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type Organization struct {
	ent.Schema
}

func (Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}

func (Organization) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("unique_name").Unique(),
	}
}

func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New).Unique().StorageKey("uuid"),
		field.String("unique_name").Match(regexp.MustCompile("[a-zA-Z_]+$")).MinLen(3),
		field.String("name"),
		field.String("description"),
		field.String("website_url"),
		field.String("logo_image_url").Optional(),
	}
}

func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("feeds", RSSFeed.Type),
		edge.To("author", RSSAuthor.Type),
	}
}
