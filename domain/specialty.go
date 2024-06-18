package domain

type Specialty struct {
	SpecialtyCode    string `json:"specialty_code"`
	SpecialtyName    string `json:"specialty_name"`
	DepartamentID    int64  `json:"departament_id"`
	EducationLevelID int64  `json:"education_level_id"`
}

type SpecialtyInfo struct {
	SpecialtySub SpecialtySub `json:"specialty_sub"`
	Specialty    Specialty    `json:"specialty"`
}

type SpecialtySub struct {
	DepartamentName    string `json:"departament_name"`
	EducationLevelName string `json:"education_level_name"`
}
