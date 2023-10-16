package main

// import (
// 	"github.com/labstack/echo"
// 	"go-module/libs/api_gateway"
// 	"go-module/config/server"
// 	"go-module/middleware"
// 	"go-module/handler"
// 	"go-module/driver"
// 	"fmt"
// )

// func main() {
	
// 	// Connect to Kong server
// 	api_gateway.RegisterKong()

// 	// Connect to Service Discovery server
// 	// registerServiceWithConsul()	

// 	// Connect to Database server
// 	db := driver.Connect()
// 	fmt.Println("db: ", db)

// 	// Echo instance
// 	e := echo.New()

// 	// Debug mode 
// 	e.Debug = true

// 	// Routes
// 	// Common endpoints
// 	e.GET("/greeting", handler.Greeting)
// 	e.POST("/register", handler.Register)
// 	e.POST("/login", handler.Login)
	
// 	// User endpoints
// 	userGroup := e.Group("/api")
// 	userGroup.Use(middleware.Authentication)
// 	userGroup.GET("/account-info", handler.GetAccountInfo)

// 	// Admin endpoints
// 	adminGroup := userGroup.Group("/admin")
// 	adminGroup.Use(middleware.Authorization)
// 	adminGroup.GET("/count", handler.Count)
// 	adminGroup.GET("/find", handler.Find)
// 	adminGroup.GET("/find/:id", handler.FindById)
// 	adminGroup.POST("/populating-data", handler.PopulatingData)
// 	adminGroup.PATCH("/update/:id", handler.UpdateById)
// 	adminGroup.DELETE("/delete-all", handler.RemoveAll)

// 	// Start server
// 	e.Logger.Fatal(e.Start(":" + server.PORT))
// }

	import "fmt"
	func main() {
		fmt.Println("Let's kill this idea!")
	}