package main

import (
	"receipt-processor-challenge/internal/delivery/router"
)

func main() {
	r := router.SetupRoutes()

	r.Run(":8080")
}
