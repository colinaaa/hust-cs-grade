package internal

import (
	"testing"
)

func TestManager_readAll(t *testing.T) {

	t.Run("default", func(t *testing.T) {
		m, err := New(Config{
			XLSXPath: "../data/out.xlsx",
			JSONPath: "../data/cs18.json",
		})
		if err != nil {
			t.Fatal(err)
		}
		m.readAll()
	})
}
