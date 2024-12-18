package handlers

import (
	"backend/models"
	"backend/services"
	"fmt"

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

// @Summary Register a new user
// @Description Register a new user with the provided details
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/register [post]
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

// Login godoc
// @Summary Login a user
// @Description Login a user with the provided credentials
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body models.User true "User credentials"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/login [post]
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

// GetProfile godoc
// @Summary Get user profile
// @Description Get the profile of the currently authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} models.User
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/auth/profile [get]
func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	// Debug print all locals
	userIDInterface := c.Locals("userId")
	if userIDInterface == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User ID not found in context",
		})
	}

	userID, ok := userIDInterface.(int64)
	if !ok {
		fmt.Printf("Type of userIDInterface: %T\n", userIDInterface)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid user ID format",
		})
	}

	user, err := h.userService.GetUserById(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}
