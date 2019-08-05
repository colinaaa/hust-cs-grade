package internal

import (
	"encoding/json"
	"io/ioutil"
)

// Huster is the struct of a huster
type Huster struct {
	ID   string `json:"XH"`
	Name string `json:"XM"`

	Class  string  `json:"class,omitempty"`
	Credit float32 `json:"credit,omitempty"`

	Term `json:"Term,omitempty"`
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

// ScoreAndRank is the struct that contains the infomation of a term
type ScoreAndRank struct {
	Score float32 `json:"score,omitempty"`
	Rank  float32 `json:"rank,omitempty"`
}

// Term is the interface that contains all the terms
type Term interface {
	Fall2018() ScoreAndRank
	Spring2019() ScoreAndRank
}

// LoadJSON read .json file from paht arg and return the HusterMap or error
func LoadJSON(path string) (HusterMap, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var husters []Huster
	if err := json.Unmarshal(data, &husters); err != nil {
		return nil, err
	}
	m := make(HusterMap, len(husters))
	for _, h := range husters {
		m[h.ID] = h
	}

	return m, nil
}
