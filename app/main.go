package main

import (
	"context"
	"log"

	"github.com/leandrohsilveira/tsm/setup"
)

func main() {
	ctx := context.Background()
	app := setup.SetupApp()

	setup.SetupLogger(app)

	ctx, err := setup.SetupDatabasePool(ctx)
	if err != nil {
		log.Fatal(err)
	}

	setup.SetupPages(ctx, app)

	log.Fatal(app.Listen(":3000"))
}
