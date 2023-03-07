package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
)

func main() {
	// 定义 GraphQL Schema
	//fields := graphql.Fields{
	//	"hello": &graphql.Field{
	//		Type: graphql.String,
	//		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	//			return "Hello, world!", nil
	//		},
	//	},
	//}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: nil}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, _ := graphql.NewSchema(schemaConfig)

	// 处理 GraphQL 请求
	//h := handler.New(&handler.Config{
	//	Schema:   &schema,
	//	Pretty:   true,
	//	GraphiQL: true,
	//})

	params := graphql.Params{
		Schema:        schema,
		RequestString: `{he}`,
		Context:       context.Background(),
	}
	result := graphql.Do(params)
	d, _ := json.Marshal(result)
	fmt.Println(string(d))

	//http.Handle("/graphql", h)
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	panic(err)
	//}
}
