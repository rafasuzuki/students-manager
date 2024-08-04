package models

import "github.com/rafasuzuki/students-manager/db"

func GetLogin(email string) (login Login, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM login WHERE email=$1`, email)

	err = row.Scan(&login.ID, &login.Email, &login.Senha)
	
	return
}