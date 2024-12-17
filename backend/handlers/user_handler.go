package handlers

import (
	"backend/models"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService services.UserServiceInterface
}

func NewUserHandler(userService services.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type UserHandlerInterface interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	GetProfile(c *fiber.Ctx) error
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate required fields
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.userService.Register(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to register user, " + err.Error(),
		})
	}

	// Clear password before sending response
	responseUser := user
	responseUser.Password = ""

	return c.Status(fiber.StatusCreated).JSON(responseUser)
}
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if credentials.Email == "" || credentials.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	token, err := h.userService.Login(credentials.Email, credentials.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int64)

	user, err := h.userService.GetUserById(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Clear password before sending response
	user.Password = ""
	return c.JSON(user)
}

// Additional example handlers...
// func (h *UserHandler) ExampleCreate(c *fiber.Ctx) error {
// 	var item models.Example
// 	if err := c.BodyParser(&item); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Invalid request body",
// 		})
// 	}

// 	// Example service call
// 	if err := h.userService.Create(&item); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to create item",
// 		})
// 	}

// 	return c.Status(fiber.StatusCreated).JSON(item)
// }

// func (h *UserHandler) ExampleList(c *fiber.Ctx) error {
// 	limit := c.QueryInt("limit", 10)
// 	offset := c.QueryInt("offset", 0)

// 	// Example service call
// 	items, err := h.userService.List(limit, offset)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to retrieve items",
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"items": items,
// 		"metadata": fiber.Map{
// 			"limit":  limit,
// 			"offset": offset,
// 		},
// 	})
// }
