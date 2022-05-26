package login

import (
	"final/pkg/sqlc/db"

	"github.com/gin-gonic/gin"
)

func GinAccounts(q *db.Queries) gin.Accounts {
	var account = make(map[string]string)

	account["jovan"] = "proev"
	account["proev"] = "jovan"
	account["jov"] = "pro"

	return account
}
