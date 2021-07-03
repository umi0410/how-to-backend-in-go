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
		// 타입을 기반으로 안전하고 편리하게 컬럼을 정의할 수 있습니다.
		field.String("id"),
		field.String("name"),
		field.Bool("isActivated").Default(true),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// ent에서는 To를 정의하는 스키마, 즉 여기선 User
		// 가 참조 관계의 주인이라고 정의합니다.
		// 일반적인 JPA의 방식과는 반대입니다.
		edge.To("products", TourProduct.Type),
	}
}
