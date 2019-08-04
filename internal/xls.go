package internal

import (
	"github.com/360EntSecGroup-Skylar/excelize"
)

const defaultXLSFile = "./data/rank2018-2019.xlsx"

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

type csRank struct {
	*excelize.File
	Path string
}
