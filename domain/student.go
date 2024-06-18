package domain

type Student struct {
	StudentID  int64  `json:"student_id"`
	GroupID    string `json:"group_id"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
}

type StudentFullName struct {
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
}
