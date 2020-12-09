package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/fiber"
	"github.com/hugoBarrios/EYProject/backend/database"
	"github.com/hugoBarrios/EYProject/backend/student"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Routes config
func setupRoutes(app *fiber.App) {
	
	// Static files (Nuxt app frontend)
	app.Static("/", "../frontend/dist")

	app.Get("/api/v1/student", student.GetStudents)
	app.Get("/api/v1/student/:id", student.GetStudent)
	app.Post("/api/v1/student", student.NewStudent)
	app.Delete("/api/v1/student/:id", student.DeleteStudent)
}

// Database
func initDatabase() {
	// Connect database
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "students.db")
	if err != nil {
		panic("Faild connect to database")
	}
	fmt.Print("\nDatabase connection successfully")
	// Migrate database
	database.DBConn.AutoMigrate(&student.Student{})

	fmt.Print("\nDatabase migrate")
}

func main() {
	// Fiber init app
	app := fiber.New()
	app.Use(cors.New())
	// Init database
	initDatabase()
	
	setupRoutes(app)
	
	log.Fatal(app.Listen(":3300"))
	
	defer database.DBConn.Close()
}
