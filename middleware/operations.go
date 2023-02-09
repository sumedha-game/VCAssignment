package middleware

func (e *Employee) AddEmp() error {
	rows, err := e.DBconn.Query("insert into Employee(Name,Department,City,Salary)values($1,$2,$3,$4);", e.Name, e.Department, e.City, e.Salary)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
func (e *Employee) UpdateEmp(empid string) error {
	rows, err := e.DBconn.Query("UPDATE Employee set Salary=$1,City=$2,Department=$3 where EmpId=$4", e.Salary, e.City, e.Department, empid)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
func (e *Employee) DeleteEmp(empid string) error {
	rows, err := e.DBconn.Query("Delete from Employee where EmpId=$1", empid)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
func (e *Employee) GetEmp(empid string) (*Employee, error) {
	rows, err := e.DBconn.Query("SELECT * from Employee where EmpId=$1", empid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var emp Employee
	for rows.Next() {
		err = rows.Scan(&emp.EmpID, &emp.Name, &emp.Department, &emp.City, &emp.Salary)
		if err != nil {
			return nil, err
		}
	}

	return &emp, nil
	//return rows.Columns()
}

func (c *Credentials) SignUp() error {
	rows, err := c.DBconn.Query("Insert into Login(UserName, Password)values($1,$2)", c.Username, c.Password)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
func (c *Credentials) SignIn() (string, error) {
	rows, err := c.DBconn.Query("Select Password from Login where UserName=$1", c.Username)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var cred Credentials
	for rows.Next() {
		err = rows.Scan(&cred.Password)
		if err != nil {
			return "", err
		}
	}
	return cred.Password, nil
}
