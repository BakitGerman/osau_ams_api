package domain

type University struct {
	UniversityID    int64  `json:"university_id"`
	UniversityName  string `json:"university_name"`
	HeadLastName    string `json:"head_last_name"`
	HeadFirstName   string `json:"head_first_name"`
	HeadMiddleName  string `json:"head_middle_name"`
	UniversityEmail string `json:"university_email"`
}
