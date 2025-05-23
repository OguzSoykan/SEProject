package Menu

import (
	"SEProject/Menu/types"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	service *Service
}

func NewHandler(e *echo.Echo, service *Service) {
	h := &Handler{service: service}
	e.GET("/restaurants/:id/menu", h.GetMenuByRestaurant)
	e.POST("/restaurant/menu", h.CreateMenuItem)
	e.PUT("/restaurant/menu/:id", h.UpdateMenuItem)
	e.DELETE("/restaurant/menu/:id", h.DeleteMenuItem)
}

func (h *Handler) GetMenuByRestaurant(c echo.Context) error {
	restID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Geçersiz restoran ID")
	}
	menu, err := h.service.GetMenuByRestaurantID(c.Request().Context(), restID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Menü getirilemedi")
	}
	return c.JSON(http.StatusOK, menu)
}

func (h *Handler) CreateMenuItem(c echo.Context) error {
	var item types.Menu
	if err := c.Bind(&item); err != nil {
		return c.String(http.StatusBadRequest, "Geçersiz menü verisi")
	}
	if err := h.service.CreateMenuItem(c.Request().Context(), &item); err != nil {
		return c.String(http.StatusInternalServerError, "Menü ürünü eklenemedi")
	}
	return c.String(http.StatusCreated, "Menü ürünü eklendi")
}

func (h *Handler) UpdateMenuItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Geçersiz ID")
	}
	var item types.Menu
	if err := c.Bind(&item); err != nil {
		return c.String(http.StatusBadRequest, "Geçersiz veri")
	}
	if err := h.service.UpdateMenuItem(c.Request().Context(), id, &item); err != nil {
		return c.String(http.StatusInternalServerError, "Menü ürünü güncellenemedi")
	}
	return c.String(http.StatusOK, "Menü ürünü güncellendi")
}

func (h *Handler) DeleteMenuItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Geçersiz ID")
	}
	if err := h.service.DeleteMenuItem(c.Request().Context(), id); err != nil {
		return c.String(http.StatusInternalServerError, "Menü ürünü silinemedi")
	}
	return c.String(http.StatusOK, "Menü ürünü silindi")
}
