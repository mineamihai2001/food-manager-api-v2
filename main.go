package main

import (
	"fmt"

	"github.com/mineamihai2001/fm/cmd/router"
	"github.com/mineamihai2001/fm/helpers"
)

func main() {
	r := router.Create()

	env := helpers.Env()

	r.Run(fmt.Sprintf(":%d", env.App.Port))
}
