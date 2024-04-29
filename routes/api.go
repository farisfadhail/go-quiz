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

	// Admin Routes
	admin := api.Group("/admin/user", middleware.Authenticated, middleware.IsAdmin)
	admin.Get("/", handler.UserHandlerGetAll).Name("user.index")
	admin.Delete("/delete?user_id=:userId", handler.UserHandlerDelete).Name("user.destroy")

	// User Routes
	user := api.Group("/user", middleware.Authenticated)
	user.Get("/:userId", handler.UserHandlerShowById).Name("user.show")
	user.Put("/update?user_id=:userId", handler.UserHandlerUpdate).Name("user.update")
	user.Put("/update-email?user_id=:userId", handler.UserHandlerUpdateEmail, middleware.Authenticated).Name("user.update-email")

	// Survey Creator Routes
	surveyCreator := api.Group("/survey/question", middleware.Authenticated, middleware.IsSurveyCreator)
	surveyCreator.Post("/create", handler.QuestionHandlerCreate).Name("question.store")
	surveyCreator.Put("/update/question_id=:questionId", handler.QuestionHandlerUpdate).Name("question.update")
	surveyCreator.Delete("/delete?question_id=:questionId", handler.QuestionHandlerDelete).Name("question.destroy")

	// Question Routes
	question := api.Group("/question", middleware.Authenticated)
	question.Get("/", handler.QuestionHandlerGetAll).Name("question.index")
	question.Get("/:questionId", handler.QuestionHandlerShowById).Name("question.show")

	// Answer Routes
	answer := api.Group("/answer", middleware.Authenticated)
	answer.Post("/create", handler.AnswerHandlerCreate).Name("answer.store")
	answer.Get("/:answerId", handler.QuestionHandlerShowById).Name("answer.show")
	answer.Put("/update?answer_id=:answerId", handler.AnswerHandlerUpdate).Name("answer.update")
	answer.Delete("/delete?answer_id=:answerId", handler.AnswerHandlerDelete).Name("answer.destroy")

	// Admin Answer Routes
	answer.Get("/", handler.AnswerHandlerGetAll, middleware.IsAdmin).Name("answer.index")
}
