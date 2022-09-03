package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var currentPost = Article{}

type Article struct {
	Id                     uint16
	Title, Anons, FullText string
}

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = 1
	dbname   = "postgres"
)

func current(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Show current post page")

	var requestPostUri = strings.Split(r.RequestURI, "/")
	var postId, _ = strconv.Atoi(requestPostUri[len(requestPostUri)-1])
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%d"+
		" dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	res, err := db.Query(fmt.Sprintf("SELECT * from articles_vals where id =%d", postId))
	if err != nil {
		panic(err)
	}
	currentPost = Article{}
	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText) //проверка на сущестование
		if err != nil {
			panic(err)
		}
		currentPost = post
	}
	defer db.Close()
	t, err := template.ParseFiles(
		"templates/current.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "current", currentPost)
}
func index(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%d"+
		" dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	res, err := db.Query("SELECT * from articles_vals")
	if err != nil {
		panic(err)
	}
	var posts = []Article{}
	for res.Next() {
		var art Article
		err = res.Scan(&art.Id, &art.Title, &art.Anons, &art.FullText) //проверка на сущестование
		if err != nil {
			panic(err)
		}
		posts = append(posts, art)
	}
	defer db.Close()
	t, err := template.ParseFiles(
		"templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", posts)

}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"templates/create.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "create", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Println("About page")
	t, err := template.ParseFiles(
		"templates/about.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "about", nil)
}
func save_article(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create page")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%d"+
		" dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	title := r.FormValue("title")
	anons := r.FormValue("anons")
	fullText := r.FormValue("full_text")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	if title == "" || anons == "" || fullText == "" {
		fmt.Fprintf(w, "Не все данные заполенны")
	} else {
		insert, err := db.Query(fmt.Sprintf("INSERT INTO articles_vals(title,anons,full_text) values('%s','%s','%s')",
			title, anons, fullText))
		if err != nil {
			log.Fatal(err)
		}
		defer insert.Close()
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
func handleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/about", about)
	http.HandleFunc("/save_article", save_article)
	http.HandleFunc("/current/", current)
	http.ListenAndServe("127.0.0.2:8000", nil)

}
func main() {
	handleFunc()
}
