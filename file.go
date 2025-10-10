package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Employee represents an employee record.
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// LoadEmployees loads employee data from a CSV file.
func LoadEmployees(filename string) ([]Employee, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var employees []Employee
	for _, record := range records[1:] { // skip header
		id, _ := strconv.Atoi(record[0])
		salary, _ := strconv.ParseFloat(record[2], 64)
		employees = append(employees, Employee{
			ID:     id,
			Name:   record[1],
			Salary: salary,
		})
	}
	return employees, nil
}

// GiveRaise increases salary by percentage for all employees.
func GiveRaise(employees []Employee, percent float64) {
	for i := range employees {
		employees[i].Salary += employees[i].Salary * percent / 100
	}
}

// SaveEmployees saves employee data to a CSV file.
func SaveEmployees(filename string, employees []Employee) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Name", "Salary"})
	for _, e := range employees {
		writer.Write([]string{
			strconv.Itoa(e.ID),
			e.Name,
			fmt.Sprintf("%.2f", e.Salary),
		})
	}
	return nil
}

func call() {
	employees, err := LoadEmployees("employees.csv")
	if err != nil {
		fmt.Println("Error loading employees:", err)
		return
	}

	GiveRaise(employees, 10) // Give 10% raise

	err = SaveEmployees("employees_updated.csv", employees)
	if err != nil {
		fmt.Println("Error saving employees:", err)
		return
	}

	fmt.Println("Employee salaries updated successfully.")
}
