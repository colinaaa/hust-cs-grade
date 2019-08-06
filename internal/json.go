package internal

import (
	"encoding/json"
	"io/ioutil"
)

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
