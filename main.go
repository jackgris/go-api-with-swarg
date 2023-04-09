package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/jackgris/go-api-with-swarg/docs"
)

//	@title			Go + Fiber Todo API
//	@version		1.0
//	@description	Sample todo server. You can visit the GitHub repository at https://github.com/jackgris/go-api-with-swarg

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host						localhost:3000
//	@BasePath					/
//	@query.collection.format	multi
func main() {
	app := fiber.New()
	app.Get("/todo", getAllTodos)
	app.Get("/todo/:id", getTodoByID)
	app.Post("/todo", createTodo)
	app.Delete("/todo/:id", deleteTodo)

	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Fatal(app.Listen(":3000"))
}

//	@Summary	get all items in the todo list
//	@ID			get-all-todos
//	@Produce	json
//	@Success	200	{object}	todo
//	@Router		/todo [get]
func getAllTodos(c *fiber.Ctx) error {
	c.Status(http.StatusOK)
	return c.JSON(todoList)
}

//	@Summary	get a todo item by ID
//	@ID			get-todo-by-id
//	@Produce	json
//	@Param		id	path		string	true	"todo ID"
//	@Success	200	{object}	todo
//	@Failure	404	{object}	message
//	@Router		/todo/{id} [get]
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

//	@Summary	add a new item to the todo list
//	@ID			create-todo
//	@Produce	json
//	@Param		data	body		todo	true	"todo data"
//	@Success	200		{object}	todo
//	@Failure	400		{object}	message
//	@Router		/todo [post]
func createTodo(c *fiber.Ctx) error {
	var newTodo todo

	// bind the received JSON data to newTodo
	if err := c.BodyParser(&newTodo); err != nil {
		r := message{"an error occurred while creating todo"}
		c.Status(http.StatusBadRequest)
		return c.JSON(r)
	}

	// add the new todo item to todoList
	todoList = append(todoList, newTodo)
	c.Status(http.StatusCreated)
	return c.JSON(newTodo)
}

//	@Summary	delete a todo item by ID
//	@ID			delete-todo-by-id
//	@Produce	json
//	@Param		id	path		string	true	"todo ID"
//	@Success	200	{object}	todo
//	@Failure	404	{object}	message
//	@Router		/todo/{id} [delete]
func deleteTodo(c *fiber.Ctx) error {
	ID := c.Params("id")

	// loop through todoList and delete item with matching ID
	for index, todo := range todoList {
		if todo.ID == ID {
			todoList = append(todoList[:index], todoList[index+1:]...)
			r := message{"successfully deleted todo"}
			c.Status(http.StatusOK)
			return c.JSON(r)
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
