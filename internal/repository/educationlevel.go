package repository

import (
	"context"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EducationLevelRepo struct {
	db *pgxpool.Pool
}

func NewEducationLevelRepo(db *pgxpool.Pool) *EducationLevelRepo {
	return &EducationLevelRepo{db: db}
}

func (r *EducationLevelRepo) Create(ctx context.Context, educationLevel domain.EducationLevel) error {
	query := `INSERT INTO educationLevels (education_level_name)
              VALUES ($1)`
	_, err := r.db.Exec(ctx, query, educationLevel.EducationLevelName)

	return err
}

func (r *EducationLevelRepo) Put(ctx context.Context, educationLevel domain.EducationLevel) error {
	query := `UPDATE educationLevels SET education_level_name=$1 WHERE education_level_id=$2`
	_, err := r.db.Exec(ctx, query, educationLevel.EducationLevelName, educationLevel.EducationLevelID)

	return err
}

func (r *EducationLevelRepo) Patch(ctx context.Context, educationLevelID int64, updates map[string]interface{}) error {
	query := `UPDATE educationLevels SET`
	args := make([]interface{}, 0, len(updates)+1)
	argsCounter := 1

	for col, val := range updates {
		query += " " + col + " = $" + strconv.Itoa(argsCounter) + ","
		args = append(args, val)
		argsCounter++
	}

	query = query[:len(query)-1]
	query += " WHERE education_level_id = $" + strconv.Itoa(argsCounter)
	args = append(args, educationLevelID)
	_, err := r.db.Exec(ctx, query, args...)

	return err
}

func (r *EducationLevelRepo) Delete(ctx context.Context, educationLevelID int64) error {
	query := `DELETE FROM educationLevels WHERE education_level_id = $1`
	_, err := r.db.Exec(ctx, query, educationLevelID)

	return err
}

func (r *EducationLevelRepo) GetByID(ctx context.Context, educationLevelID int64) (domain.EducationLevel, error) {
	query := `SELECT education_level_id, education_level_name FROM educationLevels WHERE education_level_id = $1`

	educationLevel := domain.EducationLevel{}
	err := r.db.QueryRow(ctx, query, educationLevelID).Scan(
		&educationLevel.EducationLevelID,
		&educationLevel.EducationLevelName,
	)

	return educationLevel, err
}

func (r *EducationLevelRepo) GetByName(ctx context.Context, educationLevelName string) (domain.EducationLevel, error) {
	query := `SELECT education_level_id, education_level_name FROM educationLevels WHERE education_level_name = $1`

	educationLevel := domain.EducationLevel{}
	err := r.db.QueryRow(ctx, query, educationLevelName).Scan(
		&educationLevel.EducationLevelID,
		&educationLevel.EducationLevelName,
	)

	return educationLevel, err
}

func (r *EducationLevelRepo) GetAll(ctx context.Context) ([]domain.EducationLevel, error) {
	query := `SELECT education_level_id, education_level_name FROM educationLevels`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	/*educationLevelQuantity, err := r.getCountEducationLevels(ctx)
	if err != nil {
		return nil, err
	}*/
	const defaultCapacity = 100
	educationLevels := make([]domain.EducationLevel, 0, defaultCapacity)
	for rows.Next() {
		var educationLevel domain.EducationLevel
		err := rows.Scan(
			&educationLevel.EducationLevelID,
			&educationLevel.EducationLevelName,
		)
		if err != nil {
			return nil, err
		}
		educationLevels = append(educationLevels, educationLevel)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return educationLevels, nil
}

func (r *EducationLevelRepo) getCountEducationLevels(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM educationLevels;`
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
