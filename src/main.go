package main

import (
	"fmt"
	"github.com/agedito/ag_lib/api/dice"
	"github.com/agedito/ag_lib/id"
	"github.com/agedito/ag_lib/internal/log"
)

func main() {
	log.Hi()
	id.Id()
	d := dice.New()
	fmt.Println(d)
}
