package service

import "github.com/biswash-grg/test-webservice/app/src/model"

/*
GetAllStudents Func to get all students
*/
func GetAllStudents() []*model.Student {
	return model.Students
}
