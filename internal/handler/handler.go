package handler

import (
	"GoCodeEval/internal/model"
	"GoCodeEval/internal/service"
	"GoCodeEval/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHandler handles the user registration process.
func RegisterHandler(c *gin.Context) {
	var req model.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	authService := service.NewAuthService()
	user, err := authService.Register(&req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	response.Success(c, http.StatusOK, user)
}
