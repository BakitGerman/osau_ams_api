package repository

import (
	"context"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HeadmanRepo struct {
	db *pgxpool.Pool
}

func NewHeadmanRepo(db *pgxpool.Pool) *HeadmanRepo {
	return &HeadmanRepo{db: db}
}

func (r *HeadmanRepo) Create(ctx context.Context, headman domain.Headman) error {
	query := `INSERT INTO headmans (student_id, group_id)
              VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, headman.StudentID, headman.GroupID)

	return err
}

func (r *HeadmanRepo) Put(ctx context.Context, headman domain.Headman) error {
	query := `UPDATE headmans SET student_id=$1, group_id=$2 WHERE headman_id=$3`
	_, err := r.db.Exec(ctx, query, headman.StudentID, headman.GroupID, headman.HeadmanID)

	return err
}

func (r *HeadmanRepo) Patch(ctx context.Context, headmanID int64, updates map[string]interface{}) error {
	query := `UPDATE headmans SET`
	args := make([]interface{}, 0, len(updates)+1)
	argsCounter := 1

	for col, val := range updates {
		query += " " + col + " = $" + strconv.Itoa(argsCounter) + ","
		args = append(args, val)
		argsCounter++
	}

	query = query[:len(query)-1]
	query += " WHERE headman_id = $" + strconv.Itoa(argsCounter)
	args = append(args, headmanID)
	_, err := r.db.Exec(ctx, query, args...)

	return err
}

func (r *HeadmanRepo) Delete(ctx context.Context, headmanID int64) error {
	query := `DELETE FROM headmans WHERE headman_id = $1`
	_, err := r.db.Exec(ctx, query, headmanID)

	return err
}

func (r *HeadmanRepo) GetByID(ctx context.Context, headmanID int64) (domain.HeadmanInfo, error) {
	query := `SELECT 
			h.headman_id, h.student_id, h.group_id,
			s.last_name, s.first_name, s.middle_name, 
			g.group_name
		FROM 
			headmans h
		LEFT JOIN 
			students s ON h.student_id = s.student_id
		LEFT JOIN 
			groups g ON h.group_id = g.group_id
		WHERE h.headman_id = $1`

	headmanInfo := domain.HeadmanInfo{}
	err := r.db.QueryRow(ctx, query, headmanID).Scan(
		&headmanInfo.Headman.HeadmanID,
		&headmanInfo.Headman.StudentID,
		&headmanInfo.Headman.GroupID,
		&headmanInfo.HeadmanSub.Student.LastName,
		&headmanInfo.HeadmanSub.Student.FirstName,
		&headmanInfo.HeadmanSub.Student.MiddleName,
		&headmanInfo.HeadmanSub.GroupName,
	)

	return headmanInfo, err
}

func (r *HeadmanRepo) GetByStudentID(ctx context.Context, studentID int64) (domain.HeadmanInfo, error) {
	query := `SELECT 
			h.headman_id, h.student_id, h.group_id,
			s.last_name, s.first_name, s.middle_name, 
			g.group_name
		FROM 
			headmans h
		LEFT JOIN 
			students s ON h.student_id = s.student_id
		LEFT JOIN 
			groups g ON h.group_id = g.group_id
		WHERE h.student_id = $1`

	headmanInfo := domain.HeadmanInfo{}
	err := r.db.QueryRow(ctx, query, studentID).Scan(
		&headmanInfo.Headman.HeadmanID,
		&headmanInfo.Headman.StudentID,
		&headmanInfo.Headman.GroupID,
		&headmanInfo.HeadmanSub.Student.LastName,
		&headmanInfo.HeadmanSub.Student.FirstName,
		&headmanInfo.HeadmanSub.Student.MiddleName,
		&headmanInfo.HeadmanSub.GroupName,
	)

	return headmanInfo, err
}

func (r *HeadmanRepo) GetAll(ctx context.Context) ([]domain.HeadmanInfo, error) {
	query := `SELECT 
			h.headman_id, h.student_id, h.group_id,
			s.last_name, s.first_name, s.middle_name, 
			g.group_name
		FROM 
			headmans h
		LEFT JOIN 
			students s ON h.student_id = s.student_id
		LEFT JOIN 
			groups g ON h.group_id = g.group_id`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	/*headmanQuantity, err := r.getCountHeadmans(ctx)
	if err != nil {
		return nil, err
	}*/
	const defaultCapacity = 100
	headmans := make([]domain.HeadmanInfo, 0, defaultCapacity)
	for rows.Next() {
		var headmanInfo domain.HeadmanInfo
		err := rows.Scan(
			&headmanInfo.Headman.HeadmanID,
			&headmanInfo.Headman.StudentID,
			&headmanInfo.Headman.GroupID,
			&headmanInfo.HeadmanSub.Student.LastName,
			&headmanInfo.HeadmanSub.Student.FirstName,
			&headmanInfo.HeadmanSub.Student.MiddleName,
			&headmanInfo.HeadmanSub.GroupName,
		)
		if err != nil {
			return nil, err
		}
		headmans = append(headmans, headmanInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return headmans, nil
}

func (r *HeadmanRepo) getCountHeadmans(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM headmans;`
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
