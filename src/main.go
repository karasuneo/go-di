package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/karasuneo/go-di/src/presentation/handlers"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	handlers.CreateUserHandler(r)
	r.Run()
}
