package services

import (
	user_models_domain "github.com/karasuneo/go-di/src/domain/models/user"
	"github.com/karasuneo/go-di/src/domain/repository_impl"
	"github.com/karasuneo/go-di/src/infrastructure"
)

type CreateUserService struct {
	userRepo repository_impl.UserRepositoryImpl
}

func NewCreateUserService(
	userRepo repository_impl.UserRepositoryImpl,
) *CreateUserService {
	return &CreateUserService{
		userRepo: userRepo,
	}
}

func (cus *CreateUserService) Run(
	u *user_models_domain.User,
) (
	*user_models_domain.User,
	error,
) {
	dbConnection := infrastructure.DBConnection{}
	conn, err := dbConnection.Connect()
	if err != nil {
		return nil, err
	}

	// userのインスタンスをDBに保存
	resUser, err := cus.userRepo.Save(conn, u)
	if err != nil {
		return nil, err
	}

	return resUser, nil
}
