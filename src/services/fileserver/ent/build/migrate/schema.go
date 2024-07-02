// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CheckpointsColumns holds the columns for the "checkpoints" table.
	CheckpointsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "user_uuid", Type: field.TypeString},
		{Name: "key", Type: field.TypeString},
	}
	// CheckpointsTable holds the schema information for the "checkpoints" table.
	CheckpointsTable = &schema.Table{
		Name:       "checkpoints",
		Columns:    CheckpointsColumns,
		PrimaryKey: []*schema.Column{CheckpointsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "checkpoint_user_uuid",
				Unique:  false,
				Columns: []*schema.Column{CheckpointsColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CheckpointsTable,
	}
)

func init() {
}
