package Restaurant

import (
	"SEProject/Restaurant/types"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	service *Service
}

func NewHandler(e *echo.Echo, service *Service) {
	h := &Handler{service: service}
	e.GET("/restaurants", h.GetAllRestaurants)
	e.POST("/admin/restaurant", h.CreateRestaurant)
	e.PUT("/admin/restaurant/:id", h.UpdateRestaurant)
	e.DELETE("/admin/restaurant/:id", h.DeleteRestaurant)
}

// GetAllRestaurants godoc
// @Summary Tüm restoranları getirir
// @Tags Restaurants
// @Produce json
// @Success 200 {array} types.Restaurant
// @Failure 500 {string} string "Restoranlar getirilemedi"
// @Router /restaurants [get]
func (h *Handler) GetAllRestaurants(c echo.Context) error {
	restaurants, err := h.service.GetAllRestaurants(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "Restoranlar getirilemedi")
	}
	return c.JSON(http.StatusOK, restaurants)
}

// CreateRestaurant godoc
// @Summary Yeni restoran oluştur
// @Tags Restaurants
// @Accept json
// @Produce plain
// @Param restaurant body types.Restaurant true "Yeni restoran"
// @Success 201 {string} string "Restoran başarıyla eklendi"
// @Failure 400 {string} string "Geçersiz veri"
// @Failure 500 {string} string "Restoran oluşturulamadı"
// @Router /admin/restaurant [post]
func (h *Handler) CreateRestaurant(c echo.Context) error {
	var r types.Restaurant
	if err := c.Bind(&r); err != nil {
		return c.String(http.StatusBadRequest, "Geçersiz veri")
	}
	if err := h.service.CreateRestaurant(c.Request().Context(), &r); err != nil {
		return c.String(http.StatusInternalServerError, "Restoran oluşturulamadı")
	}
	return c.String(http.StatusCreated, "Restoran başarıyla eklendi")
}

// UpdateRestaurant godoc
// @Summary Var olan restoranı güncelle
// @Tags Restaurants
// @Accept json
// @Produce plain
// @Param id path int true "Restoran ID"
// @Param restaurant body types.Restaurant true "Güncellenmiş restoran"
// @Success 200 {string} string "Restoran güncellendi"
// @Failure 400 {string} string "Geçersiz ID veya veri"
// @Failure 500 {string} string "Güncelleme başarısız"
// @Router /admin/restaurant/{id} [put]
func (h *Handler) UpdateRestaurant(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Geçersiz ID")
	}
	var r types.Restaurant
	if err := c.Bind(&r); err != nil {
		return c.String(http.StatusBadRequest, "Geçersiz veri")
	}
	if err := h.service.UpdateRestaurant(c.Request().Context(), id, &r); err != nil {
		return c.String(http.StatusInternalServerError, "Güncelleme başarısız")
	}
	return c.String(http.StatusOK, "Restoran güncellendi")
}

// DeleteRestaurant godoc
// @Summary Restoran sil
// @Tags Restaurants
// @Produce plain
// @Param id path int true "Restoran ID"
// @Success 200 {string} string "Restoran silindi"
// @Failure 400 {string} string "Geçersiz ID"
// @Failure 500 {string} string "Silme başarısız"
// @Router /admin/restaurant/{id} [delete]
func (h *Handler) DeleteRestaurant(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Geçersiz ID")
	}
	if err := h.service.DeleteRestaurant(c.Request().Context(), id); err != nil {
		return c.String(http.StatusInternalServerError, "Silme başarısız")
	}
	return c.String(http.StatusOK, "Restoran silindi")
}
