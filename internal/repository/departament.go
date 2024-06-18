package repository

import (
	"context"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DepartamentRepo struct {
	db *pgxpool.Pool
}

func NewDepartamentRepo(db *pgxpool.Pool) *DepartamentRepo {
	return &DepartamentRepo{db: db}
}

func (r *DepartamentRepo) Create(ctx context.Context, departament domain.Departament) error {
	query := `INSERT INTO departaments (faculty_id, departament_name, head_last_name, head_first_name, head_middle_name, departament_email)
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(ctx, query, departament.FacultyID, departament.DepartamentName, departament.HeadLastName, departament.HeadFirstName, departament.HeadMiddleName, departament.DepartamentEmail)

	return err
}

func (r *DepartamentRepo) Put(ctx context.Context, departament domain.Departament) error {
	query := `UPDATE departaments SET faculty_id=$1, departament_name=$2, head_last_name=$3, head_first_name=$4, head_middle_name=$5, departament_email=$6 WHERE departament_id=$7`
	_, err := r.db.Exec(ctx, query, departament.FacultyID, departament.DepartamentName, departament.HeadLastName, departament.HeadFirstName, departament.HeadMiddleName, departament.DepartamentEmail, departament.DepartamentID)

	return err
}

func (r *DepartamentRepo) Patch(ctx context.Context, departamentID int64, updates map[string]interface{}) error {
	query := `UPDATE departaments SET`
	args := make([]interface{}, 0, len(updates)+1)
	argsCounter := 1

	for col, val := range updates {
		query += " " + col + " = $" + strconv.Itoa(argsCounter) + ","
		args = append(args, val)
		argsCounter++
	}

	query = query[:len(query)-1]
	query += " WHERE departament_id = $" + strconv.Itoa(argsCounter)
	args = append(args, departamentID)
	_, err := r.db.Exec(ctx, query, args...)

	return err
}

func (r *DepartamentRepo) Delete(ctx context.Context, departamentID int64) error {
	query := `DELETE FROM departaments WHERE departament_id = $1`
	_, err := r.db.Exec(ctx, query, departamentID)

	return err
}

func (r *DepartamentRepo) GetByID(ctx context.Context, departamentID int64) (domain.DepartamentInfo, error) {
	query := `SELECT 
			d.departament_id, d.faculty_id, d.departament_name, d.head_last_name, d.head_first_name, d.head_middle_name, d.departament_email,
			f.faculty_name
		FROM 
			departaments d
		LEFT JOIN 
			faculties f ON d.faculty_id = f.faculty_id
		WHERE d.departament_id = $1`

	departamentInfo := domain.DepartamentInfo{}
	err := r.db.QueryRow(ctx, query, departamentID).Scan(
		&departamentInfo.Departament.DepartamentID,
		&departamentInfo.Departament.FacultyID,
		&departamentInfo.Departament.DepartamentName,
		&departamentInfo.Departament.HeadLastName,
		&departamentInfo.Departament.HeadFirstName,
		&departamentInfo.Departament.HeadMiddleName,
		&departamentInfo.Departament.DepartamentEmail,
		&departamentInfo.DepartamentSub.FacultyName,
	)

	return departamentInfo, err
}

func (r *DepartamentRepo) GetByName(ctx context.Context, departamentName string) (domain.DepartamentInfo, error) {
	query := `SELECT 
			d.departament_id, d.faculty_id, d.departament_name, d.head_last_name, d.head_first_name, d.head_middle_name, d.departament_email,
			f.faculty_name
		FROM 
			departaments d
		LEFT JOIN 
			faculties f ON d.faculty_id = f.faculty_id
		WHERE d.departament_name = $1`

	departamentInfo := domain.DepartamentInfo{}
	err := r.db.QueryRow(ctx, query, departamentName).Scan(
		&departamentInfo.Departament.DepartamentID,
		&departamentInfo.Departament.FacultyID,
		&departamentInfo.Departament.DepartamentName,
		&departamentInfo.Departament.HeadLastName,
		&departamentInfo.Departament.HeadFirstName,
		&departamentInfo.Departament.HeadMiddleName,
		&departamentInfo.Departament.DepartamentEmail,
		&departamentInfo.DepartamentSub.FacultyName,
	)

	return departamentInfo, err
}

func (r *DepartamentRepo) GetAll(ctx context.Context) ([]domain.DepartamentInfo, error) {
	query := `SELECT 
			d.departament_id, d.faculty_id, d.departament_name, d.head_last_name, d.head_first_name, d.head_middle_name, d.departament_email,
			f.faculty_name
		FROM 
			departaments d
		LEFT JOIN 
			faculties f ON d.faculty_id = f.faculty_id`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 100
	departaments := make([]domain.DepartamentInfo, 0, defaultCapacity)
	for rows.Next() {
		var departamentInfo domain.DepartamentInfo
		err := rows.Scan(
			&departamentInfo.Departament.DepartamentID,
			&departamentInfo.Departament.FacultyID,
			&departamentInfo.Departament.DepartamentName,
			&departamentInfo.Departament.HeadLastName,
			&departamentInfo.Departament.HeadFirstName,
			&departamentInfo.Departament.HeadMiddleName,
			&departamentInfo.Departament.DepartamentEmail,
			&departamentInfo.DepartamentSub.FacultyName,
		)
		if err != nil {
			return nil, err
		}
		departaments = append(departaments, departamentInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return departaments, nil
}

func (r *DepartamentRepo) GetAllByFacultyID(ctx context.Context, facultyID int64) ([]domain.DepartamentInfo, error) {
	query := `SELECT 
			d.departament_id, d.faculty_id, d.departament_name, d.head_last_name, d.head_first_name, d.head_middle_name, d.departament_email,
			f.faculty_name
		FROM 
			departaments d
		LEFT JOIN 
			faculties f ON d.faculty_id = f.faculty_id
		WHERE d.faculty_id = $1`

	rows, err := r.db.Query(ctx, query, facultyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 100
	departaments := make([]domain.DepartamentInfo, 0, defaultCapacity)
	for rows.Next() {
		var departamentInfo domain.DepartamentInfo
		err := rows.Scan(
			&departamentInfo.Departament.DepartamentID,
			&departamentInfo.Departament.FacultyID,
			&departamentInfo.Departament.DepartamentName,
			&departamentInfo.Departament.HeadLastName,
			&departamentInfo.Departament.HeadFirstName,
			&departamentInfo.Departament.HeadMiddleName,
			&departamentInfo.Departament.DepartamentEmail,
			&departamentInfo.DepartamentSub.FacultyName,
		)
		if err != nil {
			return nil, err
		}
		departaments = append(departaments, departamentInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return departaments, nil
}

func (r *DepartamentRepo) getCountDepartaments(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM departaments;`
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

func (r *DepartamentRepo) getCountDepartamentsByFacultyID(ctx context.Context, facultyID int64) (int64, error) {
	query := `SELECT COUNT(*) FROM departaments WHERE faculty_id = $1;`
	rows, err := r.db.Query(ctx, query, facultyID)
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
