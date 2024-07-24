package handlers

import (
	"context"
	"net/http"
	"project1/config"
	"project1/models"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
    product := new(models.Product)
    if err := c.Bind(product); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    query := `INSERT INTO products (name, description, price) VALUES (:name, :description, :price)`
    stmt, err := config.DB.PrepareNamedContext(ctx, query)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    _, err = stmt.ExecContext(ctx, product)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, product)
}
