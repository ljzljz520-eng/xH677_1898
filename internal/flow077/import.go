package flow077

import (
	"fmt"
	"social-uploader/internal/app"
	"social-uploader/internal/domain"
	"social-uploader/internal/report"
	"social-uploader/internal/validate"
)

type Processor struct{ Service *app.Service }

func New(s *app.Service) *Processor { return &Processor{Service: s} }
func (p *Processor) ImportBatch(batch string, records []domain.Record) (report.Report, error) {
	results := validate.Batch(records)
	good, bad := validate.Partition(results)
	w := domain.NewWorkflow("workflow-"+batch, batch)
	for _, r := range good {
		if e := p.Service.Register(r); e != nil {
			return report.Report{}, e
		}
		w.SuccessCount++
	}
	for _, failed := range bad {
		r := failed.Record
		// 校验失败的记录保留错误明细，但不得标为已上送；
		// 仅成功项进入上送流程，失败项标记为 failed 以正确显示。
		r.Status = domain.StatusFailed
		r.Remark = validate.ErrorText(failed)
		if e := p.Service.DB.SaveRecord(r); e != nil {
			return report.Report{}, e
		}
	}
	w.ErrorCount = len(bad)
	if len(bad) > 0 {
		w.State = "partial"
	} else {
		w.State = "submitted"
	}
	if e := p.Service.DB.SaveWorkflow(w); e != nil {
		return report.Report{}, e
	}
	out := report.Build(batch, good, bad)
	if e := p.Service.DB.SaveAttachment(report.AttachmentFor(out)); e != nil {
		return report.Report{}, e
	}
	return out, nil
}
func (p *Processor) ImportOne(batch, id, employee string, amount int64) (report.Report, error) {
	r, e := domain.NewRecord(id, batch, employee, amount)
	if e != nil {
		return report.Report{}, e
	}
	return p.ImportBatch(batch, []domain.Record{r})
}
func (p *Processor) SubmitSuccessful(batch string, ids []string) error {
	for _, id := range ids {
		r, e := p.Service.DB.GetRecord(id)
		if e != nil {
			return e
		}
		if e = r.Submit(); e != nil {
			return fmt.Errorf("%s: %w", id, e)
		}
		if e = p.Service.DB.UpdateRecord(r); e != nil {
			return e
		}
	}
	return nil
}
