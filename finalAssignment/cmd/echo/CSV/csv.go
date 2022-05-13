package final

import (
	"context"
	"encoding/csv"
	"final/cmd/echo/sqlc/db"
	"fmt"
	"log"
	"os"
)

func OpenCSV(q *db.Queries, user db.User) {
	csvFile, err := os.Create("tasks.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	tasks, err2 := q.ListTasks(context.Background())
	if err2 != nil {
		fmt.Println(err2)
	}

	records := [][]string{}
	tasksText := []string{}
	tasksText = append(tasksText, "Tasks")
	for _, task := range tasks {
		if user.ID == task.Userid {
			tasksText = append(tasksText, task.Text)
		}
	}
	records = append(records, tasksText)

	w := csv.NewWriter(csvFile)
	defer w.Flush()

	// Using Write
	for _, record := range records[0] {
		row := []string{record}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

}
