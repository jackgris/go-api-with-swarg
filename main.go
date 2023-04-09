package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/todo", getAllTodos)
	app.Get("/todo/:id", getTodoByID)
	log.Fatal(app.Listen(":3000"))
}

func getAllTodos(c *fiber.Ctx) error {
	c.Status(http.StatusOK)
	return c.JSON(todoList)
}

func getTodoByID(c *fiber.Ctx) error {
	ID := c.Params("id")

	// loop through todoList and return item with matching ID
	for _, todo := range todoList {
		if todo.ID == ID {
			c.Status(http.StatusOK)
			return c.JSON(todo)
		}
	}

	// return error message if todo is not found
	r := message{"todo not found"}
	c.Status(http.StatusNotFound)
	return c.JSON(r)
}

// todo represents data about a task in the todo list
type todo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

// message represents request response with a message
type message struct {
	Message string `json:"message"`
}

// todo slice to seed todo list data
var todoList = []todo{
	{"1", "Learn Go"},
	{"2", "Build an API with Go"},
	{"3", "Document the API with swag"},
}