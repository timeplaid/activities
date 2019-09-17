package main

import (
	"github.com/timeplaid/activities/pkg/app"
)

func main() {

	server := app.NewServer(":8080")
	server.RegisterRoutes()
	server.ListenAndServe()

}
