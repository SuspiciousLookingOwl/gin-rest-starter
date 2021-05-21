package main

import (
	"os"

	"github.com/gin-gonic/gin"
	dotenv "github.com/joho/godotenv"

	_ "example.com/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"example.com/compA"
	// "example.com/compB"
	"example.com/database"
)

// @title Swagger Example API
// @version 1.0
// @description Docs example.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host example.com
// @BasePath /api
func main() {
	// load .env
	err := dotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// set gin mode
	mode := gin.DebugMode
	if os.Getenv("GO_ENV") == "production" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	// initialize DB and run migration on all models
	db := database.Init()
	db.AutoMigrate(&compA.CompA{} /*, &compB.CompB{} */)

	// create gin instance
	app := gin.Default()
	port := os.Getenv("PORT")

	// Swagger
	url := ginSwagger.URL(port + "/swagger/doc.json")
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// health check endpoint
	app.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// register routes
	api := app.Group("/api")
	compA.RegisterRoute(api)

	// run app
	app.Run(":" + port)
}
