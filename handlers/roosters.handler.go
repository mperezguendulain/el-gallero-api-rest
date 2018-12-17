package handlers

import (
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/mperezguendulain/el-gallero-api-rest/errors"
	"github.com/mperezguendulain/el-gallero-api-rest/models"
	"github.com/mperezguendulain/el-gallero-api-rest/services/roosterserv"

	"github.com/labstack/echo"
)

// CreateRoosterHandler handle the creation of a Rooster.
func CreateRoosterHandler(ctx echo.Context) error {

	rooster := models.Rooster{}
	if ctx.Bind(&rooster) != nil {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	if roosterserv.CreateRooster(rooster) != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.NoContent(http.StatusCreated)

}

// GetAllRoostersHandler handle the obtaining of roosters
func GetAllRoostersHandler(ctx echo.Context) error {

	roosters, err := roosterserv.FindAllRoosters()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.JSON(http.StatusOK, roosters)

}

// GetRoosterByIDHandler handle the obtaining of a rooster
func GetRoosterByIDHandler(ctx echo.Context) error {

	strUID := ctx.Param("id")
	if !bson.IsObjectIdHex(strUID) {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	rooster, err := roosterserv.FindRoosterByID(strUID)
	if err != nil {

		if err == mgo.ErrNotFound {
			return ctx.NoContent(http.StatusNotFound)
		}

		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.JSON(http.StatusOK, rooster)

}

// RemoveRoosterByIDHandler handle the obtaining of a rooster
func RemoveRoosterByIDHandler(ctx echo.Context) error {

	strUID := ctx.Param("id")
	if !bson.IsObjectIdHex(strUID) {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	err := roosterserv.RemoveRoosterByID(strUID)
	if err != nil {

		if err == mgo.ErrNotFound {
			return ctx.NoContent(http.StatusNotFound)
		}

		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.NoContent(http.StatusNoContent)

}

// UpdateRoosterByIDHandler updates the rooster information
func UpdateRoosterByIDHandler(ctx echo.Context) error {

	strUID := ctx.Param("id")
	if !bson.IsObjectIdHex(strUID) {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	rooster := models.Rooster{}
	if ctx.Bind(&rooster) != nil {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	err := roosterserv.UpdateRooster(strUID, rooster)
	if err != nil {

		if err == mgo.ErrNotFound {
			return ctx.NoContent(http.StatusNotFound)
		}

		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.NoContent(http.StatusNoContent)

}
