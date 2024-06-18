package service

import (
	"time"

	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
	"github.com/BeRebornBng/OsauAmsApi/pkg/auth"
	"github.com/BeRebornBng/OsauAmsApi/pkg/myhash"
)

// support structs

type Support struct {
	Repos          *repository.Repositories
	Hasher         myhash.PasswordHasher
	TokenManager   auth.TokenManager
	AccessTokenTTL time.Duration
}

type Tokens struct {
	AccessToken string
}

type Services struct {
	ReportService         *ReportService
	HeadmanService        *HeadmanService
	StudentService        *StudentService
	ScheduleService       *ScheduleService
	AttendanceService     *AttendanceService
	UserService           *UserService
	UniversityService     *UniversityService
	FacultyService        *FacultyService
	DepartamentService    *DepartamentService
	TeacherService        *TeacherService
	DisciplineService     *DisciplineService
	DisciplineTypeService *DisciplineTypeService
	ClassroomService      *ClassroomService
	EducationLevelService *EducationLevelService
	SpecialtyService      *SpecialtyService
	ProfileService        *ProfileService
	GroupService          *GroupService
	EducationTypeService  *EducationTypeService
}

func NewServices(support Support) *Services {
	reportService := NewReportService(support.Repos.Report)
	headmanService := NewHeadmanService(support.Repos.Headman)
	studentService := NewStudentService(support.Repos.Student)
	scheduleService := NewScheduleService(support.Repos.Schedule)
	attendanceService := NewAttendanceService(support.Repos.Attendance)
	userService := NewUserService(support.TokenManager, support.Hasher, support.Repos.User, support.AccessTokenTTL)
	universityService := NewUniversityService(support.Repos.University)
	facultyService := NewFacultyService(support.Repos.Faculty)
	departamentService := NewDepartamentService(support.Repos.Departament)
	teacherService := NewTeacherService(support.Repos.Teacher)
	disciplineService := NewDisciplineService(support.Repos.Discipline)
	disciplineTypeService := NewDisciplineTypeService(support.Repos.DisciplineType)
	classroomService := NewClassroomService(support.Repos.Classroom)
	educationLevelService := NewEducationLevelService(support.Repos.EducationLevel)
	specialtyService := NewSpecialtyService(support.Repos.Specialty)
	profileService := NewProfileService(support.Repos.Profile)
	groupService := NewGroupService(support.Repos.Group)
	educationTypeService := NewEducationTypeService(support.Repos.EducationType)

	return &Services{
		ReportService:         reportService,
		HeadmanService:        headmanService,
		StudentService:        studentService,
		ScheduleService:       scheduleService,
		AttendanceService:     attendanceService,
		UserService:           userService,
		UniversityService:     universityService,
		FacultyService:        facultyService,
		DepartamentService:    departamentService,
		TeacherService:        teacherService,
		DisciplineService:     disciplineService,
		DisciplineTypeService: disciplineTypeService,
		ClassroomService:      classroomService,
		EducationLevelService: educationLevelService,
		SpecialtyService:      specialtyService,
		ProfileService:        profileService,
		GroupService:          groupService,
		EducationTypeService:  educationTypeService,
	}
}
