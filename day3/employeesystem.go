package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

// creating a slice
var (
	employees   = make(map[int]Employee)
	employeeIDs = []int{}
	nextID      = 1
	reader      = bufio.NewReader(os.Stdin)
)

func Employeesystem() {
	for {

		fmt.Println(" __employee management system__")
		fmt.Println("1.Add Employee")
		fmt.Println("2.View all Employees")
		fmt.Println("3.Update employee salary")
		fmt.Println("4.Search employee by name")
		fmt.Println("5.Delete employee")
		fmt.Println("6.Exit")
		fmt.Println("Chooose an option")

		choice := readLine()
		switch choice {
		case "1":
			addemployee()
		case "2":
			viewEmployees()
		case "3":
			updateSalary()
		case "4":
			searchByName()
		case "5":
			deleteEmployee()
		case "6":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option — try again.")

		}
	}
}
func readLine() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
func readInt(prompt string) (int, error) {
	fmt.Print(prompt)
	s := readLine()
	return strconv.Atoi(s)
}
func readFloat(prompt string) (float64, error) {
	fmt.Print(prompt)
	s := readLine()
	return strconv.ParseFloat(s, 64)
}
func addemployee() {
	fmt.Println("__Add employee__")
	fmt.Println("Name:")
	name := readLine()
	if name == "" {
		fmt.Println("name cannot be empty")
		return
	}
	age, err := readInt("Age:")
	if err != nil || age <= 0 {
		fmt.Println("Invalid age.")
		return
	}
	salary, err := readFloat("Salary:")
	if err != nil || salary < 0 {
		fmt.Println("Invalid salary")
		return
	}
	id := nextID
	nextID++

	e := Employee{
		ID:     id,
		Name:   name,
		Age:    age,
		Salary: salary,
	}
	employees[id] = e
	employeeIDs = append(employeeIDs, id)
	fmt.Printf("Employee added with %d\n", id)
}
func viewEmployees() {
	fmt.Println("__All employes__")
	if len(employeeIDs) == 0 {
		fmt.Println("No employee found")
		return
	}
	fmt.Printf("%-5s %-20s %-5s6 %-10s\n", "ID", "Name", "Age", "Salary")
	fmt.Println(strings.Repeat("-", 45))
	for _, id := range employeeIDs {
		if e, ok := employees[id]; ok {
			fmt.Printf("%-5d %-20s %-5d %-10.2f\n", e.ID, e.Name, e.Age, e.Salary)
		}
	}
}
func updateSalary() {
	fmt.Println("__Update sakary__")
	id, err := readInt("Enter employee ID:")
	if err != nil {
		fmt.Println("Invalid id")
		return
	}
	e, ok := employees[id]
	if !ok {
		fmt.Println("Employee not found")
		return
	}
	fmt.Printf("current salary for %s(ID %d): %2.f\n", e.Name, e.ID, e.Salary)
	Newsalary, err := readFloat("Enter new salary")
	if err != nil || Newsalary < 0 {
		fmt.Println("Invalid salary.")
		return
	}
	e.Salary = Newsalary
	employees[id] = e
	fmt.Println("Salary updated")
}
func searchByName() {
	fmt.Println("__Search by name__")
	fmt.Println("Enter name (partial allowed, case-intensive):")
	query := strings.ToLower(readLine())
	if query == "" {
		fmt.Println("Search query cannot be empty")
		return
	}
	found := false
	fmt.Printf("%-5s %-20s %-5s %-10s\n", "ID", "Name", "Age", "Salary")
	fmt.Println(strings.Repeat("-", 45))
	for _, e := range employees {
		if strings.Contains(strings.ToLower(e.Name), query) {
			fmt.Printf("%-5d %-20s %-5d %-10.2f\n", e.ID, e.Name, e.Age, e.Salary)
			found = true

		}
	}
	if !found {
		fmt.Println("no employee found")
	}
}
func deleteEmployee() {
	fmt.Println("__ Delete employee")
	id, err := readInt("Enter employee ID:")
	if err != nil {
		fmt.Println("Invalid ID ")
		return
	}
	_, ok := employees[id]
	if !ok {
		fmt.Println("Employee not fount.")
		return
	}
	delete(employees, id)
	for i, v := range employeeIDs {
		if v == id {
			employeeIDs = append(employeeIDs[:i], employeeIDs[i+1:]...)
			break
		}
	}
	fmt.Println("Employee deleted.")
}
