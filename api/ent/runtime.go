// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/ippsav/awesome-links/api/ent/link"
	"github.com/ippsav/awesome-links/api/ent/schema"
	"github.com/ippsav/awesome-links/api/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	linkFields := schema.Link{}.Fields()
	_ = linkFields
	// linkDescTitle is the schema descriptor for title field.
	linkDescTitle := linkFields[1].Descriptor()
	// link.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	link.TitleValidator = linkDescTitle.Validators[0].(func(string) error)
	// linkDescDescription is the schema descriptor for description field.
	linkDescDescription := linkFields[2].Descriptor()
	// link.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	link.DescriptionValidator = linkDescDescription.Validators[0].(func(string) error)
	// linkDescURL is the schema descriptor for url field.
	linkDescURL := linkFields[4].Descriptor()
	// link.URLValidator is a validator for the "url" field. It is called by the builders before save.
	link.URLValidator = linkDescURL.Validators[0].(func(string) error)
	// linkDescCreatedAt is the schema descriptor for created_at field.
	linkDescCreatedAt := linkFields[5].Descriptor()
	// link.DefaultCreatedAt holds the default value on creation for the created_at field.
	link.DefaultCreatedAt = linkDescCreatedAt.Default.(func() time.Time)
	// linkDescUpdatedAt is the schema descriptor for updated_at field.
	linkDescUpdatedAt := linkFields[6].Descriptor()
	// link.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	link.DefaultUpdatedAt = linkDescUpdatedAt.Default.(func() time.Time)
	// link.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	link.UpdateDefaultUpdatedAt = linkDescUpdatedAt.UpdateDefault.(func() time.Time)
	// linkDescID is the schema descriptor for id field.
	linkDescID := linkFields[0].Descriptor()
	// link.DefaultID holds the default value on creation for the id field.
	link.DefaultID = linkDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[6].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
