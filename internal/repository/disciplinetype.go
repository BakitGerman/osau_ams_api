package repository

import (
	"context"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DisciplineTypeRepo struct {
	db *pgxpool.Pool
}

func NewDisciplineTypeRepo(db *pgxpool.Pool) *DisciplineTypeRepo {
	return &DisciplineTypeRepo{db: db}
}

func (r *DisciplineTypeRepo) Create(ctx context.Context, disciplineType domain.DisciplineType) error {
	query := `INSERT INTO disciplineTypes (discipline_type_name)
              VALUES ($1)`
	_, err := r.db.Exec(ctx, query, disciplineType.DisciplineTypeName)

	return err
}

func (r *DisciplineTypeRepo) Put(ctx context.Context, disciplineType domain.DisciplineType) error {
	query := `UPDATE disciplineTypes SET discipline_type_name=$1 WHERE discipline_type_id=$2`
	_, err := r.db.Exec(ctx, query, disciplineType.DisciplineTypeName, disciplineType.DisciplineTypeID)

	return err
}

func (r *DisciplineTypeRepo) Patch(ctx context.Context, disciplineTypeID int64, updates map[string]interface{}) error {
	query := `UPDATE disciplineTypes SET`
	args := make([]interface{}, 0, len(updates)+1)
	argsCounter := 1

	for col, val := range updates {
		query += " " + col + " = $" + strconv.Itoa(argsCounter) + ","
		args = append(args, val)
		argsCounter++
	}

	query = query[:len(query)-1]
	query += " WHERE discipline_type_id = $" + strconv.Itoa(argsCounter)
	args = append(args, disciplineTypeID)
	_, err := r.db.Exec(ctx, query, args...)

	return err
}

func (r *DisciplineTypeRepo) Delete(ctx context.Context, disciplineTypeID int64) error {
	query := `DELETE FROM disciplineTypes WHERE discipline_type_id = $1`
	_, err := r.db.Exec(ctx, query, disciplineTypeID)

	return err
}

func (r *DisciplineTypeRepo) GetByID(ctx context.Context, disciplineTypeID int64) (domain.DisciplineType, error) {
	query := `SELECT discipline_type_id, discipline_type_name FROM disciplineTypes WHERE discipline_type_id = $1`

	disciplineType := domain.DisciplineType{}
	err := r.db.QueryRow(ctx, query, disciplineTypeID).Scan(
		&disciplineType.DisciplineTypeID,
		&disciplineType.DisciplineTypeName,
	)

	return disciplineType, err
}

func (r *DisciplineTypeRepo) GetByName(ctx context.Context, disciplineTypeName string) (domain.DisciplineType, error) {
	query := `SELECT discipline_type_id, discipline_type_name FROM disciplineTypes WHERE discipline_type_name = $1`

	disciplineType := domain.DisciplineType{}
	err := r.db.QueryRow(ctx, query, disciplineTypeName).Scan(
		&disciplineType.DisciplineTypeID,
		&disciplineType.DisciplineTypeName,
	)

	return disciplineType, err
}

func (r *DisciplineTypeRepo) GetAll(ctx context.Context) ([]domain.DisciplineType, error) {
	query := `SELECT discipline_type_id, discipline_type_name FROM disciplineTypes`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	/*disciplineTypeQuantity, err := r.getCountDisciplineTypes(ctx)
	if err != nil {
		return nil, err
	}*/
	const defaultCapacity = 100
	disciplineTypes := make([]domain.DisciplineType, 0, defaultCapacity)
	for rows.Next() {
		var disciplineType domain.DisciplineType
		err := rows.Scan(
			&disciplineType.DisciplineTypeID,
			&disciplineType.DisciplineTypeName,
		)
		if err != nil {
			return nil, err
		}
		disciplineTypes = append(disciplineTypes, disciplineType)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return disciplineTypes, nil
}

func (r *DisciplineTypeRepo) getCountDisciplineTypes(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM disciplineTypes;`
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
