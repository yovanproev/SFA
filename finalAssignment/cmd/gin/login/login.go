package final

import (
	"context"
	"crypto/subtle"
	"encoding/base64"
	"final/cmd/gin/sqlc/db"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func Login(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		userAndPass := c.Request.Header["Authorization"]
		if userAndPass != nil {
			trim := strings.Trim(userAndPass[0], "Basic")
			trim = strings.Trim(trim, " ")

			rawDecodedText, err := base64.StdEncoding.DecodeString(trim)
			if err != nil {
				panic(err)
			}

			splitToUserAndPass := strings.Split(string(rawDecodedText), ":")
			username := splitToUserAndPass[0]
			password := splitToUserAndPass[1]

			// check if the user exists in the DB
			user, err := q.GetUserByUsername(context.Background(), username)
			if err != nil {
				fmt.Println(err)
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

				// create a new login data for the user
				_, err3 := q.UpdateUsersById(context.Background(), user.ID)
				if err3 != nil {
					fmt.Println(err3)
				}
			}

			checkPass := checkPasswordHash(password, hashPassword)

			// if username and password are a match
			if subtle.ConstantTimeCompare([]byte(username), []byte(user.Username)) == 1 &&
				checkPass {

				// update the login
				_, err := q.UpdateUsersById(context.Background(), user.ID)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func GinAccounts(q *db.Queries) gin.Accounts {
	var account = make(map[string]string)

	account["jovan"] = "proev"
	account["proev"] = "jovan"

	return account
}
