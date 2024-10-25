package controller

import (
	"os"

	"github.com/ciazhar/go-start-small/pkg/qr"
	"github.com/gofiber/fiber/v2"
)

type QrController struct {}

func (q *QrController) GenerateQrCode(ctx *fiber.Ctx) error {

	url := ctx.Query("url")
	dimension := ctx.QueryInt("dimension", 255)
	base64Image := ctx.Query("base64Image", "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")
	// base64Image := ctx.Query("base64Image", "")

	fileName, err := qr.GenerateQrCode(url, dimension, base64Image)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = ctx.Status(fiber.StatusOK).SendFile(fileName)
	if err != nil {
		return err
	}

	err = os.Remove(fileName)
	if err != nil {
		return err
	}

	return nil
}

func NewQrController() *QrController {
	return &QrController{}
}
