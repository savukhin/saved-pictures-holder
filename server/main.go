package main

import (
	"fmt"
	models "saved-pictures-holder/models"
	"saved-pictures-holder/routes"
)

func main() {
	db, err := models.Connect()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	if err = models.LoadTables(db); err != nil {
		panic(err)
	}

	r := routes.SetupRouter(db)

	fmt.Println("Server is running on port 3000")
	r.Run(":3000")
}
