package repository

import (
	"context"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StudentRepo struct {
	db *pgxpool.Pool
}

func NewStudentRepo(db *pgxpool.Pool) *StudentRepo {
	return &StudentRepo{db: db}
}

func (r *StudentRepo) Create(ctx context.Context, student domain.Student) error {
	query := `INSERT INTO students (group_id, last_name, first_name, middle_name)
              VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, query, student.GroupID, student.LastName, student.FirstName, student.MiddleName)

	return err
}

func (r *StudentRepo) Put(ctx context.Context, student domain.Student) error {
	query := `UPDATE students SET group_id=$1, last_name=$2, first_name=$3, middle_name=$4 WHERE student_id=$5`
	_, err := r.db.Exec(ctx, query, student.GroupID, student.LastName, student.FirstName, student.MiddleName, student.StudentID)

	return err
}

func (r *StudentRepo) Patch(ctx context.Context, studentID int64, updates map[string]interface{}) error {
	query := `UPDATE students SET`
	args := make([]interface{}, 0, len(updates)+1)
	argsCounter := 1

	for col, val := range updates {
		query += " " + col + " = $" + strconv.Itoa(argsCounter) + ","
		args = append(args, val)
		argsCounter++
	}

	query = query[:len(query)-1]
	query += " WHERE student_id = $" + strconv.Itoa(argsCounter)
	args = append(args, studentID)
	_, err := r.db.Exec(ctx, query, args...)

	return err
}

func (r *StudentRepo) Delete(ctx context.Context, studentID int64) error {
	query := `DELETE FROM students WHERE student_id = $1`
	_, err := r.db.Exec(ctx, query, studentID)

	return err
}

func (r *StudentRepo) GetByID(ctx context.Context, studentID int64) (domain.Student, error) {
	query := `SELECT student_id, group_id, last_name, first_name, middle_name FROM students WHERE student_id = $1`

	student := domain.Student{}
	err := r.db.QueryRow(ctx, query, studentID).Scan(
		&student.StudentID,
		&student.GroupID,
		&student.LastName,
		&student.FirstName,
		&student.MiddleName,
	)

	return student, err
}

func (r *StudentRepo) GetByName(ctx context.Context, lastName, firstName, middleName string) (domain.Student, error) {
	query := `SELECT student_id, group_id, last_name, first_name, middle_name FROM students WHERE last_name = $1 AND first_name = $2 AND middle_name = $3`

	student := domain.Student{}
	err := r.db.QueryRow(ctx, query, lastName, firstName, middleName).Scan(
		&student.StudentID,
		&student.GroupID,
		&student.LastName,
		&student.FirstName,
		&student.MiddleName,
	)

	return student, err
}

func (r *StudentRepo) GetAll(ctx context.Context) ([]domain.Student, error) {
	query := `SELECT student_id, group_id, last_name, first_name, middle_name FROM students`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	/*studentQuantity, err := r.getCountStudents(ctx)
	if err != nil {
		return nil, err
	}*/
	const defaultCapacity = 100
	students := make([]domain.Student, 0, defaultCapacity)
	for rows.Next() {
		var student domain.Student
		err := rows.Scan(
			&student.StudentID,
			&student.GroupID,
			&student.LastName,
			&student.FirstName,
			&student.MiddleName,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (r *StudentRepo) GetAllByGroupID(ctx context.Context, groupID string) ([]domain.Student, error) {
	query := `SELECT student_id, group_id, last_name, first_name, middle_name FROM students WHERE group_id = $1`

	rows, err := r.db.Query(ctx, query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	/*studentQuantity, err := r.getCountStudentsByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}*/
	const defaultCapacity = 100
	students := make([]domain.Student, 0, defaultCapacity)
	for rows.Next() {
		var student domain.Student
		err := rows.Scan(
			&student.StudentID,
			&student.GroupID,
			&student.LastName,
			&student.FirstName,
			&student.MiddleName,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (r *StudentRepo) getCountStudents(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM students;`
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

// func (r *StudentRepo) getCountStudentsByGroupID(ctx context.Context, groupID string) (int64, error) {
// 	query := `SELECT COUNT(*) FROM students WHERE group_id = $1;`
// 	rows, err := r.db.Query(ctx, query, groupID)
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
