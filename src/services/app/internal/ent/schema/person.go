package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type Person struct {
	ent.Schema
}

func (Person) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
		mixin.UpdateTime{},
	}
}

func (Person) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("email").Unique(),
	}
}

func (Person) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New).Unique().StorageKey("uuid"),
		field.String("email").Unique(),
		field.String("name"),
	}
}

func (Person) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", RSSAuthor.Type).Ref("person"),
	}
}
