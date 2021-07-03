// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TourProductsColumns holds the columns for the "tour_products" table.
	TourProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "price", Type: field.TypeInt},
		{Name: "for_sale", Type: field.TypeBool, Default: true},
		{Name: "user_products", Type: field.TypeString, Nullable: true},
	}
	// TourProductsTable holds the schema information for the "tour_products" table.
	TourProductsTable = &schema.Table{
		Name:       "tour_products",
		Columns:    TourProductsColumns,
		PrimaryKey: []*schema.Column{TourProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tour_products_users_products",
				Columns:    []*schema.Column{TourProductsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "is_activated", Type: field.TypeBool, Default: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TourProductsTable,
		UsersTable,
	}
)

func init() {
	TourProductsTable.ForeignKeys[0].RefTable = UsersTable
}
