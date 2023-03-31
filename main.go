package main

import (
	author2 "crud/domain/author"
	book2 "crud/domain/book"
	reader2 "crud/domain/reader"
	database2 "crud/pkg/database"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	databaseSourceName := os.Getenv("DATABASE_URL")

	// init database instance
	postgres, err := database2.New(databaseSourceName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer postgres.Close()

	// migrate up database
	err = database2.Migrate("postgres://habrpguser:pgpwd4habr@localhost:6432/library?sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// init handlers
	authorStorage := author2.NewStorage(postgres)
	authorHandler := author2.NewHandler(authorStorage)

	bookStorage := book2.NewStorage(postgres)
	bookHandler := book2.NewHandler(bookStorage)

	readerStorage := reader2.NewStorage(postgres)
	readerHandler := reader2.NewHandler(readerStorage)

	// setup middleware
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// init routes
	apiGroup := e.Group("/api")

	authorGroup := apiGroup.Group("/authors")
	authorGroup.POST("", authorHandler.Create)
	authorGroup.GET("", authorHandler.GetAll)
	authorGroup.GET("/:id", authorHandler.Get)
	authorGroup.PUT("/:id", authorHandler.Update)
	authorGroup.DELETE("/:id", authorHandler.Delete)

	bookGroup := apiGroup.Group("/books")
	bookGroup.POST("", bookHandler.Create)
	bookGroup.GET("", bookHandler.GetAll)
	bookGroup.GET("/:id", bookHandler.Get)
	bookGroup.PUT("/:id", bookHandler.Update)
	bookGroup.DELETE("/:id", bookHandler.Delete)

	readerGroup := apiGroup.Group("/readers")
	readerGroup.POST("", readerHandler.Create)
	readerGroup.GET("", readerHandler.GetAll)
	readerGroup.GET("/:id", readerHandler.Get)
	readerGroup.PUT("/:id", readerHandler.Update)
	readerGroup.DELETE("/:id", readerHandler.Delete)

	// start server
	fmt.Println("server running on: http://localhost")
	e.Logger.Fatal(e.Start(":81"))
}
