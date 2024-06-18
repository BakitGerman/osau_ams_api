package repository

import (
	"context"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DisciplineRepo struct {
	db *pgxpool.Pool
}

func NewDisciplineRepo(db *pgxpool.Pool) *DisciplineRepo {
	return &DisciplineRepo{db: db}
}

func (r *DisciplineRepo) Create(ctx context.Context, discipline domain.Discipline) error {
	query := `INSERT INTO disciplines (departament_id, discipline_name)
              VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, discipline.DepartamentID, discipline.DisciplineName)

	return err
}

func (r *DisciplineRepo) Put(ctx context.Context, discipline domain.Discipline) error {
	query := `UPDATE disciplines SET departament_id=$1, discipline_name=$2 WHERE discipline_id=$3`
	_, err := r.db.Exec(ctx, query, discipline.DepartamentID, discipline.DisciplineName, discipline.DisciplineID)

	return err
}

func (r *DisciplineRepo) Patch(ctx context.Context, disciplineID int64, updates map[string]interface{}) error {
	query := `UPDATE disciplines SET`
	args := make([]interface{}, 0, len(updates)+1)
	argsCounter := 1

	for col, val := range updates {
		query += " " + col + " = $" + strconv.Itoa(argsCounter) + ","
		args = append(args, val)
		argsCounter++
	}

	query = query[:len(query)-1]
	query += " WHERE discipline_id = $" + strconv.Itoa(argsCounter)
	args = append(args, disciplineID)
	_, err := r.db.Exec(ctx, query, args...)

	return err
}

func (r *DisciplineRepo) Delete(ctx context.Context, disciplineID int64) error {
	query := `DELETE FROM disciplines WHERE discipline_id = $1`
	_, err := r.db.Exec(ctx, query, disciplineID)

	return err
}

func (r *DisciplineRepo) GetByID(ctx context.Context, disciplineID int64) (domain.DisciplineInfo, error) {
	query := `SELECT 
			d.discipline_id, d.departament_id, d.discipline_name,
			dp.departament_name
		FROM 
			disciplines d
		LEFT JOIN 
			departaments dp ON d.departament_id = dp.departament_id
		WHERE d.discipline_id = $1`

	disciplineInfo := domain.DisciplineInfo{}
	err := r.db.QueryRow(ctx, query, disciplineID).Scan(
		&disciplineInfo.Discipline.DisciplineID,
		&disciplineInfo.Discipline.DepartamentID,
		&disciplineInfo.Discipline.DisciplineName,
		&disciplineInfo.DisciplineSub.DepartamentName,
	)

	return disciplineInfo, err
}

func (r *DisciplineRepo) GetByName(ctx context.Context, disciplineName string) (domain.DisciplineInfo, error) {
	query := `SELECT 
			d.discipline_id, d.departament_id, d.discipline_name,
			dp.departament_name
		FROM 
			disciplines d
		LEFT JOIN 
			departaments dp ON d.departament_id = dp.departament_id
		WHERE d.discipline_name = $1`

	disciplineInfo := domain.DisciplineInfo{}
	err := r.db.QueryRow(ctx, query, disciplineName).Scan(
		&disciplineInfo.Discipline.DisciplineID,
		&disciplineInfo.Discipline.DepartamentID,
		&disciplineInfo.Discipline.DisciplineName,
		&disciplineInfo.DisciplineSub.DepartamentName,
	)

	return disciplineInfo, err
}

func (r *DisciplineRepo) GetAll(ctx context.Context) ([]domain.DisciplineInfo, error) {
	query := `SELECT 
			d.discipline_id, d.departament_id, d.discipline_name,
			dp.departament_name
		FROM 
			disciplines d
		LEFT JOIN 
			departaments dp ON d.departament_id = dp.departament_id`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	/*disciplineQuantity, err := r.getCountDisciplines(ctx)
	if err != nil {
		return nil, err
	}*/
	const defaultCapacity = 100
	disciplines := make([]domain.DisciplineInfo, 0, defaultCapacity)
	for rows.Next() {
		var disciplineInfo domain.DisciplineInfo
		err := rows.Scan(
			&disciplineInfo.Discipline.DisciplineID,
			&disciplineInfo.Discipline.DepartamentID,
			&disciplineInfo.Discipline.DisciplineName,
			&disciplineInfo.DisciplineSub.DepartamentName,
		)
		if err != nil {
			return nil, err
		}
		disciplines = append(disciplines, disciplineInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return disciplines, nil
}

func (r *DisciplineRepo) GetAllByDepartamentID(ctx context.Context, departamentID int64) ([]domain.DisciplineInfo, error) {
	query := `SELECT 
			d.discipline_id, d.departament_id, d.discipline_name,
			dp.departament_name
		FROM 
			disciplines d
		LEFT JOIN 
			departaments dp ON d.departament_id = dp.departament_id
		WHERE d.departament_id = $1`

	rows, err := r.db.Query(ctx, query, departamentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 100
	disciplines := make([]domain.DisciplineInfo, 0, defaultCapacity)
	for rows.Next() {
		var disciplineInfo domain.DisciplineInfo
		err := rows.Scan(
			&disciplineInfo.Discipline.DisciplineID,
			&disciplineInfo.Discipline.DepartamentID,
			&disciplineInfo.Discipline.DisciplineName,
			&disciplineInfo.DisciplineSub.DepartamentName,
		)
		if err != nil {
			return nil, err
		}
		disciplines = append(disciplines, disciplineInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return disciplines, nil
}

func (r *DisciplineRepo) getCountDisciplines(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM disciplines;`
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

func (r *DisciplineRepo) getCountDisciplinesByDepartamentID(ctx context.Context, departamentID int64) (int64, error) {
	query := `SELECT COUNT(*) FROM disciplines WHERE departament_id = $1;`
	rows, err := r.db.Query(ctx, query, departamentID)
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
