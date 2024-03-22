package user

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	errorsModel "github.com/MrSossa/AeroAccess/internal/model/errors"
	userModel "github.com/MrSossa/AeroAccess/internal/model/user"
)

type UserService interface {
	Login(user userModel.RequestLogin) (uint, error)
}

type userService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (s *userService) Login(user userModel.RequestLogin) (uint, error) {
	_, pass, err := s.repository.GetUserPassword(user.User)
	if err != nil {
		return 0, err
	}
	passHash := sha256.Sum256([]byte(user.Password))
	passString := hex.EncodeToString(passHash[:])
	if pass != passString {
		return 0, errors.New(errorsModel.ErrInvalidLogin)
	}
	//level, err := s.repository.GetAccessLevel(id)
	return 1, nil
}
