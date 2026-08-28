package app

import "social-uploader/internal/domain"

func NextState(status string) string {
	switch status {
	case domain.StatusDraft:
		return "review"
	case domain.StatusApproved:
		return "submit"
	case domain.StatusSubmitted:
		return "archive"
	default:
		return "closed"
	}
}
func CanChange(status string) bool { return status != domain.StatusArchived }
func BatchReady(records []domain.Record) bool {
	if len(records) == 0 {
		return false
	}
	for _, r := range records {
		if r.Status == domain.StatusFailed {
			return false
		}
	}
	return true
}
