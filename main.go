package main

import (
	"context"
	"net/http"

	"github.com/emili-grant/mockgen-demo/internal"
)

func main() {
	// The usual DI pattern

	// Construct components
	comp := internal.NewSomeComponent()

	// Inject
	svc := internal.NewService(comp, &http.Client{})

	// Use
	_ = svc.DoAThing(context.Background())
}
