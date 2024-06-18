package repository

import (
	"context"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SpecialtyRepo struct {
	db *pgxpool.Pool
}

func NewSpecialtyRepo(db *pgxpool.Pool) *SpecialtyRepo {
	return &SpecialtyRepo{db: db}
}

func (r *SpecialtyRepo) Create(ctx context.Context, specialty domain.Specialty) error {
	query := `INSERT INTO specialties (specialty_code, specialty_name, departament_id, education_level_id)
              VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, query, specialty.SpecialtyCode, specialty.SpecialtyName, specialty.DepartamentID, specialty.EducationLevelID)

	return err
}

func (r *SpecialtyRepo) Put(ctx context.Context, specialty domain.Specialty) error {
	query := `UPDATE specialties SET specialty_name=$1, departament_id=$2, education_level_id=$3 WHERE specialty_code=$4`
	_, err := r.db.Exec(ctx, query, specialty.SpecialtyName, specialty.DepartamentID, specialty.EducationLevelID, specialty.SpecialtyCode)

	return err
}

func (r *SpecialtyRepo) Patch(ctx context.Context, specialtyCode string, updates map[string]interface{}) error {
	query := `UPDATE specialties SET`
	args := make([]interface{}, 0, len(updates)+1)
	argsCounter := 1

	for col, val := range updates {
		query += " " + col + " = $" + strconv.Itoa(argsCounter) + ","
		args = append(args, val)
		argsCounter++
	}

	query = query[:len(query)-1]
	query += " WHERE specialty_code = $" + strconv.Itoa(argsCounter)
	args = append(args, specialtyCode)
	_, err := r.db.Exec(ctx, query, args...)

	return err
}

func (r *SpecialtyRepo) Delete(ctx context.Context, specialtyCode string) error {
	query := `DELETE FROM specialties WHERE specialty_code = $1`
	_, err := r.db.Exec(ctx, query, specialtyCode)

	return err
}

func (r *SpecialtyRepo) GetByCode(ctx context.Context, specialtyCode string) (domain.SpecialtyInfo, error) {
	query := `SELECT 
			s.specialty_code, s.specialty_name, s.departament_id, s.education_level_id,
			d.departament_name, e.education_level_name
		FROM 
			specialties s
		LEFT JOIN 
			departaments d ON s.departament_id = d.departament_id
		LEFT JOIN 
			educationLevels e ON s.education_level_id = e.education_level_id
		WHERE s.specialty_code = $1`

	specialtyInfo := domain.SpecialtyInfo{}
	err := r.db.QueryRow(ctx, query, specialtyCode).Scan(
		&specialtyInfo.Specialty.SpecialtyCode,
		&specialtyInfo.Specialty.SpecialtyName,
		&specialtyInfo.Specialty.DepartamentID,
		&specialtyInfo.Specialty.EducationLevelID,
		&specialtyInfo.SpecialtySub.DepartamentName,
		&specialtyInfo.SpecialtySub.EducationLevelName,
	)

	return specialtyInfo, err
}

func (r *SpecialtyRepo) GetByName(ctx context.Context, specialtyName string) (domain.SpecialtyInfo, error) {
	query := `SELECT 
			s.specialty_code, s.specialty_name, s.departament_id, s.education_level_id,
			d.departament_name, e.education_level_name
		FROM 
			specialties s
		LEFT JOIN 
			departaments d ON s.departament_id = d.departament_id
		LEFT JOIN 
			educationLevels e ON s.education_level_id = e.education_level_id
		WHERE s.specialty_name = $1`

	specialtyInfo := domain.SpecialtyInfo{}
	err := r.db.QueryRow(ctx, query, specialtyName).Scan(
		&specialtyInfo.Specialty.SpecialtyCode,
		&specialtyInfo.Specialty.SpecialtyName,
		&specialtyInfo.Specialty.DepartamentID,
		&specialtyInfo.Specialty.EducationLevelID,
		&specialtyInfo.SpecialtySub.DepartamentName,
		&specialtyInfo.SpecialtySub.EducationLevelName,
	)

	return specialtyInfo, err
}

func (r *SpecialtyRepo) GetAll(ctx context.Context) ([]domain.SpecialtyInfo, error) {
	query := `SELECT 
			s.specialty_code, s.specialty_name, s.departament_id, s.education_level_id,
			d.departament_name, e.education_level_name
		FROM 
			specialties s
		LEFT JOIN 
			departaments d ON s.departament_id = d.departament_id
		LEFT JOIN 
			educationLevels e ON s.education_level_id = e.education_level_id`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	/*specialtyQuantity, err := r.getCountSpecialties(ctx)
	if err != nil {
		return nil, err
	}*/
	const defaultCapacity = 100
	specialties := make([]domain.SpecialtyInfo, 0, defaultCapacity)
	for rows.Next() {
		var specialtyInfo domain.SpecialtyInfo
		err := rows.Scan(
			&specialtyInfo.Specialty.SpecialtyCode,
			&specialtyInfo.Specialty.SpecialtyName,
			&specialtyInfo.Specialty.DepartamentID,
			&specialtyInfo.Specialty.EducationLevelID,
			&specialtyInfo.SpecialtySub.DepartamentName,
			&specialtyInfo.SpecialtySub.EducationLevelName,
		)
		if err != nil {
			return nil, err
		}
		specialties = append(specialties, specialtyInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return specialties, nil
}

func (r *SpecialtyRepo) GetAllByDepartamentID(ctx context.Context, departamentID int64) ([]domain.SpecialtyInfo, error) {
	query := `SELECT 
			s.specialty_code, s.specialty_name, s.departament_id, s.education_level_id,
			d.departament_name, e.education_level_name
		FROM 
			specialties s
		LEFT JOIN 
			departaments d ON s.departament_id = d.departament_id
		LEFT JOIN 
			educationLevels e ON s.education_level_id = e.education_level_id
		WHERE s.departament_id = $1`

	rows, err := r.db.Query(ctx, query, departamentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	/*specialtyQuantity, err := r.getCountSpecialtiesByDepartamentID(ctx, departamentID)
	if err != nil {
		return nil, err
	}*/
	const defaultCapacity = 100
	specialties := make([]domain.SpecialtyInfo, 0, defaultCapacity)
	for rows.Next() {
		var specialtyInfo domain.SpecialtyInfo
		err := rows.Scan(
			&specialtyInfo.Specialty.SpecialtyCode,
			&specialtyInfo.Specialty.SpecialtyName,
			&specialtyInfo.Specialty.DepartamentID,
			&specialtyInfo.Specialty.EducationLevelID,
			&specialtyInfo.SpecialtySub.DepartamentName,
			&specialtyInfo.SpecialtySub.EducationLevelName,
		)
		if err != nil {
			return nil, err
		}
		specialties = append(specialties, specialtyInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return specialties, nil
}

func (r *SpecialtyRepo) getCountSpecialties(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM specialties;`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var count int64
	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}

// func (r *SpecialtyRepo) getCountSpecialtiesByDepartamentID(ctx context.Context, departamentID int64) (int64, error) {
// 	query := `SELECT COUNT(*) FROM specialties WHERE departament_id = $1;`
// 	rows, err := r.db.Query(ctx, query, departamentID)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer rows.Close()

// 	var count int64
// 	if rows.Next() {
// 		err := rows.Scan(&count)
// 		if err != nil {
// 			return 0, err
// 		}
// 	}

// 	return count, nil
// }
