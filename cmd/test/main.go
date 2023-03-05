package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql/language/parser"
	"github.com/graphql-go/graphql/language/source"
	"io/ioutil"
	"net/url"
	"os"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	//解析这个 URL 并确保解析没有出错。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)

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
		return
	}
	data, _ := json.Marshal(AST)
	fmt.Println(string(data))
}
