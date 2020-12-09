module github.com/hugoBarrios/EYProject

go 1.15

require (
	// github.com/gofiber/fiber v1.14.6
	github.com/gofiber/fiber/v2 v2.2.4 // indirect
	github.com/hugoBarrios/EYProject/backend/database v0.0.0
	github.com/hugoBarrios/EYProject/backend/student v0.0.0
	github.com/jinzhu/gorm v1.9.16 // indirect
)

replace (
	github.com/hugoBarrios/EYProject/backend/database => ./database
	github.com/hugoBarrios/EYProject/backend/student => ./student
)
