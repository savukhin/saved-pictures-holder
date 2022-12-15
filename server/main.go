package main

import (
	models "saved-pictures-holder/models"
	"saved-pictures-holder/routes"
)

func main() {
	db, err := models.Connect()

	if err != nil {
		panic(err)
	}

	r := routes.SetupRouter(db)

	r.Run(":3000")
}
