package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ldaysjun/unbff/config"
	"github.com/ldaysjun/unbff/engine"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	kernel, err := engine.NewKernel(config.ForTest())
	if err != nil {
		panic(err)
	}
	s := "{\n  getCustomerById(id:1){\n    email\n    id\n    name\n  }\n}"
	r := kernel.Do(context.Background(), engine.Params{
		DSL:            s,
		App:            "test_app",
		VariableValues: nil,
	})

	d, _ := json.Marshal(r)
	fmt.Println(string(d))
}
