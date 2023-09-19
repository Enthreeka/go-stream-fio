package http

import (
	"context"
	"github.com/Enthreeka/go-stream-fio/internal/apperror"
	"github.com/Enthreeka/go-stream-fio/internal/entity"
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
	u.log.Info("search user handler is running with filtering")

	name := c.Query("name")
	if name == "" {
		u.log.Error("name is empty")
	}

	users, err := u.userUsecase.FilteredUser(context.Background(), name)
	if err != nil {
		u.log.Error("failed to get user: %v", err)
		c.Status(fiber.StatusInternalServerError).JSON(apperror.NewAppError(err, "failed to get user with filter"))
	}

	u.log.Info("search user handler with filtering is executed")
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"users": users,
	})
}

func (u *userHandler) DeleteUserHandler(c *fiber.Ctx) error {
	u.log.Info("delete user handler is running")

	id := dto.IdUserRequest{}
	err := c.BodyParser(&id)
	if err != nil {
		u.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewAppError(err, "Invalid request body"))
	}

	err = u.userUsecase.DeleteUser(context.Background(), id.ID)
	if err != nil {
		u.log.Error("failed to delete user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	u.log.Info("delete user handler is executed")
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message":    "completed successfully",
		"deleted id": id,
	})
}

func (u *userHandler) CreatePersonHandler(c *fiber.Ctx) error {
	u.log.Info("create person handler is running")

	userRequest := dto.FioRequest{}
	err := c.BodyParser(&userRequest)
	if err != nil {
		u.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewAppError(err, "Invalid request body"))
	}

	user := &entity.User{
		Firstname: userRequest.Name,
		Lastname:  userRequest.Surname,
	}
	err = u.userUsecase.CreatePerson(context.Background(), user)
	if err != nil {
		u.log.Error("failed to create person: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	u.log.Info("create person handler is executed")
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message":        "completed successfully",
		"created person": userRequest,
	})
}
