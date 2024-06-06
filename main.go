package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/qsmsoft/go-rest-api/db"
	"github.com/qsmsoft/go-rest-api/routes"
)

func main() {
	db.InitDB()
	defer func() {
		if err := db.DB.Close(); err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}()
	fmt.Println("Database initialized and tables created successfully.")

	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
