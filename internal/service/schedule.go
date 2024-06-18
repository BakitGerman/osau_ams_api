package service

import (
	"context"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
)

type ScheduleService struct {
	ScheduleRepo repository.ISchedule
}

func NewScheduleService(scheduleRepo repository.ISchedule) *ScheduleService {
	return &ScheduleService{ScheduleRepo: scheduleRepo}
}

func (s *ScheduleService) Create(ctx context.Context, schedule domain.Schedule) error {
	return s.ScheduleRepo.Create(ctx, schedule)
}

func (s *ScheduleService) Put(ctx context.Context, schedule domain.Schedule) error {
	return s.ScheduleRepo.Put(ctx, schedule)
}

func (s *ScheduleService) Patch(ctx context.Context, schedule domain.Schedule) error {
	updates := make(map[string]interface{})
	if schedule.GroupID != "" {
		updates["group_id"] = schedule.GroupID
	}
	if schedule.DisciplineID != 0 {
		updates["discipline_id"] = schedule.DisciplineID
	}
	if schedule.TeacherID != 0 {
		updates["teacher_id"] = schedule.TeacherID
	}
	if schedule.DisciplineTypeID != 0 {
		updates["discipline_type_id"] = schedule.DisciplineTypeID
	}
	if schedule.ClassroomID != 0 {
		updates["classroom_id"] = schedule.ClassroomID
	}
	if schedule.Semester != 0 {
		updates["semester"] = schedule.Semester
	}
	if !schedule.BeginStudies.IsZero() {
		updates["begin_studies"] = schedule.BeginStudies
	}
	if schedule.WeekType != "" {
		updates["week_type"] = schedule.WeekType
	}
	if schedule.DayOfWeek != "" {
		updates["day_of_week"] = schedule.DayOfWeek
	}
	if !schedule.StartTime.IsZero() {
		updates["start_time"] = schedule.StartTime
	}
	if schedule.IsActual != nil {
		updates["is_actual"] = schedule.IsActual
	}
	if len(updates) == 0 {
		return ErrNoUpdates
	}
	return s.ScheduleRepo.Patch(ctx, schedule.ScheduleID, updates)
}

func (s *ScheduleService) Delete(ctx context.Context, scheduleID int64) error {
	return s.ScheduleRepo.Delete(ctx, scheduleID)
}

func (s *ScheduleService) GetByID(ctx context.Context, scheduleID int64) (domain.ScheduleInfo, error) {
	return s.ScheduleRepo.GetByID(ctx, scheduleID)
}

func (s *ScheduleService) GetAll(ctx context.Context) ([]domain.ScheduleInfo, error) {
	return s.ScheduleRepo.GetAll(ctx)
}

func (s *ScheduleService) GetByGroupID(ctx context.Context, groupID string) ([]domain.ScheduleInfo, error) {
	return s.ScheduleRepo.GetByGroupID(ctx, groupID)
}

func (s *ScheduleService) GetByTeacherID(ctx context.Context, teacherID int64) ([]domain.ScheduleInfo, error) {
	return s.ScheduleRepo.GetByTeacherID(ctx, teacherID)
}

func (s *ScheduleService) GetByGroupAndWeekType(ctx context.Context, groupID, weekType string) ([]domain.ScheduleInfo, error) {
	return s.ScheduleRepo.GetByGroupAndWeekType(ctx, groupID, weekType)
}

func (s *ScheduleService) GetByTeacherAndWeekType(ctx context.Context, teacherID int64, weekType string) ([]domain.ScheduleInfo, error) {
	return s.ScheduleRepo.GetByTeacherAndWeekType(ctx, teacherID, weekType)
}

func (s *ScheduleService) GetByGroupWeekTypeAndDay(ctx context.Context, groupID, weekType, dayOfWeek string) ([]domain.ScheduleInfo, error) {
	return s.ScheduleRepo.GetByGroupWeekTypeAndDay(ctx, groupID, weekType, dayOfWeek)
}

func (s *ScheduleService) GetByTeacherWeekTypeAndDay(ctx context.Context, teacherID int64, weekType, dayOfWeek string) ([]domain.ScheduleInfo, error) {
	return s.ScheduleRepo.GetByTeacherWeekTypeAndDay(ctx, teacherID, weekType, dayOfWeek)
}

func (s *ScheduleService) GetActualByGroupID(ctx context.Context, groupID string) ([]domain.ScheduleInfo, error) {
	return s.ScheduleRepo.GetActualByGroupID(ctx, groupID)
}

func (s *ScheduleService) GetActualByTeacherID(ctx context.Context, teacherID int64) ([]domain.ScheduleInfo, error) {
	return s.ScheduleRepo.GetActualByTeacherID(ctx, teacherID)
}

func (s *ScheduleService) GetActualByGroupAndWeekType(ctx context.Context, groupID, weekType string) ([]domain.ScheduleInfo, error) {
	return s.ScheduleRepo.GetActualByGroupAndWeekType(ctx, groupID, weekType)
}

func (s *ScheduleService) GetActualByGroupWeekTypeAndDay(ctx context.Context, groupID, weekType, dayOfWeek string) ([]domain.ScheduleInfo, error) {
	return s.ScheduleRepo.GetActualByGroupWeekTypeAndDay(ctx, groupID, weekType, dayOfWeek)
}

func (s *ScheduleService) GetActualByTeacherAndWeekType(ctx context.Context, teacherID int64, weekType string) ([]domain.ScheduleInfo, error) {
	return s.ScheduleRepo.GetActualByTeacherAndWeekType(ctx, teacherID, weekType)
}

func (s *ScheduleService) GetActualByTeacherWeekTypeAndDay(ctx context.Context, teacherID int64, weekType, dayOfWeek string) ([]domain.ScheduleInfo, error) {
	return s.ScheduleRepo.GetActualByTeacherWeekTypeAndDay(ctx, teacherID, weekType, dayOfWeek)
}
