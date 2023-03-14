package reader

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func CreateReader(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	r := &Reader{
		ID: seq,
	}
	if err := c.Bind(r); err != nil {
		return err
	}
	readers[r.ID] = r
	seq++
	return c.JSON(http.StatusCreated, r)
}
func GetReader(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, readers[id])
}
func GetReadersList(c echo.Context) error {
	return c.JSON(http.StatusOK, readers)
}
func UpdateReader(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := new(Reader)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	readers[id].FullName = u.FullName
	return c.JSON(http.StatusOK, readers[id])
}
func DeleteReader(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	delete(readers, id)
	return c.NoContent(http.StatusNoContent)
}
