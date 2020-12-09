package student

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hugoBarrios/EYProject/backend/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Student table database struct
type Student struct {
	gorm.Model
	Name         string `json:"Name"`
	Age          int    `json:"Age"`
	InitialLevel string `json:"InitialLevel"`
	CurrentLevel string `json:"CurrentLevel"`
}

// GetStudents - get all students in the table student
func GetStudents(c *fiber.Ctx) error {
	db := database.DBConn
	var students []Student
	db.Find(&students)
	return c.JSON(students)
}

// GetStudent - get single student
func GetStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var student Student
	if student.Name == "" {
		return c.Status(500).SendString("No student found")
	}
	db.Find(&student, id)
	return c.JSON(student)
}

// NewStudent create a new student
func NewStudent(c *fiber.Ctx) error{
	db := database.DBConn
	student := new(Student)
	if err := c.BodyParser(student); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&student)
	return c.JSON(student)

	// db := database.DBConn
	// student := new(Student)
	// if err := c.BodyParser(student); err != nil {
	// 	return c.Status(503).SendString(err)
	// }
	// db.Create(&student)
	// return c.JSON(student)
}

// DeleteStudent delete a student
func DeleteStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var student Student
	db.First(&student, id)
	if student.Name == "" {
		return c.Status(500).SendString("No student found")
	}
	db.Delete(&student)
	return c.SendString("Student successfully deleted")
}
