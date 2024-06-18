package service

import (
	"context"
	"time"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
)

type AttendanceService struct {
	AttendanceRepo repository.IAttendance
}

func NewAttendanceService(attendanceRepo repository.IAttendance) *AttendanceService {
	return &AttendanceService{AttendanceRepo: attendanceRepo}
}

func (s *AttendanceService) Create(ctx context.Context, attendance domain.Attendance) error {
	return s.AttendanceRepo.Create(ctx, attendance)
}

func (s *AttendanceService) Put(ctx context.Context, attendance domain.Attendance) error {
	return s.AttendanceRepo.Put(ctx, attendance)
}

func (s *AttendanceService) Patch(ctx context.Context, attendance domain.Attendance) error {
	updates := make(map[string]interface{})
	if attendance.StudentID != 0 {
		updates["student_id"] = attendance.StudentID
	}
	if attendance.ScheduleID != 0 {
		updates["schedule_id"] = attendance.ScheduleID
	}
	if attendance.Presence != nil {
		updates["presence"] = attendance.Presence
	}
	if attendance.LateArrival != nil {
		updates["late_arrival"] = attendance.LateArrival
	}
	if attendance.Respectfulness != nil {
		updates["respectfulness"] = attendance.Respectfulness
	}
	if attendance.Reason != nil {
		updates["reason"] = attendance.Reason
	}
	if !attendance.Created.IsZero() {
		updates["created"] = attendance.Created
	}
	if len(updates) == 0 {
		return ErrNoUpdates
	}
	return s.AttendanceRepo.Patch(ctx, attendance.AttendanceID, updates)
}

func (s *AttendanceService) Delete(ctx context.Context, attendanceID int64) error {
	return s.AttendanceRepo.Delete(ctx, attendanceID)
}

func (s *AttendanceService) GetByID(ctx context.Context, attendanceID int64) (domain.AttendanceInfo, error) {
	return s.AttendanceRepo.GetByID(ctx, attendanceID)
}

func (s *AttendanceService) GetByStudentID(ctx context.Context, studentID int64) ([]domain.AttendanceInfo, error) {
	return s.AttendanceRepo.GetByStudentID(ctx, studentID)
}

func (s *AttendanceService) GetAll(ctx context.Context) ([]domain.AttendanceInfo, error) {
	return s.AttendanceRepo.GetAll(ctx)
}

func (s *AttendanceService) GetAllByGroupIDAndCreated(ctx context.Context, groupID string, scheduleID int64, created time.Time) ([]domain.GroupAttendanceInfo, error) {
	return s.AttendanceRepo.GetAllByGroupIDAndCreated(ctx, groupID, scheduleID, created)
}
