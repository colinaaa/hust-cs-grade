package internal

import (
	"github.com/360EntSecGroup-Skylar/excelize"
)

const defaultXLSFile = "./data/rank2018-2019.xlsx"

// LoadXLS read xlsx file from path arg and return the *excelize.File or error
func LoadXLS(path string) (*excelize.File, error) {
	if path == "" {
		path = defaultXLSFile
	}

	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// Manager is the root struct that contains all the thing needed
type Manager struct {
	*excelize.File
	Path string
}
