package domain

type Faculty struct {
	FacultyID      int64  `json:"faculty_id"`
	UniversityID   int64  `json:"university_id"`
	FacultyName    string `json:"faculty_name"`
	HeadLastName   string `json:"head_last_name"`
	HeadFirstName  string `json:"head_first_name"`
	HeadMiddleName string `json:"head_middle_name"`
	FacultyEmail   string `json:"faculty_email"`
}

type FacultyInfo struct {
	FacultySub FacultySub `json:"faculty_sub"`
	Faculty    Faculty    `json:"faculty"`
}

type FacultySub struct {
	UniversityName string `json:"university_name"`
}
