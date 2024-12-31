package main

import (
	"errors"
	"fmt"
	"strings"
)

type employee struct {
	id         int
	name       string
	age        int
	department string
}

var employees []employee

func add_employee(id int, name string, age int, department string) error {
	for _, emp := range employees {
		if emp.id == id {
			return errors.New("employee id must be unique")
		}
	}
	if age <= 18 {
		return errors.New("employee age must be greater than 18")
	}
	employees = append(employees, employee{id, name, age, department})
	return nil
}

func search_employee(term string) (*employee, error) {
	for _, emp := range employees {
		if fmt.Sprintf("%d", emp.id) == term || strings.EqualFold(emp.name, term) {
			return &emp, nil
		}
	}
	return nil, errors.New("employee not found")
}

func list_by_department(department string) ([]employee, error) {
	var result []employee
	for _, emp := range employees {
		if strings.EqualFold(emp.department, department) {
			result = append(result, emp)
		}
	}
	if len(result) == 0 {
		return nil, errors.New("no employees in this department")
	}
	return result, nil
}

func count_by_department(department string) int {
	count := 0
	for _, emp := range employees {
		if strings.EqualFold(emp.department, department) {
			count++
		}
	}
	return count
}

func main() {
	add_employee(1, "Alice", 25, "HR")
	add_employee(2, "Bob", 30, "IT")
	add_employee(3, "Charlie", 22, "HR")

	if emp, err := search_employee("Alice"); err == nil {
		fmt.Println("Found:", emp)
	} else {
		fmt.Println("Error:", err)
	}

	if list, err := list_by_department("HR"); err == nil {
		fmt.Println("HR Employees:", list)
	} else {
		fmt.Println("Error:", err)
	}

	fmt.Println("IT Employee Count:", count_by_department("IT"))

	if err := add_employee(2, "Eve", 17, "IT"); err != nil {
		fmt.Println("Error:", err)
	}

	if _, err := search_employee("John"); err != nil {
		fmt.Println("Error:", err)
	}

	if _, err := list_by_department("Finance"); err != nil {
		fmt.Println("Error:", err)
	}
}
