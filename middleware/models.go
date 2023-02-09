package middleware

import "database/sql"

type Employee struct {
	Name       string  `json:"name"`
	EmpID      int     `json:"empID"`
	City       string  `json:"city"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
	DBconn     *sql.DB
}
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DBconn   *sql.DB
}
