package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-quiz/models/entity"
	"go-quiz/models/request"
)

func QuestionHandlerGetAll(ctx *fiber.Ctx) error {
	var questions []entity.Question

	result := db.Debug().Find(&questions)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "FAILED TO GET ALL DATAS",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "SUCCESS GET ALL DATAS",
		"data":    questions,
	})
}

func QuestionHandlerCreate(ctx *fiber.Ctx) error {
	question := new(request.QuestionRequest)
	err := ctx.BodyParser(question)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "BAD REQUEST",
			"error":   err.Error(),
		})
	}

	err = validate.Struct(question)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "BAD REQUEST",
			"error":   err.Error(),
		})
	}

	newQuestion := entity.Question{
		Question: question.Question,
		Points:   question.Points,
	}

	result := db.Debug().Create(&newQuestion)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "FAILED TO STORE DATA",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "SUCCESS STORE DATA",
		"data":    newQuestion,
	})
}

func QuestionHandlerShowById(ctx *fiber.Ctx) error {
	questionId := ctx.Params("questionId")

	var question entity.Question

	result := db.Debug().First(&question, questionId)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "DATA NOT FOUND",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "DATA HAS FOUNDED",
		"data":    question,
	})
}

func QuestionHandlerUpdate(ctx *fiber.Ctx) error {
	questionRequest := new(request.QuestionUpdateRequest)
	err := ctx.BodyParser(questionRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "BAD REQUEST",
		})
	}

	err = validate.Struct(questionRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "BAD REQUEST",
			"error":   err.Error(),
		})
	}

	questionId := ctx.Params("questionId")

	var question entity.Question

	result := db.First(&question, questionId)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "DATA NOT FOUND",
		})
	}

	question.Question = questionRequest.Question
	question.Points = questionRequest.Points

	result = db.Debug().Save(&question)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "FAILED TO UPDATE DATA",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "QUESTION UPDATED",
		"data":    question,
	})
}

func QuestionHandlerDelete(ctx *fiber.Ctx) error {
	questionId := ctx.Params("questionId")

	var question entity.Question

	result := db.Debug().First(&question, questionId)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "DATA NOT FOUND",
		})
	}

	result = db.Debug().Delete(&question, questionId)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "FAILED TO DELETE DATA",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "DATA HAS BEEN DELETED",
	})
}
