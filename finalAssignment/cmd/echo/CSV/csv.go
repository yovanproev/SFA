package final

import (
	"context"
	"encoding/csv"
	"final/cmd/echo/sqlc/db"
	"fmt"
	"log"
	"os"
)

func GetTasksByUser(q *db.Queries, user db.User) [][]string {
	listTasks, err2 := q.ListTasksByUserId(context.Background(), user.ID)
	if err2 != nil {
		fmt.Println(err2)
	}

	records := [][]string{}
	tasks := []string{}
	tasks = append(tasks, "Tasks")
	for _, task := range listTasks {
		if user.ID == task.Userid {
			tasks = append(tasks, task.Text)
		}
	}
	records = append(records, tasks)

	return records
}

func OpenCSV(records [][]string, filename string) {
	csvFile, err := os.Create(filename)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	w := csv.NewWriter(csvFile)
	defer w.Flush()

	for _, record := range records[0] {
		row := []string{record}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

}
