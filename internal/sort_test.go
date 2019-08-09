package internal

import "testing"

func TestByTerm(t *testing.T) {
	t.Run("by overall", func(t *testing.T) {
		m, err := New(Config{
			XLSXPath: "../data/out.xlsx",
			JSONPath: "../data/cs18.json",
		})
		if err != nil {
			t.Fatal(err)
		}
		m.ByTerm("overall")
		t.Log(m.Slice)
	})
}
