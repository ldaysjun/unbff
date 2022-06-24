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

type kernel struct {
	// graphql schema set
	schemas map[string]graphql.Schema

	// resource node pool
	np *nodePool

	// processor parse SDL
	p *processor
}

// Do execute query logic
func (k *kernel) Do(ctx context.Context, p Params) *graphql.Result {
	params := graphql.Params{
		Schema:         k.schemas[p.App],
		RequestString:  p.DSL,
		VariableValues: p.VariableValues,
		Context:        ctx,
	}
	return graphql.Do(params)
}

type option func(k *kernel) error

// NewKernel new kernel
func NewKernel(conf *config.Config) (*kernel, error) {
	k := &kernel{}
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
	return func(k *kernel) error {
		var err error
		k.np, err = NewNodePool()
		if err != nil {
			return err
		}
		return nil
	}
}

func withProcessor(conf *config.Config) option {
	return func(k *kernel) error {
		k.p = newProcessor(conf)
		return nil
	}
}

func withSchema() option {
	return func(k *kernel) error {
		k.schemas = make(map[string]graphql.Schema)
		appFields, err := k.p.generateRootFields()
		if err != nil {
			return nil
		}
		for app, fields := range appFields {
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
