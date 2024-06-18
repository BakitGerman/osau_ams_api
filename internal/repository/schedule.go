package repository

import (
	"context"
	"strconv"
	"sync"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ScheduleRepo struct {
	db *pgxpool.Pool
}

func NewScheduleRepo(db *pgxpool.Pool) *ScheduleRepo {
	return &ScheduleRepo{db: db}
}

func (r *ScheduleRepo) Create(ctx context.Context, schedule domain.Schedule) error {
	query := `INSERT INTO schedules (
		group_id, discipline_id, teacher_id, discipline_type_id, classroom_id, semester, week_type, day_of_week, start_time, is_actual
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := r.db.Exec(ctx, query,
		schedule.GroupID, schedule.DisciplineID, schedule.TeacherID, schedule.DisciplineTypeID, schedule.ClassroomID, schedule.Semester, schedule.WeekType, schedule.DayOfWeek, schedule.StartTime, schedule.IsActual)
	return err
}

func (r *ScheduleRepo) Put(ctx context.Context, schedule domain.Schedule) error {
	query := `UPDATE schedules SET 
		group_id=$1, discipline_id=$2, teacher_id=$3, discipline_type_id=$4, classroom_id=$5, semester=$6, week_type=$7, day_of_week=$8, start_time=$9, is_actual=$10
		WHERE schedule_id=$11`
	_, err := r.db.Exec(ctx, query,
		schedule.GroupID, schedule.DisciplineID, schedule.TeacherID, schedule.DisciplineTypeID, schedule.ClassroomID, schedule.Semester, schedule.WeekType, schedule.DayOfWeek, schedule.StartTime, schedule.IsActual, schedule.ScheduleID)
	return err
}

func (r *ScheduleRepo) Patch(ctx context.Context, scheduleID int64, updates map[string]interface{}) error {
	query := `UPDATE schedules SET`
	args := make([]interface{}, 0, len(updates)+1)
	argsCounter := 1

	for col, val := range updates {
		query += " " + col + " = $" + strconv.Itoa(argsCounter) + ","
		args = append(args, val)
		argsCounter++
	}

	query = query[:len(query)-1]
	query += " WHERE schedule_id = $" + strconv.Itoa(argsCounter)
	args = append(args, scheduleID)
	_, err := r.db.Exec(ctx, query, args...)
	return err
}

func (r *ScheduleRepo) Delete(ctx context.Context, scheduleID int64) error {
	query := `DELETE FROM schedules WHERE schedule_id = $1`
	_, err := r.db.Exec(ctx, query, scheduleID)
	return err
}

func (r *ScheduleRepo) GetByID(ctx context.Context, scheduleID int64) (domain.ScheduleInfo, error) {
	query := `SELECT 
		s.schedule_id, s.group_id, s.discipline_id, s.teacher_id, s.discipline_type_id, s.classroom_id, s.semester, s.week_type, s.day_of_week, s.start_time, s.is_actual,
		d.discipline_name, t.last_name, t.first_name, t.middle_name, dt.discipline_type_name, c.classroom_name
	FROM schedules s
	LEFT JOIN disciplines d ON s.discipline_id = d.discipline_id
	LEFT JOIN teachers t ON s.teacher_id = t.teacher_id
	LEFT JOIN disciplineTypes dt ON s.discipline_type_id = dt.discipline_type_id
	LEFT JOIN classrooms c ON s.classroom_id = c.classroom_id
	WHERE s.schedule_id = $1`

	scheduleInfo := domain.ScheduleInfo{}
	err := r.db.QueryRow(ctx, query, scheduleID).Scan(
		&scheduleInfo.Schedule.ScheduleID, &scheduleInfo.Schedule.GroupID, &scheduleInfo.Schedule.DisciplineID, &scheduleInfo.Schedule.TeacherID, &scheduleInfo.Schedule.DisciplineTypeID, &scheduleInfo.Schedule.ClassroomID, &scheduleInfo.Schedule.Semester, &scheduleInfo.Schedule.WeekType, &scheduleInfo.Schedule.DayOfWeek, &scheduleInfo.Schedule.StartTime, &scheduleInfo.Schedule.IsActual,
		&scheduleInfo.ScheduleSub.DisciplineName, &scheduleInfo.ScheduleSub.TeacherFullName.LastName, &scheduleInfo.ScheduleSub.TeacherFullName.FirstName, &scheduleInfo.ScheduleSub.TeacherFullName.MiddleName, &scheduleInfo.ScheduleSub.DisciplineTypeName, &scheduleInfo.ScheduleSub.ClassroomName)

	return scheduleInfo, err
}

func (r *ScheduleRepo) getCountSchedules(ctx context.Context, query string, args ...interface{}) (int64, error) {
	row := r.db.QueryRow(ctx, query, args...)
	var count int64
	err := row.Scan(&count)
	return count, err
}

func (r *ScheduleRepo) GetAll(ctx context.Context) ([]domain.ScheduleInfo, error) {
	query := `SELECT 
		s.schedule_id, s.group_id, s.discipline_id, s.teacher_id, s.discipline_type_id, s.classroom_id, s.semester, s.week_type, s.day_of_week, s.start_time, s.is_actual,
		d.discipline_name, t.last_name, t.first_name, t.middle_name, dt.discipline_type_name, c.classroom_name
	FROM schedules s
	LEFT JOIN disciplines d ON s.discipline_id = d.discipline_id
	LEFT JOIN teachers t ON s.teacher_id = t.teacher_id
	LEFT JOIN disciplineTypes dt ON s.discipline_type_id = dt.discipline_type_id
	LEFT JOIN classrooms c ON s.classroom_id = c.classroom_id`

	countQuery := `SELECT COUNT(*) FROM schedules`

	var count int64
	var countErr error

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		count, countErr = r.getCountSchedules(ctx, countQuery)
	}()

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	wg.Wait()

	if countErr != nil {
		return nil, countErr
	}

	const defaultCapacity = 200
	schedules := make([]domain.ScheduleInfo, 0, count)
	for rows.Next() {
		var scheduleInfo domain.ScheduleInfo
		err := rows.Scan(
			&scheduleInfo.Schedule.ScheduleID, &scheduleInfo.Schedule.GroupID, &scheduleInfo.Schedule.DisciplineID, &scheduleInfo.Schedule.TeacherID, &scheduleInfo.Schedule.DisciplineTypeID, &scheduleInfo.Schedule.ClassroomID, &scheduleInfo.Schedule.Semester, &scheduleInfo.Schedule.WeekType, &scheduleInfo.Schedule.DayOfWeek, &scheduleInfo.Schedule.StartTime, &scheduleInfo.Schedule.IsActual,
			&scheduleInfo.ScheduleSub.DisciplineName, &scheduleInfo.ScheduleSub.TeacherFullName.LastName, &scheduleInfo.ScheduleSub.TeacherFullName.FirstName, &scheduleInfo.ScheduleSub.TeacherFullName.MiddleName, &scheduleInfo.ScheduleSub.DisciplineTypeName, &scheduleInfo.ScheduleSub.ClassroomName)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, scheduleInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *ScheduleRepo) GetByGroupID(ctx context.Context, groupID string) ([]domain.ScheduleInfo, error) {
	query := `SELECT 
		s.schedule_id, s.group_id, s.discipline_id, s.teacher_id, s.discipline_type_id, s.classroom_id, s.semester, s.week_type, s.day_of_week, s.start_time, s.is_actual,
		d.discipline_name, t.last_name, t.first_name, t.middle_name, dt.discipline_type_name, c.classroom_name
	FROM schedules s
	LEFT JOIN disciplines d ON s.discipline_id = d.discipline_id
	LEFT JOIN teachers t ON s.teacher_id = t.teacher_id
	LEFT JOIN disciplineTypes dt ON s.discipline_type_id = dt.discipline_type_id
	LEFT JOIN classrooms c ON s.classroom_id = c.classroom_id
	WHERE s.group_id = $1`

	rows, err := r.db.Query(ctx, query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 200
	schedules := make([]domain.ScheduleInfo, 0, defaultCapacity)
	for rows.Next() {
		var scheduleInfo domain.ScheduleInfo
		err := rows.Scan(
			&scheduleInfo.Schedule.ScheduleID, &scheduleInfo.Schedule.GroupID, &scheduleInfo.Schedule.DisciplineID, &scheduleInfo.Schedule.TeacherID, &scheduleInfo.Schedule.DisciplineTypeID, &scheduleInfo.Schedule.ClassroomID, &scheduleInfo.Schedule.Semester, &scheduleInfo.Schedule.WeekType, &scheduleInfo.Schedule.DayOfWeek, &scheduleInfo.Schedule.StartTime, &scheduleInfo.Schedule.IsActual,
			&scheduleInfo.ScheduleSub.DisciplineName, &scheduleInfo.ScheduleSub.TeacherFullName.LastName, &scheduleInfo.ScheduleSub.TeacherFullName.FirstName, &scheduleInfo.ScheduleSub.TeacherFullName.MiddleName, &scheduleInfo.ScheduleSub.DisciplineTypeName, &scheduleInfo.ScheduleSub.ClassroomName)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, scheduleInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *ScheduleRepo) GetByTeacherID(ctx context.Context, teacherID int64) ([]domain.ScheduleInfo, error) {
	query := `SELECT 
		s.schedule_id, s.group_id, s.discipline_id, s.teacher_id, s.discipline_type_id, s.classroom_id, s.semester, s.week_type, s.day_of_week, s.start_time, s.is_actual,
		d.discipline_name, t.last_name, t.first_name, t.middle_name, dt.discipline_type_name, c.classroom_name
	FROM schedules s
	LEFT JOIN disciplines d ON s.discipline_id = d.discipline_id
	LEFT JOIN teachers t ON s.teacher_id = t.teacher_id
	LEFT JOIN disciplineTypes dt ON s.discipline_type_id = dt.discipline_type_id
	LEFT JOIN classrooms c ON s.classroom_id = c.classroom_id
	WHERE s.teacher_id = $1`

	rows, err := r.db.Query(ctx, query, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 200
	schedules := make([]domain.ScheduleInfo, 0, defaultCapacity)
	for rows.Next() {
		var scheduleInfo domain.ScheduleInfo
		err := rows.Scan(
			&scheduleInfo.Schedule.ScheduleID, &scheduleInfo.Schedule.GroupID, &scheduleInfo.Schedule.DisciplineID, &scheduleInfo.Schedule.TeacherID, &scheduleInfo.Schedule.DisciplineTypeID, &scheduleInfo.Schedule.ClassroomID, &scheduleInfo.Schedule.Semester, &scheduleInfo.Schedule.WeekType, &scheduleInfo.Schedule.DayOfWeek, &scheduleInfo.Schedule.StartTime, &scheduleInfo.Schedule.IsActual,
			&scheduleInfo.ScheduleSub.DisciplineName, &scheduleInfo.ScheduleSub.TeacherFullName.LastName, &scheduleInfo.ScheduleSub.TeacherFullName.FirstName, &scheduleInfo.ScheduleSub.TeacherFullName.MiddleName, &scheduleInfo.ScheduleSub.DisciplineTypeName, &scheduleInfo.ScheduleSub.ClassroomName)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, scheduleInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *ScheduleRepo) GetByGroupAndWeekType(ctx context.Context, groupID string, weekType string) ([]domain.ScheduleInfo, error) {
	query := `SELECT 
		s.schedule_id, s.group_id, s.discipline_id, s.teacher_id, s.discipline_type_id, s.classroom_id, s.semester, s.week_type, s.day_of_week, s.start_time, s.is_actual,
		d.discipline_name, t.last_name, t.first_name, t.middle_name, dt.discipline_type_name, c.classroom_name
	FROM schedules s
	LEFT JOIN disciplines d ON s.discipline_id = d.discipline_id
	LEFT JOIN teachers t ON s.teacher_id = t.teacher_id
	LEFT JOIN disciplineTypes dt ON s.discipline_type_id = dt.discipline_type_id
	LEFT JOIN classrooms c ON s.classroom_id = c.classroom_id
	WHERE s.group_id = $1 AND s.week_type = $2`

	rows, err := r.db.Query(ctx, query, groupID, weekType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 200
	schedules := make([]domain.ScheduleInfo, 0, defaultCapacity)
	for rows.Next() {
		var scheduleInfo domain.ScheduleInfo
		err := rows.Scan(
			&scheduleInfo.Schedule.ScheduleID, &scheduleInfo.Schedule.GroupID, &scheduleInfo.Schedule.DisciplineID, &scheduleInfo.Schedule.TeacherID, &scheduleInfo.Schedule.DisciplineTypeID, &scheduleInfo.Schedule.ClassroomID, &scheduleInfo.Schedule.Semester, &scheduleInfo.Schedule.WeekType, &scheduleInfo.Schedule.DayOfWeek, &scheduleInfo.Schedule.StartTime, &scheduleInfo.Schedule.IsActual,
			&scheduleInfo.ScheduleSub.DisciplineName, &scheduleInfo.ScheduleSub.TeacherFullName.LastName, &scheduleInfo.ScheduleSub.TeacherFullName.FirstName, &scheduleInfo.ScheduleSub.TeacherFullName.MiddleName, &scheduleInfo.ScheduleSub.DisciplineTypeName, &scheduleInfo.ScheduleSub.ClassroomName)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, scheduleInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *ScheduleRepo) GetByTeacherAndWeekType(ctx context.Context, teacherID int64, weekType string) ([]domain.ScheduleInfo, error) {
	query := `SELECT 
		s.schedule_id, s.group_id, s.discipline_id, s.teacher_id, s.discipline_type_id, s.classroom_id, s.semester, s.week_type, s.day_of_week, s.start_time, s.is_actual,
		d.discipline_name, t.last_name, t.first_name, t.middle_name, dt.discipline_type_name, c.classroom_name
	FROM schedules s
	LEFT JOIN disciplines d ON s.discipline_id = d.discipline_id
	LEFT JOIN teachers t ON s.teacher_id = t.teacher_id
	LEFT JOIN disciplineTypes dt ON s.discipline_type_id = dt.discipline_type_id
	LEFT JOIN classrooms c ON s.classroom_id = c.classroom_id
	WHERE s.teacher_id = $1 AND s.week_type = $2`

	rows, err := r.db.Query(ctx, query, teacherID, weekType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 200
	schedules := make([]domain.ScheduleInfo, 0, defaultCapacity)
	for rows.Next() {
		var scheduleInfo domain.ScheduleInfo
		err := rows.Scan(
			&scheduleInfo.Schedule.ScheduleID, &scheduleInfo.Schedule.GroupID, &scheduleInfo.Schedule.DisciplineID, &scheduleInfo.Schedule.TeacherID, &scheduleInfo.Schedule.DisciplineTypeID, &scheduleInfo.Schedule.ClassroomID, &scheduleInfo.Schedule.Semester, &scheduleInfo.Schedule.WeekType, &scheduleInfo.Schedule.DayOfWeek, &scheduleInfo.Schedule.StartTime, &scheduleInfo.Schedule.IsActual,
			&scheduleInfo.ScheduleSub.DisciplineName, &scheduleInfo.ScheduleSub.TeacherFullName.LastName, &scheduleInfo.ScheduleSub.TeacherFullName.FirstName, &scheduleInfo.ScheduleSub.TeacherFullName.MiddleName, &scheduleInfo.ScheduleSub.DisciplineTypeName, &scheduleInfo.ScheduleSub.ClassroomName)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, scheduleInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *ScheduleRepo) GetByGroupWeekTypeAndDay(ctx context.Context, groupID, weekType, dayOfWeek string) ([]domain.ScheduleInfo, error) {
	query := `SELECT 
		s.schedule_id, s.group_id, s.discipline_id, s.teacher_id, s.discipline_type_id, s.classroom_id, s.semester, s.week_type, s.day_of_week, s.start_time, s.is_actual,
		d.discipline_name, t.last_name, t.first_name, t.middle_name, dt.discipline_type_name, c.classroom_name
	FROM schedules s
	LEFT JOIN disciplines d ON s.discipline_id = d.discipline_id
	LEFT JOIN teachers t ON s.teacher_id = t.teacher_id
	LEFT JOIN disciplineTypes dt ON s.discipline_type_id = dt.discipline_type_id
	LEFT JOIN classrooms c ON s.classroom_id = c.classroom_id
	WHERE s.group_id = $1 AND s.week_type = $2 AND s.day_of_week = $3`

	rows, err := r.db.Query(ctx, query, groupID, weekType, dayOfWeek)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 200
	schedules := make([]domain.ScheduleInfo, 0, defaultCapacity)
	for rows.Next() {
		var scheduleInfo domain.ScheduleInfo
		err := rows.Scan(
			&scheduleInfo.Schedule.ScheduleID, &scheduleInfo.Schedule.GroupID, &scheduleInfo.Schedule.DisciplineID, &scheduleInfo.Schedule.TeacherID, &scheduleInfo.Schedule.DisciplineTypeID, &scheduleInfo.Schedule.ClassroomID, &scheduleInfo.Schedule.Semester, &scheduleInfo.Schedule.WeekType, &scheduleInfo.Schedule.DayOfWeek, &scheduleInfo.Schedule.StartTime, &scheduleInfo.Schedule.IsActual,
			&scheduleInfo.ScheduleSub.DisciplineName, &scheduleInfo.ScheduleSub.TeacherFullName.LastName, &scheduleInfo.ScheduleSub.TeacherFullName.FirstName, &scheduleInfo.ScheduleSub.TeacherFullName.MiddleName, &scheduleInfo.ScheduleSub.DisciplineTypeName, &scheduleInfo.ScheduleSub.ClassroomName)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, scheduleInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *ScheduleRepo) GetByTeacherWeekTypeAndDay(ctx context.Context, teacherID int64, weekType, dayOfWeek string) ([]domain.ScheduleInfo, error) {
	query := `SELECT 
		s.schedule_id, s.group_id, s.discipline_id, s.teacher_id, s.discipline_type_id, s.classroom_id, s.semester, s.week_type, s.day_of_week, s.start_time, s.is_actual,
		d.discipline_name, t.last_name, t.first_name, t.middle_name, dt.discipline_type_name, c.classroom_name
	FROM schedules s
	LEFT JOIN disciplines d ON s.discipline_id = d.discipline_id
	LEFT JOIN teachers t ON s.teacher_id = t.teacher_id
	LEFT JOIN disciplineTypes dt ON s.discipline_type_id = dt.discipline_type_id
	LEFT JOIN classrooms c ON s.classroom_id = c.classroom_id
	WHERE s.teacher_id = $1 AND s.week_type = $2 AND s.day_of_week = $3`

	rows, err := r.db.Query(ctx, query, teacherID, weekType, dayOfWeek)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 200
	schedules := make([]domain.ScheduleInfo, 0, defaultCapacity)
	for rows.Next() {
		var scheduleInfo domain.ScheduleInfo
		err := rows.Scan(
			&scheduleInfo.Schedule.ScheduleID, &scheduleInfo.Schedule.GroupID, &scheduleInfo.Schedule.DisciplineID, &scheduleInfo.Schedule.TeacherID, &scheduleInfo.Schedule.DisciplineTypeID, &scheduleInfo.Schedule.ClassroomID, &scheduleInfo.Schedule.Semester, &scheduleInfo.Schedule.WeekType, &scheduleInfo.Schedule.DayOfWeek, &scheduleInfo.Schedule.StartTime, &scheduleInfo.Schedule.IsActual,
			&scheduleInfo.ScheduleSub.DisciplineName, &scheduleInfo.ScheduleSub.TeacherFullName.LastName, &scheduleInfo.ScheduleSub.TeacherFullName.FirstName, &scheduleInfo.ScheduleSub.TeacherFullName.MiddleName, &scheduleInfo.ScheduleSub.DisciplineTypeName, &scheduleInfo.ScheduleSub.ClassroomName)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, scheduleInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *ScheduleRepo) GetActualByGroupID(ctx context.Context, groupID string) ([]domain.ScheduleInfo, error) {
	query := `SELECT 
		s.schedule_id, s.group_id, s.discipline_id, s.teacher_id, s.discipline_type_id, s.classroom_id, s.semester, s.week_type, s.day_of_week, s.start_time, s.is_actual,
		d.discipline_name, t.last_name, t.first_name, t.middle_name, dt.discipline_type_name, c.classroom_name
	FROM schedules s
	LEFT JOIN disciplines d ON s.discipline_id = d.discipline_id
	LEFT JOIN teachers t ON s.teacher_id = t.teacher_id
	LEFT JOIN disciplineTypes dt ON s.discipline_type_id = dt.discipline_type_id
	LEFT JOIN classrooms c ON s.classroom_id = c.classroom_id
	WHERE s.group_id = $1 AND s.is_actual = TRUE`

	rows, err := r.db.Query(ctx, query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 200
	schedules := make([]domain.ScheduleInfo, 0, defaultCapacity)
	for rows.Next() {
		var scheduleInfo domain.ScheduleInfo
		err := rows.Scan(
			&scheduleInfo.Schedule.ScheduleID, &scheduleInfo.Schedule.GroupID, &scheduleInfo.Schedule.DisciplineID, &scheduleInfo.Schedule.TeacherID, &scheduleInfo.Schedule.DisciplineTypeID, &scheduleInfo.Schedule.ClassroomID, &scheduleInfo.Schedule.Semester, &scheduleInfo.Schedule.WeekType, &scheduleInfo.Schedule.DayOfWeek, &scheduleInfo.Schedule.StartTime, &scheduleInfo.Schedule.IsActual,
			&scheduleInfo.ScheduleSub.DisciplineName, &scheduleInfo.ScheduleSub.TeacherFullName.LastName, &scheduleInfo.ScheduleSub.TeacherFullName.FirstName, &scheduleInfo.ScheduleSub.TeacherFullName.MiddleName, &scheduleInfo.ScheduleSub.DisciplineTypeName, &scheduleInfo.ScheduleSub.ClassroomName)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, scheduleInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *ScheduleRepo) GetActualByTeacherID(ctx context.Context, teacherID int64) ([]domain.ScheduleInfo, error) {
	query := `SELECT 
		s.schedule_id, s.group_id, s.discipline_id, s.teacher_id, s.discipline_type_id, s.classroom_id, s.semester, s.week_type, s.day_of_week, s.start_time, s.is_actual,
		d.discipline_name, t.last_name, t.first_name, t.middle_name, dt.discipline_type_name, c.classroom_name
	FROM schedules s
	LEFT JOIN disciplines d ON s.discipline_id = d.discipline_id
	LEFT JOIN teachers t ON s.teacher_id = t.teacher_id
	LEFT JOIN disciplineTypes dt ON s.discipline_type_id = dt.discipline_type_id
	LEFT JOIN classrooms c ON s.classroom_id = c.classroom_id
	WHERE s.teacher_id = $1 AND s.is_actual = TRUE`

	rows, err := r.db.Query(ctx, query, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 200
	schedules := make([]domain.ScheduleInfo, 0, defaultCapacity)
	for rows.Next() {
		var scheduleInfo domain.ScheduleInfo
		err := rows.Scan(
			&scheduleInfo.Schedule.ScheduleID, &scheduleInfo.Schedule.GroupID, &scheduleInfo.Schedule.DisciplineID, &scheduleInfo.Schedule.TeacherID, &scheduleInfo.Schedule.DisciplineTypeID, &scheduleInfo.Schedule.ClassroomID, &scheduleInfo.Schedule.Semester, &scheduleInfo.Schedule.WeekType, &scheduleInfo.Schedule.DayOfWeek, &scheduleInfo.Schedule.StartTime, &scheduleInfo.Schedule.IsActual,
			&scheduleInfo.ScheduleSub.DisciplineName, &scheduleInfo.ScheduleSub.TeacherFullName.LastName, &scheduleInfo.ScheduleSub.TeacherFullName.FirstName, &scheduleInfo.ScheduleSub.TeacherFullName.MiddleName, &scheduleInfo.ScheduleSub.DisciplineTypeName, &scheduleInfo.ScheduleSub.ClassroomName)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, scheduleInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

// Методы для получения расписания, сгруппированного по семестру, типу недели и дню недели

func (r *ScheduleRepo) GetGroupedByGroupID(ctx context.Context, groupID string) (map[int]map[string]map[string][]domain.ScheduleInfo, error) {
	query := `SELECT 
		s.schedule_id, s.group_id, s.discipline_id, s.teacher_id, s.discipline_type_id, s.classroom_id, s.semester, s.week_type, s.day_of_week, s.start_time, s.is_actual,
		d.discipline_name, t.last_name, t.first_name, t.middle_name, dt.discipline_type_name, c.classroom_name
	FROM schedules s
	LEFT JOIN disciplines d ON s.discipline_id = d.discipline_id
	LEFT JOIN teachers t ON s.teacher_id = t.teacher_id
	LEFT JOIN disciplineTypes dt ON s.discipline_type_id = dt.discipline_type_id
	LEFT JOIN classrooms c ON s.classroom_id = c.classroom_id
	WHERE s.group_id = $1
	ORDER BY s.semester, s.week_type, s.day_of_week`

	rows, err := r.db.Query(ctx, query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int]map[string]map[string][]domain.ScheduleInfo)
	for rows.Next() {
		var scheduleInfo domain.ScheduleInfo
		err := rows.Scan(
			&scheduleInfo.Schedule.ScheduleID, &scheduleInfo.Schedule.GroupID, &scheduleInfo.Schedule.DisciplineID, &scheduleInfo.Schedule.TeacherID, &scheduleInfo.Schedule.DisciplineTypeID, &scheduleInfo.Schedule.ClassroomID, &scheduleInfo.Schedule.Semester, &scheduleInfo.Schedule.WeekType, &scheduleInfo.Schedule.DayOfWeek, &scheduleInfo.Schedule.StartTime, &scheduleInfo.Schedule.IsActual,
			&scheduleInfo.ScheduleSub.DisciplineName, &scheduleInfo.ScheduleSub.TeacherFullName.LastName, &scheduleInfo.ScheduleSub.TeacherFullName.FirstName, &scheduleInfo.ScheduleSub.TeacherFullName.MiddleName, &scheduleInfo.ScheduleSub.DisciplineTypeName, &scheduleInfo.ScheduleSub.ClassroomName)
		if err != nil {
			return nil, err
		}

		if _, ok := result[scheduleInfo.Schedule.Semester]; !ok {
			result[scheduleInfo.Schedule.Semester] = make(map[string]map[string][]domain.ScheduleInfo)
		}
		if _, ok := result[scheduleInfo.Schedule.Semester][scheduleInfo.Schedule.WeekType]; !ok {
			result[scheduleInfo.Schedule.Semester][scheduleInfo.Schedule.WeekType] = make(map[string][]domain.ScheduleInfo)
		}
		result[scheduleInfo.Schedule.Semester][scheduleInfo.Schedule.WeekType][scheduleInfo.Schedule.DayOfWeek] = append(result[scheduleInfo.Schedule.Semester][scheduleInfo.Schedule.WeekType][scheduleInfo.Schedule.DayOfWeek], scheduleInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *ScheduleRepo) GetGroupedByTeacherID(ctx context.Context, teacherID int64) (map[int]map[string]map[string][]domain.ScheduleInfo, error) {
	query := `SELECT 
		s.schedule_id, s.group_id, s.discipline_id, s.teacher_id, s.discipline_type_id, s.classroom_id, s.semester, s.week_type, s.day_of_week, s.start_time, s.is_actual,
		d.discipline_name, t.last_name, t.first_name, t.middle_name, dt.discipline_type_name, c.classroom_name
	FROM schedules s
	LEFT JOIN disciplines d ON s.discipline_id = d.discipline_id
	LEFT JOIN teachers t ON s.teacher_id = t.teacher_id
	LEFT JOIN disciplineTypes dt ON s.discipline_type_id = dt.discipline_type_id
	LEFT JOIN classrooms c ON s.classroom_id = c.classroom_id
	WHERE s.teacher_id = $1
	ORDER BY s.semester, s.week_type, s.day_of_week`

	rows, err := r.db.Query(ctx, query, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int]map[string]map[string][]domain.ScheduleInfo)
	for rows.Next() {
		var scheduleInfo domain.ScheduleInfo
		err := rows.Scan(
			&scheduleInfo.Schedule.ScheduleID, &scheduleInfo.Schedule.GroupID, &scheduleInfo.Schedule.DisciplineID, &scheduleInfo.Schedule.TeacherID, &scheduleInfo.Schedule.DisciplineTypeID, &scheduleInfo.Schedule.ClassroomID, &scheduleInfo.Schedule.Semester, &scheduleInfo.Schedule.WeekType, &scheduleInfo.Schedule.DayOfWeek, &scheduleInfo.Schedule.StartTime, &scheduleInfo.Schedule.IsActual,
			&scheduleInfo.ScheduleSub.DisciplineName, &scheduleInfo.ScheduleSub.TeacherFullName.LastName, &scheduleInfo.ScheduleSub.TeacherFullName.FirstName, &scheduleInfo.ScheduleSub.TeacherFullName.MiddleName, &scheduleInfo.ScheduleSub.DisciplineTypeName, &scheduleInfo.ScheduleSub.ClassroomName)
		if err != nil {
			return nil, err
		}

		if _, ok := result[scheduleInfo.Schedule.Semester]; !ok {
			result[scheduleInfo.Schedule.Semester] = make(map[string]map[string][]domain.ScheduleInfo)
		}
		if _, ok := result[scheduleInfo.Schedule.Semester][scheduleInfo.Schedule.WeekType]; !ok {
			result[scheduleInfo.Schedule.Semester][scheduleInfo.Schedule.WeekType] = make(map[string][]domain.ScheduleInfo)
		}
		result[scheduleInfo.Schedule.Semester][scheduleInfo.Schedule.WeekType][scheduleInfo.Schedule.DayOfWeek] = append(result[scheduleInfo.Schedule.Semester][scheduleInfo.Schedule.WeekType][scheduleInfo.Schedule.DayOfWeek], scheduleInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *ScheduleRepo) GetActualByGroupAndWeekType(ctx context.Context, groupID string, weekType string) ([]domain.ScheduleInfo, error) {
	query := `SELECT 
		s.schedule_id, s.group_id, s.discipline_id, s.teacher_id, s.discipline_type_id, s.classroom_id, s.semester, s.week_type, s.day_of_week, s.start_time, s.is_actual,
		d.discipline_name, t.last_name, t.first_name, t.middle_name, dt.discipline_type_name, c.classroom_name
	FROM schedules s
	LEFT JOIN disciplines d ON s.discipline_id = d.discipline_id
	LEFT JOIN teachers t ON s.teacher_id = t.teacher_id
	LEFT JOIN disciplineTypes dt ON s.discipline_type_id = dt.discipline_type_id
	LEFT JOIN classrooms c ON s.classroom_id = c.classroom_id
	WHERE s.group_id = $1 AND s.week_type = $2 AND s.is_actual = TRUE`

	rows, err := r.db.Query(ctx, query, groupID, weekType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 200
	schedules := make([]domain.ScheduleInfo, 0, defaultCapacity)
	for rows.Next() {
		var scheduleInfo domain.ScheduleInfo
		err := rows.Scan(
			&scheduleInfo.Schedule.ScheduleID, &scheduleInfo.Schedule.GroupID, &scheduleInfo.Schedule.DisciplineID, &scheduleInfo.Schedule.TeacherID, &scheduleInfo.Schedule.DisciplineTypeID, &scheduleInfo.Schedule.ClassroomID, &scheduleInfo.Schedule.Semester, &scheduleInfo.Schedule.WeekType, &scheduleInfo.Schedule.DayOfWeek, &scheduleInfo.Schedule.StartTime, &scheduleInfo.Schedule.IsActual,
			&scheduleInfo.ScheduleSub.DisciplineName, &scheduleInfo.ScheduleSub.TeacherFullName.LastName, &scheduleInfo.ScheduleSub.TeacherFullName.FirstName, &scheduleInfo.ScheduleSub.TeacherFullName.MiddleName, &scheduleInfo.ScheduleSub.DisciplineTypeName, &scheduleInfo.ScheduleSub.ClassroomName)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, scheduleInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *ScheduleRepo) GetActualByTeacherAndWeekType(ctx context.Context, teacherID int64, weekType string) ([]domain.ScheduleInfo, error) {
	query := `SELECT 
		s.schedule_id, s.group_id, s.discipline_id, s.teacher_id, s.discipline_type_id, s.classroom_id, s.semester, s.week_type, s.day_of_week, s.start_time, s.is_actual,
		d.discipline_name, t.last_name, t.first_name, t.middle_name, dt.discipline_type_name, c.classroom_name
	FROM schedules s
	LEFT JOIN disciplines d ON s.discipline_id = d.discipline_id
	LEFT JOIN teachers t ON s.teacher_id = t.teacher_id
	LEFT JOIN disciplineTypes dt ON s.discipline_type_id = dt.discipline_type_id
	LEFT JOIN classrooms c ON s.classroom_id = c.classroom_id
	WHERE s.teacher_id = $1 AND s.week_type = $2 AND s.is_actual = TRUE`

	rows, err := r.db.Query(ctx, query, teacherID, weekType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 200
	schedules := make([]domain.ScheduleInfo, 0, defaultCapacity)
	for rows.Next() {
		var scheduleInfo domain.ScheduleInfo
		err := rows.Scan(
			&scheduleInfo.Schedule.ScheduleID, &scheduleInfo.Schedule.GroupID, &scheduleInfo.Schedule.DisciplineID, &scheduleInfo.Schedule.TeacherID, &scheduleInfo.Schedule.DisciplineTypeID, &scheduleInfo.Schedule.ClassroomID, &scheduleInfo.Schedule.Semester, &scheduleInfo.Schedule.WeekType, &scheduleInfo.Schedule.DayOfWeek, &scheduleInfo.Schedule.StartTime, &scheduleInfo.Schedule.IsActual,
			&scheduleInfo.ScheduleSub.DisciplineName, &scheduleInfo.ScheduleSub.TeacherFullName.LastName, &scheduleInfo.ScheduleSub.TeacherFullName.FirstName, &scheduleInfo.ScheduleSub.TeacherFullName.MiddleName, &scheduleInfo.ScheduleSub.DisciplineTypeName, &scheduleInfo.ScheduleSub.ClassroomName)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, scheduleInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *ScheduleRepo) GetActualByGroupWeekTypeAndDay(ctx context.Context, groupID, weekType, dayOfWeek string) ([]domain.ScheduleInfo, error) {
	query := `SELECT 
		s.schedule_id, s.group_id, s.discipline_id, s.teacher_id, s.discipline_type_id, s.classroom_id, s.semester, s.week_type, s.day_of_week, s.start_time, s.is_actual,
		d.discipline_name, t.last_name, t.first_name, t.middle_name, dt.discipline_type_name, c.classroom_name
	FROM schedules s
	LEFT JOIN disciplines d ON s.discipline_id = d.discipline_id
	LEFT JOIN teachers t ON s.teacher_id = t.teacher_id
	LEFT JOIN disciplineTypes dt ON s.discipline_type_id = dt.discipline_type_id
	LEFT JOIN classrooms c ON s.classroom_id = c.classroom_id
	WHERE s.group_id = $1 AND s.week_type = $2 AND s.day_of_week = $3 AND s.is_actual = TRUE`

	rows, err := r.db.Query(ctx, query, groupID, weekType, dayOfWeek)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 200
	schedules := make([]domain.ScheduleInfo, 0, defaultCapacity)
	for rows.Next() {
		var scheduleInfo domain.ScheduleInfo
		err := rows.Scan(
			&scheduleInfo.Schedule.ScheduleID, &scheduleInfo.Schedule.GroupID, &scheduleInfo.Schedule.DisciplineID, &scheduleInfo.Schedule.TeacherID, &scheduleInfo.Schedule.DisciplineTypeID, &scheduleInfo.Schedule.ClassroomID, &scheduleInfo.Schedule.Semester, &scheduleInfo.Schedule.WeekType, &scheduleInfo.Schedule.DayOfWeek, &scheduleInfo.Schedule.StartTime, &scheduleInfo.Schedule.IsActual,
			&scheduleInfo.ScheduleSub.DisciplineName, &scheduleInfo.ScheduleSub.TeacherFullName.LastName, &scheduleInfo.ScheduleSub.TeacherFullName.FirstName, &scheduleInfo.ScheduleSub.TeacherFullName.MiddleName, &scheduleInfo.ScheduleSub.DisciplineTypeName, &scheduleInfo.ScheduleSub.ClassroomName)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, scheduleInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *ScheduleRepo) GetActualByTeacherWeekTypeAndDay(ctx context.Context, teacherID int64, weekType, dayOfWeek string) ([]domain.ScheduleInfo, error) {
	query := `SELECT 
		s.schedule_id, s.group_id, s.discipline_id, s.teacher_id, s.discipline_type_id, s.classroom_id, s.semester, s.week_type, s.day_of_week, s.start_time, s.is_actual,
		d.discipline_name, t.last_name, t.first_name, t.middle_name, dt.discipline_type_name, c.classroom_name
	FROM schedules s
	LEFT JOIN disciplines d ON s.discipline_id = d.discipline_id
	LEFT JOIN teachers t ON s.teacher_id = t.teacher_id
	LEFT JOIN disciplineTypes dt ON s.discipline_type_id = dt.discipline_type_id
	LEFT JOIN classrooms c ON s.classroom_id = c.classroom_id
	WHERE s.teacher_id = $1 AND s.week_type = $2 AND s.day_of_week = $3 AND s.is_actual = TRUE`

	rows, err := r.db.Query(ctx, query, teacherID, weekType, dayOfWeek)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 200
	schedules := make([]domain.ScheduleInfo, 0, defaultCapacity)
	for rows.Next() {
		var scheduleInfo domain.ScheduleInfo
		err := rows.Scan(
			&scheduleInfo.Schedule.ScheduleID, &scheduleInfo.Schedule.GroupID, &scheduleInfo.Schedule.DisciplineID, &scheduleInfo.Schedule.TeacherID, &scheduleInfo.Schedule.DisciplineTypeID, &scheduleInfo.Schedule.ClassroomID, &scheduleInfo.Schedule.Semester, &scheduleInfo.Schedule.WeekType, &scheduleInfo.Schedule.DayOfWeek, &scheduleInfo.Schedule.StartTime, &scheduleInfo.Schedule.IsActual,
			&scheduleInfo.ScheduleSub.DisciplineName, &scheduleInfo.ScheduleSub.TeacherFullName.LastName, &scheduleInfo.ScheduleSub.TeacherFullName.FirstName, &scheduleInfo.ScheduleSub.TeacherFullName.MiddleName, &scheduleInfo.ScheduleSub.DisciplineTypeName, &scheduleInfo.ScheduleSub.ClassroomName)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, scheduleInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}
