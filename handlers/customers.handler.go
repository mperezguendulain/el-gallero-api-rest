package handlers

import (
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/mperezguendulain/el-gallero-api-rest/errors"
	"github.com/mperezguendulain/el-gallero-api-rest/models"
	"github.com/mperezguendulain/el-gallero-api-rest/services/customerserv"

	"github.com/labstack/echo"
)

// CreateCustomerHandler handle the creation of a Customer.
func CreateCustomerHandler(ctx echo.Context) error {

	customer := models.Customer{}
	if ctx.Bind(&customer) != nil {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	if customerserv.CreateCustomer(customer) != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.NoContent(http.StatusCreated)

}

// GetAllCustomersHandler handle the obtaining of customers
func GetAllCustomersHandler(ctx echo.Context) error {

	customers, err := customerserv.FindAllCustomers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.JSON(http.StatusOK, customers)

}

// GetCustomerByIDHandler handle the obtaining of a customer
func GetCustomerByIDHandler(ctx echo.Context) error {

	strUID := ctx.Param("id")
	if !bson.IsObjectIdHex(strUID) {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	customer, err := customerserv.FindCustomerByID(strUID)
	if err != nil {

		if err == mgo.ErrNotFound {
			return ctx.NoContent(http.StatusNotFound)
		}

		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.JSON(http.StatusOK, customer)

}

// RemoveCustomerByIDHandler handle the obtaining of a customer
func RemoveCustomerByIDHandler(ctx echo.Context) error {

	strUID := ctx.Param("id")
	if !bson.IsObjectIdHex(strUID) {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	err := customerserv.RemoveCustomerByID(strUID)
	if err != nil {

		if err == mgo.ErrNotFound {
			return ctx.NoContent(http.StatusNotFound)
		}

		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.NoContent(http.StatusNoContent)

}

// UpdateCustomerByIDHandler updates the customer information
func UpdateCustomerByIDHandler(ctx echo.Context) error {

	strUID := ctx.Param("id")
	if !bson.IsObjectIdHex(strUID) {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	customer := models.Customer{}
	if ctx.Bind(&customer) != nil {
		return ctx.JSON(http.StatusBadRequest, errors.GetBadRequestBodyError())
	}

	err := customerserv.UpdateCustomer(strUID, customer)
	if err != nil {

		if err == mgo.ErrNotFound {
			return ctx.NoContent(http.StatusNotFound)
		}

		return ctx.JSON(http.StatusInternalServerError, errors.GetInternalServerError())
	}

	return ctx.NoContent(http.StatusNoContent)

}
