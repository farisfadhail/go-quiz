package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-quiz/handler"
	"go-quiz/middleware"
)

func RouteInit(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to Go Quiz API")
	})

	// Auth Routes
	api.Post("/login", handler.LoginHandler).Name("login")
	api.Post("/register", handler.UserHandlerCreate).Name("register")

	// User Routes
	user := api.Group("/user", middleware.Authenticated, middleware.IsAdmin)
	user.Get("/", handler.UserHandlerGetAll).Name("user.index")
	user.Get("/:userId", handler.UserHandlerShowById).Name("user.show")
	user.Put("/:userId", handler.UserHandlerUpdate).Name("user.update")
	user.Delete("/:userId", handler.UserHandlerDelete).Name("user.destroy")
	user.Put("/:userId/update-email", handler.UserHandlerUpdateEmail, middleware.Authenticated).Name("user.update-email")

	// Question Routes
	question := api.Group("/question", middleware.Authenticated)
	question.Get("/", handler.QuestionHandlerGetAll).Name("question.index")
	question.Post("/create", handler.QuestionHandlerCreate).Name("question.store")
	question.Get("/:questionId", handler.QuestionHandlerShowById).Name("question.show")
	question.Put("/:questionId/update", handler.QuestionHandlerUpdate).Name("question.update")
	question.Delete("/:questionId", handler.QuestionHandlerDelete).Name("question.destroy")

	// Answer Routes
	answer := api.Group("/answer", middleware.Authenticated)
	answer.Get("/", handler.AnswerHandlerGetAll).Name("answer.index")
	answer.Post("/create", handler.AnswerHandlerCreate).Name("answer.store")
	answer.Get("/:answerId", handler.QuestionHandlerShowById).Name("answer.show")
	answer.Put("/:answerId/update", handler.AnswerHandlerUpdate).Name("answer.update")
	answer.Delete(":answerId", handler.AnswerHandlerDelete).Name("answer.destroy")
}
