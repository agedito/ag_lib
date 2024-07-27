package main

import (
	"fmt"
	"github.com/agedito/ag_lib/api/dice"
	"github.com/agedito/ag_lib/internal/log"
)

func main() {
	fmt.Println("Hi")
	log.Hi()
	d := dice.New()
	fmt.Println("D", d)
	//d := dice.New()
	//d.Launch()
}
