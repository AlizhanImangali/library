package author

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	Create(c echo.Context) (err error)
	Get(c echo.Context) (err error)
	GetAll(c echo.Context) error
	Update(c echo.Context) (err error)
	Delete(c echo.Context) (err error)
}

type handler struct {
	storage Storage
}

func NewHandler(storage Storage) Handler {
	return &handler{
		storage: storage,
	}
}

func (h *handler) Create(c echo.Context) (err error) {
	// Binding body to struct
	data := Author{}
	if err = c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Implementation of storage method
	data.ID, err = h.storage.CreateRow(data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, data)
}

func (h *handler) Get(c echo.Context) (err error) {
	id := c.Param("id")

	data, err := h.storage.GetRowByID(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}
func (h *handler) GetAll(c echo.Context) error {
	data, err := h.storage.SelectRows()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) Update(c echo.Context) (err error) {
	// Binding body to struct
	data := Author{}
	if err = c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Binding uri to parameter
	data.ID = c.Param("id")

	// Implementation of storage method
	err = h.storage.UpdateRow(data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) Delete(c echo.Context) (err error) {
	// Binding uri to parameter
	id := c.Param("id")

	// Implementation of storage method
	err = h.storage.DeleteRow(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
