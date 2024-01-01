package handler

import (
	"net/http"

	"github.com/ahmtsenlik/UserNotification/internal/userservice/service"

	"github.com/gin-gonic/gin"
)

// ListUsersResponse is the struct that defines the response body for the list users handler
type ListUsersResponse struct {
	Users []*service.User `json:"users"` // The users
}

// ListUsersHandler is the handler that returns all the users
func ListUsersHandler(s service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Call the user service to list all the users
		users, err := s.ListUsers(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Create a ListUsersResponse with the users
		res := &ListUsersResponse{
			Users: users,
		}

		// Return the response with status code 200
		c.JSON(http.StatusOK, res)
	}
}
