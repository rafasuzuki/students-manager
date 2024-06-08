package models

import "github.com/rafasuzuki/students-manager/db"

func Get(id int64) (student Student, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM aula WHERE dia=$1`, id)

	err = row.Scan(&student.Nome)
	
	return
}