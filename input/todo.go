package input

type InputIDTodo struct {
	ID int `uri:"id" binding:"required"`
}

type TodoInput struct {
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

//Generated by Micagen at 11 Desember 2021
