package user_models_domain

import "github.com/kajiLabTeam/xr-project-relay-server/utils"

type UserId struct {
	value string
}

func NewUserId(
	value *string,
) (*UserId, error) {
	if value == nil {
		ulid, err := utils.GenerateUlid()
		if err != nil {
			return nil, err
		}
		return &UserId{
			value: ulid,
		}, nil
	}

	return &UserId{
		value: *value,
	}, nil
}

func (s *UserId) GetValue() string {
	return s.value
}
