package main

// CRUD
import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type UserInfo struct {
	User_id    int    `json:"user_id"`
	FullName   string `json:"fullName"`
	Role       string `json:"role"`
	Created_at string `json:"created_at"`
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

func (s *server) selectUsers() []UserInfo {
	rows, err := s.db.Query("select id, fullName, role, created_at from users;")
	if err != nil {
		log.Fatal(err)
	}

	var users []UserInfo
	for rows.Next() {
		var user UserInfo
		err := rows.Scan(&user.User_id, &user.FullName, &user.Role, &user.Created_at)
		if err != nil {
			log.Fatal("selectUsers", err)
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		log.Fatal("selectUsers2", err)
	}

	// fmt.Println(users)

	return users
}

func (s *server) selectUser(id int) UserInfo {
	rows := s.db.QueryRow("select id, fullName, role, created_at from users where id=?;", id)

	var user UserInfo
	err := rows.Scan(&user.User_id, &user.FullName, &user.Role, &user.Created_at)
	if err != nil {
		log.Fatal("selectUsers", err)
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
	fmt.Println(allUsers[0].FullName)
	if errExecute != nil {
		log.Fatal("allUsersHandle2", err)
	}
}

func (s *server) updateUserByID(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	id := r.FormValue("id")
	idInt, err := strconv.Atoi(id)
	fullName := r.FormValue("name")
	role := r.FormValue("role")
	updateUser(fullName, role, idInt, s)
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

func (s *server) allUserChangeHandle(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/updateUsers.html")
	if err != nil {
		log.Fatal("allUsersHandle", err)
	}

	allUsers := s.selectUsers()
	errExecute := t.Execute(w, allUsers)
	// fmt.Println(allUsers[0].FullName)
	if errExecute != nil {
		log.Fatal("allUsersHandle2", err)
	}
}

func (s *server) deleteUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	id := r.FormValue("id")
	idInt, err := strconv.Atoi(id)
	deleteUser(idInt, s)
	http.Redirect(w, r, "/index.html", http.StatusSeeOther)
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
	// fmt.Println(person)

	outputHTML(w, "./static/formComplete.html", person)
}
func (s *server) createUser(w http.ResponseWriter, r *http.Request) {
	// Get form values
	name := r.FormValue("name")
	role := r.FormValue("role")

	// Insert data into database
	result, err := s.db.Exec("INSERT INTO users (name, role) VALUES (?, ?)", name, role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected > 0 {
		fmt.Fprintf(w, "User created successfully")
	} else {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
	}
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

func updateUser(fullName string, role string, id int, s *server) int {
	res, err := s.db.Exec("update users set fullName=?, role=? where id=?", fullName, role, id)
	if err != nil {
		log.Fatal(err)
	}

	user_id, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return int(user_id)
}

func deleteUser(id int, s *server) {
	_, err := s.db.Exec("delete from users where id=?", id)
	if err != nil {
		log.Fatal(err)
	}
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
	// Connecting database
	s := dbConnect()
	defer s.db.Close()
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", s.formHandle)
	http.HandleFunc("/users", s.allUsersHandle)
	http.HandleFunc("/change", s.allUserChangeHandle)
	http.HandleFunc("/update", s.updateUserForm)
	http.HandleFunc("/delete", s.deleteUser)
	http.HandleFunc("/create", s.createUser)
	http.HandleFunc("/updateUserByID", s.updateUserByID)
	fmt.Println("Server running...")
	http.ListenAndServe(":8080", nil)
}