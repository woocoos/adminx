package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/woocoos/adminx/service/snowflake"
)

// SnowFlakeID 是采用雪花算法生成的ID.
type SnowFlakeID struct {
	// ID is the unique identifier of the user in the database.
	mixin.Schema
}

func (id SnowFlakeID) Fields() []ent.Field {
	Incremental := false
	return []ent.Field{
		field.Int("id").SchemaType(id.SchemaType()).Annotations(entsql.Annotation{
			Incremental: &Incremental,
		}).DefaultFunc(func() int {
			return int(snowflake.New().Int64())
		}),
	}
}

func (SnowFlakeID) SchemaType() map[string]string {
	return map[string]string{
		"mysql": "bigint",
	}
}

// IntID 是采用常规int型自增字型.
type IntID struct {
	// ID is the unique identifier of the user in the database.
	mixin.Schema
}

func (id IntID) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").SchemaType(id.SchemaType()),
	}
}

func (IntID) SchemaType() map[string]string {
	return map[string]string{
		"mysql": "int",
	}
}
