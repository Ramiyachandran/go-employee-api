package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

var employees = []Employee{
	{ID: 1, Name: "Ramiya", Role: "Developer"},
	{ID: 2, Name: "Haju", Role: "Tester"},
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Employee API")

}
func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	var newEmployee Employee
	err := json.NewDecoder(r.Body).Decode(&newEmployee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	employees = append(employees, newEmployee)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(newEmployee)
}
func updateEmployee(w http.ResponseWriter, r *http.Request) {
	var updatedEmployee Employee
	err := json.NewDecoder(r.Body).Decode(&updatedEmployee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, emp := range employees {
		if emp.ID == updatedEmployee.ID {
			employees[i] = updatedEmployee
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(updatedEmployee)

			return
		}

	}
	http.Error(w, "Employee not found", http.StatusNotFound)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	var empToDelete Employee
	err := json.NewDecoder(r.Body).Decode(&empToDelete)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, emp := range employees {
		if emp.ID == empToDelete.ID {
			employees = append(employees[:i], employees[i+1:]...)
			w.Header().Set("Content-Type", "application/json")

			w.Write([]byte("Employee deleted successfully"))

			return
		}

	}
	http.Error(w, "Employee not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/add", addEmployee)
	http.HandleFunc("/update", updateEmployee)
	http.HandleFunc("/delete", deleteEmployee)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
