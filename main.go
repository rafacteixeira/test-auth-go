package main

import (
	"test-auth/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Overload(".env")

	routes.HandleRequests()
}
