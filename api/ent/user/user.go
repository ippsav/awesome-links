// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeLinks holds the string denoting the links edge name in mutations.
	EdgeLinks = "links"
	// EdgeBookmarks holds the string denoting the bookmarks edge name in mutations.
	EdgeBookmarks = "bookmarks"
	// Table holds the table name of the user in the database.
	Table = "users"
	// LinksTable is the table that holds the links relation/edge.
	LinksTable = "links"
	// LinksInverseTable is the table name for the Link entity.
	// It exists in this package in order to avoid circular dependency with the "link" package.
	LinksInverseTable = "links"
	// LinksColumn is the table column denoting the links relation/edge.
	LinksColumn = "user_links"
	// BookmarksTable is the table that holds the bookmarks relation/edge. The primary key declared below.
	BookmarksTable = "user_bookmarks"
	// BookmarksInverseTable is the table name for the Link entity.
	// It exists in this package in order to avoid circular dependency with the "link" package.
	BookmarksInverseTable = "links"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldPassword,
	FieldImageURL,
	FieldRole,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// BookmarksPrimaryKey and BookmarksColumn2 are the table columns denoting the
	// primary key for the bookmarks relation (M2M).
	BookmarksPrimaryKey = []string{"user_id", "link_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Role defines the type for the "role" enum field.
type Role string

// RoleUSER is the default value of the Role enum.
const DefaultRole = RoleUSER

// Role values.
const (
	RoleADMIN Role = "ADMIN"
	RoleUSER  Role = "USER"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleADMIN, RoleUSER:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for role field: %q", r)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (r Role) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(r.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (r *Role) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*r = Role(str)
	if err := RoleValidator(*r); err != nil {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}
