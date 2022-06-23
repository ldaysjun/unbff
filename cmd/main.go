package main

import (
	"context"
	"github.com/ldaysjun/unbff/engine"
)

func main() {
	c := engine.NewCore()
	c.Do(context.Background())
}
