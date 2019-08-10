package internal

import "fmt"

// Term is the struct that contains the infomation of a term
type Term struct {
	Name  string  `json:"name,omitempty"`
	Score float32 `json:"score,omitempty"`
	Rank  int     `json:"rank,omitempty"`
}

func (t Term) String() string {
	return fmt.Sprint("score:", t.Score, "rank:", t.Rank)
}

// Terms is the interface that contains all the terms
type Terms interface {
	All() []Term
	Overall() Term
	Fall18() Term
	Spring19() Term
}
