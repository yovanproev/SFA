package final

import (
	"context"
	handlers "final/cmd/gin/handlers"
	"final/cmd/gin/sqlc/db"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
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

func Login(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password := handlers.GetUserAndPassFromHeader(c)

		// check if the user exists in the DB
		user, err := q.GetUserByUsername(context.Background(), username)
		if err != nil {
			fmt.Println(err)
		}

		// create new user if it doesn't exist
		if user.Username != username {
			// hash password
			hashPassword, err := hashPassword(password)
			if err != nil {
				fmt.Println(err)
			}

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
		}
	}
}

func GinAccounts(q *db.Queries) gin.Accounts {
	var account = make(map[string]string)

	account["jovan"] = "proev"
	account["proev"] = "jovan"
	account["jov"] = "pro"

	return account
}
