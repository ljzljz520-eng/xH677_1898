package validate

import (
	"fmt"
	"social-uploader/internal/domain"
)

func EligibleForSubmit(r domain.Record) bool {
	return r.Status == domain.StatusApproved && r.Amount > 0
}
func StatusMessage(r domain.Record) string {
	switch r.Status {
	case domain.StatusDraft:
		return "待审核"
	case domain.StatusApproved:
		return "已审核"
	case domain.StatusSubmitted:
		return "已上送"
	case domain.StatusFailed:
		return "失败"
	default:
		return "已归档"
	}
}
func RequireBatch(batch string) error {
	if batch == "" {
		return fmt.Errorf("batch required")
	}
	return nil
}
