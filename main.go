package main

import (
	"crud/author"
	"crud/book"
	"crud/reader"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/*
Автора характеризуют следующие параметры:
- Идентификатор
- ФИО
- Псевдоним
- Специализация

Книга:
- Иденификатор
- Название
- Жанр
- Код ISBN

Читатель:
- Идентификатор
- ФИО
- Список взятых книг
*/
/*
- /authors/{id}/books - список автора, которые есь в наличии
- /members/{id}/books - список взятых читателем книг
*/
func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes

	e.POST("/authors", author.CreateAuthor)
	e.POST("/books", book.CreateBook)
	e.POST("/members", reader.CreateReader)

	e.GET("/authors/:id", author.GetAuthor)
	e.GET("/books/:id", book.GetBook)
	e.GET("/members/:id", reader.GetReader)
	//List all
	e.GET("/authors", author.GetAuthorsList)
	e.GET("/books", book.GetBooksList)
	e.GET("members", reader.GetReadersList)
	//Special Routes
	//e.GET("/authors/{id}/books", getAuthorInStock)
	//e.GET("/members/{id}/books", getReaderBook)

	e.PUT("/authors/:id", author.UpdateAuthor)
	e.PUT("/books/:id", book.UpdateBook)
	e.PUT("/members/:id", reader.UpdateReader)

	e.DELETE("/authors/:id", author.DeleteAuthor)
	e.DELETE("/books/:id", book.DeleteBook)
	e.DELETE("/members/:id", reader.DeleteReader)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
