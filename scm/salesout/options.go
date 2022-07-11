package salesout

import "time"

type ListRequestOption func(req *ListRequest)

func WithVouchDateBetween(start, end time.Time) ListRequestOption {
	return func(req *ListRequest) {
		newVO := SimpleVO{
			Field:  "vouchdate",
			Op:     "between",
			Value1: start.Local().Format("2006-01-02 15:04:05"),
			Value2: end.Local().Format("2006-01-02 15:04:05"),
		}

		has := false
		for i, vo := range req.SimpleVOs {
			if vo.Field == "vouchdate" {
				req.SimpleVOs[i] = newVO
				has = true
			}
		}

		if !has {
			req.SimpleVOs = append(req.SimpleVOs, newVO)
		}
	}
}

func WithPubTsBetween(start, end time.Time) ListRequestOption {
	return func(req *ListRequest) {
		newVO := SimpleVO{
			Field:  "pubts",
			Op:     "between",
			Value1: start.Local().Format("2006-01-02 15:04:05"),
			Value2: end.Local().Format("2006-01-02 15:04:05"),
		}

		has := false
		for i, vo := range req.SimpleVOs {
			if vo.Field == "pubts" {
				req.SimpleVOs[i] = newVO
				has = true
			}
		}

		if !has {
			req.SimpleVOs = append(req.SimpleVOs, newVO)
		}
	}
}
