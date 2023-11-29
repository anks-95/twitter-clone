// // controllers.go

// package controllers

// import (
// 	"twitter/models"

// 	"github.com/gofiber/fiber/v2"
// )

// Signup handles user signup
// func Signup(c *fiber.Ctx) error {
// 	var user models.User
// 	if err := c.BodyParser(&user); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
// 	}

// 	// Add logic to create a new user in the database (use models.CreateUser)

// 	return c.SendString("Signup successful")
// }

// // Login handles user login
// func Login(c *fiber.Ctx) error {
// 	var user models.User
// 	if err := c.BodyParser(&user); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
// 	}

// 	// Add logic to authenticate the user (use models.GetUserByEmail)

// 	return c.SendString("Login successful")
// }

package controllers
