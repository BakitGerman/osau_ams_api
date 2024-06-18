package domain

import "github.com/google/uuid"

type User struct {
	UserID    uuid.UUID `json:"user_id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Role      string    `json:"user_role"`
	HeadmanID *int64    `json:"headman_id"`
	StudentID *int64    `json:"student_id"`
	TeacherID *int64    `json:"teacher_id"`
}

type UserInfo struct {
	UserSub UserSub `json:"user_sub"`
	User    User    `json:"user"`
}

type UserSub struct {
	StudentFullName *StudentFullName `json:"student_full_name"`
	TeacherFullName *TeacherFullName `json:"teacher_full_name"`
	GroupID         *string          `json:"group_id"`
}
