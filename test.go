package main

import (
	"fmt"
	"html/template"
	"net/http"
) // не понятно что ты импортируешь - какой ужас

type User struct { // аналог класса , какой ужас :(
	Name                  string
	Ages                  uint16
	Money                 int16
	Avg_grades, Happiness float64
	Hobbies               []string
}

func (u User) getAllInfo() string { //спереди указывается используемый тип данных , * это ссылка
	return fmt.Sprintf("Username is : %s . He is %d years old, and has money : %d",
		u.Name, u.Ages, u.Money)
}

func (u *User) setNewName(newName string) { // * - это ссылка
	u.Name = newName
}

func home_page(page http.ResponseWriter, r *http.Request) {
	bob := User{Name: "Bob", Ages: 17, Money: -50, Avg_grades: 4.2,
		Happiness: 0.8, Hobbies: []string{"Footbal", "Skiing", "Dancing"}}
	//bob.setNewName("Alex")
	//fmt.Fprintf(page, bob.getAllInfo()) // f-стринга
	//fmt.Fprintf(page, "<h1>Main text</h1><b>Main text</b>")

	temp, _ := template.ParseFiles("templates/home_page.html")
	temp.Execute(page, bob)
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
