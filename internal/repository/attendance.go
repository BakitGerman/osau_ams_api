package repository

import (
	"context"
	"strconv"
	"time"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AttendanceRepo struct {
	db *pgxpool.Pool
}

func NewAttendanceRepo(db *pgxpool.Pool) *AttendanceRepo {
	return &AttendanceRepo{db: db}
}

func (r *AttendanceRepo) Create(ctx context.Context, attendance domain.Attendance) error {
	query := `INSERT INTO attendance (student_id, schedule_id, presence, late_arrival, respectfulness, reason, created)
              VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(ctx, query, attendance.StudentID, attendance.ScheduleID, attendance.Presence, attendance.LateArrival, attendance.Respectfulness, attendance.Reason, attendance.Created)

	return err
}

func (r *AttendanceRepo) Put(ctx context.Context, attendance domain.Attendance) error {
	query := `UPDATE attendance SET presence=$1, late_arrival=$2, respectfulness=$3, reason=$4 WHERE attendance_id=$5`
	_, err := r.db.Exec(ctx, query, attendance.Presence, attendance.LateArrival, attendance.Respectfulness, attendance.Reason, attendance.AttendanceID)

	return err
}

func (r *AttendanceRepo) Patch(ctx context.Context, attendanceID int64, updates map[string]interface{}) error {
	query := `UPDATE attendance SET`
	args := make([]interface{}, 0, len(updates)+1)
	argsCounter := 1

	for col, val := range updates {
		query += " " + col + " = $" + strconv.Itoa(argsCounter) + ","
		args = append(args, val)
		argsCounter++
	}

	query = query[:len(query)-1]
	query += " WHERE attendance_id = $" + strconv.Itoa(argsCounter)
	args = append(args, attendanceID)
	_, err := r.db.Exec(ctx, query, args...)

	return err
}

func (r *AttendanceRepo) Delete(ctx context.Context, attendanceID int64) error {
	query := `DELETE FROM attendance WHERE attendance_id = $1`
	_, err := r.db.Exec(ctx, query, attendanceID)

	return err
}

func (r *AttendanceRepo) GetByID(ctx context.Context, attendanceID int64) (domain.AttendanceInfo, error) {
	query := `SELECT 
			a.attendance_id, a.student_id, a.schedule_id, a.presence, a.late_arrival, a.respectfulness, a.reason, a.created,
			s.last_name, s.first_name, s.middle_name
		FROM 
			attendance a
		LEFT JOIN 
			students s ON a.student_id = s.student_id
		WHERE a.attendance_id = $1`

	attendanceInfo := domain.AttendanceInfo{}
	err := r.db.QueryRow(ctx, query, attendanceID).Scan(
		&attendanceInfo.Attendance.AttendanceID,
		&attendanceInfo.Attendance.StudentID,
		&attendanceInfo.Attendance.ScheduleID,
		&attendanceInfo.Attendance.Presence,
		&attendanceInfo.Attendance.LateArrival,
		&attendanceInfo.Attendance.Respectfulness,
		&attendanceInfo.Attendance.Reason,
		&attendanceInfo.Attendance.Created,
		&attendanceInfo.AttendanceSub.Student.LastName,
		&attendanceInfo.AttendanceSub.Student.FirstName,
		&attendanceInfo.AttendanceSub.Student.MiddleName,
	)

	return attendanceInfo, err
}

func (r *AttendanceRepo) GetByStudentID(ctx context.Context, studentID int64) ([]domain.AttendanceInfo, error) {
	query := `SELECT 
			a.attendance_id, a.student_id, a.schedule_id, a.presence, a.late_arrival, a.respectfulness, a.reason, a.created,
			s.last_name, s.first_name, s.middle_name
		FROM 
			attendance a
		LEFT JOIN 
			students s ON a.student_id = s.student_id
		WHERE a.student_id = $1`

	rows, err := r.db.Query(ctx, query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 100
	attendances := make([]domain.AttendanceInfo, 0, defaultCapacity)
	for rows.Next() {
		var attendanceInfo domain.AttendanceInfo
		err := rows.Scan(
			&attendanceInfo.Attendance.AttendanceID,
			&attendanceInfo.Attendance.StudentID,
			&attendanceInfo.Attendance.ScheduleID,
			&attendanceInfo.Attendance.Presence,
			&attendanceInfo.Attendance.LateArrival,
			&attendanceInfo.Attendance.Respectfulness,
			&attendanceInfo.Attendance.Reason,
			&attendanceInfo.Attendance.Created,
			&attendanceInfo.AttendanceSub.Student.LastName,
			&attendanceInfo.AttendanceSub.Student.FirstName,
			&attendanceInfo.AttendanceSub.Student.MiddleName,
		)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, attendanceInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return attendances, nil
}

func (r *AttendanceRepo) GetAll(ctx context.Context) ([]domain.AttendanceInfo, error) {
	query := `SELECT 
			a.attendance_id, a.student_id, a.schedule_id, a.presence, a.late_arrival, a.respectfulness, a.reason, a.created,
			s.last_name, s.first_name, s.middle_name
		FROM 
			attendance a
		LEFT JOIN 
			students s ON a.student_id = s.student_id`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 100
	attendances := make([]domain.AttendanceInfo, 0, defaultCapacity)
	for rows.Next() {
		var attendanceInfo domain.AttendanceInfo
		err := rows.Scan(
			&attendanceInfo.Attendance.AttendanceID,
			&attendanceInfo.Attendance.StudentID,
			&attendanceInfo.Attendance.ScheduleID,
			&attendanceInfo.Attendance.Presence,
			&attendanceInfo.Attendance.LateArrival,
			&attendanceInfo.Attendance.Respectfulness,
			&attendanceInfo.Attendance.Reason,
			&attendanceInfo.Attendance.Created,
			&attendanceInfo.AttendanceSub.Student.LastName,
			&attendanceInfo.AttendanceSub.Student.FirstName,
			&attendanceInfo.AttendanceSub.Student.MiddleName,
		)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, attendanceInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return attendances, nil
}

func (r *AttendanceRepo) GetAllByGroupIDAndCreated(ctx context.Context, groupID string, scheduleID int64, created time.Time) ([]domain.GroupAttendanceInfo, error) {
	query := `SELECT
	a.attendance_id,
	s.student_id,
	sc.schedule_id,
	a.presence,
	a.late_arrival,
	a.respectfulness,
	a.reason,
	a.created,
    s.last_name,
    s.first_name,
    s.middle_name,
	s.group_id
	FROM
		attendance a
	JOIN
		schedules sc
		ON sc.schedule_id = a.schedule_id AND sc.schedule_id = $1
	RIGHT JOIN
		students s ON s.student_id = a.student_id AND a.created = $2
	WHERE
		s.group_id = $3`

	rows, err := r.db.Query(ctx, query, scheduleID, created, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	const defaultCapacity = 100
	attendances := make([]domain.GroupAttendanceInfo, 0, defaultCapacity)
	for rows.Next() {
		var attendanceInfo domain.GroupAttendanceInfo
		err := rows.Scan(
			&attendanceInfo.Attendance.AttendanceID,
			&attendanceInfo.Attendance.StudentID,
			&attendanceInfo.Attendance.ScheduleID,
			&attendanceInfo.Attendance.Presence,
			&attendanceInfo.Attendance.LateArrival,
			&attendanceInfo.Attendance.Respectfulness,
			&attendanceInfo.Attendance.Reason,
			&attendanceInfo.Attendance.Created,
			&attendanceInfo.AttendanceSub.Student.LastName,
			&attendanceInfo.AttendanceSub.Student.FirstName,
			&attendanceInfo.AttendanceSub.Student.MiddleName,
			&attendanceInfo.AttendanceSub.Student.GroupID,
		)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, attendanceInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return attendances, nil
}

func (r *AttendanceRepo) getCountAttendance(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM attendance;`
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

// func (r *AttendanceRepo) getCountAttendanceByStudentID(ctx context.Context, studentID int64) (int64, error) {
// 	query := `SELECT COUNT(*) FROM attendance WHERE student_id = $1;`
// 	rows, err := r.db.Query(ctx, query, studentID)
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
