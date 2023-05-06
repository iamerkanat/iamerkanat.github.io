package main

import (
	"html/template"
	"database/sql"
	"log"
	"net/http"
	"fmt"
	"strconv"
	_"github.com/mattn/go-sqlite3"
)

type UserInfo struct {
	User_id int `json:"user_id"`
	Name string `json:"name"`
	Joined_at string `json:"joined_at"`
}

type BookInfo struct {
	Book_id int `json:"book_id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

type server struct{
	db *sql.DB 
}

func dbConnect() server {
	db, err := sql.Open("sqlite3", "./mylib.db")
	fmt.Println("Opening database")
	if err != nil {
		log.Fatal(err)
	}

	s := server{db: db}
	return s
}

func (s *server) selectUsers() []UserInfo {
	rows, err := s.db.Query("select id, name, joined_at from users;")
	if err != nil {
		log.fatal(err)
	}

	var users []UserInfo
	for rows.Next() {
		var user UserInfo
		err := rows.Scan(&user.User_id, &user.Name, &user.Joined_at)
		if err != nil {
			log.Fatal("selectUsers", err)
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		log.Fatal("selectUsers2", err)
	}

	//fmt.Println(users)
	return users
}
func (s *server) selectBooks() []BookInfo {
	rows, err := s.db.Query("select id, name, author from books;")
	if err != nil {
		log.fatal(err)
	}

	var books []BookInfo
	for rows.Next() {
		var book BookInfo
		err := rows.Scan(&book.User_id, &book.Title, &book.Author)
		if err != nil {
			log.Fatal("selectBooks", err)
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		log.Fatal("selectUsers2", err)
	}

	//fmt.Println(users)
	return books
}

func (s *server) selectUser(id int) UserInfo {
	rows := s.db.QueryRow("select id, name, joined_at from users where id=?;", id)

	var user UserInfo
	err := rows.Scan(&user.User_id, &user.Name, &user.Joined_at)
	if err != nil {
		log.Fatal("selectUsers", err)
	}

	return user
}

func (s *server) selectBook(id int) BookInfo {
	rows := s.db.QueryRow("select id, title, author from books where id=?;", id)

	var user BookInfo
	err := rows.Scan(&book.Book_id, &book.Title, &book.Author)
	if err != nil {
		log.Fatal("selectBooks", err)
	}

	return user
}

func (s *server) allUsersHandle(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/users.html")
	if err != nil {
		log.Fatal("allUsersHandle", err)
	}

	allUsers := s.selectUsers()
	errExecute := t.Execute(w, allUsers)
	fmt.Println(allUsers[0].Name)
	if errExecute != nil {
		log.Fatal("allUsersHandle2", err)
	}
}
func (s *server) allBooksHandle(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/books.html")
	if err != nil {
		log.Fatal("allBooksHandle", err)
	}

	allBooks := s.selectBooks()
	errExecute := t.Execute(w, allBooks)
	fmt.Println(allBooks[0].Title)
	if errExecute != nil {
		log.Fatal("allBooksHandle2", err)
	}
}


func (s *server) updateUserByID(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	id := r.FormValue("id")
	idInt, err := strconv.Atoi(id)
	Name := r.FormVlaue("name")
	updateUser(Name, idInt, s)
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (s *server) updateBookByID(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	id := r.FormValue("id")
	idInt, err := strconv.Atoi(id)
	Name := r.FormVlaue("name")
	updateUser(Name, idInt, s)
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (s *server) updateUserForm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/updateUser.html")
	if err != nil {
		log.Fatal("allUsersHandle", err)
	}

	err = r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	id := r.FormValue("id")
	idInt, err := strconv.Atoi(id)
	user := s.selectUser(idInt)

	t.Execute(w, user)
}
func (s *server) updateBookForm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/updateBook.html")
	if err != nil {
		log.Fatal("allBooksHandle", err)
	}

	err = r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	id := r.FormValue("id")
	idInt, err := strconv.Atoi(id)
	book := s.selectBook(idInt)

	t.Execute(w, book)
}

func (s *server) allUserChangeHandle(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/updateUsers.html")
	if err != nil {
		log.Fatal("allUsersHandle", err)
	}

	allUsers := s.selectUsers()
	errExecute := t.Execute(w.allUsers)
	fmt.Println(allUsers[0].Name)
	if errExecute != nil {
		log.Fatal("allUsersHandle2", err)
	}
}

func (s *server) allBookChangeHandle(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/updateBooks.html")
	if err != nil {
		log.Fatal("allBooksHandle", err)
	}

	allBooks := s.selectBooks()
	errExecute := t.Execute(w.allBooks)
	fmt.Println(allBooks[0].Title)
	if errExecute != nil {
		log.Fatal("allBooksHandle2", err)
	}
}

func (s *server) deleteUser(w http.responseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	id := r.FormValue("id")
	idInt, err := strconv.Atoi(id)
	deleteUser(idInt, s)
	http.Redirect(w, r, "./index.html", http.StatusSeeOther)
}

func (s *server) deleteBook(w http.responseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	id := r.FormValue("id")
	idInt, err := strconv.Atoi(id)
	deleteBook(idInt, s)
	http.Redirect(w, r, "./index.html", http.StatusSeeOther)
}

func (s *server) formHandle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	Name := r.FormValue("name")
	userId := createUser(Name, s)
	//myVar := map[string]interface{}{"name": nameForm, "addr": address}

	person := userInfo{
		User_id: userId,
		Name: Name,
	}
	//fmt.Println(person)

	outputHTML(w, "./static/formComplete.html", person)
}
//func formHandle for books
func (s *server) formHandleB(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	Title := r.FormValue("title")
	bookId := createBook(Title, s)
	//myVar := map[string]interface{}{"name": nameForm, "addr": address}

	booking := BookInfo{
		Book_id: bookId,
		Title: Title,
		Author: Author,
	}
	//fmt.Println(book)

	outputHTML(w, "./static/formCompleteB.html", booking)
}

func createUser(Name string, s *server) int {
	res, err := s.db.Exec("insert into users(Name) value ($1)", Name)
	if err != nil {
		log.Fatal(err)
	}

	user_id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return int(user_id)
}
func createBook(Title string, Author string s *server) int {
	res, err := s.db.Exec("insert into books(Title, Author) values ($1, $2)", Title, Author)
	if err != nil {
		log.Fatal(err)
	}

	book_id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return int(book_id)
}

func updateUser(Name string, id int, s *server) int{
	res, err := s.db.exec("update users set Name=? where id=?", Name, id)
	if err != nil {
		log.Fatal(err)
	}

	user_id, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return int(user_id)
}

func updateBook(Title string, id int, Author string s *server) int{
	res, err := s.db.exec("update users set Title=?, Author=? where id=?", Name, Author, id)
	if err != nil {
		log.Fatal(err)
	}

	book_id, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return int(book_id)
}

func deleteUser(id int, s *server) {
	_, err := s.db.Exec("delete from users where id=?", id)
	if err != nil {
		log.Fatal(err)
	}
}
func deleteBook(id int, s *server) {
	_, err := s.db.Exec("delete from books where id=?", id)
	if err != nil {
		log.Fatal(err)
	}
}

func outputHTML(w http.ResponseWriter, filename string, person UserInfo, booking BookInfo) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatal(err)
	}

	//b, _ := json.Marshal(&person)
	errExecute := t.execute(w, person)
	errExecute := t.execute(w, booking)

	if errExecute != nil {
		log.Fatal(err)
	}
}

func main() {
	s := dbConnect()
	defer s.db.Close()
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", s.formHandle)
	http.HandleFunc("/form", s.formHandle)
	http.HandleFunc("/form", s.formHandleB)
	http.HandleFunc("/users", s.allUsers.Handle)
	http.HandleFunc("/books", s.allBooks.Handle)
	http.HandleFunc("/change", s.allUserChangeHandle)
	http.HandleFunc("/change", s.allBookChangeHandle)
	http.HandleFunc("/update", s.updateUserForm)
	http.HandleFunc("/update", s.updateBookForm)
	http.HandleFunc("/delete", s.deleteUser)
	http.HandleFunc("/delete", s.deleteBook)
	http.HandleFunc("/updateUserByID", s.updateUserByID)
	http.HandleFunc("/updateBookByID", s.updateBookByID)
	fmt.Println("Server running.....")
	http.ListenAndServe(":8080", nil)
}