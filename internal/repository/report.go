package repository

import (
	"context"
	"time"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ReportRepo struct {
	db *pgxpool.Pool
}

func NewReportRepo(db *pgxpool.Pool) *ReportRepo {
	return &ReportRepo{db: db}
}

func (r *ReportRepo) GetActualReportByGroupIDCreated(ctx context.Context, groupID string, startRange time.Time, endRange time.Time) (*domain.AttendanceReport, error) {
	query := `	SELECT 
	u.university_name,
	u.head_last_name || ' ' || u.head_first_name || ' ' || u.head_middle_name AS university_head,
	f.faculty_name,
	f.head_last_name || ' ' || f.head_first_name || ' ' || f.head_middle_name AS faculty_head,
	d.departament_name,
	d.head_last_name || ' ' || d.head_first_name || ' ' || d.head_middle_name AS departament_head,
	g.group_id,
	s.specialty_name,
	e.education_level_name,
	p.profile_name,
	et.education_type_name
	FROM
	university u INNER JOIN faculties f
	ON u.university_id = f.university_id
	INNER JOIN departaments d
	ON d.faculty_id = f.faculty_id
	INNER JOIN specialties s
	ON s.departament_id = d.departament_id
	INNER JOIN educationLevels e
	ON e.education_level_id = s.education_level_id
	INNER JOIN profiles p
	ON p.specialty_code = s.specialty_code
	INNER JOIN educationTypes et
	ON et.education_type_id = p.education_type_id
	INNER JOIN groups g
	ON g.profile_id = p.profile_id AND g.group_id = $1`
	var reportHead domain.ReportHead = domain.ReportHead{}
	if err := r.db.QueryRow(ctx, query, groupID).Scan(&reportHead.UniversityName, &reportHead.UniversityHead,
		&reportHead.FacultyName, &reportHead.FacultyHead,
		&reportHead.DepartamentName, &reportHead.DepartamentHead,
		&reportHead.GroupID, &reportHead.SpecialtyName,
		&reportHead.EducationLevelName, &reportHead.ProfileName,
		&reportHead.EducationTypeName); err != nil {
		return nil, err
	}
	query = `WITH subatt AS (
		SELECT
			t.student_id,
			t.schedule_id,
			COUNT(CASE t.presence WHEN false THEN 1 END) AS passes,
			COUNT(CASE t.presence WHEN true THEN 1 END) AS visits,
			COUNT(t.presence) AS total,
			ROUND(CAST(COUNT(CASE t.presence WHEN true THEN 1 END) * 100.0 / COUNT(t.presence) AS NUMERIC), 2) AS percentage_of_visits
		FROM attendance t
		GROUP BY t.student_id, t.schedule_id
	)
	SELECT
		sch.semester,
		sch.week_type,
		sch.day_of_week,
		dis.discipline_name,
		dtype.discipline_type_name,
		sch.start_time,
		cr.classroom_name,
		teach.last_name || ' ' || teach.first_name || ' ' || teach.middle_name AS teacher_name,
		st.last_name || ' ' || st.first_name || ' ' || st.middle_name AS student_name,
		at.presence,
		at.late_arrival,
		at.respectfulness,
		at.reason,
		COALESCE(subatt.visits, 0) AS visits,
		COALESCE(subatt.passes, 0) AS passes,
		COALESCE(subatt.total, 0) AS total,
		COALESCE(subatt.percentage_of_visits, 0) AS percentage_of_visits,
		at.created
	FROM attendance at
	INNER JOIN students st ON st.student_id = at.student_id AND st.group_id = $1
	INNER JOIN schedules sch ON at.schedule_id = sch.schedule_id
	INNER JOIN disciplineTypes dtype ON sch.discipline_type_id = dtype.discipline_type_id
	INNER JOIN classrooms cr ON sch.classroom_id = cr.classroom_id
	INNER JOIN disciplines dis ON dis.discipline_id = sch.discipline_id
	INNER JOIN teachers teach ON teach.teacher_id = sch.teacher_id
	INNER JOIN subatt ON at.student_id = subatt.student_id AND at.schedule_id = subatt.schedule_id
	WHERE at.created >= $2 and at.created <= $3
	GROUP BY
		sch.semester,
		sch.week_type,
		sch.day_of_week,
		sch.start_time,
		dis.discipline_name,
		dtype.discipline_type_name,
		cr.classroom_name,
		teach.last_name, teach.first_name, teach.middle_name,
		st.last_name, st.first_name, st.middle_name,
		at.presence,
		at.late_arrival,
		at.respectfulness,
		at.reason,
		subatt.visits,
		subatt.passes,
		subatt.total,
		subatt.percentage_of_visits,
		at.created
	ORDER BY st.last_name ASC;`
	rows, err := r.db.Query(ctx, query, groupID, startRange, endRange)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reportData := make([]domain.ReportData, 0, 1000)
	for rows.Next() {
		var report domain.ReportData
		err := rows.Scan(
			&report.Semester,
			&report.WeekType,
			&report.DayOfWeek,
			&report.DisciplineName,
			&report.DisciplineTypeName,
			&report.StartTime,
			&report.ClassroomName,
			&report.TeacherName,
			&report.StudentName,
			&report.Presence,
			&report.LateArrival,
			&report.Respectfulness,
			&report.Reason,
			&report.Visits,
			&report.Passes,
			&report.Total,
			&report.PercentageOfVisits,
			&report.Created,
		)
		if err != nil {
			return nil, err
		}
		reportData = append(reportData, report)

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	var attendanceReport domain.AttendanceReport = domain.AttendanceReport{}
	attendanceReport.ReportHead = reportHead
	attendanceReport.ReportData = reportData
	return &attendanceReport, nil
}
