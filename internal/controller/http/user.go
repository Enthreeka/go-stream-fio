package http

import (
	"context"
	"github.com/Enthreeka/go-stream-fio/internal/apperror"
	"github.com/Enthreeka/go-stream-fio/internal/entity/dto"
	"github.com/Enthreeka/go-stream-fio/internal/usecase"
	"github.com/Enthreeka/go-stream-fio/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userUsecase usecase.User

	log *logger.Logger
}

func NewUserHandler(userUsecase usecase.User, log *logger.Logger) *userHandler {
	return &userHandler{
		userUsecase: userUsecase,
		log:         log,
	}
}

func (u *userHandler) UserHandler(c *fiber.Ctx) error {
	name := c.Query("name")
	if name == "" {
		u.log.Error("name is empty")
	}

	u.log.Info("%s", name)

	users, err := u.userUsecase.FilteredUser(context.Background(), name)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(apperror.NewAppError(err, "failed to get user with filter"))
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"users": users,
	})
}

func (u *userHandler) DeleteUserHandler(c *fiber.Ctx) error {

	id := dto.IdUserRequest{}
	err := c.BodyParser(&id)
	if err != nil {
		u.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewAppError(err, "Invalid request body"))
	}

	err = u.userUsecase.DeleteUser(context.Background(), id.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message":    "completed successfully",
		"deleted id": id,
	})
}
