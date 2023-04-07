package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"crud/domain/author"
	"crud/domain/book"
	"crud/domain/health"
	"crud/domain/reader"
	"crud/pkg/database"
)

func main() {
	Postgres := os.Getenv("POSTGRES_URL")

	googleURL := os.Getenv("GOOGLE_URL")

	// init database instance
	postgres, err := database.New(Postgres)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer postgres.Close()
	//postgres://postgres:postgrespw@localhost:6432/library?sslmode=disable
	//postgres://habrpguser:pgpwd4habr@localhost:6432/library?sslmode=disable
	// migrate up databasegit
	err = database.Migrate(Postgres)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// init handlers
	authorStorage := author.NewStorage(postgres)
	authorHandler := author.NewHandler(authorStorage)

	bookStorage := book.NewStorage(postgres)
	bookHandler := book.NewHandler(bookStorage)

	readerStorage := reader.NewStorage(postgres)
	readerHandler := reader.NewHandler(readerStorage)

	healthHandler := health.NewHandler(googleURL, Postgres)

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

	//health
	health := apiGroup.Group("/health")
	health.GET("", healthHandler.Health)

	//special routes

	// start server
	e.Logger.Fatal(e.Start(":8081"))
}
