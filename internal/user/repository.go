package user

import (
	"database/sql"
	"errors"

	errorsModel "github.com/MrSossa/AeroAccess/internal/model/errors"
)

type UserRepository interface {
	GetUserPassword(user string) (uint, string, error)
	GetAccessLevel(id uint) (uint, error)
	SaveUser(user, password, name string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserPassword(user string) (uint, string, error) {
	row := r.db.QueryRow("SELECT id,pass FROM usuarios WHERE user = ?", user)
	var password string
	var id uint
	err := row.Scan(&id, &password)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, "", errors.New(errorsModel.ErrInvalidLogin)
		}
		return 0, "", err
	}
	return id, password, nil
}

func (r *userRepository) SaveUser(user, password, name string) error {
	_, err := r.db.Exec("INSERT INTO usuarios (user, pass,nombre) VALUES (?, ?,?)", user, password, name)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetAccessLevel(id uint) (uint, error) {
	row := r.db.QueryRow("SELECT nivel_acceso FROM roles r RIGHT JOIN usuarios_roles u ON u.rol_id = r.id WHERE u.usuario_id = ?", id)
	var accessLevel uint
	err := row.Scan(&accessLevel)
	if err != nil {
		return 0, err
	}
	return accessLevel, nil
}
