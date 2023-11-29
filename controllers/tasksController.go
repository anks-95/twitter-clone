package controllers

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// User struct to represent the data structure
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NewUser handles the signup logic
func NewUser(c *fiber.Ctx, db *sql.DB) error {
	// Parse JSON from the request body into a User struct
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Insert data into the database
	sqlStatement := "INSERT INTO users (name,email,password) VALUES (?,?,?)"
	_, err := db.Exec(sqlStatement, user.Name, user.Email, user.Password)
	if err != nil {
		fmt.Println("Error inserting user into the database:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.JSON(fiber.Map{"message": "User created successfully"})

}

// func LoggedInUser(c *fiber.Ctx, db *sql.DB) error {
// 	// Parse JSON from the request body into a User struct
// 	var user User
// 	if err := c.BodyParser(&user); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": "Invalid request payload"})
// 	}

// 	// Insert data into the database
// 	sqlStatement := "SELSECT * FROM users WHERE `email`=? and `password`=?"
// 	_, err := db.Exec(sqlStatement, user.Email, user.Password)
// 	if err != nil {
// 		fmt.Println("Error inserting user into the database:", err)
// 		return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
// 	}
// 	return c.JSON(fiber.Map{"message": "success"})

// }

// //Manages CRUD (Create, Read, Update, Delete) operations for tasks.
// //FetchTasks: Retrieves all tasks from the database and returns them as JSON.
// //CreateTask: Parses the request body to create a new task and returns it as JSON.
// //FetchTask: Retrieves a specific task based on the provided ID and returns it as JSON.
// //DeleteTask: Deletes a task based on the provided ID and returns a success message as JSON.

// package controllers

// import (
// 	"twitter/initializers"
// 	"twitter/models"

// 	"github.com/gofiber/fiber/v2"
// )

// func FetchTasks(c *fiber.Ctx) error {
// 	// Get all tasks
// 	var tasks []models.Task
// 	initializers.DB.Order("created_at desc").Find(&tasks)

// 	return c.JSON(fiber.Map{
// 		"tasks": tasks,
// 	})
// }

// func CreateTask(c *fiber.Ctx) error {
// 	// Parse body
// 	var body struct {
// 		Title string
// 		Body  string
// 	}

// 	if err := c.BodyParser(&body); err != nil {
// 		return err
// 	}

// 	// Create record
// 	task := models.Task{Title: body.Title, Body: body.Body}
// 	result := initializers.DB.Create(&task)

// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	// return it
// 	return c.JSON(fiber.Map{
// 		"task": task,
// 	})
// }

// func FetchTask(c *fiber.Ctx) error {
// 	// Fetch task
// 	var task models.Task
// 	taskId := c.Params("id")

// 	result := initializers.DB.First(&task, taskId)

// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	// return it
// 	return c.JSON(fiber.Map{
// 		"task": task,
// 	})
// }

// func DeleteTask(c *fiber.Ctx) error {
// 	taskId := c.Params("id")
// 	result := initializers.DB.Delete(&models.Task{}, taskId)

// 	if result.Error != nil {
// 		return result.Error
// 	}

//		// return it
//		return c.JSON(fiber.Map{
//			"success": "Task deleted",
//		})
//	}
