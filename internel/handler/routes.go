package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "timertracker/docs"
	"timertracker/internel/handler/tasks"
	"timertracker/internel/handler/users"
	"timertracker/internel/handler/works"
	"timertracker/internel/service"
)

func RegisterHandlers(app *fiber.App, serviceContext service.ServiceContext) {
	v1 := app.Group("/v1")

	//GET Запросы
	v1.Get("/users", func(ctx *fiber.Ctx) error {
		return users.GetUsersHandler(ctx, serviceContext)
	})
	v1.Get("/user/:userId", func(ctx *fiber.Ctx) error {
		return users.GetUserHandler(ctx, serviceContext)
	})

	v1.Get("/tasks", func(ctx *fiber.Ctx) error {
		return tasks.GetTasksHandler(ctx, serviceContext)
	})
	v1.Get("/tasks/:taskID", func(ctx *fiber.Ctx) error {
		return tasks.GetTaskHandler(ctx, serviceContext)
	})

	v1.Get("/works", func(ctx *fiber.Ctx) error {
		return works.GetWorksHandler(ctx, serviceContext)
	})
	v1.Get("/work/:workID", func(ctx *fiber.Ctx) error {
		return works.GetWorkHandler(ctx, serviceContext)
	})
	v1.Get("/work/period/:userID", func(ctx *fiber.Ctx) error {
		return works.GetPeriodWorkHandler(ctx, serviceContext)
	})

	//POST Запросы
	v1.Post("/user", func(ctx *fiber.Ctx) error {
		return users.AddUserHandler(ctx, serviceContext)
	})

	v1.Post("/task", func(ctx *fiber.Ctx) error {
		return tasks.AddTaskHandler(ctx, serviceContext)
	})

	v1.Post("/work", func(ctx *fiber.Ctx) error {
		return works.AddWorkHandler(ctx, serviceContext)
	})

	//PUT Запросы
	v1.Put("/user/edit", func(ctx *fiber.Ctx) error {
		return users.EditUserHandler(ctx, serviceContext)
	})

	v1.Put("/tasks/edit", func(ctx *fiber.Ctx) error {
		return tasks.EditTaskHandler(ctx, serviceContext)
	})

	v1.Put("/work/edit", func(ctx *fiber.Ctx) error {
		return works.EditWorkHandler(ctx, serviceContext)
	})

	//DELETE Запросы
	v1.Delete("/user/delete/:userId", func(ctx *fiber.Ctx) error {
		return users.DeleteUserHandler(ctx, serviceContext)
	})

	v1.Delete("/tasks/delete/:taskID", func(ctx *fiber.Ctx) error {
		return tasks.DeleteTaskHandler(ctx, serviceContext)
	})

	v1.Delete("/work/delete/:workID", func(ctx *fiber.Ctx) error {
		return works.DeleteWorkHandler(ctx, serviceContext)
	})

	//SWAGGER
	app.Get("/swagger/*", swagger.HandlerDefault)
}
