package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
		field.String("password").Sensitive(),
		field.Int("age").NonNegative().Optional(),
		field.String("phone").Optional(),
		field.String("email").Optional(),
		field.String("addr").Optional(),
		field.String("avatar").Default("http://cdn.redrock.team/web_app_default.jpeg"),
		field.String("career").Default("自由工作者"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sent", Pet.Type),
		edge.To("token", Token.Type).
			Unique(),
		edge.To("adopted", Pet.Type),
	}
}
