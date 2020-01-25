package models

type Student struct {
	Age int
	Name string
	Username string
}

func AddStudent(s Student) Student  {
	var student = Student{25,"Hedi","Stiv"}
	return student
}
