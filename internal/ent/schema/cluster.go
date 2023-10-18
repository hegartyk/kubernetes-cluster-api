package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/gidx"
)

// Cluster holds the schema definition for the Cluster entity.
type Cluster struct {
	ent.Schema
}

// Mixin of the Cluster type
func (Cluster) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the Cluster.
func (Cluster) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the Cluster.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ClusterPrefix) }),
		field.Text("name").
			NotEmpty().
			Comment("The name of the Cluster.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Text("description").
			Optional().
			Comment("The description of the Cluster.").
			Annotations(
				entgql.OrderField("DESCRIPTION"),
			),
		field.String("tenant_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the tenant for this Cluster.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("TENANT"),
			),
	}
}

// Edges of the Cluster.
func (Cluster) Edges() []ent.Edge {
	return nil
}

// Indexes of the Cluster.
func (Cluster) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tenant_id"),
	}
}

// Annotations of the Cluster
func (Cluster) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Represents a Cluster on the graph."),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Create a new Cluster."),
			entgql.MutationUpdate().Description("Update an existing Cluster."),
		),
	}
}
