package main

import (
	"github.com/colinaaa/hust-cs-grade/internal"
)

const output = "./data/out.xlsx"

func main() {
	m := internal.New(internal.DefaultConfig)
	m.JoinName(output)
}
