package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go-quiz/utils"
)

func Authenticated(ctx *fiber.Ctx) error {
	token := ctx.Get("x-jwt-token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "UNAUTHENTICATED",
		})
	}

	_, err := utils.VerifyTokenJwt(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "UNAUTHENTICATED",
		})
	}

	return ctx.Next()
}

func IsAdmin(ctx *fiber.Ctx) error {
	token := ctx.Get("x-jwt-token")
	claims, err := utils.DecodeToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthenticated.",
		})
	}

	role := claims["role"].(string)
	if role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Forbidden access.",
		})
	}

	return ctx.Next()
}

func IsSurveyCreator(ctx *fiber.Ctx) error {
	token := ctx.Get("x-jwt-token")
	claims, err := utils.DecodeToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthenticated.",
		})
	}

	role := claims["role"].(string)
	if role != "survey_creator" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Forbidden access.",
		})
	}

	return ctx.Next()
}
