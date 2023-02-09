package main

import (
	"VCAssignment/middleware"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	jwtKey = []byte("secretAssignment#@!123")
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "DBroot#123"
	dbname   = "postgres"
)

func main() {
	fmt.Println("main")
	setupDB()
	http.HandleFunc("/signUP", SignUp)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/get", GetEmployee)
	http.HandleFunc("/createNew", CreateEmployee)
	http.HandleFunc("/delete", DeleteEmployee)
	http.HandleFunc("/update", UpdateEmployee)
	fmt.Println("server started at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprint(w, "ALLOWED METHOD: GET")
	}
	//r.URL.
	var cred middleware.Credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		fmt.Fprint(w, "Bad Request")
		return
	}

	cred.Password = middleware.HashPwds([]byte(cred.Password))
	cred.DBconn = setupDB()
	err = cred.SignUp()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, "User Added!")
}
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprint(w, "ALLOWED METHOD: GET")
	}
	var cred middleware.Credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		fmt.Fprint(w, "Bad Request!")
		return
	}
	cred.DBconn = setupDB()
	passwd, err := cred.SignIn()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	if middleware.ComparePwds(passwd, cred.Password) {
		fmt.Println("Login Successful!")

	} else {
		fmt.Fprint(w, "Wrong Credentials!")
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = cred.Username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		fmt.Fprint(w, err)
	}

	fmt.Fprint(w, fmt.Sprintln("Token:", tokenString))
}
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprint(w, "ALLOWED METHOD: GET")
	}
	var emp middleware.Employee
	if ok, err := ValidateToken(r); !ok {
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		fmt.Fprintf(w, "Not Authorized")
	}
	fmt.Println("authorized")
	id := r.URL.Query().Get("id")
	emp.DBconn = setupDB()
	output, err := emp.GetEmp(id)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	newop, _ := json.Marshal(output)
	fmt.Fprint(w, string(newop))
}
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprint(w, "ALLOWED METHOD: POST")
	}
	var emp middleware.Employee
	json.NewDecoder(r.Body).Decode(&emp)
	fmt.Println("chk:", emp)

	if ok, err := ValidateToken(r); !ok {
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		fmt.Fprintf(w, "Not Authorized")
	}
	fmt.Println("authorized")
	emp.DBconn = setupDB()
	err := emp.AddEmp()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, "Employee Successfully Added")
}
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		fmt.Fprint(w, "ALLOWED METHOD: DELETE")
	}
	var emp middleware.Employee
	if ok, err := ValidateToken(r); !ok {
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		fmt.Fprintf(w, "Not Authorized")
	}
	fmt.Println("authorized")
	emp.DBconn = setupDB()
	id := r.URL.Query().Get("id")
	err := emp.DeleteEmp(id)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, "Employee Successfully Deleted")
}
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		fmt.Fprint(w, "ALLOWED METHOD: PUT")
	}
	var emp middleware.Employee
	json.NewDecoder(r.Body).Decode(&emp)
	if ok, err := ValidateToken(r); !ok {
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		fmt.Fprintf(w, "Not Authorized")
	}
	fmt.Println("authorized")
	fmt.Println("chk:", emp)
	emp.DBconn = setupDB()
	id := r.URL.Query().Get("id")
	err := emp.UpdateEmp(id)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, "Employee Successfully Updated")

}

//create table Employee(EmpId SERIAL, Name varchar(100),Department varchar(30),City varchar(30), Salary Float(2), PRIMARY KEY (EmpId));
//insert into Employee(Name,Department,City,Salary)values('Tony','Dev','Nagpur',32550.65);

func setupDB() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	// this will be printed in the terminal, confirming the connection to the database
	fmt.Println("The database is connected")

	_, err = db.Query("create table Employee(EmpId SERIAL, Name varchar(100),Department varchar(30),City varchar(30), Salary Float(2), PRIMARY KEY (EmpId));")
	if err != nil {
		fmt.Println("create query error", err)
	}

	_, err = db.Query("create table Login(UserName varchar(100) UNIQUE ,Password varchar(100));")
	if err != nil {
		fmt.Println("create query error", err)
	}
	//if rows.Err()
	//fmt.Println(rows.Err())
	return db
}
func ValidateToken(r *http.Request) (bool, error) {
	tokenString := r.Header.Get("Authorization")
	token := strings.Split(tokenString, " ")[1]

	if r.Header["Authorization"] != nil {

		token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return jwtKey, nil
		})

		if err != nil {
			//fmt.Fprintf(w, err.Error())
			return false, err
		}

		if token.Valid {
			//endpoint(w, r)
			fmt.Println("token validated")
			return true, nil
		}
	} else {

		return false, nil
	}
	return false, fmt.Errorf("No Authorization Found")
}
