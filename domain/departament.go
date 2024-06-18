package domain

type Departament struct {
	DepartamentID    int64  `json:"departament_id"`
	FacultyID        int64  `json:"faculty_id"`
	DepartamentName  string `json:"departament_name"`
	HeadLastName     string `json:"head_last_name"`
	HeadFirstName    string `json:"head_first_name"`
	HeadMiddleName   string `json:"head_middle_name"`
	DepartamentEmail string `json:"departament_email"`
}

type DepartamentInfo struct {
	DepartamentSub DepartamentSub `json:"departament_sub"`
	Departament    Departament    `json:"departament"`
}

type DepartamentSub struct {
	FacultyName string `json:"faculty_name"`
}
