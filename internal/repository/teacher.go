package repository

import (
	"context"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TeacherRepo struct {
	db *pgxpool.Pool
}

func NewTeacherRepo(db *pgxpool.Pool) *TeacherRepo {
	return &TeacherRepo{db: db}
}

func (r *TeacherRepo) Create(ctx context.Context, teacher domain.Teacher) error {
	query := `INSERT INTO teachers (departament_id, last_name, first_name, middle_name, teacher_email)
              VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(ctx, query, teacher.DepartamentID, teacher.LastName, teacher.FirstName, teacher.MiddleName, teacher.TeacherEmail)

	return err
}

func (r *TeacherRepo) Put(ctx context.Context, teacher domain.Teacher) error {
	query := `UPDATE teachers SET departament_id=$1, last_name=$2, first_name=$3, middle_name=$4, teacher_email=$5 WHERE teacher_id=$6`
	_, err := r.db.Exec(ctx, query, teacher.DepartamentID, teacher.LastName, teacher.FirstName, teacher.MiddleName, teacher.TeacherEmail, teacher.TeacherID)

	return err
}

func (r *TeacherRepo) Patch(ctx context.Context, teacherID int64, updates map[string]interface{}) error {
	query := `UPDATE teachers SET`
	args := make([]interface{}, 0, len(updates)+1)
	argsCounter := 1

	for col, val := range updates {
		query += " " + col + " = $" + strconv.Itoa(argsCounter) + ","
		args = append(args, val)
		argsCounter++
	}

	query = query[:len(query)-1]
	query += " WHERE teacher_id = $" + strconv.Itoa(argsCounter)
	args = append(args, teacherID)
	_, err := r.db.Exec(ctx, query, args...)

	return err
}

func (r *TeacherRepo) Delete(ctx context.Context, teacherID int64) error {
	query := `DELETE FROM teachers WHERE teacher_id = $1`
	_, err := r.db.Exec(ctx, query, teacherID)

	return err
}

func (r *TeacherRepo) GetByID(ctx context.Context, teacherID int64) (domain.TeacherInfo, error) {
	query := `SELECT 
			t.teacher_id, t.departament_id, t.last_name, t.first_name, t.middle_name, t.teacher_email,
			d.departament_name
		FROM 
			teachers t
		LEFT JOIN 
			departaments d ON t.departament_id = d.departament_id
		WHERE t.teacher_id = $1`

	teacherInfo := domain.TeacherInfo{}
	err := r.db.QueryRow(ctx, query, teacherID).Scan(
		&teacherInfo.Teacher.TeacherID,
		&teacherInfo.Teacher.DepartamentID,
		&teacherInfo.Teacher.LastName,
		&teacherInfo.Teacher.FirstName,
		&teacherInfo.Teacher.MiddleName,
		&teacherInfo.Teacher.TeacherEmail,
		&teacherInfo.TeacherSub.DepartamentName,
	)

	return teacherInfo, err
}

func (r *TeacherRepo) GetByEmail(ctx context.Context, teacherEmail string) (domain.TeacherInfo, error) {
	query := `SELECT 
			t.teacher_id, t.departament_id, t.last_name, t.first_name, t.middle_name, t.teacher_email,
			d.departament_name
		FROM 
			teachers t
		LEFT JOIN 
			departaments d ON t.departament_id = d.departament_id
		WHERE t.teacher_email = $1`

	teacherInfo := domain.TeacherInfo{}
	err := r.db.QueryRow(ctx, query, teacherEmail).Scan(
		&teacherInfo.Teacher.TeacherID,
		&teacherInfo.Teacher.DepartamentID,
		&teacherInfo.Teacher.LastName,
		&teacherInfo.Teacher.FirstName,
		&teacherInfo.Teacher.MiddleName,
		&teacherInfo.Teacher.TeacherEmail,
		&teacherInfo.TeacherSub.DepartamentName,
	)

	return teacherInfo, err
}

func (r *TeacherRepo) GetAll(ctx context.Context) ([]domain.TeacherInfo, error) {
	query := `SELECT 
			t.teacher_id, t.departament_id, t.last_name, t.first_name, t.middle_name, t.teacher_email,
			d.departament_name
		FROM 
			teachers t
		LEFT JOIN 
			departaments d ON t.departament_id = d.departament_id`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 100
	teachers := make([]domain.TeacherInfo, 0, defaultCapacity)
	for rows.Next() {
		var teacherInfo domain.TeacherInfo
		err := rows.Scan(
			&teacherInfo.Teacher.TeacherID,
			&teacherInfo.Teacher.DepartamentID,
			&teacherInfo.Teacher.LastName,
			&teacherInfo.Teacher.FirstName,
			&teacherInfo.Teacher.MiddleName,
			&teacherInfo.Teacher.TeacherEmail,
			&teacherInfo.TeacherSub.DepartamentName,
		)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, teacherInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return teachers, nil
}

func (r *TeacherRepo) GetAllByDepartamentID(ctx context.Context, departamentID int64) ([]domain.TeacherInfo, error) {
	query := `SELECT 
			t.teacher_id, t.departament_id, t.last_name, t.first_name, t.middle_name, t.teacher_email,
			d.departament_name
		FROM 
			teachers t
		LEFT JOIN 
			departaments d ON t.departament_id = d.departament_id
		WHERE t.departament_id = $1`

	rows, err := r.db.Query(ctx, query, departamentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 100
	teachers := make([]domain.TeacherInfo, 0, defaultCapacity)
	for rows.Next() {
		var teacherInfo domain.TeacherInfo
		err := rows.Scan(
			&teacherInfo.Teacher.TeacherID,
			&teacherInfo.Teacher.DepartamentID,
			&teacherInfo.Teacher.LastName,
			&teacherInfo.Teacher.FirstName,
			&teacherInfo.Teacher.MiddleName,
			&teacherInfo.Teacher.TeacherEmail,
			&teacherInfo.TeacherSub.DepartamentName,
		)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, teacherInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return teachers, nil
}

func (r *TeacherRepo) getCountTeachers(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM teachers;`
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

// func (r *TeacherRepo) getCountTeachersByDepartamentID(ctx context.Context, departamentID int64) (int64, error) {
// 	query := `SELECT COUNT(*) FROM teachers WHERE departament_id = $1;`
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
