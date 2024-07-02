package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Checkpoint struct {
	ent.Schema
}

func (Checkpoint) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_uuid"),
	}
}

func (Checkpoint) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_time").Default(time.Now),
		field.String("user_uuid"),
		field.String("key"),
	}
}

func (Checkpoint) Edges() []ent.Edge {
	return nil
}
