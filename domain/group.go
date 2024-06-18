package domain

type Group struct {
	GroupID   string `json:"group_id"`
	ProfileID int64  `json:"profile_id"`
}

type GroupInfo struct {
	GroupSub GroupSub `json:"group_sub"`
	Group    Group    `json:"group"`
}

type GroupSub struct {
	ProfileName string `json:"profile_name"`
}
