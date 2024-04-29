package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-quiz/database"
	"go-quiz/models/entity"
	"go-quiz/models/request"
	"go-quiz/models/response"
	"go-quiz/utils"
)

var db = database.DatabaseInit()
var validate = validator.New()

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User

	result := db.Debug().Find(&users)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "FAILED TO GET ALL DATAS",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "SUCCESS GET ALL DATAS",
		"data":    users,
	})
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserRequest)
	err := ctx.BodyParser(user)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "BAD REQUEST",
			"error":   err.Error(),
		})
	}

	err = validate.Struct(user)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "BAD REQUEST",
			"error":   err.Error(),
		})
	}

	newUser := entity.User{
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}

	hashedPassword, err := utils.HashingPassword(user.Password)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "INTERNAL SERVER ERROR",
		})
	}

	newUser.Password = hashedPassword

	result := db.Debug().Create(&newUser)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "FAILED TO STORE DATA",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "SUCCESS CREATE DATA",
		"data":    newUser,
	})
}

func UserHandlerShowById(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")

	var user entity.User

	result := db.Debug().First(&user, userId)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "DATA NOT FOUND",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "DATA HAS FOUNDED",
		"data":    user,
	})
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	err := ctx.BodyParser(userRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "BAD REQUEST",
		})
	}

	var user entity.User

	userId := ctx.Params("userId")

	result := db.First(&user, userId)
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "DATA NOT FOUND",
		})
	}

	result = db.Debug().Model(&user).Updates(userRequest)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "FAILED TO UPDATE DATA",
		})
	}

	//userResponse := response.UserResponse{
	//	ID:        user.ID,
	//	Username:  user.Username,
	//	Email:     user.Email,
	//	Role:      user.Role,
	//	CreatedAt: user.CreatedAt,
	//	UpdatedAt: user.UpdatedAt,
	//}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "USER UPDATED",
		"data":    user,
	})
}

func UserHandlerUpdateEmail(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateEmailRequest)
	err := ctx.BodyParser(userRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "BAD REQUEST",
		})
	}

	var user entity.User

	userId := ctx.Params("userId")
	result := db.Debug().First(&user, userId)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "DATA NOT FOUND",
		})
	}

	if userRequest.Email != "" {
		err := validate.Struct(userRequest)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		user.Email = userRequest.Email
	}

	result = db.Debug().Save(&user)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "FAILED TO UPDATE EMAIL",
		})
	}

	userResponse := response.UpdateEmailUserResponse{
		Username:  user.Username,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "EMAIL UPDATED",
		"data":    userResponse,
	})
}

func UserHandlerDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")

	var user entity.User

	result := db.Debug().First(&user)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "DATA NOT FOUND",
		})
	}

	result = db.Debug().Delete(&user, userId)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "FAILED TO DELETE DATA",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "DATA HAS BEEN DELETED",
	})
}
