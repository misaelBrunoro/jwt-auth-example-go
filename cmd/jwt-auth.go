package main

import (
	"context"

	"github.com/jwt-auth/pkg/server"
)

func main() {
	ctx := context.Background()
	err := server.Execute(ctx)

	if err != nil {
		panic(err)
	}
}
