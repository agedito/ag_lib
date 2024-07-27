package dice

import (
	"fmt"
	"math/rand"
)

type Dice struct {
	number int8
}

func New() *Dice {
	d := &Dice{number: 0}
	d.Launch()
	return d
}

func (d *Dice) Launch() int8 {
	number := int8(1 + rand.Intn(6))
	d.number = number
	return number
}

func (d *Dice) Value() int8 {
	return d.number
}

func (d *Dice) String() string {
	return fmt.Sprintf("Dice (%d)", d.number)
}
