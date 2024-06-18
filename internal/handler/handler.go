package handler

import (
	"log/slog"
	"time"

	"github.com/BeRebornBng/OsauAmsApi/internal/service"
	"github.com/BeRebornBng/OsauAmsApi/pkg/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/ru"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	TokenManager auth.TokenManager
	services     *service.Services
	logger       *slog.Logger
	validate     *validator.Validate
	translator   ut.Translator
}

func NewHandler(TokenManager auth.TokenManager, services *service.Services, logger *slog.Logger) *Handler {
	validate := validator.New()

	uni := ut.New(ru.New())
	trans, _ := uni.GetTranslator("ru")

	translations.RegisterDefaultTranslations(validate, trans)
	validate.RegisterTranslation("custompasswordregex", trans, func(ut ut.Translator) error {
		return ut.Add("custompasswordregex", "{0} должен содержать латинские буквы, цифры и символы $!%*#?&@0-9", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("custompasswordregex", fe.Field())
		return t
	})

	translations.RegisterDefaultTranslations(validate, trans)
	validate.RegisterTranslation("alphanum", trans, func(ut ut.Translator) error {
		return ut.Add("alphanum", "{0} должен содержать латиницские буквы и цифры", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("alphanum", fe.Field())
		return t
	})

	translations.RegisterDefaultTranslations(validate, trans)
	validate.RegisterTranslation("customfieldrusregex", trans, func(ut ut.Translator) error {
		return ut.Add("customfieldrusregex", "{0} должен содержать только кирилицу", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("customfieldrusregex", fe.Field())
		return t
	})

	validate.RegisterValidation("customspecialtycoderegex", CustomSpecialtyCodeRegex)
	validate.RegisterValidation("customgroupidregex", CustomGroupIDRegex)
	validate.RegisterValidation("custompasswordregex", CustomPasswordRegex)
	validate.RegisterValidation("customfieldrusregex", CustomFieldRusRegex)
	validate.RegisterValidation("time", ValidateTime)
	validate.RegisterValidation("roledependentfields", roleDependentFields)

	return &Handler{
		TokenManager: TokenManager,
		services:     services,
		logger:       logger,
		validate:     validate,
		translator:   trans,
	}
}

func Logger(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path

		// Выполняем запрос
		c.Next()

		// Логируем детали запроса и статус ответа
		status := c.Writer.Status()
		protocol := "HTTP"
		if c.Request.TLS != nil {
			protocol = "HTTPS"
		}
		log.Info("Request", slog.String("method", method), slog.String("path", path), slog.String("protocol", protocol), slog.Int("status", status))
		// if c.Request.Method == "OPTIONS" || c.Request.Method == "GET" {
		// 	log.Info("Request", slog.String("method", method), slog.String("path", path), slog.String("protocol", protocol), slog.Int("status", status))
		// }
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	config := cors.Config{
		AllowOrigins:     []string{"", "", ""},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"},
		ExposeHeaders:    []string{"Content-Length", "Connection"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))
	router.Use(Logger(h.logger))

	api := router.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/signin", h.SignInUser)
		auth.POST("/signup", h.CreateUser)
	}

	authorized := api.Group("/")
	authorized.Use(h.userIdentity)
	{
		admin := authorized.Group("/admins")
		admin.Use(RoleMiddleware("Админ"))
		{
			admin.POST("/universities", h.CreateUniversity)
			admin.PUT("/universities", h.PutUniversity)
			admin.PATCH("/universities", h.PatchUniversity)
			admin.DELETE("/universities/id/:id", h.DeleteUniversity)
			admin.GET("/universities/id/:id", h.GetUniversityByID)
			admin.GET("/universities/name/:name", h.GetUniversityByName)
			admin.GET("/universities", h.GetAllUniversities)

			admin.POST("/faculties", h.CreateFaculty)
			admin.PUT("/faculties", h.PutFaculty)
			admin.PATCH("/faculties", h.PatchFaculty)
			admin.DELETE("/faculties/id/:id", h.DeleteFaculty)
			admin.GET("/faculties/id/:id", h.GetFacultyByID)
			admin.GET("/faculties/name/:name", h.GetFacultyByName)
			admin.GET("/faculties", h.GetAllFaculties)
			admin.GET("/faculties/university/id/:id", h.GetFacultiesByUniversityID)

			admin.POST("/users", h.CreateUser)
			admin.PUT("/users", h.PutUser)
			admin.PATCH("/users", h.PatchUser)
			admin.DELETE("/users/:id", h.DeleteUser)
			admin.GET("/users/:id", h.GetUserByID)
			admin.GET("/users/name/:name", h.GetUserByName)
			admin.GET("/users/student/:id", h.GetUserByStudentID)
			admin.GET("/users/headman/:id", h.GetUserByHeadmanID)
			admin.GET("/users/teacher/:id", h.GetUserByTeacherID)
			admin.GET("/users/role/:role", h.GetAllUsersByRole)
			admin.GET("/users", h.GetAll)

			admin.POST("/departaments", h.CreateDepartament)
			admin.PUT("/departaments", h.PutDepartament)
			admin.PATCH("/departaments", h.PatchDepartament)
			admin.DELETE("/departaments/:id", h.DeleteDepartament)
			admin.GET("/departaments/:id", h.GetDepartamentByID)
			admin.GET("/departaments/name/:name", h.GetDepartamentByName)
			admin.GET("/departaments", h.GetAllDepartaments)

			admin.POST("/teachers", h.CreateTeacher)
			admin.PUT("/teachers", h.PutTeacher)
			admin.PATCH("/teachers", h.PatchTeacher)
			admin.DELETE("/teachers/:id", h.DeleteTeacher)
			admin.GET("/teachers/:id", h.GetTeacherByID)
			admin.GET("/teachers/email/:email", h.GetTeacherByEmail)
			admin.GET("/teachers", h.GetAllTeachers)

			admin.POST("/disciplines", h.CreateDiscipline)
			admin.PUT("/disciplines", h.PutDiscipline)
			admin.PATCH("/disciplines", h.PatchDiscipline)
			admin.DELETE("/disciplines/:id", h.DeleteDiscipline)
			admin.GET("/disciplines/:id", h.GetDisciplineByID)
			admin.GET("/disciplines/name/:name", h.GetDisciplineByName)
			admin.GET("/disciplines", h.GetAllDisciplines)

			admin.POST("/discipline_types", h.CreateDisciplineType)
			admin.PUT("/discipline_types", h.PutDisciplineType)
			admin.PATCH("/discipline_types", h.PatchDisciplineType)
			admin.DELETE("/discipline_types/:id", h.DeleteDisciplineType)
			admin.GET("/discipline_types/:id", h.GetDisciplineTypeByID)
			admin.GET("/discipline_types", h.GetAllDisciplineTypes)

			admin.POST("/classrooms", h.CreateClassroom)
			admin.PUT("/classrooms", h.PutClassroom)
			admin.PATCH("/classrooms", h.PatchClassroom)
			admin.DELETE("/classrooms/:id", h.DeleteClassroom)
			admin.GET("/classrooms/:id", h.GetClassroomByID)
			admin.GET("/classrooms", h.GetAllClassrooms)

			admin.POST("/education_levels", h.CreateEducationLevel)
			admin.PUT("/education_levels", h.PutEducationLevel)
			admin.PATCH("/education_levels", h.PatchEducationLevel)
			admin.DELETE("/education_levels/:id", h.DeleteEducationLevel)
			admin.GET("/education_levels/:id", h.GetEducationLevelByID)
			admin.GET("/education_levels", h.GetAllEducationLevels)

			admin.POST("/specialties", h.CreateSpecialty)
			admin.PUT("/specialties", h.PutSpecialty)
			admin.PATCH("/specialties", h.PatchSpecialty)
			admin.DELETE("/specialties/:code", h.DeleteSpecialty)
			admin.GET("/specialties/code/:code", h.GetSpecialtyByCode)
			admin.GET("/specialties/name/:name", h.GetSpecialtyByName)
			admin.GET("/specialties", h.GetAllSpecialties)

			admin.POST("/profiles", h.CreateProfile)
			admin.PUT("/profiles", h.PutProfile)
			admin.PATCH("/profiles", h.PatchProfile)
			admin.DELETE("/profiles/:id", h.DeleteProfile)
			admin.GET("/profiles/:id", h.GetProfileByID)
			admin.GET("/profiles/name/:name", h.GetProfileByName)
			admin.GET("/profiles", h.GetAllProfiles)

			admin.POST("/groups", h.CreateGroup)
			admin.PUT("/groups", h.PutGroup)
			admin.PATCH("/groups", h.PatchGroup)
			admin.DELETE("/groups/:id", h.DeleteGroup)
			admin.GET("/groups/:id", h.GetGroupByID)
			admin.GET("/groups/name/:name", h.GetGroupByName)
			admin.GET("/groups", h.GetAllGroups)

			admin.POST("/education_types", h.CreateEducationType)
			admin.PUT("/education_types", h.PutEducationType)
			admin.PATCH("/education_types", h.PatchEducationType)
			admin.DELETE("/education_types/:id", h.DeleteEducationType)
			admin.GET("/education_types/:id", h.GetEducationTypeByID)
			admin.GET("/education_types", h.GetAllEducationTypes)
		}

		headman := authorized.Group("/headmans")
		headman.Use(RoleMiddleware("Староста"))
		{
			headman.POST("/attendances", h.CreateAttendances)
			headman.PUT("/attendances", h.PutAttendances)
			headman.GET("/schedules/week/:week", h.GetActualSchedulesByGroupAndWeekType)
			headman.GET("/schedules/week/:week/day/:day", h.GetActualSchedulesByGroupWeekTypeAndDay)
			headman.GET("/attendances/schedule/:id/date/:date", h.GetHeadmanAllAttendances)
			headman.GET("/reports/start/:start_date/end/:end_date", h.GetActualReportByGroupIDAndCreated)
		}

		student := authorized.Group("/students")
		student.Use(RoleMiddleware("Студент"))
		{
			student.GET("/schedules/week/:week", h.GetActualSchedulesByGroupAndWeekType)
			student.GET("/schedules/week/:week/day/:day", h.GetActualSchedulesByGroupWeekTypeAndDay)
		}

		teacher := authorized.Group("/teachers")
		teacher.Use(RoleMiddleware("Преподаватель"))
		{
			teacher.POST("/attendances", h.CreateAttendances)
			teacher.PUT("/attendances", h.PutAttendances)
			teacher.GET("/schedules/week/:week", h.GetActualSchedulesByTeacherIDWeekType)
			teacher.GET("/schedules/week/:week/day/:day", h.GetActualSchedulesByTeacherIDWeekTypeAndDay)
			teacher.GET("/attendances/group/:group_id/schedule/:id/date/:date", h.GetTeacherAllAttendances)
		}

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
