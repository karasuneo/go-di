package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/karasuneo/go-di/src/application/services"
	user_models_domain "github.com/karasuneo/go-di/src/domain/models/user"
	"github.com/karasuneo/go-di/src/infrastructure/repository"
)

type CreateUserRequest struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
}

type CreateUserResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Mail string `json:"mail"`
}

func CreateUserHandler(r *gin.Engine) {
	r.POST("api/user/create", func(c *gin.Context) {
		var req CreateUserRequest
		// リクエストのバリデーション
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		cus := services.NewCreateUserService(
			repository.NewUserRepository(),
		)

		user, err := user_models_domain.NewUser(
			nil,
			req.Name,
			req.Mail,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		resUser, err := cus.Run(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res := CreateUserResponse{
			Id:   resUser.GetId(),
			Name: resUser.GetName(),
			Mail: resUser.GetMail(),
		}

		// レスポンスを返却
		c.JSON(http.StatusCreated, res)
	})
}
