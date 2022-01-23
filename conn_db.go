package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
) // не понятно что ты импортируешь - какой ужас

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = 1
	dbname   = "postgres"
)

type Users struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func main() {
	fmt.Println("Работа с Postgres")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%d"+
		" dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	//Установка данных

	//insert, err := db.Query("INSERT INTO go_users(name,age) values('Ayrat',25);")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer insert.Close()
	//err = db.Ping()
	//if err != nil {
	//	panic(err)
	//}

	//Выборка данных
	res, err := db.Query("SELECT name, age from go_users")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		var user Users
		err = res.Scan(&user.Name, &user.Age) //проверка на сущестование
		if err != nil {
			panic(err)
		}

		fmt.Println(
			fmt.Sprintf("User:%s with age: %d!", user.Name, user.Age))
	}
	fmt.Println("Подключение успешно завершено!")
}
