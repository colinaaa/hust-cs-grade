package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Huster is the struct of a huster
type Huster struct {
	ID   string `json:"XH"`
	Name string `json:"XM"`

	Class  string  `json:"class,omitempty"`
	Credit float32 `json:"credit,omitempty"`

	Terms `json:"terms,omitempty"`

	rankFall2018 int
	rankSpr2019  int

	scoreFall2018 float32
	scoreSpr2019  float32
}

func (h Huster) String() string {
	s := fmt.Sprint(h.ID, h.Name, h.Class, h.Credit, h.All())
	return s
}

// Fall2018 implements the Terms interface
func (h Huster) Fall2018() Term {
	return Term{
		Name:  "fall-2018",
		Score: h.scoreFall2018,
		Rank:  h.rankFall2018,
	}
}

// Spring2019 implements the Terms interface
func (h Huster) Spring2019() Term {
	return Term{
		Name:  "Spring2019",
		Score: h.scoreSpr2019,
		Rank:  h.rankSpr2019,
	}
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
