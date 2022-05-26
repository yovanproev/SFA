// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"time"
)

type List struct {
	ID     int32  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Userid int32  `json:"userid,omitempty"`
}

type Task struct {
	ID        int32       `json:"id,omitempty"`
	Text      string      `json:"text,omitempty"`
	Listid    int32       `json:"listid,omitempty"`
	Userid    int32       `json:"userid,omitempty"`
	Completed interface{} `json:"completed,omitempty"`
}

type User struct {
	ID        int32     `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
	Datestamp time.Time `json:"datestamp,omitempty"`
}