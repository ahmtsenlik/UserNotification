package handler

import (
	"net/http"

	"github.com/ahmtsenlik/UserNotification/internal/userservice/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CreateUserRequest is the struct that defines the request body for the create user handler
type CreateUserRequest struct {
	Name  string `json:"name" validate:"required"`  // The user name
	Email string `json:"email" validate:"required"` // The user email
}

// CreateUserResponse is the struct that defines the response body for the create user handler
type CreateUserResponse struct {
	User *service.User `json:"user"` // The user
}

// CreateUserHandler is the handler that creates a new user
func CreateUserHandler(s service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind the request body to a CreateUserRequest struct
		var req CreateUserRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the request body
		validate := validator.New()
		if err := validate.Struct(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the user service to create a new user with the given name and email
		user, err := s.CreateUser(c.Request.Context(), req.Name, req.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Create a CreateUserResponse with the user
		res := &CreateUserResponse{
			User: user,
		}

		// Return the response with status code 201
		c.JSON(http.StatusCreated, res)
	}
}
