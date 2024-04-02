package repository

import (
	"database/sql"
	"time"

	user_models_domain "github.com/karasuneo/go-di/src/domain/models/user"
	"github.com/karasuneo/go-di/src/domain/repository_impl"
)

type UserRepository struct{}

func NewUserRepository() repository_impl.UserRepositoryImpl {
	return &UserRepository{}
}

func (ur *UserRepository) Save(
	db *sql.DB,
	u *user_models_domain.User,
) (*user_models_domain.User, error) {
	var id string
	var name string
	var mail string
	var createdAt, updatedAt time.Time
	err := db.QueryRow(`
        INSERT INTO users (id, name, mail) 
        VALUES ($1, $2, $3) 
        RETURNING id, name, mail, created_at, updated_at
	`, u.GetId(), u.GetName(), u.GetMail()).Scan(&id, &name, &mail, &createdAt, &updatedAt)
	if err != nil {
		panic(err)
	}

	resUser, err := user_models_domain.NewUser(
		&id,
		name,
		mail,
	)
	if err != nil {
		return nil, err
	}

	return resUser, nil
}
