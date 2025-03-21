package puppy

import (
	"fmt"

	"github.com/aelhadyy/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func Fromv12() {
	fmt.Println("I'm from v1.2.0")
}

func Fromv13() {
	fmt.Println("I'm from v1.3.0")
}
