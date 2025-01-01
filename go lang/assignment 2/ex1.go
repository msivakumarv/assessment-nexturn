package main

import (
	"errors"
	"fmt"
)

type Emp struct {
	ID   int
	Name string
	Age  int
	Dept string
}

var employees []Emp

func addEmp(id int, name string, age int, dept string) error {
	for _, e := range employees {
		if e.ID == id {
			return errors.New("ID must be unique")
		}
	}
	if age <= 18 {
		return errors.New("Age must be greater than 18")
	}
	employees = append(employees, Emp{
		ID:   id,
		Name: name,
		Age:  age,
		Dept: dept,
	})
	return nil
}

func searchEmpByID(id int) (Emp, error) {
	for _, e := range employees {
		if e.ID == id {
			return e, nil
		}
	}
	return Emp{}, errors.New("Employee not found")
}

func searchEmpByName(name string) (Emp, error) {
	for _, e := range employees {
		if e.Name == name {
			return e, nil
		}
	}
	return Emp{}, errors.New("Employee not found")
}

func listEmpByDept(dept string) []Emp {
	var res []Emp
	for _, e := range employees {
		if e.Dept == dept {
			res = append(res, e)
		}
	}
	return res
}

func countEmpByDept(dept string) int {
	count := 0
	for _, e := range employees {
		if e.Dept == dept {
			count++
		}
	}
	return count
}

func main() {
	addEmp(1, "Siva", 21, "HR")
	addEmp(2, "Shamu", 21, "HR")
	addEmp(3, "Selva", 20, "IT")

	emp, err := searchEmpByID(1)
	if err == nil {
		fmt.Println(emp)
	} else {
		fmt.Println(err)
	}

	itEmps := listEmpByDept("HR")
	for _, e := range itEmps {
		fmt.Println(e)
	}

	fmt.Println("Count in HR:", countEmpByDept("HR"))
}
