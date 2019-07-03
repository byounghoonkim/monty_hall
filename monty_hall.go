package main

import (
	"fmt"
	"math/rand"
	"time"
)

func montyHall(changeMind bool) bool {
	rand.Seed(time.Now().UnixNano())

	doors := [3]byte{'G', 'G', 'G'} // Goat
	doors[rand.Intn(3)] = 'C'       // Car

	firstSelection := rand.Intn(3)

	openDoor := -1
	if doors[firstSelection] == 'C' {
		r := rand.Intn(2)

		for i := 0; i < 3; i++ {
			if firstSelection == i {
				continue
			} else {
				if r == 0 {
					openDoor = i
					break
				}
				r--
			}
		}

	} else {
		for i := 0; i < 3; i++ {
			if i != firstSelection && doors[i] == 'G' {
				openDoor = i
				break
			}
		}
	}

	if true == changeMind {
		secondSelection := -1
		for i := 0; i < 3; i++ {
			if i != firstSelection && i != openDoor {
				secondSelection = i
				break
			}

		}
		return doors[secondSelection] == 'C'

	} else {
		return doors[firstSelection] == 'C'
	}
}

func simulate(n int, changeMind bool) {
	count := 0

	for i := 0; i < n; i++ {
		if true == montyHall(changeMind) {
			count++
		}
	}

	fmt.Printf("Change mind %5t => %d / %d (%0.3f %%)\n", changeMind, count, n, float64(count*100)/float64(n))
}

func main() {
	fmt.Println("Monty Hall")

	simulate(10000, false)
	simulate(10000, true)
}
