package handler

import (
	"net/http"
	"strconv"
	"get-getById-with-clean-architecture/repository"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users := h.repo.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user := h.repo.GetUserByID(id)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
