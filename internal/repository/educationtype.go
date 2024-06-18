package repository

import (
	"context"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EducationTypeRepo struct {
	db *pgxpool.Pool
}

func NewEducationTypeRepo(db *pgxpool.Pool) *EducationTypeRepo {
	return &EducationTypeRepo{db: db}
}

func (r *EducationTypeRepo) Create(ctx context.Context, educationType domain.EducationType) error {
	query := `INSERT INTO educationTypes (education_type_name)
              VALUES ($1)`
	_, err := r.db.Exec(ctx, query, educationType.EducationTypeName)

	return err
}

func (r *EducationTypeRepo) Put(ctx context.Context, educationType domain.EducationType) error {
	query := `UPDATE educationTypes SET education_type_name=$1 WHERE education_type_id=$2`
	_, err := r.db.Exec(ctx, query, educationType.EducationTypeName, educationType.EducationTypeID)

	return err
}

func (r *EducationTypeRepo) Patch(ctx context.Context, educationTypeID int64, updates map[string]interface{}) error {
	query := `UPDATE educationTypes SET`
	args := make([]interface{}, 0, len(updates)+1)
	argsCounter := 1

	for col, val := range updates {
		query += " " + col + " = $" + strconv.Itoa(argsCounter) + ","
		args = append(args, val)
		argsCounter++
	}

	query = query[:len(query)-1]
	query += " WHERE education_type_id = $" + strconv.Itoa(argsCounter)
	args = append(args, educationTypeID)
	_, err := r.db.Exec(ctx, query, args...)

	return err
}

func (r *EducationTypeRepo) Delete(ctx context.Context, educationTypeID int64) error {
	query := `DELETE FROM educationTypes WHERE education_type_id = $1`
	_, err := r.db.Exec(ctx, query, educationTypeID)

	return err
}

func (r *EducationTypeRepo) GetByID(ctx context.Context, educationTypeID int64) (domain.EducationType, error) {
	query := `SELECT education_type_id, education_type_name FROM educationTypes WHERE education_type_id = $1`

	educationType := domain.EducationType{}
	err := r.db.QueryRow(ctx, query, educationTypeID).Scan(
		&educationType.EducationTypeID,
		&educationType.EducationTypeName,
	)

	return educationType, err
}

func (r *EducationTypeRepo) GetByName(ctx context.Context, educationTypeName string) (domain.EducationType, error) {
	query := `SELECT education_type_id, education_type_name FROM educationTypes WHERE education_type_name = $1`

	educationType := domain.EducationType{}
	err := r.db.QueryRow(ctx, query, educationTypeName).Scan(
		&educationType.EducationTypeID,
		&educationType.EducationTypeName,
	)

	return educationType, err
}

func (r *EducationTypeRepo) GetAll(ctx context.Context) ([]domain.EducationType, error) {
	query := `SELECT education_type_id, education_type_name FROM educationTypes`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	educationTypeQuantity, err := r.getCountEducationTypes(ctx)
	if err != nil {
		return nil, err
	}
	const defaultCapacity = 100
	educationTypes := make([]domain.EducationType, 0, educationTypeQuantity)
	for rows.Next() {
		var educationType domain.EducationType
		err := rows.Scan(
			&educationType.EducationTypeID,
			&educationType.EducationTypeName,
		)
		if err != nil {
			return nil, err
		}
		educationTypes = append(educationTypes, educationType)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return educationTypes, nil
}

func (r *EducationTypeRepo) getCountEducationTypes(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM educationTypes;`
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
