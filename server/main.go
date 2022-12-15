package main

import (
	"saved-pictures-holder/routes"
)

func main() {
	r := routes.SetupRouter()

	r.Run(":3000")
}
