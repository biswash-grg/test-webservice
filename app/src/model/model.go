package model

/*
Students A list of student
*/
var Students = []*Student{
	{
		ID:        1,
		Firstname: "Mike",
		Lastname:  "Tyson",
	},
	{
		ID:        2,
		Firstname: "Conor",
		Lastname:  "McGregor",
	},
}

/*
Student struct representing student model
*/
type Student struct {
	ID        int
	Firstname string
	Lastname  string
}
