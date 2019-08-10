package main

import (
	"fmt"

	"github.com/colinaaa/hust-cs-grade/internal"
)

const output = "./data/out.xlsx"

func main() {
	m, err := internal.New(internal.Config{
		XLSXPath: output,
		JSONPath: "./data/cs18.json",
	})
	if err != nil {
		panic(err)
	}
	m.ByTerm("overall")
	for _, h := range m.Slice {
		fmt.Println(h)
	}
}
