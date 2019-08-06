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

	Terms `json:"terms,omitempty"`

	rankFall18 int
	rankSpr19  int

	scoreFall18 float32
	scoreSpr19  float32
}

func (h Huster) String() string {
	s := fmt.Sprint(h.ID, h.Name, h.Class, h.Credit, h.All())
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

// All returns all the terms
func (h Huster) All() []Term {
	return []Term{h.Fall18(), h.Spring19()}
}

// HusterMap is the map with id as key and Huster struct as value
type HusterMap map[string]Huster

// FindNameByID is the method to find names
// returns "" if not found
func (hm HusterMap) FindNameByID(id string) string {
	h, ok := hm[id]
	if !ok {
		return ""
	}

	return h.Name
}
