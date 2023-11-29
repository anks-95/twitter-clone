// main.go

package main

import (
	"database/sql"
	"fmt"
	"os"
	"twitter/controllers"
	"twitter/initializers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

var db *sql.DB // Declare a global variable to store the database connection
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {

	initializers.LoadEnvVariables()

	// Establish a database connection
	var err error
	db, err = initializers.ConnectToDatabase()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// Load templates
	engine := html.New("./views", ".html")

	// Create app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Configure app
	app.Static("/", "./public")
	app.Use(cors.New())

	app.Post("/signup", func(c *fiber.Ctx) error {
		if err := controllers.NewUser(c, db); err != nil {
			fmt.Printf("Error creating user: %v\n", err)
			return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
		}
		return nil
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		if err := controllers.LoggedInUser(c, db); err != nil {
			fmt.Printf("User not logged in: %v\n", err)
			return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
		}
		return nil
	})

	// Frontend routes
	frontendRoutes := []string{
		"/",
		"/about",
		"/signup",
		"/login",
	}

	for _, route := range frontendRoutes {
		app.Get(route, controllers.Home)
	}

	// Start app
	app.Listen(":" + os.Getenv("PORT"))
}

// func init() {
// 	initializers.LoadEnvVariables()
// 	var err error
// 	db, err = initializers.ConnectToDatabase()
// 	if err != nil {
// 		fmt.Println("Error connecting to the database:", err)
// 		return
// 	}
// 	defer db.Close() // Close the database connection when the program exits

// }

//Example: Executing a SQL statement using the imported query
// _, err = db.Exec(helpers.InsertTaskQuery(), 4, "executed", "nice")
// if err != nil {
// 	fmt.Println("Error executing SQL statement:", err)
// 	return
// }

//creating a database connection

// Routing
// app.Get("/api/tasks", controllers.FetchTasks)
// app.Post("/api/tasks", controllers.CreateTask)
// app.Get("/api/tasks/:id", controllers.FetchTask)
// app.Delete("/api/tasks/:id", controllers.DeleteTask)

// New routes for user authentication
// app.Post("/api/signup", controllers.Signup)
// app.Post("/api/login", controllers.Login)

// Example: Querying data using the imported query
// rows, err := db.Query(helpers.SelectAllTasksQuery())
// if err != nil {
// 	fmt.Println("Error querying data:", err)
// 	return
// }
// defer rows.Close()

// Iterate through the result set
// for rows.Next() {
// 	var id int
// 	var title string
// 	var body string
// 	err := rows.Scan(&id, &title, &body)
// 	if err != nil {
// 		fmt.Println("Error scanning row:", err)
// 		return
// 	}
// 	fmt.Printf("ID: %d, TITLE: %s, DEMO: %s\n", id, title, body)
// }
