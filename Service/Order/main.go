package main

import (
	"github.com/labstack/echo"
	"go-module/libs/api_gateway"
	"go-module/config/server"
	"go-module/middleware"
	"go-module/handler"
	"go-module/driver"
	"fmt"
)

func main() {
	
	// Connecting to Kong server
	api_gateway.RegisterKong()

	// Conencting to database
	db := driver.Connect()
	// err := db.SQL.Ping()
	fmt.Println(db)

	// Echo instance
	e := echo.New()

	// Routes
	// Common endpoints
	e.GET("/greeting", handler.Greeting)

	// User endpoints
	userGroup := e.Group("/api")
	userGroup.Use(middleware.Authentication)
	userGroup.POST("/create", handler.Create)
	userGroup.PATCH("/update/:id", handler.UpdateById)

	// Admin endpoints
	// adminGroup := e.Group("/admin")
	adminGroup := userGroup.Group("/admin")
	adminGroup.Use(middleware.Authorization)
	adminGroup.GET("/find/:id", handler.Find)
	adminGroup.GET("/find", handler.Find)
	adminGroup.DELETE("/delete/:id", handler.DeleteById)

	// Start server
	e.Logger.Fatal(e.Start(":" + server.PORT))
}
