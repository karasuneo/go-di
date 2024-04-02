package repository_impl

import (
	"database/sql"

	user_models_domain "github.com/karasuneo/go-di/src/domain/models/user"
)

type UserRepositoryImpl interface {
	Save(
		db *sql.DB,
		u *user_models_domain.User,
	) (*user_models_domain.User, error)
}
