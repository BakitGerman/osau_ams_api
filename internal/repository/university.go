package repository

import (
	"context"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UniversityRepo struct {
	db *pgxpool.Pool
}

func NewUniversityRepo(db *pgxpool.Pool) *UniversityRepo {
	return &UniversityRepo{db: db}
}

func (r *UniversityRepo) Create(ctx context.Context, university domain.University) error {
	query := `INSERT INTO university (university_name, head_last_name, head_first_name, head_middle_name, university_email)
              VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(ctx, query, university.UniversityName, university.HeadLastName, university.HeadFirstName, university.HeadMiddleName, university.UniversityEmail)

	return err
}

func (r *UniversityRepo) Put(ctx context.Context, university domain.University) error {
	query := `UPDATE university SET university_name=$1, head_last_name=$2, head_first_name=$3, head_middle_name=$4, university_email=$5 WHERE university_id=$6`
	_, err := r.db.Exec(ctx, query, university.UniversityName, university.HeadLastName, university.HeadFirstName, university.HeadMiddleName, university.UniversityEmail, university.UniversityID)

	return err
}

func (r *UniversityRepo) Patch(ctx context.Context, universityID int64, updates map[string]interface{}) error {
	query := `UPDATE university SET`
	args := make([]interface{}, 0, len(updates)+1)
	argsCounter := 1

	for col, val := range updates {
		query += " " + col + " = $" + strconv.Itoa(argsCounter) + ","
		args = append(args, val)
		argsCounter++
	}

	query = query[:len(query)-1]
	query += " WHERE university_id = $" + strconv.Itoa(argsCounter)
	args = append(args, universityID)
	_, err := r.db.Exec(ctx, query, args...)

	return err
}

func (r *UniversityRepo) Delete(ctx context.Context, universityID int64) error {
	query := `DELETE FROM university WHERE university_id = $1`
	_, err := r.db.Exec(ctx, query, universityID)

	return err
}

func (r *UniversityRepo) GetByID(ctx context.Context, universityID int64) (domain.University, error) {
	query := `SELECT university_id, university_name, head_last_name, head_first_name, head_middle_name, university_email FROM university WHERE university_id = $1`
	university := domain.University{}
	err := r.db.QueryRow(ctx, query, universityID).Scan(&university.UniversityID, &university.UniversityName, &university.HeadLastName, &university.HeadFirstName, &university.HeadMiddleName, &university.UniversityEmail)

	return university, err
}

func (r *UniversityRepo) GetByName(ctx context.Context, universityName string) (domain.University, error) {
	query := `SELECT university_id, university_name, head_last_name, head_first_name, head_middle_name, university_email FROM university WHERE university_name = $1`
	university := domain.University{}
	err := r.db.QueryRow(ctx, query, universityName).Scan(&university.UniversityID, &university.UniversityName, &university.HeadLastName, &university.HeadFirstName, &university.HeadMiddleName, &university.UniversityEmail)

	return university, err
}

func (r *UniversityRepo) GetByEmail(ctx context.Context, universityName string) (domain.University, error) {
	query := `SELECT university_id, university_name, head_last_name, head_first_name, head_middle_name, university_email FROM university WHERE university_name = $1`
	university := domain.University{}
	err := r.db.QueryRow(ctx, query, universityName).Scan(&university.UniversityID, &university.UniversityName, &university.HeadLastName, &university.HeadFirstName, &university.HeadMiddleName, &university.UniversityEmail)

	return university, err
}

func (r *UniversityRepo) GetAll(ctx context.Context) ([]domain.University, error) {
	query := `SELECT university_id, university_name, head_last_name, head_first_name, head_middle_name, university_email FROM university`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 100
	universities := make([]domain.University, 0, defaultCapacity)
	for rows.Next() {
		var university domain.University
		err := rows.Scan(&university.UniversityID, &university.UniversityName, &university.HeadLastName, &university.HeadFirstName, &university.HeadMiddleName, &university.UniversityEmail)
		if err != nil {
			return nil, err
		}
		universities = append(universities, university)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return universities, nil
}

// func (r *UniversityRepo) getCountUniversities(ctx context.Context) (int64, error) {
// 	query := `SELECT COUNT(*) FROM university;`
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
