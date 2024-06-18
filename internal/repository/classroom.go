package repository

import (
	"context"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClassroomRepo struct {
	db *pgxpool.Pool
}

func NewClassroomRepo(db *pgxpool.Pool) *ClassroomRepo {
	return &ClassroomRepo{db: db}
}

func (r *ClassroomRepo) Create(ctx context.Context, classroom domain.Classroom) error {
	query := `INSERT INTO classrooms (classroom_name)
              VALUES ($1)`
	_, err := r.db.Exec(ctx, query, classroom.ClassroomName)

	return err
}

func (r *ClassroomRepo) Put(ctx context.Context, classroom domain.Classroom) error {
	query := `UPDATE classrooms SET classroom_name=$1 WHERE classroom_id=$2`
	_, err := r.db.Exec(ctx, query, classroom.ClassroomName, classroom.ClassroomID)

	return err
}

func (r *ClassroomRepo) Patch(ctx context.Context, classroomID int64, updates map[string]interface{}) error {
	query := `UPDATE classrooms SET`
	args := make([]interface{}, 0, len(updates)+1)
	argsCounter := 1

	for col, val := range updates {
		query += " " + col + " = $" + strconv.Itoa(argsCounter) + ","
		args = append(args, val)
		argsCounter++
	}

	query = query[:len(query)-1]
	query += " WHERE classroom_id = $" + strconv.Itoa(argsCounter)
	args = append(args, classroomID)
	_, err := r.db.Exec(ctx, query, args...)

	return err
}

func (r *ClassroomRepo) Delete(ctx context.Context, classroomID int64) error {
	query := `DELETE FROM classrooms WHERE classroom_id = $1`
	_, err := r.db.Exec(ctx, query, classroomID)

	return err
}

func (r *ClassroomRepo) GetByID(ctx context.Context, classroomID int64) (domain.Classroom, error) {
	query := `SELECT classroom_id, classroom_name FROM classrooms WHERE classroom_id = $1`

	classroom := domain.Classroom{}
	err := r.db.QueryRow(ctx, query, classroomID).Scan(
		&classroom.ClassroomID,
		&classroom.ClassroomName,
	)

	return classroom, err
}

func (r *ClassroomRepo) GetByName(ctx context.Context, classroomName string) (domain.Classroom, error) {
	query := `SELECT classroom_id, classroom_name FROM classrooms WHERE classroom_name = $1`

	classroom := domain.Classroom{}
	err := r.db.QueryRow(ctx, query, classroomName).Scan(
		&classroom.ClassroomID,
		&classroom.ClassroomName,
	)

	return classroom, err
}

func (r *ClassroomRepo) GetAll(ctx context.Context) ([]domain.Classroom, error) {
	query := `SELECT classroom_id, classroom_name FROM classrooms`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// classroomQuantity, err := r.getCountClassrooms(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	const defaultCapacity = 100
	classrooms := make([]domain.Classroom, 0, defaultCapacity)
	for rows.Next() {
		var classroom domain.Classroom
		err := rows.Scan(
			&classroom.ClassroomID,
			&classroom.ClassroomName,
		)
		if err != nil {
			return nil, err
		}
		classrooms = append(classrooms, classroom)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return classrooms, nil
}

// func (r *ClassroomRepo) getCountClassrooms(ctx context.Context) (int64, error) {
// 	query := `SELECT COUNT(*) FROM classrooms;`
// 	rows, err := r.db.Query(ctx, query)
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
