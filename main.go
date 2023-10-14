package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Student struct {
	ID     int
	Name   string
	Domain string
}

func main() {
	var choice int
	setupDatabase()
	fmt.Println("Welcome to the Student Database Management System")
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Insert Student Record")
		fmt.Println("2. View Student Records")
		fmt.Println("3. Update Student Record")
		fmt.Println("4. Delete Student Record")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		fmt.Scan(&choice)
		switch choice {
		case 1:
			Insert()
		case 2:
			View()
		case 3:
			Update()
		case 4:
			Delete()
		case 5:
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func setupDatabase() {
	dsn := "user=ar dbname=studinfo password=1234 host=localhost port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db.AutoMigrate(&Student{})
}

// INSERT
func Insert() {
	var student Student

	fmt.Println("\nEnter Student Information:")
	fmt.Print("Student ID: ")
	fmt.Scan(&student.ID)

	fmt.Print("Student Name: ")
	fmt.Scan(&student.Name)

	fmt.Print("Student Domain: ")
	fmt.Scan(&student.Domain)

	db.Create(&student)
	fmt.Println("Student record inserted successfully!")
}

// VIEW
func View() {
	var students []Student
	db.Find(&students)
	fmt.Println("\nStudent Records:")
	fmt.Println("ID  Name       Domain")
	for _, student := range students {
		fmt.Printf("%-3d %-10s %s\n", student.ID, student.Name, student.Domain)
	}
}

// UPDATE
func Update() {
	var id int
	var choice int

	View() // Display current student records.

	fmt.Print("\nEnter the ID of the student you want to update: ")
	fmt.Scan(&id)

	fmt.Println("Choose what to update:")
	fmt.Println("1. Update Name")
	fmt.Println("2. Update Domain")
	fmt.Println("3. Cancel")
	fmt.Print("Enter your choice: ")

	fmt.Scan(&choice)
	switch choice {
	case 1:
		updatename(id)
	case 2:
		updatelang(id)
	case 3:
		fmt.Println("Update canceled.")
	default:
		fmt.Println("Invalid choice. Update canceled.")
	}
}

func updatename(id int) {
	var newName string
	fmt.Println("Enter new name:")
	fmt.Scan(&newName)

	db.Model(&Student{}).Where("id = ?", id).Update("name", newName)
	fmt.Println("Data updated")
}

func updatelang(id int) {
	var newDomain string
	fmt.Println("Enter new domain:")
	fmt.Scan(&newDomain)

	db.Model(&Student{}).Where("id = ?", id).Update("domain", newDomain)
	fmt.Println("Data updated")
}

// DELETE
func Delete() {
	var id int

	View() // Display current student records.

	fmt.Print("\nEnter the ID of the student you want to delete: ")
	fmt.Scan(&id)

	db.Delete(&Student{}, id)
	fmt.Println("Student record deleted successfully!")
}
