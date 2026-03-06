package main

func main() {
	type student struct {
		firstName string
		lastName  string
		age       int
		classes   []string
	}

	var student1 student = student{
		firstName: "John",
		lastName:  "Doe",
		age:       20,
		classes:   []string{"Math", "Science", "History"},
	}

	println("Student Name:", student1.firstName, student1.lastName)
	println("Age:", student1.age)
	println("Classes:")
	for _, class := range student1.classes {
		println("-", class)
	}

	student2 := student{
		firstName: "Jane",
		lastName:  "Smith",
		age:       22,
		classes:   []string{"English", "Art", "Physical Education"},
	}

	println("\nStudent Name:", student2.firstName, student2.lastName)
	println("Age:", student2.age)
	println("Classes:")
	for _, class := range student2.classes {
		println("-", class)
	}

	// Anonymous struct for a guardian
	guardian := struct {
		firstName    string
		lastName     string
		relationship string
	}{
		firstName:    "Alice",
		lastName:     "Johnson",
		relationship: "Mother",
	}

	println("\nGuardian Name:", guardian.firstName, guardian.lastName)
	println("Relationship:", guardian.relationship)
}
