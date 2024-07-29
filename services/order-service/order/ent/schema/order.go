package schema

import "entgo.io/ent"
import "time"
import "entgo.io/ent/schema/field"
// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
        field.Int("user_id"),
        field.Int("product_id"),
        field.Int("quantity"),
        field.String("status"),
        field.Time("created_at").Default(time.Now),
        field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
    }
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
        edge.From("user", User.Type).Ref("orders").Unique().Field("user_id"),
        edge.From("product", Product.Type).Ref("orders").Unique().Field("product_id"),
    }
}
