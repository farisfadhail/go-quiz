package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go-quiz/models/entity"
	"go-quiz/models/request"
	"go-quiz/utils"
	"time"
)

func LoginHandler(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	err := ctx.BodyParser(loginRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "BAD REQUEST",
			"error":   err.Error(),
		})
	}

	err = validate.Struct(loginRequest)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "BAD REQUEST",
			"error":   err.Error(),
		})
	}

	// Check Availability User
	var user entity.User
	result := db.First(&user, "email = ?", loginRequest.Email)

	if result.Error != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "WRONG CREDENTIAL",
		})
	}

	// Check Validation Password
	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)

	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "WRONG CREDENTIAL",
		})
	}

	// Generate JWT
	claims := jwt.MapClaims{}
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	token, err := utils.GenerateJwtToken(&claims)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "WRONG CREDENTIAL",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
