package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Role struct {
	ent.Schema
}

func (Role) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uuid").Unique(),
		index.Fields("key").Unique(),
	}
}

func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.New()).Default(uuid.New).Unique(),
		field.String("key").Unique(),
		field.Strings("permissions"),
	}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("roles"),
	}
}
