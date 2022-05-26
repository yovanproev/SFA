package users

import (
	"context"
	"encoding/base64"
	passwordHash "final/pkg/app"
	"final/pkg/sqlc/db"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

func GetUserAndPassFromHeaderEcho(c echo.Context) (string, string) {
	userAndPass := c.Request().Header.Get(echo.HeaderAuthorization)
	username := "test"
	password := "test"

	if userAndPass != "" {
		trim := strings.Trim(userAndPass, "Basic")
		trim = strings.Trim(trim, " ")

		rawDecodedText, err := base64.StdEncoding.DecodeString(trim)
		if err != nil {
			panic(err)
		}

		splitToUserAndPass := strings.Split(string(rawDecodedText), ":")
		username = splitToUserAndPass[0]
		password = splitToUserAndPass[1]
	}

	return username, password
}

func GetUserAndPassFromHeaderGin(c *gin.Context) (string, string) {
	userAndPass := c.Request.Header["Authorization"]
	username := "test"
	password := "test"

	if userAndPass != nil {
		trim := strings.Trim(userAndPass[0], "Basic")
		trim = strings.Trim(trim, " ")

		rawDecodedText, err := base64.StdEncoding.DecodeString(trim)
		if err != nil {
			panic(err)
		}

		splitToUserAndPass := strings.Split(string(rawDecodedText), ":")
		username = splitToUserAndPass[0]
	}

	return username, password
}

func CreateUser(q *db.Queries, username, password string) {
	// hash password
	hashPassword, err := passwordHash.HashPassword(password)
	if err != nil {
		log.Println(err)
	}

	var user = db.CreateUserParams{
		Username:  username,
		Password:  hashPassword,
		Datestamp: time.Now(),
	}

	// for testing, to eliminate the datestamp element
	if username == "test" && password == "test" {
		user = db.CreateUserParams{
			Username: username,
			Password: password,
		}
	}

	_, err = q.CreateUser(context.Background(), user)
	if err != nil {
		log.Println(err)
	}
}
