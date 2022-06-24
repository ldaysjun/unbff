package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql/language/parser"
	"github.com/graphql-go/graphql/language/source"
	"io/ioutil"
	"os"
)

func main() {
	r, err := os.Open("/Users/jun/Desktop/work/dream/unbff/cmd/test/example.graphql")
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(r)
	AST, err := parser.Parse(parser.ParseParams{Source: &source.Source{
		Body: content,
		Name: "GraphQL",
	}})
	if err != nil {
		fmt.Println(err)
	}
	data, _ := json.Marshal(AST)
	fmt.Println(string(data))
}
