package main

import (
	"github.com/joho/godotenv"
	"github.com/t-shah02/mochi/internal"
)

func main() {
	godotenv.Load()

	server := internal.NewMochiServer()
	server.Serve()
}
