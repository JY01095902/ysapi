package salesout

import (
	"time"
)

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

func WithMerchantIdsIn(ids []string) ListRequestOption {
	return func(req *ListRequest) {
		if len(ids) == 0 {
			return
		}

		newVO := SimpleVO{
			Field:  "cust",
			Op:     "in",
			Value1: ids,
		}

		has := false
		for i, vo := range req.SimpleVOs {
			if vo.Field == "cust" {
				req.SimpleVOs[i] = newVO
				has = true
			}
		}

		if !has {
			req.SimpleVOs = append(req.SimpleVOs, newVO)
		}
	}
}

func WithSKUCodesIn(codes []string) ListRequestOption {
	return func(req *ListRequest) {
		if len(codes) == 0 {
			return
		}

		newVO := SimpleVO{
			Field:  "details.product.cCode",
			Op:     "in",
			Value1: codes,
		}

		has := false
		for i, vo := range req.SimpleVOs {
			if vo.Field == "details.product.cCode" {
				req.SimpleVOs[i] = newVO
				has = true
			}
		}

		if !has {
			req.SimpleVOs = append(req.SimpleVOs, newVO)
		}
	}
}
