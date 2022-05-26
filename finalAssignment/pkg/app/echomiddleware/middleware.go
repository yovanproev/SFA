package login

import (
	"context"
	"crypto/subtle"
	passwordHash "final/pkg/app"
	handleErrors "final/pkg/app/errors"
	"final/pkg/app/users"
	"final/pkg/sqlc/db"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"tawesoft.co.uk/go/dialog"
)

func Authenticate(q *db.Queries, e handleErrors.Error) middleware.BasicAuthValidator {
	return func(username, password string, c echo.Context) (bool, error) {

		// check if the user exists in the DB
		user, err := q.GetUserByUsername(context.Background(), username)
		if err != nil {
			log.Println(e.DatabaseError, err)
		}

		// if no username is provided
		if username == "" {
			log.Println("No user provided")
			dialog.Alert("You must provide a username!")
			return false, nil
		}

		// create new user if it doesn't exist
		if user.Username != username {
			users.CreateUser(q, username, password)

			// check for the newly created user and delete the double user creation
			user2, err := q.GetUserByUsername(context.Background(), username)
			if err != nil {
				log.Println(e.DatabaseError, err)
			}

			err2 := q.DeleteUserById(context.Background(), user2.ID+1)
			if err2 != nil {
				log.Println(e.DatabaseError, err)
			}

			return true, nil
		}

		// hash password
		hashPassword, err := passwordHash.HashPassword(password)
		if err != nil {
			fmt.Println(err)
		}
		checkPass := passwordHash.CheckPasswordHash(password, hashPassword)

		// if username exists and the password is incorrect
		if user.Username == username && !checkPass {
			log.Println("You password is incorrect, try again!")
			dialog.Alert("You password is incorrect, try again!")
			return false, e.InvalidCredentialsError
		}

		// if username and password are a match
		if subtle.ConstantTimeCompare([]byte(username), []byte(user.Username)) == 1 &&
			checkPass {
			return true, nil
		}

		return false, e.InvalidCredentialsError
	}
}
