package repo

import (
	"context"
	"main/models"
)

func (repo *PGRepo) GetUsers() ([]models.Users, error) {
	rows, err := repo.pool.Query(context.Background(),
		`SELECT UserId, Name, SecondName, MiddleName, Email, PhoneNumber, Password FROM Users;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []models.Users

	for rows.Next() {
		var item models.Users
		err = rows.Scan(
			&item.UserId,
			&item.Name,
			&item.SecondName,
			&item.MiddleName,
			&item.Email,
			&item.PhoneNumber,
			&item.Password,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, item)
	}
	return data, nil
}

func (repo *PGRepo) NewUsers(item models.Users) (id int, err error) {
	err = repo.pool.QueryRow(context.Background(),
		`INSERT INTO Users (Name, SecondName, MiddleName, Email, PhoneNumber, Password)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING UserId;`,
		&item.Name,
		&item.SecondName,
		&item.MiddleName,
		&item.Email,
		&item.PhoneNumber,
		&item.Password,
	).Scan(&id)

	return id, err
}
