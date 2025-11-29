package repository

import (
	"Api-Aula1/models"
	"database/sql"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db}
}

func (u UsersRepo) Create(user models.Users) (int8, error) {
	query := `INSERT INTO treehousedb.users(
                            name,
                            email,
                            password,
                            cpf)
                            VALUES (?,?,?,?)`

	statement, err := u.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email, user.Password, user.CPF)
	if err != nil {
		return 0, err
	}
	lastid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int8(uint64(lastid)), nil
}
