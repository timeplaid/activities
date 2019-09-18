package main

import (
	"github.com/timeplaid/activities/pkg/app"
)

func main() {

	server := app.NewServer()
	server.RegisterRoutes()
	server.ListenAndServe()

}
