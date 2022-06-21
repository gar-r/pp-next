package model

import (
	"fmt"
	"sort"
)

// Average returns a string representation of the average value of current votes.
// Special vote options are skipped and not counted.
func (r *Room) Average() string {
	sum := 0.0
	cnt := 0
	for _, v := range r.Votes {
		if v.IsCoffee() || v.IsLarge() || v.IsNothing() || v.IsQuestion() {
			continue
		}
		sum += v.Vote
		cnt += 1
	}
	var avg float64
	if cnt == 0 {
		avg = 0
	} else {
		avg = sum / float64(cnt)
	}
	return fmt.Sprintf("%.2f", avg)
}

// Summary counts votes by occurrence and returns a slice of SummaryItems
// representing the groups. The slice is sorted by category in ascending order.
func (r *Room) Summary() []*SummaryItem {
	result := make([]*SummaryItem, 0)
	m := r.summaryMap()
	cat := make([]float64, 0, len(m))
	for k := range m {
		cat = append(cat, k)
	}
	sort.Float64s(cat)
	for _, k := range cat {
		result = append(result, &SummaryItem{
			Category: k,
			Count:    m[k],
		})
	}
	return result
}

func (r *Room) summaryMap() map[float64]int {
	m := make(map[float64]int)
	for _, v := range r.Votes {
		vote := v.Vote
		i, ok := m[vote]
		if !ok {
			m[vote] = 1
		} else {
			m[vote] = i + 1
		}
	}
	return m
}

// SummaryItem represents a single row of the summary table
type SummaryItem struct {
	Category float64
	Count    int
}
