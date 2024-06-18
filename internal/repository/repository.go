package repository

import (
	"context"
	"time"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IStudent interface {
	Create(ctx context.Context, student domain.Student) error
	Put(ctx context.Context, student domain.Student) error
	Patch(ctx context.Context, studentID int64, updates map[string]interface{}) error
	Delete(ctx context.Context, studentID int64) error
	GetByID(ctx context.Context, studentID int64) (domain.Student, error)
	GetByName(ctx context.Context, lastName, firstName, middleName string) (domain.Student, error)
	GetAll(ctx context.Context) ([]domain.Student, error)
	GetAllByGroupID(ctx context.Context, groupID string) ([]domain.Student, error)
}

type ISchedule interface {
	Create(ctx context.Context, schedule domain.Schedule) error
	Put(ctx context.Context, schedule domain.Schedule) error
	Patch(ctx context.Context, scheduleID int64, updates map[string]interface{}) error
	Delete(ctx context.Context, scheduleID int64) error
	GetByID(ctx context.Context, scheduleID int64) (domain.ScheduleInfo, error)
	GetAll(ctx context.Context) ([]domain.ScheduleInfo, error)
	GetByGroupID(ctx context.Context, groupID string) ([]domain.ScheduleInfo, error)
	GetByTeacherID(ctx context.Context, teacherID int64) ([]domain.ScheduleInfo, error)
	GetByGroupAndWeekType(ctx context.Context, groupID, weekType string) ([]domain.ScheduleInfo, error)
	GetByTeacherAndWeekType(ctx context.Context, teacherID int64, weekType string) ([]domain.ScheduleInfo, error)
	GetByGroupWeekTypeAndDay(ctx context.Context, groupID, weekType, dayOfWeek string) ([]domain.ScheduleInfo, error)
	GetByTeacherWeekTypeAndDay(ctx context.Context, teacherID int64, weekType, dayOfWeek string) ([]domain.ScheduleInfo, error)
	GetActualByGroupID(ctx context.Context, groupID string) ([]domain.ScheduleInfo, error)
	GetActualByTeacherID(ctx context.Context, teacherID int64) ([]domain.ScheduleInfo, error)
	GetActualByTeacherWeekTypeAndDay(ctx context.Context, teacherID int64, weekType, dayOfWeek string) ([]domain.ScheduleInfo, error)
	GetActualByTeacherAndWeekType(ctx context.Context, teacherID int64, weekType string) ([]domain.ScheduleInfo, error)
	GetActualByGroupWeekTypeAndDay(ctx context.Context, groupID, weekType, dayOfWeek string) ([]domain.ScheduleInfo, error)
	GetActualByGroupAndWeekType(ctx context.Context, groupID string, weekType string) ([]domain.ScheduleInfo, error)
}

type IAttendance interface {
	Create(ctx context.Context, attendance domain.Attendance) error
	Put(ctx context.Context, attendance domain.Attendance) error
	Patch(ctx context.Context, attendanceID int64, updates map[string]interface{}) error
	Delete(ctx context.Context, attendanceID int64) error
	GetByID(ctx context.Context, attendanceID int64) (domain.AttendanceInfo, error)
	GetByStudentID(ctx context.Context, studentID int64) ([]domain.AttendanceInfo, error)
	GetAll(ctx context.Context) ([]domain.AttendanceInfo, error)
	GetAllByGroupIDAndCreated(ctx context.Context, groupID string, scheduleID int64, created time.Time) ([]domain.GroupAttendanceInfo, error)
}

type IUser interface {
	Create(ctx context.Context, user domain.User) error
	Put(ctx context.Context, user domain.User) error
	Patch(ctx context.Context, userID uuid.UUID, updates map[string]interface{}) error
	Delete(ctx context.Context, userID uuid.UUID) error
	GetByID(ctx context.Context, userID uuid.UUID) (domain.UserInfo, error)
	GetByName(ctx context.Context, username string) (domain.UserInfo, error)
	GetByStudentID(ctx context.Context, studentID int64) (domain.UserInfo, error)
	GetByTeacherID(ctx context.Context, teacherID int64) (domain.UserInfo, error)
	GetByHeadmanID(ctx context.Context, headmanID int64) (domain.UserInfo, error)
	GetAllByRole(ctx context.Context, role string) ([]domain.UserInfo, error)
	GetAll(ctx context.Context) ([]domain.UserInfo, error)
}

type IHeadman interface {
	Create(ctx context.Context, headman domain.Headman) error
	Put(ctx context.Context, headman domain.Headman) error
	Patch(ctx context.Context, headmanID int64, updates map[string]interface{}) error
	Delete(ctx context.Context, headmanID int64) error
	GetByID(ctx context.Context, headmanID int64) (domain.HeadmanInfo, error)
	GetByStudentID(ctx context.Context, studentID int64) (domain.HeadmanInfo, error)
	GetAll(ctx context.Context) ([]domain.HeadmanInfo, error)
}

type IUniversity interface {
	Create(ctx context.Context, university domain.University) error
	Put(ctx context.Context, university domain.University) error
	Patch(ctx context.Context, universityID int64, updates map[string]interface{}) error
	Delete(ctx context.Context, universityID int64) error
	GetByID(ctx context.Context, universityID int64) (domain.University, error)
	GetByName(ctx context.Context, universityName string) (domain.University, error)
	GetAll(ctx context.Context) ([]domain.University, error)
}

type IFaculty interface {
	Create(ctx context.Context, faculty domain.Faculty) error
	Put(ctx context.Context, faculty domain.Faculty) error
	Patch(ctx context.Context, facultyID int64, updates map[string]interface{}) error
	Delete(ctx context.Context, facultyID int64) error
	GetByID(ctx context.Context, facultyID int64) (domain.FacultyInfo, error)
	GetByName(ctx context.Context, facultyName string) (domain.FacultyInfo, error)
	GetAll(ctx context.Context) ([]domain.FacultyInfo, error)
	GetAllByUniversityID(ctx context.Context, universityID int64) ([]domain.FacultyInfo, error)
}

type IDepartament interface {
	Create(ctx context.Context, departament domain.Departament) error
	Put(ctx context.Context, departament domain.Departament) error
	Patch(ctx context.Context, departamentID int64, updates map[string]interface{}) error
	Delete(ctx context.Context, departamentID int64) error
	GetByID(ctx context.Context, departamentID int64) (domain.DepartamentInfo, error)
	GetByName(ctx context.Context, departamentName string) (domain.DepartamentInfo, error)
	GetAll(ctx context.Context) ([]domain.DepartamentInfo, error)
	GetAllByFacultyID(ctx context.Context, facultyID int64) ([]domain.DepartamentInfo, error)
}

type ITeacher interface {
	Create(ctx context.Context, teacher domain.Teacher) error
	Put(ctx context.Context, teacher domain.Teacher) error
	Patch(ctx context.Context, teacherID int64, updates map[string]interface{}) error
	Delete(ctx context.Context, teacherID int64) error
	GetByID(ctx context.Context, teacherID int64) (domain.TeacherInfo, error)
	GetByEmail(ctx context.Context, teacherEmail string) (domain.TeacherInfo, error)
	GetAll(ctx context.Context) ([]domain.TeacherInfo, error)
	GetAllByDepartamentID(ctx context.Context, departamentID int64) ([]domain.TeacherInfo, error)
}

type IDiscipline interface {
	Create(ctx context.Context, discipline domain.Discipline) error
	Put(ctx context.Context, discipline domain.Discipline) error
	Patch(ctx context.Context, disciplineID int64, updates map[string]interface{}) error
	Delete(ctx context.Context, disciplineID int64) error
	GetByID(ctx context.Context, disciplineID int64) (domain.DisciplineInfo, error)
	GetByName(ctx context.Context, disciplineName string) (domain.DisciplineInfo, error)
	GetAll(ctx context.Context) ([]domain.DisciplineInfo, error)
	GetAllByDepartamentID(ctx context.Context, departamentID int64) ([]domain.DisciplineInfo, error)
}

type IDisciplineType interface {
	Create(ctx context.Context, disciplineType domain.DisciplineType) error
	Put(ctx context.Context, disciplineType domain.DisciplineType) error
	Patch(ctx context.Context, disciplineTypeID int64, updates map[string]interface{}) error
	Delete(ctx context.Context, disciplineTypeID int64) error
	GetByID(ctx context.Context, disciplineTypeID int64) (domain.DisciplineType, error)
	GetAll(ctx context.Context) ([]domain.DisciplineType, error)
}

type IClassroom interface {
	Create(ctx context.Context, classroom domain.Classroom) error
	Put(ctx context.Context, classroom domain.Classroom) error
	Patch(ctx context.Context, classroomID int64, updates map[string]interface{}) error
	Delete(ctx context.Context, classroomID int64) error
	GetByID(ctx context.Context, classroomID int64) (domain.Classroom, error)
	GetAll(ctx context.Context) ([]domain.Classroom, error)
}

type IEducationLevel interface {
	Create(ctx context.Context, educationLevel domain.EducationLevel) error
	Put(ctx context.Context, educationLevel domain.EducationLevel) error
	Patch(ctx context.Context, educationLevelID int64, updates map[string]interface{}) error
	Delete(ctx context.Context, educationLevelID int64) error
	GetByID(ctx context.Context, educationLevelID int64) (domain.EducationLevel, error)
	GetAll(ctx context.Context) ([]domain.EducationLevel, error)
}

type ISpecialty interface {
	Create(ctx context.Context, specialty domain.Specialty) error
	Put(ctx context.Context, specialty domain.Specialty) error
	Patch(ctx context.Context, specialtyCode string, updates map[string]interface{}) error
	Delete(ctx context.Context, specialtyCode string) error
	GetByCode(ctx context.Context, specialtyCode string) (domain.SpecialtyInfo, error)
	GetByName(ctx context.Context, specialtyName string) (domain.SpecialtyInfo, error)
	GetAll(ctx context.Context) ([]domain.SpecialtyInfo, error)
	GetAllByDepartamentID(ctx context.Context, departamentID int64) ([]domain.SpecialtyInfo, error)
}

type IProfile interface {
	Create(ctx context.Context, profile domain.Profile) error
	Put(ctx context.Context, profile domain.Profile) error
	Patch(ctx context.Context, profileID int64, updates map[string]interface{}) error
	Delete(ctx context.Context, profileID int64) error
	GetByID(ctx context.Context, profileID int64) (domain.ProfileInfo, error)
	GetByName(ctx context.Context, profileName string) (domain.ProfileInfo, error)
	GetAll(ctx context.Context) ([]domain.ProfileInfo, error)
	GetAllBySpecialtyCode(ctx context.Context, specialtyCode string) ([]domain.ProfileInfo, error)
	GetByEducationTypeID(ctx context.Context, educationTypeID int64) ([]domain.ProfileInfo, error)
}

type IGroup interface {
	Create(ctx context.Context, group domain.Group) error
	Put(ctx context.Context, group domain.Group) error
	Patch(ctx context.Context, groupID string, updates map[string]interface{}) error
	Delete(ctx context.Context, groupID string) error
	GetByID(ctx context.Context, groupID string) (domain.GroupInfo, error)
	GetByName(ctx context.Context, profileName string) (domain.GroupInfo, error)
	GetAll(ctx context.Context) ([]domain.GroupInfo, error)
	GetAllByProfileID(ctx context.Context, profileID int64) ([]domain.GroupInfo, error)
}

type IEducationType interface {
	Create(ctx context.Context, educationType domain.EducationType) error
	Put(ctx context.Context, educationType domain.EducationType) error
	Patch(ctx context.Context, educationTypeID int64, updates map[string]interface{}) error
	Delete(ctx context.Context, educationTypeID int64) error
	GetByID(ctx context.Context, educationTypeID int64) (domain.EducationType, error)
	GetByName(ctx context.Context, educationTypeName string) (domain.EducationType, error)
	GetAll(ctx context.Context) ([]domain.EducationType, error)
}

type IReport interface {
	GetActualReportByGroupIDCreated(ctx context.Context, groupID string, startRange time.Time, endRange time.Time) (*domain.AttendanceReport, error)
}

type Repositories struct {
	Student        IStudent
	Schedule       ISchedule
	Headman        IHeadman
	Attendance     IAttendance
	User           IUser
	University     IUniversity
	Faculty        IFaculty
	Departament    IDepartament
	Teacher        ITeacher
	Discipline     IDiscipline
	DisciplineType IDisciplineType
	Classroom      IClassroom
	EducationLevel IEducationLevel
	Specialty      ISpecialty
	Profile        IProfile
	Group          IGroup
	EducationType  IEducationType
	Report         IReport
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		Student:        NewStudentRepo(db),
		Schedule:       NewScheduleRepo(db),
		Headman:        NewHeadmanRepo(db),
		Attendance:     NewAttendanceRepo(db),
		User:           NewUserRepo(db),
		University:     NewUniversityRepo(db),
		Faculty:        NewFacultyRepo(db),
		Departament:    NewDepartamentRepo(db),
		Teacher:        NewTeacherRepo(db),
		Discipline:     NewDisciplineRepo(db),
		DisciplineType: NewDisciplineTypeRepo(db),
		Classroom:      NewClassroomRepo(db),
		EducationLevel: NewEducationLevelRepo(db),
		Specialty:      NewSpecialtyRepo(db),
		Profile:        NewProfileRepo(db),
		Group:          NewGroupRepo(db),
		EducationType:  NewEducationTypeRepo(db),
		Report:         NewReportRepo(db),
	}
}
