package service

import (
	"context"
	"time"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
)

type ReportService struct {
	ReportRepo repository.IReport
}

func NewReportService(reportRepo repository.IReport) *ReportService {
	return &ReportService{ReportRepo: reportRepo}
}

func (s *ReportService) GetActualReportByGroupIDCreated(ctx context.Context, groupID string, startRange time.Time, endRange time.Time) (*domain.AttendanceReport, error) {
	return s.ReportRepo.GetActualReportByGroupIDCreated(ctx, groupID, startRange, endRange)
}
