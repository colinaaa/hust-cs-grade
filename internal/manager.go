package internal

import (
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// Manager is the root struct that contains all the thing needed
type Manager struct {
	HusterMap
	Slice []Huster

	*excelize.File
}

// Config contains the file path (json & xlsx)
type Config struct {
	XLSXPath string
	JSONPath string
}

const (
	defaultXLSXPath = "./data/rank2018-2019.xlsx"
	defaultJSONPath = "./data/cs18.json"
)

// DefaultConfig is the Config using default path
var DefaultConfig = Config{
	XLSXPath: defaultXLSXPath,
	JSONPath: defaultJSONPath,
}

// New gets a new manager by config arg
func New(cfg Config) (*Manager, error) {
	xlsx, err := excelize.OpenFile(cfg.XLSXPath)
	if err != nil {
		return nil, err
	}

	hm, err := LoadJSON(cfg.JSONPath)
	if err != nil {
		return nil, err
	}

	m := Manager{File: xlsx, HusterMap: hm}
	m.readAll()
	return &m, nil
}

const idIndex = 1

// JoinName is the method that join the name and score and save into output
func (m *Manager) JoinName(output string) {
	rows := m.GetRows("sheet0")
	for index, cols := range rows[4:] {
		name := m.FindNameByID(strings.TrimSpace(cols[idIndex]))
		m.SetCellStr("sheet0", "C"+fmt.Sprint(index+5), name)
	}
	if err := m.SaveAs(output); err != nil {
		panic(err.Error())
	}
}

func (m *Manager) readAll() {
	rows := m.GetRows("sheet0")
	for _, cols := range rows[4:] {
		h := Huster{}
		fmt.Sscanf(
			strings.Join(cols, " "),
			"%d %s %s %f %f %f %f %s",
			&h.rankOverall,
			&h.ID,
			&h.Name,
			&h.scoreFall18,
			&h.scoreSpr19,
			&h.scoreOverall,
			&h.Credit,
			&h.Class,
		)
		fmt.Println(h)
		m.Slice = append(m.Slice, h)
	}
}
