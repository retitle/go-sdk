package core

import (
	"strconv"
)

type PageParams struct {
	StartingAfter string
	Limit         int
}

func (p PageParams) GetQueryParams() map[string]string {
	qp := map[string]string{}
	if p.StartingAfter != "" {
		qp["starting_after"] = p.StartingAfter
	}
	if p.Limit != 0 {
		qp["limit"] = strconv.Itoa(p.Limit)
	}
	return qp
}

func WithPage(pageParams *PageParams) RequestOption {
	var pp PageParams
	if pageParams == nil {
		pp = PageParams{}
	} else {
		pp = *pageParams
	}
	return withQueryParams(pp.GetQueryParams())
}

func WithPageParams(limit int, startingAfter string) RequestOption {
	return WithPage(&PageParams{
		Limit:         limit,
		StartingAfter: startingAfter,
	})
}

func WithPageLimit(limit int) RequestOption {
	return WithPageParams(limit, "")
}

func WithPageStartingAfter(startingAfter string) RequestOption {
	return WithPageParams(0, startingAfter)
}
