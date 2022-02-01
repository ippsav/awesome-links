// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/ippsav/awesome-links/api/ent/link"
	"github.com/ippsav/awesome-links/api/ent/user"
)

// Link is the model entity for the Link schema.
type Link struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LinkQuery when eager-loading is set.
	Edges          LinkEdges `json:"edges"`
	user_bookmarks *uuid.UUID
}

// LinkEdges holds the relations/edges for other nodes in the graph.
type LinkEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LinkEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Link) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case link.FieldTitle, link.FieldDescription, link.FieldImageURL, link.FieldURL:
			values[i] = new(sql.NullString)
		case link.FieldCreatedAt, link.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case link.FieldID:
			values[i] = new(uuid.UUID)
		case link.ForeignKeys[0]: // user_bookmarks
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Link", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Link fields.
func (l *Link) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case link.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				l.ID = *value
			}
		case link.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				l.Title = value.String
			}
		case link.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				l.Description = value.String
			}
		case link.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				l.ImageURL = value.String
			}
		case link.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				l.URL = value.String
			}
		case link.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		case link.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				l.UpdatedAt = value.Time
			}
		case link.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_bookmarks", values[i])
			} else if value.Valid {
				l.user_bookmarks = new(uuid.UUID)
				*l.user_bookmarks = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Link entity.
func (l *Link) QueryOwner() *UserQuery {
	return (&LinkClient{config: l.config}).QueryOwner(l)
}

// Update returns a builder for updating this Link.
// Note that you need to call Link.Unwrap() before calling this method if this Link
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Link) Update() *LinkUpdateOne {
	return (&LinkClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the Link entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Link) Unwrap() *Link {
	tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Link is not a transactional entity")
	}
	l.config.driver = tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Link) String() string {
	var builder strings.Builder
	builder.WriteString("Link(")
	builder.WriteString(fmt.Sprintf("id=%v", l.ID))
	builder.WriteString(", title=")
	builder.WriteString(l.Title)
	builder.WriteString(", description=")
	builder.WriteString(l.Description)
	builder.WriteString(", image_url=")
	builder.WriteString(l.ImageURL)
	builder.WriteString(", url=")
	builder.WriteString(l.URL)
	builder.WriteString(", created_at=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(l.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Links is a parsable slice of Link.
type Links []*Link

func (l Links) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}
