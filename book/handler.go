package book

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func CreateBook(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	b := &Book{
		ID: seq,
	}
	if err := c.Bind(b); err != nil {
		return err
	}
	books[b.ID] = b
	seq++
	return c.JSON(http.StatusCreated, b)
}

func GetBook(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, books[id])
}
func GetBooksList(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}
func UpdateBook(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := new(Book)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	books[id].Name = u.Name
	return c.JSON(http.StatusOK, books[id])
}
func DeleteBook(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	delete(books, id)
	return c.NoContent(http.StatusNoContent)
}
