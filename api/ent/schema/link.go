package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Link holds the schema definition for the Link entity.
type Link struct {
	ent.Schema
}

// Fields of the Link.
func (Link) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("title").NotEmpty(),
		field.String("description").NotEmpty(),
		field.String("image_url"),
		field.String("url").NotEmpty(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Link.
func (Link) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("links").Unique(),
		edge.From("users", User.Type).Ref("bookmarks"),
	}
}
