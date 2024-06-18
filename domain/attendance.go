package domain

import (
	"time"
)

type Attendance struct {
	AttendanceID   int64     `json:"attendance_id"`
	StudentID      int64     `json:"student_id"`
	ScheduleID     int64     `json:"schedule_id"`
	Presence       *bool     `json:"presence"`
	LateArrival    *bool     `json:"late_arrival"`
	Respectfulness *bool     `json:"respectfulness"`
	Reason         *string   `json:"reason"`
	Created        time.Time `json:"created"`
}

type AttendanceInfo struct {
	AttendanceSub AttendanceSub `json:"attendance_sub"`
	Attendance    Attendance    `json:"attendance"`
}

type AttendanceSub struct {
	Student StudentFullName `json:"student_full_name"`
}

type GroupAttendance struct {
	AttendanceID   *int64     `json:"attendance_id"`
	StudentID      *int64     `json:"student_id"`
	ScheduleID     *int64     `json:"schedule_id"`
	Presence       *bool      `json:"presence"`
	LateArrival    *bool      `json:"late_arrival"`
	Respectfulness *bool      `json:"respectfulness"`
	Reason         *string    `json:"reason"`
	Created        *time.Time `json:"created"`
}

type GroupAttendanceInfo struct {
	AttendanceSub GroupAttendanceSub `json:"attendance_sub"`
	Attendance    GroupAttendance    `json:"attendance"`
}

type GroupAttendanceSub struct {
	Student AttendanceStudent `json:"student_info"`
}

type AttendanceStudent struct {
	GroupID    string `json:"group_id"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
}
