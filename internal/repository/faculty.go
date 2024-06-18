package repository

import (
	"context"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FacultyRepo struct {
	db *pgxpool.Pool
}

func NewFacultyRepo(db *pgxpool.Pool) *FacultyRepo {
	return &FacultyRepo{db: db}
}

func (r *FacultyRepo) Create(ctx context.Context, faculty domain.Faculty) error {
	query := `INSERT INTO faculties (university_id, faculty_name, head_last_name, head_first_name, head_middle_name, faculty_email)
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(ctx, query, faculty.UniversityID, faculty.FacultyName, faculty.HeadLastName, faculty.HeadFirstName, faculty.HeadMiddleName, faculty.FacultyEmail)

	return err
}

func (r *FacultyRepo) Put(ctx context.Context, faculty domain.Faculty) error {
	query := `UPDATE faculties SET university_id=$1, faculty_name=$2, head_last_name=$3, head_first_name=$4, head_middle_name=$5, faculty_email=$6 WHERE faculty_id=$7`
	_, err := r.db.Exec(ctx, query, faculty.UniversityID, faculty.FacultyName, faculty.HeadLastName, faculty.HeadFirstName, faculty.HeadMiddleName, faculty.FacultyEmail, faculty.FacultyID)

	return err
}

func (r *FacultyRepo) Patch(ctx context.Context, facultyID int64, updates map[string]interface{}) error {
	query := `UPDATE faculties SET`
	args := make([]interface{}, 0, len(updates)+1)
	argsCounter := 1

	for col, val := range updates {
		query += " " + col + " = $" + strconv.Itoa(argsCounter) + ","
		args = append(args, val)
		argsCounter++
	}

	query = query[:len(query)-1]
	query += " WHERE faculty_id = $" + strconv.Itoa(argsCounter)
	args = append(args, facultyID)
	_, err := r.db.Exec(ctx, query, args...)

	return err
}

func (r *FacultyRepo) Delete(ctx context.Context, facultyID int64) error {
	query := `DELETE FROM faculties WHERE faculty_id = $1`
	_, err := r.db.Exec(ctx, query, facultyID)

	return err
}

func (r *FacultyRepo) GetByID(ctx context.Context, facultyID int64) (domain.FacultyInfo, error) {
	query := `SELECT 
			f.faculty_id, f.university_id, f.faculty_name, f.head_last_name, f.head_first_name, f.head_middle_name, f.faculty_email,
			u.university_name
		FROM 
			faculties f
		LEFT JOIN 
			university u ON f.university_id = u.university_id
		WHERE f.faculty_id = $1`

	facultyInfo := domain.FacultyInfo{}
	err := r.db.QueryRow(ctx, query, facultyID).Scan(
		&facultyInfo.Faculty.FacultyID,
		&facultyInfo.Faculty.UniversityID,
		&facultyInfo.Faculty.FacultyName,
		&facultyInfo.Faculty.HeadLastName,
		&facultyInfo.Faculty.HeadFirstName,
		&facultyInfo.Faculty.HeadMiddleName,
		&facultyInfo.Faculty.FacultyEmail,
		&facultyInfo.FacultySub.UniversityName,
	)

	return facultyInfo, err
}

func (r *FacultyRepo) GetByName(ctx context.Context, facultyName string) (domain.FacultyInfo, error) {
	query := `SELECT 
			f.faculty_id, f.university_id, f.faculty_name, f.head_last_name, f.head_first_name, f.head_middle_name, f.faculty_email,
			u.university_name
		FROM 
			faculties f
		LEFT JOIN 
			university u ON f.university_id = u.university_id
		WHERE f.faculty_name = $1`

	facultyInfo := domain.FacultyInfo{}
	err := r.db.QueryRow(ctx, query, facultyName).Scan(
		&facultyInfo.Faculty.FacultyID,
		&facultyInfo.Faculty.UniversityID,
		&facultyInfo.Faculty.FacultyName,
		&facultyInfo.Faculty.HeadLastName,
		&facultyInfo.Faculty.HeadFirstName,
		&facultyInfo.Faculty.HeadMiddleName,
		&facultyInfo.Faculty.FacultyEmail,
		&facultyInfo.FacultySub.UniversityName,
	)

	return facultyInfo, err
}

func (r *FacultyRepo) GetAll(ctx context.Context) ([]domain.FacultyInfo, error) {
	query := `SELECT 
			f.faculty_id, f.university_id, f.faculty_name, f.head_last_name, f.head_first_name, f.head_middle_name, f.faculty_email,
			u.university_name
		FROM 
			faculties f
		LEFT JOIN 
			university u ON f.university_id = u.university_id`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 200
	faculties := make([]domain.FacultyInfo, 0, defaultCapacity)
	for rows.Next() {
		var facultyInfo domain.FacultyInfo
		err := rows.Scan(
			&facultyInfo.Faculty.FacultyID,
			&facultyInfo.Faculty.UniversityID,
			&facultyInfo.Faculty.FacultyName,
			&facultyInfo.Faculty.HeadLastName,
			&facultyInfo.Faculty.HeadFirstName,
			&facultyInfo.Faculty.HeadMiddleName,
			&facultyInfo.Faculty.FacultyEmail,
			&facultyInfo.FacultySub.UniversityName,
		)
		if err != nil {
			return nil, err
		}
		faculties = append(faculties, facultyInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return faculties, nil
}

func (r *FacultyRepo) GetAllByUniversityID(ctx context.Context, universityID int64) ([]domain.FacultyInfo, error) {
	query := `SELECT 
			f.faculty_id, f.university_id, f.faculty_name, f.head_last_name, f.head_first_name, f.head_middle_name, f.faculty_email,
			u.university_name
		FROM 
			faculties f
		LEFT JOIN 
			university u ON f.university_id = u.university_id
		WHERE f.university_id = $1`

	rows, err := r.db.Query(ctx, query, universityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 200
	faculties := make([]domain.FacultyInfo, 0, defaultCapacity)
	for rows.Next() {
		var facultyInfo domain.FacultyInfo
		err := rows.Scan(
			&facultyInfo.Faculty.FacultyID,
			&facultyInfo.Faculty.UniversityID,
			&facultyInfo.Faculty.FacultyName,
			&facultyInfo.Faculty.HeadLastName,
			&facultyInfo.Faculty.HeadFirstName,
			&facultyInfo.Faculty.HeadMiddleName,
			&facultyInfo.Faculty.FacultyEmail,
			&facultyInfo.FacultySub.UniversityName,
		)
		if err != nil {
			return nil, err
		}
		faculties = append(faculties, facultyInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return faculties, nil
}

func (r *FacultyRepo) getCountFaculties(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM faculties;`
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

func (r *FacultyRepo) getCountFacultiesByUniversityID(ctx context.Context, universityID int64) (int64, error) {
	query := `SELECT COUNT(*) FROM faculties WHERE university_id = $1;`
	rows, err := r.db.Query(ctx, query, universityID)
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
