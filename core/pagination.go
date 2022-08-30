package core

import (
	"strconv"
)

type PageParams interface {
	GetQueryParams() map[string]string
}

type PageParamsImpl struct {
	StartingAfter string
	Limit         int
}

func (p *PageParamsImpl) GetQueryParams() map[string]string {
	qp := map[string]string{}
	if p.StartingAfter != "" {
		qp["starting_after"] = p.StartingAfter
	}
	if p.Limit != 0 {
		qp["limit"] = strconv.Itoa(p.Limit)
	}
	return qp
}

func WithPageParams(limit int, startingAfter string) RequestOption {
	return WithPage(&PageParamsImpl{
		Limit:         limit,
		StartingAfter: startingAfter,
	})
}

func WithPage(pageParams PageParams) RequestOption {
	var pp PageParams
	if pageParams == nil {
		pp = &PageParamsImpl{}
	} else {
		pp = pageParams
	}
	return withQueryParams(pp.GetQueryParams())
}

func NewPageParamsWith(limit int, startingAfter string) PageParams {
	return &PageParamsImpl{
		Limit:         limit,
		StartingAfter: startingAfter,
	}
}
