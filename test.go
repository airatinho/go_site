package main

import (
	"fmt"
	"net/http"
)

type User struct { // аналог класса , какой ужас :(
	name                  string
	ages                  uint16
	money                 int16
	avg_grades, happiness float64
}

func (u User) getAllInfo() string { //спереди указывается используемый тип данных , * это ссылка
	return fmt.Sprintf("Username is : %s . He is %d years old, and has money : %d",
		u.name, u.ages, u.money)
}

func (u *User) setNewName(newName string) { // * - это ссылка
	u.name = newName
}

func home_page(page http.ResponseWriter, r *http.Request) {
	bob := User{name: "Bob", ages: 25, money: -50, avg_grades: 4.2, happiness: 0.8}
	bob.setNewName("Alex")

	fmt.Fprintf(page, bob.getAllInfo()) // f-стринга

}

func contacts_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Contacts page!") // f-стринга
}

func handle_request() {
	http.HandleFunc("/", home_page)             // аналог render
	http.HandleFunc("/contacts", contacts_page) // аналог render

	http.ListenAndServe(":5000", nil) // слушаем порт и создаем и запускаем сервер
}

func main() {

	handle_request()
}
