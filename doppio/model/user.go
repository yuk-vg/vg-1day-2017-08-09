package model

import (
	"database/sql"
)

type User struct {
	ID   int64  `json:"id"`
	Username string `json:"username"`
	// 1-1. ユーザー名を表示しよう
    Sex string `json:"sex"`
    Age int64 `json:"age"`
}


func UsersAll(db *sql.DB) ([]*User, error) {

	rows, err := db.Query(`select * from user`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var us []*User
	for rows.Next() {
		u := &User{}
		// 1-1. ユーザー名を表示しよう
		if err := rows.Scan(&u.ID, &u.Username , &u.Sex, &u.Age); err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return us, nil
}


func UserByID(db *sql.DB, id string) (*User, error) {
	u := &User{}

	// 1-1. ユーザー名を表示しよう
	if err := db.QueryRow(`select * from user where id = ?`, id).Scan(&u.ID, &u.Username , &u.Sex, &u.Age); err != nil {
		return nil, err
	}

	return u, nil
}