// model.go

package main

import (
  "fmt"
	"database/sql"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *user) getUser(db *sql.DB) error {
	statement := `SELECT name, age FROM users WHERE id=$1`
	return db.QueryRow(statement,u.ID).Scan(&u.Name, &u.Age)
}

func (u *user) updateUser(db *sql.DB) error {
	statement := `UPDATE users SET name=$2, age=$3 WHERE id=$1`
	_, err := db.Exec(statement,u.ID, u.Name, u.Age)
	return err
}

func (u *user) deleteUser(db *sql.DB) error {
	statement := `DELETE FROM users WHERE id=$1`
	_, err := db.Exec(statement,u.ID)
	return err
}

func (u *user) createUser(db *sql.DB) error {
	
	statement := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING userid`
	

	if err != nil {
		return err
	}

	err =  db.QueryRow(sqlStatement, user.Name, user.Age).Scan(&u.id)

	if err != nil {
		return err
	}

	return nil
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	statement := `SELECT id, name, age FROM users LIMIT $1 OFFSET $2`,count, start)
	rows, err := db.Query( "SELECT id, name, age FROM users LIMIT $1 OFFSET $2", count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []user{}

	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
