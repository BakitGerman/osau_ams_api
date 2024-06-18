package domain

import "time"

type Schedule struct {
	ScheduleID       int64     `json:"schedule_id"`
	GroupID          string    `json:"group_id"`
	DisciplineID     int64     `json:"discipline_id"`
	TeacherID        int64     `json:"teacher_id"`
	DisciplineTypeID int64     `json:"discipline_type_id"`
	ClassroomID      int64     `json:"classroom_id"`
	Semester         int       `json:"semester"`
	BeginStudies     time.Time `json:"begin_studies"`
	WeekType         string    `json:"week_type"`
	DayOfWeek        string    `json:"day_of_week"`
	StartTime        time.Time `json:"start_time"`
	IsActual         *bool     `json:"is_actual"`
}

type ScheduleInfo struct {
	ScheduleSub ScheduleSub `json:"schedule_sub"`
	Schedule    Schedule    `json:"schedule"`
}

type ScheduleSub struct {
	DisciplineName     string          `json:"discipline_name"`
	TeacherFullName    TeacherFullName `json:"teacher_full_name"`
	DisciplineTypeName string          `json:"discipline_type_name"`
	ClassroomName      string          `json:"classroom_name"`
}
