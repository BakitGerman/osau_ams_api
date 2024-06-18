package repository

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user domain.User) error {
	query := `INSERT INTO users (username, password, user_role, headman_id, student_id, teacher_id)
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(ctx, query, user.Username, user.Password, user.Role, user.HeadmanID, user.StudentID, user.TeacherID)

	return err
}

func (r *UserRepo) Put(ctx context.Context, user domain.User) error {
	query := `UPDATE users SET username=$1, password=$2, user_role=$3, headman_id=$4, student_id=$5, teacher_id=$6 WHERE user_id=$7`
	_, err := r.db.Exec(ctx, query, user.Username, user.Password, user.Role, user.HeadmanID, user.StudentID, user.TeacherID, user.UserID)

	return err
}

func (r *UserRepo) Patch(ctx context.Context, userID uuid.UUID, updates map[string]interface{}) error {
	query := `UPDATE users SET`
	args := make([]interface{}, 0, len(updates)+1)
	argsCounter := 1

	for col, val := range updates {
		query += " " + col + " = $" + strconv.Itoa(argsCounter) + ","
		args = append(args, val)
		argsCounter++
	}

	query = query[:len(query)-1]
	query += " WHERE user_id = $" + strconv.Itoa(argsCounter)
	args = append(args, userID)
	_, err := r.db.Exec(ctx, query, args...)

	return err
}

func (r *UserRepo) Delete(ctx context.Context, userID uuid.UUID) error {
	query := `DELETE FROM users WHERE user_id = $1`
	_, err := r.db.Exec(ctx, query, userID)

	return err
}

func (r *UserRepo) GetByID(ctx context.Context, userID uuid.UUID) (domain.UserInfo, error) {
	query := `SELECT 
			u.user_id,
			u.username,
			u.password,
			u.user_role,
			u.headman_id,
			u.student_id,
			u.teacher_id,
			s.last_name,
			s.first_name,
			s.middle_name,
			s.group_id,
			t.last_name,
			t.first_name,
			t.middle_name
		FROM 
			users u
		LEFT JOIN 
			headmans h ON u.headman_id = h.headman_id
		LEFT JOIN 
			students s ON u.student_id = s.student_id OR h.student_id = s.student_id
		LEFT JOIN 
			teachers t ON u.teacher_id = t.teacher_id
		WHERE u.user_id = $1`
	user := domain.UserInfo{
		UserSub: domain.UserSub{
			StudentFullName: &domain.StudentFullName{},
			TeacherFullName: &domain.TeacherFullName{},
		},
	}
	var studentLastName sql.NullString
	var studentFirstName sql.NullString
	var studentMiddleName sql.NullString
	var teacherLastName sql.NullString
	var teacherFirstName sql.NullString
	var teacherMiddleName sql.NullString
	err := r.db.QueryRow(ctx, query, userID).Scan(
		&user.User.UserID,
		&user.User.Username,
		&user.User.Password,
		&user.User.Role,
		&user.User.HeadmanID,
		&user.User.StudentID,
		&user.User.TeacherID,
		&studentLastName,
		&studentFirstName,
		&studentMiddleName,
		&user.UserSub.GroupID,
		&teacherLastName,
		&teacherFirstName,
		&teacherMiddleName,
	)

	if err != nil {
		return user, err
	}

	if studentFirstName.Valid && studentLastName.Valid && studentMiddleName.Valid {
		user.UserSub.StudentFullName.FirstName = studentFirstName.String
		user.UserSub.StudentFullName.LastName = studentLastName.String
		user.UserSub.StudentFullName.MiddleName = studentMiddleName.String
	} else {
		user.UserSub.StudentFullName = nil
	}
	if teacherFirstName.Valid && teacherLastName.Valid && teacherMiddleName.Valid {
		user.UserSub.TeacherFullName.FirstName = teacherFirstName.String
		user.UserSub.TeacherFullName.LastName = teacherLastName.String
		user.UserSub.TeacherFullName.MiddleName = teacherMiddleName.String
	} else {
		user.UserSub.TeacherFullName = nil
	}

	return user, nil
}

func (r *UserRepo) GetByName(ctx context.Context, username string) (domain.UserInfo, error) {
	query := `SELECT 
    u.user_id,
    u.username,
    u.password,
    u.user_role,
    u.headman_id,
    u.student_id,
    u.teacher_id,
	s.last_name,
	s.first_name,
	s.middle_name,
	h.group_id, 
	t.last_name,
	t.first_name,
	t.middle_name
	FROM 
		users u
	LEFT JOIN 
		headmans h ON u.headman_id = h.headman_id
	LEFT JOIN 
		students s ON u.student_id = s.student_id OR h.student_id = s.student_id
	LEFT JOIN 
		teachers t ON u.teacher_id = t.teacher_id
	WHERE 
		u.username = $1`
	user := domain.UserInfo{
		UserSub: domain.UserSub{
			StudentFullName: &domain.StudentFullName{},
			TeacherFullName: &domain.TeacherFullName{},
		},
	}

	var studentLastName sql.NullString
	var studentFirstName sql.NullString
	var studentMiddleName sql.NullString
	var teacherLastName sql.NullString
	var teacherFirstName sql.NullString
	var teacherMiddleName sql.NullString
	err := r.db.QueryRow(ctx, query, username).Scan(
		&user.User.UserID,
		&user.User.Username,
		&user.User.Password,
		&user.User.Role,
		&user.User.HeadmanID,
		&user.User.StudentID,
		&user.User.TeacherID,
		&studentLastName,
		&studentFirstName,
		&studentMiddleName,
		&user.UserSub.GroupID,
		&teacherLastName,
		&teacherFirstName,
		&teacherMiddleName,
	)

	if err != nil {
		return user, err
	}

	if studentFirstName.Valid && studentLastName.Valid && studentMiddleName.Valid {
		user.UserSub.StudentFullName.FirstName = studentFirstName.String
		user.UserSub.StudentFullName.LastName = studentLastName.String
		user.UserSub.StudentFullName.MiddleName = studentMiddleName.String
	} else {
		user.UserSub.StudentFullName = nil
	}
	if teacherFirstName.Valid && teacherLastName.Valid && teacherMiddleName.Valid {
		user.UserSub.TeacherFullName.FirstName = teacherFirstName.String
		user.UserSub.TeacherFullName.LastName = teacherLastName.String
		user.UserSub.TeacherFullName.MiddleName = teacherMiddleName.String
	} else {
		user.UserSub.TeacherFullName = nil
	}

	return user, nil
}

func (r *UserRepo) GetByStudentID(ctx context.Context, studentID int64) (domain.UserInfo, error) {
	query := `SELECT 
			  u.user_id,
			  u.username,
			  u.password,
			  u.user_role,
			  u.headman_id,
			  u.student_id,
			  u.teacher_id,
			  s.last_name,
			  s.first_name,
			  s.middle_name,
			  h.group_id,
			  t.last_name,
			  t.first_name,
			  t.middle_name
		  FROM 
			  users u
		  LEFT JOIN 
			  headmans h ON u.headman_id = h.headman_id
		  LEFT JOIN 
			  students s ON u.student_id = s.student_id OR h.student_id = s.student_id
		  LEFT JOIN 
			  teachers t ON u.teacher_id = t.teacher_id
			  WHERE s.student_id = $1`
	user := domain.UserInfo{
		UserSub: domain.UserSub{
			StudentFullName: &domain.StudentFullName{},
			TeacherFullName: &domain.TeacherFullName{},
		},
	}

	var studentLastName sql.NullString
	var studentFirstName sql.NullString
	var studentMiddleName sql.NullString
	var teacherLastName sql.NullString
	var teacherFirstName sql.NullString
	var teacherMiddleName sql.NullString
	err := r.db.QueryRow(ctx, query, studentID).Scan(
		&user.User.UserID,
		&user.User.Username,
		&user.User.Password,
		&user.User.Role,
		&user.User.HeadmanID,
		&user.User.StudentID,
		&user.User.TeacherID,
		&studentLastName,
		&studentFirstName,
		&studentMiddleName,
		&user.UserSub.GroupID,
		&teacherLastName,
		&teacherFirstName,
		&teacherMiddleName,
	)

	if err != nil {
		return user, err
	}

	if studentFirstName.Valid && studentLastName.Valid && studentMiddleName.Valid {
		user.UserSub.StudentFullName.FirstName = studentFirstName.String
		user.UserSub.StudentFullName.LastName = studentLastName.String
		user.UserSub.StudentFullName.MiddleName = studentMiddleName.String
	} else {
		user.UserSub.StudentFullName = nil
	}
	if teacherFirstName.Valid && teacherLastName.Valid && teacherMiddleName.Valid {
		user.UserSub.TeacherFullName.FirstName = teacherFirstName.String
		user.UserSub.TeacherFullName.LastName = teacherLastName.String
		user.UserSub.TeacherFullName.MiddleName = teacherMiddleName.String
	} else {
		user.UserSub.TeacherFullName = nil
	}

	return user, nil
}

func (r *UserRepo) GetByTeacherID(ctx context.Context, teacherID int64) (domain.UserInfo, error) {
	query := `SELECT 
			  u.user_id,
			  u.username,
			  u.password,
			  u.user_role,
			  u.headman_id,
			  u.student_id,
			  u.teacher_id,
			  s.last_name,
			  s.first_name,
			  s.middle_name,
			  h.group_id,
			  t.last_name,
			  t.first_name,
			  t.middle_name
		  FROM 
			  users u
		  LEFT JOIN 
			  headmans h ON u.headman_id = h.headman_id
		  LEFT JOIN 
			  students s ON u.student_id = s.student_id OR h.student_id = s.student_id
		  LEFT JOIN 
			  teachers t ON u.teacher_id = t.teacher_id
			  WHERE t.teacher_id = $1`
	user := domain.UserInfo{
		UserSub: domain.UserSub{
			StudentFullName: &domain.StudentFullName{},
			TeacherFullName: &domain.TeacherFullName{},
		},
	}

	var studentLastName sql.NullString
	var studentFirstName sql.NullString
	var studentMiddleName sql.NullString
	var teacherLastName sql.NullString
	var teacherFirstName sql.NullString
	var teacherMiddleName sql.NullString
	err := r.db.QueryRow(ctx, query, teacherID).Scan(
		&user.User.UserID,
		&user.User.Username,
		&user.User.Password,
		&user.User.Role,
		&user.User.HeadmanID,
		&user.User.StudentID,
		&user.User.TeacherID,
		&studentLastName,
		&studentFirstName,
		&studentMiddleName,
		&user.UserSub.GroupID,
		&teacherLastName,
		&teacherFirstName,
		&teacherMiddleName,
	)

	if err != nil {
		return user, err
	}

	if studentFirstName.Valid && studentLastName.Valid && studentMiddleName.Valid {
		user.UserSub.StudentFullName.FirstName = studentFirstName.String
		user.UserSub.StudentFullName.LastName = studentLastName.String
		user.UserSub.StudentFullName.MiddleName = studentMiddleName.String
	} else {
		user.UserSub.StudentFullName = nil
	}
	if teacherFirstName.Valid && teacherLastName.Valid && teacherMiddleName.Valid {
		user.UserSub.TeacherFullName.FirstName = teacherFirstName.String
		user.UserSub.TeacherFullName.LastName = teacherLastName.String
		user.UserSub.TeacherFullName.MiddleName = teacherMiddleName.String
	} else {
		user.UserSub.TeacherFullName = nil
	}

	return user, nil
}

func (r *UserRepo) GetByHeadmanID(ctx context.Context, headmanID int64) (domain.UserInfo, error) {
	query := `SELECT 
			  u.user_id,
			  u.username,
			  u.password,
			  u.user_role,
			  u.headman_id,
			  u.student_id,
			  u.teacher_id,
			  s.last_name,
			  s.first_name,
			  s.middle_name,
			  h.group_id,
			  t.last_name,
			  t.first_name,
			  t.middle_name
		  FROM 
			  users u
		  LEFT JOIN 
			  headmans h ON u.headman_id = h.headman_id
		  LEFT JOIN 
			  students s ON h.student_id = s.student_id
		  LEFT JOIN 
			  teachers t ON u.teacher_id = t.teacher_id
			  WHERE h.headman_id = $1`
	user := domain.UserInfo{
		UserSub: domain.UserSub{
			StudentFullName: &domain.StudentFullName{},
			TeacherFullName: &domain.TeacherFullName{},
		},
	}

	var studentLastName sql.NullString
	var studentFirstName sql.NullString
	var studentMiddleName sql.NullString
	var teacherLastName sql.NullString
	var teacherFirstName sql.NullString
	var teacherMiddleName sql.NullString
	err := r.db.QueryRow(ctx, query, headmanID).Scan(
		&user.User.UserID,
		&user.User.Username,
		&user.User.Password,
		&user.User.Role,
		&user.User.HeadmanID,
		&user.User.StudentID,
		&user.User.TeacherID,
		&studentLastName,
		&studentFirstName,
		&studentMiddleName,
		&user.UserSub.GroupID,
		&teacherLastName,
		&teacherFirstName,
		&teacherMiddleName,
	)

	if err != nil {
		return user, err
	}

	if studentFirstName.Valid && studentLastName.Valid && studentMiddleName.Valid {
		user.UserSub.StudentFullName.FirstName = studentFirstName.String
		user.UserSub.StudentFullName.LastName = studentLastName.String
		user.UserSub.StudentFullName.MiddleName = studentMiddleName.String
	} else {
		user.UserSub.StudentFullName = nil
	}
	if teacherFirstName.Valid && teacherLastName.Valid && teacherMiddleName.Valid {
		user.UserSub.TeacherFullName.FirstName = teacherFirstName.String
		user.UserSub.TeacherFullName.LastName = teacherLastName.String
		user.UserSub.TeacherFullName.MiddleName = teacherMiddleName.String
	} else {
		user.UserSub.TeacherFullName = nil
	}

	return user, nil
}

func (r *UserRepo) GetAllByRole(ctx context.Context, role string) ([]domain.UserInfo, error) {
	query := `SELECT 
	u.user_id,
	u.username,
	u.password,
	u.user_role,
	u.headman_id,
	u.student_id,
	u.teacher_id,
	s.last_name,
	s.first_name,
	s.middle_name,
	h.group_id, 
	t.last_name,
	t.first_name,
	t.middle_name
	FROM 
		users u
	LEFT JOIN 
		headmans h ON u.headman_id = h.headman_id
	LEFT JOIN 
		students s ON u.student_id = s.student_id OR h.student_id = s.student_id
	LEFT JOIN 
		teachers t ON u.teacher_id = t.teacher_id
		WHERE user_role = $1`
	rows, err := r.db.Query(ctx, query, role)
	if err != nil {
		return nil, err
	}

	const defaultCapacity = 200

	users := make([]domain.UserInfo, 0, defaultCapacity)
	for rows.Next() {
		user := domain.UserInfo{
			UserSub: domain.UserSub{
				StudentFullName: &domain.StudentFullName{},
				TeacherFullName: &domain.TeacherFullName{},
			},
		}

		var studentLastName sql.NullString
		var studentFirstName sql.NullString
		var studentMiddleName sql.NullString
		var teacherLastName sql.NullString
		var teacherFirstName sql.NullString
		var teacherMiddleName sql.NullString
		err := rows.Scan(&user.User.UserID,
			&user.User.Username,
			&user.User.Password,
			&user.User.Role,
			&user.User.HeadmanID,
			&user.User.StudentID,
			&user.User.TeacherID,
			&studentLastName,
			&studentFirstName,
			&studentMiddleName,
			&user.UserSub.GroupID,
			&teacherLastName,
			&teacherFirstName,
			&teacherMiddleName)

		if err != nil {
			return nil, err
		}

		if studentFirstName.Valid && studentLastName.Valid && studentMiddleName.Valid {
			user.UserSub.StudentFullName.FirstName = studentFirstName.String
			user.UserSub.StudentFullName.LastName = studentLastName.String
			user.UserSub.StudentFullName.MiddleName = studentMiddleName.String
		} else {
			user.UserSub.StudentFullName = nil
		}
		if teacherFirstName.Valid && teacherLastName.Valid && teacherMiddleName.Valid {
			user.UserSub.TeacherFullName.FirstName = teacherFirstName.String
			user.UserSub.TeacherFullName.LastName = teacherLastName.String
			user.UserSub.TeacherFullName.MiddleName = teacherMiddleName.String
		} else {
			user.UserSub.TeacherFullName = nil
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepo) GetAll(ctx context.Context) ([]domain.UserInfo, error) {
	query := `SELECT 
	u.user_id,
	u.username,
	u.password,
	u.user_role,
	u.headman_id,
	u.student_id,
	u.teacher_id,
	s.last_name,
	s.first_name,
	s.middle_name,
	h.group_id, 
	t.last_name,
	t.first_name,
	t.middle_name
	FROM 
		users u
	LEFT JOIN 
		headmans h ON u.headman_id = h.headman_id
	LEFT JOIN 
		students s ON u.student_id = s.student_id OR h.student_id = s.student_id
	LEFT JOIN 
		teachers t ON u.teacher_id = t.teacher_id`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	const defaultCapacity = 200

	users := make([]domain.UserInfo, 0, defaultCapacity)
	for rows.Next() {
		user := domain.UserInfo{
			UserSub: domain.UserSub{
				StudentFullName: &domain.StudentFullName{},
				TeacherFullName: &domain.TeacherFullName{},
			},
		}

		var studentLastName sql.NullString
		var studentFirstName sql.NullString
		var studentMiddleName sql.NullString
		var teacherLastName sql.NullString
		var teacherFirstName sql.NullString
		var teacherMiddleName sql.NullString
		err := rows.Scan(&user.User.UserID,
			&user.User.Username,
			&user.User.Password,
			&user.User.Role,
			&user.User.HeadmanID,
			&user.User.StudentID,
			&user.User.TeacherID,
			&studentLastName,
			&studentFirstName,
			&studentMiddleName,
			&user.UserSub.GroupID,
			&teacherLastName,
			&teacherFirstName,
			&teacherMiddleName)

		if err != nil {
			return nil, err
		}

		if studentFirstName.Valid && studentLastName.Valid && studentMiddleName.Valid {
			user.UserSub.StudentFullName.FirstName = studentFirstName.String
			user.UserSub.StudentFullName.LastName = studentLastName.String
			user.UserSub.StudentFullName.MiddleName = studentMiddleName.String
		} else {
			user.UserSub.StudentFullName = nil
		}
		if teacherFirstName.Valid && teacherLastName.Valid && teacherMiddleName.Valid {
			user.UserSub.TeacherFullName.FirstName = teacherFirstName.String
			user.UserSub.TeacherFullName.LastName = teacherLastName.String
			user.UserSub.TeacherFullName.MiddleName = teacherMiddleName.String
		} else {
			user.UserSub.TeacherFullName = nil
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// func (r *UserRepo) getCountUsers(ctx context.Context) (int64, error) {
// 	query := `SELECT COUNT(*) FROM users;`
// 	rows, err := r.db.Query(ctx, query)
// 	if err != nil {
// 		return 0, err
// 	}
// 	var count int64
// 	rows.Scan(&count)

// 	return count, nil
// }

// func (r *UserRepo) GetProfileByUsername(ctx context.Context, username string) (domain.UserProfile, error) {
// 	query := `SELECT
// 		u.username,
// 		u.user_role,
// 		CASE
// 			WHEN u.user_role = 'Студент' THEN CONCAT(s.last_name, ' ', s.first_name, ' ', s.middle_name)
// 			WHEN u.user_role = 'Староста' THEN CONCAT(s.last_name, ' ', s.first_name, ' ', s.middle_name)
// 			WHEN u.user_role = 'Преподаватель' THEN CONCAT(t.last_name, ' ', t.first_name, ' ', t.middle_name)
// 			ELSE 'Админ'
// 		END AS full_name,
// 		CASE
// 			WHEN u.user_role = 'Студент' THEN sp.specialty_name
// 			WHEN u.user_role = 'Староста' THEN sp.specialty_name
// 			ELSE NULL
// 		END AS specialty,
// 		CASE
// 			WHEN u.user_role = 'Студент' THEN p.profile_name
// 			WHEN u.user_role = 'Староста' THEN p.profile_name
// 			ELSE NULL
// 		END AS profile,
// 		CASE
// 			WHEN u.user_role = 'Студент' THEN g.group_id
// 			WHEN u.user_role = 'Староста' THEN g.group_id
// 			ELSE NULL
// 		END AS group_id,
// 		CASE
// 			WHEN u.user_role = 'Преподаватель' THEN d.departament_name
// 			ELSE NULL
// 		END AS departament,
// 		f.faculty_name,
// 		un.university_name
// 	FROM
// 		users u
// 	LEFT JOIN
// 		students s ON u.student_id = s.student_id
// 	LEFT JOIN
// 		headmans h ON u.headman_id = h.headman_id
// 	LEFT JOIN
// 		teachers t ON u.teacher_id = t.teacher_id
// 	LEFT JOIN
// 		groups g ON s.group_id = g.group_id OR h.group_id = g.group_id
// 	LEFT JOIN
// 		profiles p ON g.profile_id = p.profile_id
// 	LEFT JOIN
// 		specialties sp ON p.specialty_code = sp.specialty_code
// 	LEFT JOIN
// 		faculties f ON sp.departament_id = f.faculty_id
// 	LEFT JOIN
// 		departaments d ON t.departament_id = d.departament_id
// 	LEFT JOIN
// 		university un ON f.university_id = un.university_id
// 	WHERE
// 		u.username = $1`

// 	profile := domain.UserProfile{}
// 	err := r.db.QueryRow(ctx, query, username).Scan(
// 		&profile.Username,
// 		&profile.UserRole,
// 		&profile.FullName,
// 		&profile.Specialty,
// 		&profile.Profile,
// 		&profile.GroupID,
// 		&profile.Departament,
// 		&profile.FacultyName,
// 		&profile.UniversityName,
// 	)

// 	if err != nil {
// 		return profile, err
// 	}

// 	return profile, nil
// }
