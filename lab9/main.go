package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type UserInfo struct {
	User_id  int    `json:"user_id"`
	FullName string `json:"fullName"`
	Role     string `json:"role"`
}

type server struct {
	db *sql.DB
}

func dbConnect() server {
	db, err := sql.Open("sqlite3", "database.sql")
	fmt.Println("Opening database")
	if err != nil {
		log.Fatal(err)
	}

	s := server{db: db}

	return s
}

func (s *server) formHandle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	fullName := r.FormValue("name")
	role := r.FormValue("role")
	userId := createUser(fullName, role, s)
	// myVar := map[string]interface{}{"name": nameForm, "addr": address}

	person := UserInfo{
		User_id:  userId,
		FullName: fullName,
		Role:     role,
	}
	fmt.Println(person)

	outputHTML(w, "./static/formComplete.html", person)
}

func createUser(fullName string, role string, s *server) int {
	res, err := s.db.Exec("insert into users(fullName, role) values ($1, $2)", fullName, role)
	if err != nil {
		log.Fatal(err)
	}

	user_id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return int(user_id)
}

func outputHTML(w http.ResponseWriter, filename string, person UserInfo) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatal(err)
	}

	// b, _ := json.Marshal(&person)
	errExecute := t.Execute(w, person)

	if errExecute != nil {
		log.Fatal(err)
	}

}

func main() {
	// Connecting the database SQL
	s := dbConnect()
	defer s.db.Close()
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", s.formHandle)
	fmt.Println("Server running...")
	http.ListenAndServe(":8080", nil)
}