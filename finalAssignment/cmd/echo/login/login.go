package final

import (
	"context"
	"crypto/subtle"
	"final/cmd/echo/sqlc/db"
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
	"tawesoft.co.uk/go/dialog"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(q *db.Queries, username, password string) {
	var user = db.CreateUserParams{
		Username:  username,
		Password:  password,
		Datestamp: time.Now(),
	}

	// for testing, to eliminate the datestamp element
	if username == "test" && password == "test" {
		user = db.CreateUserParams{
			Username: username,
			Password: password,
		}
	}

	_, err := q.CreateUser(context.Background(), user)
	if err != nil {
		log.Println(err)
	}
}

func Login(q *db.Queries) middleware.BasicAuthValidator {
	return func(username, password string, c echo.Context) (bool, error) {

		// check if the user exists in the DB
		user, err := q.GetUserByUsername(context.Background(), username)
		if err != nil {
			fmt.Println(err)
		}

		// if no username is provided
		if username == "" {
			log.Println("No user provided")
			dialog.Alert("You must provide a username!")
			return false, nil
		}

		// hash password
		hashPassword, err := hashPassword(password)
		if err != nil {
			fmt.Println(err)
		}

		// create new user if it doesn't exist
		if user.Username != username {
			CreateUser(q, username, hashPassword)

			// check for the newly created user and delete the double user creation
			user2, err := q.GetUserByUsername(context.Background(), username)
			if err != nil {
				fmt.Println(err)
			}
			err2 := q.DeleteUserById(context.Background(), user2.ID+1)
			if err2 != nil {
				fmt.Println(err2)
			}

			return true, nil
		}

		checkPass := checkPasswordHash(password, hashPassword)

		// if username exists and the password is incorrect
		if user.Username == username && !checkPass {
			log.Println("You password is incorrect, try again!")
			dialog.Alert("You password is incorrect, try again!")
			return false, nil
		}

		// if username and password are a match
		if subtle.ConstantTimeCompare([]byte(username), []byte(user.Username)) == 1 &&
			checkPass {
			return true, nil
		}

		return false, nil
	}
}
