package controller

import (
	"database/sql"
	"energyApp_site/controller/structs"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var d structs.Db

func AddUser(username string, password string, email string) int {
	d.Password = "2008"
	d.DBname = "postgres"
	d.Name = "postgres"
	d.Host = "localhost"
	d.Port = 5432

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s "+
		"sslmode=disable", d.Host, d.Port, d.Name, d.Password, d.DBname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users(username , email ,password) VALUES ($1 , $2 ,$3)", username, email, password)
	if err != nil {
		log.Fatal(err)
	}

	return 0
}

func CheckUser(username string, password string) int {
	d.Password = "2008"
	d.DBname = "postgres"
	d.Name = "postgres"
	d.Host = "localhost"
	d.Port = 5432

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s "+
		"sslmode=disable", d.Host, d.Port, d.Name, d.Password, d.DBname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE username=$1 AND password=$2", username, password)

	count := 0
	for rows.Next() {
		count++
	}

	if count == 1 {
		return 0
	} else {
		return 1
	}
}
