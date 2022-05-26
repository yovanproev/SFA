package handleErrors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Error struct {
	DatabaseError           string
	StatusNotFoundError     string
	InvalidCredentialsError *echo.HTTPError
	DatabaseInit            string
	HTTPRequest             string
	HTTPResponse            string
	JSONMarshalling         string
}

func (e Error) SetErrors() Error {
	e.DatabaseError = "Problem with DB!, %+v"
	e.StatusNotFoundError = "Entry with the provided id not found!"
	e.InvalidCredentialsError = echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
	e.DatabaseInit = "Database cannot be created, %+v"
	e.HTTPRequest = "Request denied %+v"
	e.HTTPResponse = "Response unavailable %+v"
	e.JSONMarshalling = "Cannot unmarshal %+v"

	return e
}
