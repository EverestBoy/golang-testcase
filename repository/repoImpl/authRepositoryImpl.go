package repoImpl

import (
	"cf-service/database"
	"cf-service/model"
	"cf-service/repository"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type authRepository struct {
	connection database.Database
}

func (a authRepository) FindUserByEmail(email string) (*model.User, error) {
	log.Println("FindUserByEmail: entered")
	// getting the db
	var db, err = a.connection.GetDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	//getting records with either email or user name or phone
	var user model.User
	var query = fmt.Sprintf("SELECT user_id, username, email, phone, password, status, type, created_on, updated_on  FROM   where email=$1")
	log.Printf("Sql query is: %s", query)
	row := db.QueryRow(query, email)
	err = row.Scan(&user.Id, &user.Username, &user.Email,
		&user.Phone, &user.Password, &user.Status, &user.Type, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("user authentication failed")
		}
		log.Println(err)
		return nil, err
	}

	log.Println("FindUserByEmail: exiting")
	return &user, nil
}

func NewAuthRepository(connection database.Database) repository.AuthRepository {
	return &authRepository{connection: connection}
}
