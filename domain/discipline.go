package domain

type Discipline struct {
	DisciplineID   int64  `json:"discipline_id"`
	DepartamentID  int64  `json:"departament_id"`
	DisciplineName string `json:"discipline_name"`
}

type DisciplineInfo struct {
	DisciplineSub DisciplineSub `json:"discipline_sub"`
	Discipline    Discipline    `json:"discipline"`
}

type DisciplineSub struct {
	DepartamentName string `json:"departament_name"`
}
