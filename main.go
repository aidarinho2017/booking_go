package main

import (
	"fmt"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func (u User) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Email: %s", u.ID, u.Name, u.Email)
}

type Student struct {
	User
	GPA   float64
	Marks map[string]float64
}

type Master struct {
	Student
	ResearchTopic string
}

type Teacher struct {
	User
	Subject string
}

func (s Student) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Email: %s", s.ID, s.Name, s.Email)
}

func (s Student) GetMarks() {
	for mark, grade := range s.Marks {
		fmt.Println(mark, grade)
	}
}

func (s Student) GetGPA() float64 {
	var GPA float64 = 0
	var count float64 = 0
	for _, grade := range s.Marks {
		count++
		GPA += grade
	}
	return GPA / count
}

func (t Teacher) GetDetails() string {
	return fmt.Sprintf("Id: %d, Name: %s", t.ID, t.Name)
}

func (t Teacher) mark(s *Student, subject string, mark float64) {
	if s.Marks == nil {
		s.Marks = make(map[string]float64)
	}
	s.Marks[subject] = mark
}

var students = map[int]Student{
	1: {User: User{ID: 1, Name: "Alice", Email: "alice@example.com"}, GPA: 3.5, Marks: map[string]float64{"Math": 85}},
	2: {User: User{ID: 2, Name: "Bob", Email: "bob@example.com"}, GPA: 3.8, Marks: map[string]float64{"Physics": 90}},
}

func AddStudent(s *Student) map[int]Student {
	students[s.ID] = *s
	return students
}
func GetStudents() map[int]Student {
	return students
}

func DeleteStudent(id int) {
	if students[id].Marks != nil {
		delete(students[id].Marks, "Math")
	}
}

//
//func GetStudents(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(students)
//}

func main() {

	//student := Student{
	//	User:  User{ID: 1, Name: "Alice", Email: "alice@example.com"},
	//	GPA:   3.67,
	//	Marks: make(map[string]float64),
	//}
	//
	//AddStudent(&student)
	//fmt.Println(GetStudents())
	//
	//teacher := Teacher{
	//	User:    User{ID: 1, Name: "Alice", Email: "alice@example.com"},
	//	Subject: "Math",
	//}
	//
	//teacher.mark(&student, "Literature", 5.45)
	//teacher.mark(&student, "Literature2", 1)
	//teacher.mark(&student, "Literature3", 5.45)
	//teacher.mark(&student, "Literature4", 5.45)
	//fmt.Println(student.GetGPA())
	//
	//student.GetMarks()

	//http.HandleFunc("/students", GetStudents)
	//http.ListenAndServe(":8080", nil)

	var s = [3]string{"aa", "bb", "cc"}

	for i, w := range s {
		fmt.Println(i, w)
	}
}
