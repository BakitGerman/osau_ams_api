package domain

type Profile struct {
	ProfileID       int64  `json:"profile_id"`
	SpecialtyCode   string `json:"specialty_code"`
	EducationTypeID int64  `json:"education_type_id"`
	ProfileName     string `json:"profile_name"`
}

type ProfileInfo struct {
	ProfileSub ProfileSub `json:"profile_sub"`
	Profile    Profile    `json:"profile"`
}

type ProfileSub struct {
	EducationTypeName string `json:"education_type_name"`
}
