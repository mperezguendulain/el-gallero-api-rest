package handlers

import (
	"fmt"
	"net/http"

	"github.com/mperezguendulain/el-gallero-api-rest/errors"
	"github.com/mperezguendulain/el-gallero-api-rest/models"
	"github.com/mperezguendulain/el-gallero-api-rest/services/jwtserv"

	"github.com/labstack/echo"
)

// LoginHandler handle the login request.
func LoginHandler(c echo.Context) error {

	fmt.Println(`================`)
	fmt.Println(`LOGIN HANDLER`)
	fmt.Println(`================`)

	credential := new(jwtserv.Credential)
	if c.Bind(credential) != nil {
		return c.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	if credential.Username == "mperezguendulain" && credential.Password == "12345" {

		token, err := jwtserv.GetJWT(credential.Username, "admin")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
		}
		responseToken := models.ResponseToken{
			Token: token,
		}
		return c.JSON(http.StatusOK, responseToken)

	}
	return c.JSON(http.StatusUnauthorized, errors.GetUnauthorizedError())

}
