package domain

import "time"

type AttendanceReport struct {
	ReportHead ReportHead   `json:"reportHead"`
	ReportData []ReportData `json:"ReportData"`
}

type ReportData struct {
	Semester           int64     `json:"semester"`
	WeekType           string    `json:"week_type"`
	DayOfWeek          string    `json:"day_of_week"`
	DisciplineName     string    `json:"discipline_name"`
	DisciplineTypeName string    `json:"discipline_type_name"`
	StartTime          time.Time `json:"start_time"`
	ClassroomName      string    `json:"classroom_name"`
	TeacherName        string    `json:"teacher_name"`
	StudentName        string    `json:"student_name"`
	Presence           *bool     `json:"presence"`
	LateArrival        *bool     `json:"late_arrival"`
	Respectfulness     *bool     `json:"respectfulness"`
	Reason             *string   `json:"reason"`
	Visits             int64     `json:"visits"`
	Passes             int64     `json:"passes"`
	Total              int64     `json:"total"`
	PercentageOfVisits float64   `json:"percentage_of_visits"`
	Created            time.Time `json:"created"`
}

type ReportHead struct {
	UniversityName     string `json:"university_name"`
	UniversityHead     string `json:"university_head"`
	FacultyName        string `json:"faculty_name"`
	FacultyHead        string `json:"faculty_head"`
	DepartamentName    string `json:"departament_name"`
	DepartamentHead    string `json:"departament_head"`
	GroupID            string `json:"group_id"`
	SpecialtyName      string `json:"specialty_name"`
	EducationLevelName string `json:"education_level_name"`
	ProfileName        string `json:"profile_name"`
	EducationTypeName  string `json:"education_type_name"`
}
