//go:build ignore

package main

import (
	"github.com/vektah/gqlparser/v2/ast"
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

// work directory 设置为项目根目录.可直接调试.
func main() {
	ex, err := entgql.NewExtension(
		entgql.WithSchemaGenerator(),
		entgql.WithWhereInputs(true),
		entgql.WithConfigPath("graph/gqlgen/gqlgen.yaml"),
		entgql.WithSchemaPath("graph/doc/ent.graphql"),
		entgql.WithSchemaHook(changeRelayNodeType()),
		entgql.WithTemplates(entgqlTemplate()...),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	opts := []entc.Option{
		entc.Extensions(ex),
		globalID(),
		//entc.TemplateDir("./ent/template"),
	}
	err = entc.Generate("./graph/entgen/schema", &gen.Config{
		Package:  "github.com/woocoos/adminx/ent",
		Features: []gen.Feature{gen.FeatureVersionedMigration, gen.FeatureIntercept},
		Target:   "./ent",
	}, opts...)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

func entgqlTemplate() []*gen.Template {
	return []*gen.Template{
		entgql.CollectionTemplate,
		entgql.EnumTemplate,
		entgql.NodeTemplate,
		entgql.PaginationTemplate,
		entgql.TransactionTemplate,
		entgql.EdgeTemplate,
		mutationInput(),
	}
}

func globalID() func(g *gen.Config) error {
	return func(g *gen.Config) error {
		g.Templates = append(g.Templates, gen.MustParse(gen.NewTemplate("gql_globalid").
			Funcs(entgql.TemplateFuncs).
			ParseFiles("./graph/entgen/template/globalid.tmpl")))
		return nil
	}
}

func mutationInput() *gen.Template {
	return gen.MustParse(gen.NewTemplate("gql_mutation_input").
		Funcs(entgql.TemplateFuncs).
		ParseFiles("./graph/entgen/template/gql_mutation_input.tmpl")).SkipIf(skipMutationTemplate)
}

// annotation extracts the entgql.Annotation or returns its empty value.
func annotation(ants gen.Annotations) (*entgql.Annotation, error) {
	ant := &entgql.Annotation{}
	if ants != nil && ants[ant.Name()] != nil {
		if err := ant.Decode(ants[ant.Name()]); err != nil {
			return nil, err
		}
	}
	return ant, nil
}

func skipMutationTemplate(g *gen.Graph) bool {
	for _, n := range g.Nodes {
		ant, err := annotation(n.Annotations)
		if err != nil {
			continue
		}
		for _, i := range ant.MutationInputs {
			if (i.IsCreate && !ant.Skip.Is(entgql.SkipMutationCreateInput)) ||
				(!i.IsCreate && !ant.Skip.Is(entgql.SkipMutationUpdateInput)) {
				return false
			}
		}
	}
	return true
}

func changeRelayNodeType() entgql.SchemaHook {
	idType := ast.NonNullNamedType("GID", nil)
	return func(graph *gen.Graph, schema *ast.Schema) error {
		for _, field := range schema.Types["Query"].Fields {
			if field.Name == "node" {
				field.Arguments[0].Type = idType
			}
			if field.Name == "nodes" {
				field.Arguments[0].Type = ast.NonNullListType(idType, nil)
			}
		}
		return nil
	}
}
