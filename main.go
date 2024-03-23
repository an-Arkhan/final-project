package main

import (
	// "fmt"
	"fmt"
	"log"
	"os"

	"final-project/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	port := os.Getenv("PORT")

	router.UserRouter(r)
	router.PhotoRouter(r)
	router.CommentRouter(r)
	router.SocialMediaRouter(r)

	fmt.Println("Service running at port :" + port)
	r.Run(":" + port)
}
