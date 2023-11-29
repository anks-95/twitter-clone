//This controller handles requests related to the home page.
//It uses the Fiber framework to render the "index" view.

package controllers

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}
