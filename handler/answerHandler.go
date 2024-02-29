package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go-quiz/models/entity"
	"go-quiz/models/request"
)

func AnswerHandlerGetAll(ctx *fiber.Ctx) error {
	var answers []entity.Answer

	result := db.Debug().Find(&answers)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "FAILED TO GET ALL DATAS",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "SUCCESS GET ALL DATAS",
		"data":    answers,
	})
}

func AnswerHandlerCreate(ctx *fiber.Ctx) error {
	answer := new(request.AnswerRequest)
	err := ctx.BodyParser(answer)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "BAD REQUEST",
			"error":   err.Error(),
		})
	}

	err = validate.Struct(answer)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "BAD REQUEST",
			"error":   err.Error(),
		})
	}

	marshal, err := json.Marshal(answer.Answers)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "FAILED TO CONVERT DATA ANSWERS",
			"error":   err.Error(),
		})
	}

	newAnswer := entity.Answer{
		UserId:  answer.UserId,
		Answers: string(marshal),
	}

	result := db.Debug().Create(&newAnswer)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "FAILED TO STORE DATA",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "SUCCESS STORE DATA",
		"data":    newAnswer,
	})
}

func AnswerHandlerShowById(ctx *fiber.Ctx) error {
	answerId := ctx.Params("answerId")

	var answer entity.Answer

	result := db.Debug().First(&answer, answerId)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "DATA NOT FOUND",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "DATA HAS FOUNDED",
		"data":    answer,
	})
}

func AnswerHandlerUpdate(ctx *fiber.Ctx) error {
	answerRequest := new(request.AnswerRequest)
	err := ctx.BodyParser(answerRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "BAD REQUEST",
		})
	}

	err = validate.Struct(answerRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "BAD REQUEST",
			"error":   err.Error(),
		})
	}

	answerId := ctx.Params("answerId")

	var answer entity.Answer

	result := db.First(&answer, answerId)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "DATA NOT FOUND",
		})
	}

	marshal, err := json.Marshal(answerRequest.Answers)

	answer.UserId = answerRequest.UserId
	answer.Answers = string(marshal)

	result = db.Debug().Save(&answer)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "FAILED TO UPDATE DATA",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "SUCCESS UPDATE DATA",
		"data":    answer,
	})
}

func AnswerHandlerDelete(ctx *fiber.Ctx) error {
	answerId := ctx.Params("answerId")

	var answer entity.Answer

	result := db.Debug().First(&answer, answerId)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "DATA NOT FOUND",
		})
	}

	result = db.Debug().Delete(&answer)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "FAILED TO DELETE DATA",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "DATA HAS BEEN DELETED",
	})
}
