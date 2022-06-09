package apiecho

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

//Generates a Header in every request
func RequestIdMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {

		requestId := c.Request().Header.Get("X-Request-ID") //Create header id
		if len(requestId) == 0 {
			requestId = uuid.New().String()
		}

		c.Response().Header().Set("X-Request-ID", requestId)

		return next(c) //When finish this function, we proced to the next one

	})

}

//Define validUsers for endpoints
var validUsers = map[string]string{"user1": "password1", "user2": "password2"}

//Control authorization for certian endpoints
func AuthorizationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {

		user := c.Request().Header.Get("Authorization")

		//Check if user is in the list of valid users
		if validUsers[user] == "" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")

		}

		return next(c) //When finish this function, we proced to the next one

	})
}
