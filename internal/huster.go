package internal

import (
	"fmt"
)

// Huster is the struct of a huster
type Huster struct {
	ID   string `json:"XH"`
	Name string `json:"XM"`

	Class  string  `json:"class,omitempty"`
	Credit float32 `json:"credit,omitempty"`

	rankFall18  int
	rankSpr19   int
	rankOverall int

	scoreFall18  float32
	scoreSpr19   float32
	scoreOverall float32
}

func (h Huster) String() string {
	sep := " "
	s := fmt.Sprint(h.ID, sep, h.Name, sep, h.Class, sep, h.Credit, sep, h.All())
	return s
}

// Fall18 implements the Terms interface
func (h Huster) Fall18() Term {
	return Term{
		Name:  "fall18",
		Score: h.scoreFall18,
		Rank:  h.rankFall18,
	}
}

// Spring19 implements the Terms interface
func (h Huster) Spring19() Term {
	return Term{
		Name:  "Spring19",
		Score: h.scoreSpr19,
		Rank:  h.rankSpr19,
	}
}

// Overall impplements the Terms interface
func (h Huster) Overall() Term {
	return Term{
		Name:  "Overall",
		Score: h.scoreOverall,
		Rank:  h.rankOverall,
	}
}

// All returns all the terms
func (h Huster) All() []Term {
	return []Term{h.Fall18(), h.Spring19(), h.Overall()}
}

// ToStringSlice turns Huster into []string
func (h Huster) ToStringSlice() *[]string {
	s := []string{
		h.ID,
		h.Name,
		h.Class,
		fmt.Sprint(h.Credit),
	}
	for _, t := range h.All() {
		s = append(s, fmt.Sprint(t.Score), fmt.Sprint(t.Rank))
	}
	return &s
}
