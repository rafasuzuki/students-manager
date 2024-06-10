package models

import "github.com/rafasuzuki/students-manager/db"

func Insert(s Student) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql_register := "INSERT INTO alunos (nome, graduacao, idade) VALUES ($1, $2, $3) RETURNING id"

	sql_login := "INSERT INTO login (email, senha) VALUES ($1, $2) RETURNING id"

	err = conn.QueryRow(sql_register, s.Nome, s.Graduacao, s.Idade).Scan(&id)
	err = conn.QueryRow(sql_login, s.Email, s.Senha).Scan(&id)

	return
}
