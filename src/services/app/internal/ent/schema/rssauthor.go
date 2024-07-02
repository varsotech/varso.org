package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type RSSAuthor struct {
	ent.Schema
}

func (RSSAuthor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}

func (RSSAuthor) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("name"),
		index.Edges("organization").Fields("name").Unique(),
	}
}

func (RSSAuthor) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New).Unique().StorageKey("uuid"),
		field.String("email").Optional(),
		field.String("name"),
	}
}

func (RSSAuthor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("person", Person.Type).Unique(),
		edge.To("organization", Organization.Type).Unique(),
		edge.From("newsitem", NewsItem.Type).Ref("authors"),
	}
}
