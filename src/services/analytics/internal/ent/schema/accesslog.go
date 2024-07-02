package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type AccessLog struct {
	ent.Schema
}

func (AccessLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("ip"),
		index.Fields("uri"),
	}
}

func (AccessLog) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_time").Default(time.Now).Immutable(),
		field.String("ip"),
		field.String("uri"),

		field.String("forwarded_for"),
		field.String("forwarded_proto"),
		field.String("forwarded_host"),
		field.String("forwarded_port"),
		field.String("forwarded_server"),
		field.String("request_id"),
		field.String("user_agent"),
	}
}
