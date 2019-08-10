package internal

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
