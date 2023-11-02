package api

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"unit_test/models"
	"unit_test/repo"
)

func CreateUser(c echo.Context) error {

	req := &models.CreateUserRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(400, http.StatusBadRequest)
	}
	if req.ID == "" {
		return c.JSON(400, http.StatusBadRequest)
	}
	if err := repo.CreateUser(context.Background(), req.ID); err != nil {
		return c.JSON(500, http.StatusInternalServerError)
	}
	fmt.Println("Nothing's gonna change my love for you")

	return c.JSON(200, http.StatusOK)
}

func UpdateUser(c echo.Context) error {

	req := &models.CreateUserRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(400, http.StatusBadRequest)
	}
	if req.ID == "" {
		return c.JSON(400, http.StatusBadRequest)
	}
	if err := repo.Updateuser(context.Background(), req.ID); err != nil {
		return c.JSON(500, http.StatusInternalServerError)
	}
	fmt.Println("Nothing's gonna change my love for you")

	return c.JSON(200, http.StatusOK)
}
