package main

import (
	"os"

	"github.com/nextadm/cmd/auth/app"
)

func main() {
	run := app.NewAuthServer()
	code := run()
	os.Exit(code)
}
