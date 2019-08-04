package internal

import (
	"encoding/json"
	"io/ioutil"
)

type Huster struct {
	ID   string `json:"XH"`
	Name string `json:"XM"`

	Class  string  `json:"class,omitempty"`
	Credit float32 `json:"credit,omitempty"`

	term `json:"term,omitempty"`
}

type HusterMap map[string]Huster

func (hm HusterMap) FindNameByID(id string) string {
	h, ok := hm[id]
	if !ok {
		return ""
	}

	return h.Name
}

type scoreAndRank struct {
	Score float32 `json:"score,omitempty"`
	Rank  float32 `json:"rank,omitempty"`
}

type term interface {
	Fall2018() scoreAndRank
	Spring2019() scoreAndRank
}

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
