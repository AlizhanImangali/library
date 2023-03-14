package author

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateAuthor(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	a := &Author{}
	if err := c.Bind(a); err != nil {
		return err
	}

	a.ID = seq
	seq++
	authors[a.ID] = a

	return c.JSON(http.StatusCreated, a)
}

func GetAuthor(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	return c.JSON(http.StatusOK, authors[id])
}

func GetAuthorsList(c echo.Context) error {
	return c.JSON(http.StatusOK, authors)
}

func UpdateAuthor(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	u := new(Author)
	if err := c.Bind(u); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	authors[id].FullName = u.FullName

	return c.JSON(http.StatusOK, authors[id])
}

func DeleteAuthor(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	id, _ := strconv.Atoi(c.Param("id"))
	delete(authors, id)

	return c.NoContent(http.StatusNoContent)
}
