package controllers

import (
	"math"
)

type Pagination struct {
	Page         int
	PerPage      int
	Total        int
	LeftEdge     int
	LeftCurrent  int
	RightEdge    int
	RightCurrent int
}

func NewPagination(page, perPage, total int) *Pagination {
	return &Pagination{
		Page:         page,
		PerPage:      perPage,
		Total:        total,
		LeftEdge:     2,
		LeftCurrent:  2,
		RightCurrent: 5,
		RightEdge:    2,
	}
}

func (p *Pagination) Pages() int {
	totalPages := float64(p.Total) / float64(p.PerPage)
	return int(math.Ceil(totalPages))
}

func (p *Pagination) HasPrev() bool {
	return p.Page > 1
}

func (p *Pagination) HasNext() bool {
	return p.Page < p.Pages()
}

func (p *Pagination) IterationSet() []int {
	set := make([]int, 0)

	last := 0
	for i := 1; i <= p.Pages(); i++ {
		if (i <= p.LeftEdge) ||
			(i > p.Page-p.LeftCurrent-1 && i < p.Page+p.RightCurrent) ||
			(i > p.Pages()-p.RightEdge) {
			if last+1 != i {
				set = append(set, 0)
			}
			set = append(set, i)
			last = i
		}
	}

	return set
}
