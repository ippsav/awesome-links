// Code generated by entc, DO NOT EDIT.

package link

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the link type in the database.
	Label = "link"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the link in the database.
	Table = "links"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "links"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_links"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "user_bookmarks"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
)

// Columns holds all SQL columns for link fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescription,
	FieldImageURL,
	FieldURL,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "links"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_links",
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "link_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
