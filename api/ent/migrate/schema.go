// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// LinksColumns holds the columns for the "links" table.
	LinksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "image_url", Type: field.TypeString, Nullable: true},
		{Name: "url", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_links", Type: field.TypeUUID, Nullable: true},
	}
	// LinksTable holds the schema information for the "links" table.
	LinksTable = &schema.Table{
		Name:       "links",
		Columns:    LinksColumns,
		PrimaryKey: []*schema.Column{LinksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "links_users_links",
				Columns:    []*schema.Column{LinksColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "image_url", Type: field.TypeString, Nullable: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"ADMIN", "USER"}, Default: "USER"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserBookmarksColumns holds the columns for the "user_bookmarks" table.
	UserBookmarksColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "link_id", Type: field.TypeUUID},
	}
	// UserBookmarksTable holds the schema information for the "user_bookmarks" table.
	UserBookmarksTable = &schema.Table{
		Name:       "user_bookmarks",
		Columns:    UserBookmarksColumns,
		PrimaryKey: []*schema.Column{UserBookmarksColumns[0], UserBookmarksColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_bookmarks_user_id",
				Columns:    []*schema.Column{UserBookmarksColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_bookmarks_link_id",
				Columns:    []*schema.Column{UserBookmarksColumns[1]},
				RefColumns: []*schema.Column{LinksColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LinksTable,
		UsersTable,
		UserBookmarksTable,
	}
)

func init() {
	LinksTable.ForeignKeys[0].RefTable = UsersTable
	UserBookmarksTable.ForeignKeys[0].RefTable = UsersTable
	UserBookmarksTable.ForeignKeys[1].RefTable = LinksTable
}
