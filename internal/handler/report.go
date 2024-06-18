package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HeadmanReport struct {
	GroupID    string `json:"group_id" validate:"required,min=15,max=20,customgroupidregex"`
	StartRange string `json:"start_range" validate:"required,datetime=2006-01-02"`
	EndRange   string `json:"end_range" validate:"required,datetime=2006-01-02"`
}

func (h *Handler) GetActualReportByGroupIDAndCreated(c *gin.Context) {

	start_range, err := time.Parse("2006-01-02", c.Param("start_date"))
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}
	end_range, err := time.Parse("2006-01-02", c.Param("end_date"))
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	data, ok := c.Get(groupCtx)
	if !ok {
		respondWithError(h.logger, c, http.StatusUnauthorized, "Group ID not found in context")
		return
	}

	groupID, ok := data.(string)
	if !ok {
		respondWithError(h.logger, c, http.StatusInternalServerError, "Failed to convert group ID")
		return
	}

	report := HeadmanReport{GroupID: groupID, StartRange: c.Param("start_date"), EndRange: c.Param("end_date")}

	if err := h.validate.Struct(report); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if attendanceData, err := h.services.ReportService.GetActualReportByGroupIDCreated(c.Request.Context(), report.GroupID, start_range, end_range); err == nil {
		c.JSON(http.StatusOK, attendanceData)
	} else {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}
}
