package internal

import (
	"log"
	"sort"
)

// ByID is the sort meth that sort the slice inside manager by id
// it also return m for chain calling
func (m *Manager) ByID() *Manager {
	sort.Slice(m.Slice, func(i, j int) bool {
		return m.Slice[i].ID < m.Slice[j].ID
	})
	log.Println("sorted by id")
	return m
}

// ByTerm is the method sorting the slice by score of specific term
func (m *Manager) ByTerm(term string) *Manager {
	switch term {
	case "fall18":
		sort.Slice(m.Slice, func(i, j int) bool {
			return m.Slice[i].Fall18().Score < m.Slice[j].Fall18().Score
		})
	case "spr19":
		sort.Slice(m.Slice, func(i, j int) bool {
			return m.Slice[i].Spring2019().Score < m.Slice[j].Spring2019().Score
		})
	default:
		log.Println("sortByTerm: term not support")
		return m
	}
	log.Println("sorted by term", term)
	return m
}
