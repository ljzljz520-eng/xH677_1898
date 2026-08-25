package report

import (
	"fmt"
	"social-uploader/internal/domain"
)

func AttachmentFor(r Report) domain.Attachment {
	return domain.Attachment{ID: fmt.Sprintf("attachment-%s", r.BatchID), BatchID: r.BatchID, Name: "submission-report.txt", Content: r.Text}
}
func IsEmpty(r Report) bool { return len(r.Submitted) == 0 && len(r.Errors) == 0 }
func Summary(r Report) string {
	return fmt.Sprintf("batch=%s submitted=%d errors=%d", r.BatchID, len(r.Submitted), len(r.Errors))
}
