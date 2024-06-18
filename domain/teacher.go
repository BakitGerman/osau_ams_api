package domain

type Teacher struct {
	TeacherID     int64  `json:"teacher_id"`
	DepartamentID int64  `json:"departament_id"`
	LastName      string `json:"last_name"`
	FirstName     string `json:"first_name"`
	MiddleName    string `json:"middle_name"`
	TeacherEmail  string `json:"teacher_email"`
}

type TeacherFullName struct {
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
}

type TeacherInfo struct {
	TeacherSub TeacherSub `json:"teacher_sub"`
	Teacher    Teacher    `json:"teacher"`
}

type TeacherSub struct {
	DepartamentName string `json:"departament_name"`
}
