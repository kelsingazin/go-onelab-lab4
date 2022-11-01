package utils

import (
	"github.com/labstack/echo/v4"
	user "lab4-kelsingazin/models"
	"net/http"
	"strconv"
)

func CreateUser(c echo.Context) error {
	u := &user.User{ID: user.Seq}
	if err := c.Bind(u); err != nil {
		return err
	}
	user.Users[u.ID] = u
	user.Seq++
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, user.Users[id])
}

func GetAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, user.Users)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(user.Users, id)
	return c.NoContent(http.StatusNoContent)
}
