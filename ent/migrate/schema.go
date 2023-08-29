// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// RemotesColumns holds the columns for the "remotes" table.
	RemotesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "url", Type: field.TypeString},
	}
	// RemotesTable holds the schema information for the "remotes" table.
	RemotesTable = &schema.Table{
		Name:       "remotes",
		Columns:    RemotesColumns,
		PrimaryKey: []*schema.Column{RemotesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		RemotesTable,
	}
)

func init() {
}