package engine

import (
	"context"
	"github.com/graphql-go/graphql"
	"github.com/ldaysjun/unbff/config"
)

type Params struct {
	// GraphQL DSL
	DSL string

	// Application
	App string

	// A mapping of variable name to runtime value to use for all variables defined in the requestString.
	VariableValues map[string]interface{}
}

type Kernel struct {
	// graphql schema set
	schemas map[string]graphql.Schema

	// resource node pool
	np *nodePool

	// processor parse SDL
	p *processor
}

// Do execute query logic
func (k *Kernel) Do(ctx context.Context, p Params) *graphql.Result {
	s := k.schemas[p.App]
	params := graphql.Params{
		Schema:         s,
		RequestString:  p.DSL,
		VariableValues: p.VariableValues,
		Context:        ctx,
	}
	r := graphql.Do(params)
	return r
}

// NewSchema execute query logic
func (k *Kernel) NewSchema(ctx context.Context, p Params) graphql.Schema {
	return k.schemas[p.App]
}

// NewSchemaWithApp new schema
func (k *Kernel) NewSchemaWithApp(ctx context.Context, app string) (graphql.Schema, error) {
	if v, ok := k.schemas[app]; ok {
		return v, nil
	}
	return baseQuerySchema()
}

func baseQuerySchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "QUERY",
			Fields: graphql.Fields{
				"ping": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return "ping ok", nil
					},
				},
			},
			Description: "ALL QUERY",
		}),
	})
}

type option func(k *Kernel) error

// NewKernel new kernel
func NewKernel(conf *config.Config) (*Kernel, error) {
	k := &Kernel{}
	// strict order dependency
	opts := []option{
		//withNodePool(),
		withProcessor(conf),
		withSchema(),
	}
	for _, opt := range opts {
		if err := opt(k); err != nil {
			return nil, err
		}
	}
	return k, nil
}

func withNodePool() option {
	return func(k *Kernel) error {
		var err error
		k.np, err = NewNodePool()
		if err != nil {
			return err
		}
		return nil
	}
}

func withProcessor(conf *config.Config) option {
	return func(k *Kernel) error {
		k.p = newProcessor(conf)
		return nil
	}
}

func withSchema() option {
	return func(k *Kernel) error {
		k.schemas = make(map[string]graphql.Schema)
		appFields, err := k.p.generateRootFields()
		if err != nil {
			return nil
		}
		for app, fields := range appFields {
			fields["ping"] = &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "ping ok", nil
				},
			}
			if k.schemas[app], err = graphql.NewSchema(graphql.SchemaConfig{
				Query: graphql.NewObject(graphql.ObjectConfig{
					Name:        "QUERY",
					Fields:      fields,
					Description: "ALL QUERY",
				}),
			}); err != nil {
				return err
			}
		}
		return nil
	}
}
