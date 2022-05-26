package final

import (
	"bytes"
	"context"
	"encoding/csv"
	handleErrors "final/pkg/app/errors"
	"final/pkg/sqlc/db"

	"log"
)

func GetTasksByUser(q *db.Queries, user db.User, e handleErrors.Error) [][]string {
	listTasks, err := q.ListTasksByUserId(context.Background(), user.ID)
	if err != nil {
		log.Println(e.DatabaseError, err)
	}

	records := [][]string{}
	records = append(records, []string{"Tasks"})
	for _, task := range listTasks {
		if user.ID == task.Userid {
			records = append(records, []string{task.Text})
		}
	}

	return records
}

func CreateBytesFromTasks(records [][]string) string {
	b := new(bytes.Buffer)
	w := csv.NewWriter(b)

	w.WriteAll(records)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	s := b.String()

	return s
}
