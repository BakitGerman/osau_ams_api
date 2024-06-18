package domain

type Headman struct {
	HeadmanID int64  `json:"headman_id"`
	StudentID int64  `json:"student_id"`
	GroupID   string `json:"group_id"`
}

type HeadmanInfo struct {
	HeadmanSub HeadmanSub `json:"headman_sub"`
	Headman    Headman    `json:"headman"`
}

type HeadmanSub struct {
	Student   StudentFullName `json:"student_full_name"`
	GroupName string          `json:"group_name"`
}
