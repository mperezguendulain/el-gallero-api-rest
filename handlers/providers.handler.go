package handlers

import (
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/mperezguendulain/el-gallero-api-rest/errors"
	"github.com/mperezguendulain/el-gallero-api-rest/models"
	"github.com/mperezguendulain/el-gallero-api-rest/services/providerserv"

	"github.com/labstack/echo"
)

// CreateProviderHandler handle the creation of a Provider.
func CreateProviderHandler(ctx echo.Context) error {

	provider := models.Provider{}
	if ctx.Bind(&provider) != nil {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	if providerserv.CreateProvider(provider) != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.NoContent(http.StatusCreated)

}

// GetAllProvidersHandler handle the obtaining of providers
func GetAllProvidersHandler(ctx echo.Context) error {

	providers, err := providerserv.FindAllProviders()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.JSON(http.StatusOK, providers)

}

// GetProviderByIDHandler handle the obtaining of a provider
func GetProviderByIDHandler(ctx echo.Context) error {

	strUID := ctx.Param("id")
	if !bson.IsObjectIdHex(strUID) {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	provider, err := providerserv.FindProviderByID(strUID)
	if err != nil {

		if err == mgo.ErrNotFound {
			return ctx.NoContent(http.StatusNotFound)
		}

		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.JSON(http.StatusOK, provider)

}

// RemoveProviderByIDHandler handle the obtaining of a provider
func RemoveProviderByIDHandler(ctx echo.Context) error {

	strUID := ctx.Param("id")
	if !bson.IsObjectIdHex(strUID) {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	err := providerserv.RemoveProviderByID(strUID)
	if err != nil {

		if err == mgo.ErrNotFound {
			return ctx.NoContent(http.StatusNotFound)
		}

		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.NoContent(http.StatusNoContent)

}

// UpdateProviderByIDHandler updates the provider information
func UpdateProviderByIDHandler(ctx echo.Context) error {

	strUID := ctx.Param("id")
	if !bson.IsObjectIdHex(strUID) {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	provider := models.Provider{}
	if ctx.Bind(&provider) != nil {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	err := providerserv.UpdateProvider(strUID, provider)
	if err != nil {

		if err == mgo.ErrNotFound {
			return ctx.NoContent(http.StatusNotFound)
		}

		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.NoContent(http.StatusNoContent)

}
