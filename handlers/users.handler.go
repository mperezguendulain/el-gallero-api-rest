package handlers

import (
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/mperezguendulain/el-gallero-api-rest/errors"
	"github.com/mperezguendulain/el-gallero-api-rest/models"
	"github.com/mperezguendulain/el-gallero-api-rest/services/userserv"

	"github.com/labstack/echo"
)

// CreateUserHandler handle the creation of a User.
func CreateUserHandler(ctx echo.Context) error {

	user := models.User{}
	if ctx.Bind(&user) != nil {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	if userserv.CreateUser(user) != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.NoContent(http.StatusCreated)

}

// GetAllUsersHandler handle the obtaining of users
func GetAllUsersHandler(ctx echo.Context) error {

	users, err := userserv.FindAllUsers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.JSON(http.StatusOK, users)

}

// GetUserByIDHandler handle the obtaining of a user
func GetUserByIDHandler(ctx echo.Context) error {

	strUID := ctx.Param("id")
	if !bson.IsObjectIdHex(strUID) {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	user, err := userserv.FindUserByID(strUID)
	if err != nil {

		if err == mgo.ErrNotFound {
			return ctx.NoContent(http.StatusNotFound)
		}

		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.JSON(http.StatusOK, user)

}

// RemoveUserByIDHandler handle the obtaining of a user
func RemoveUserByIDHandler(ctx echo.Context) error {

	strUID := ctx.Param("id")
	if !bson.IsObjectIdHex(strUID) {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	err := userserv.RemoveUserByID(strUID)
	if err != nil {

		if err == mgo.ErrNotFound {
			return ctx.NoContent(http.StatusNotFound)
		}

		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.NoContent(http.StatusNoContent)

}

// UpdateUserByIDHandler updates the user information
func UpdateUserByIDHandler(ctx echo.Context) error {

	strUID := ctx.Param("id")
	if !bson.IsObjectIdHex(strUID) {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	user := models.User{}
	if ctx.Bind(&user) != nil {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	err := userserv.UpdateUser(strUID, user)
	if err != nil {

		if err == mgo.ErrNotFound {
			return ctx.NoContent(http.StatusNotFound)
		}

		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.NoContent(http.StatusNoContent)

}
