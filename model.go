package main

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

// MODEL DATA STRUCT
type Todo struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdateAt    string `json:"updated_at"`
}

type Data []Todo // state variabel penyimpan data struct
