package main

import (
	"log"
	"myapp/handlers/bot"
	"myapp/handlers/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// init вызовится перед main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	config := config.New()
	router := gin.Default()
	bot.CreateBot()
	router.Run(":" + config.PORT)

}
