package controllers

import (
	"awesomeProject/src/models"
	"fmt"
)
func CreateNewStudent()  {
	var s = models.Student{24,"Amani","test"}
	fmt.Println(models.AddStudent(s))
	fmt.Println(s)
}