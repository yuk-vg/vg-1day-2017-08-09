package model

import (
	"database/sql"
)

// User はユーザーの構造体です
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

// UserAll は全てのユーザーを返します
func UserAll(db *sql.DB) ([]*User, error) {

	rows, err := db.Query(`select id, name, age from user`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var us []*User
	for rows.Next() {
		u := &User{}
		// 1-1. ユーザー名を表示しよう
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return us, nil
}
