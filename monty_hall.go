package main

import (
	"fmt"
	"math/rand"
	"time"
)

func monty_hall(change_mind bool) bool {
	rand.Seed(time.Now().UnixNano())

	doors := [3]byte{'G', 'G', 'G'} // Goat
	doors[rand.Intn(3)] = 'C'       // Car

	first_selection := rand.Intn(3)

	open_door := -1
	if doors[first_selection] == 'C' {
		r := rand.Intn(2)

		for i := 0; i < 3; i++ {
			if first_selection == i {
				continue
			} else {
				if r == 0 {
					open_door = i
					break
				}
				r -= 1
			}
		}

	} else {
		for i := 0; i < 3; i++ {
			if i != first_selection && doors[i] == 'G' {
				open_door = i
				break
			}
		}
	}

	if true == change_mind {
		second_selection := -1
		for i := 0; i < 3; i++ {
			if i != first_selection && i != open_door {
				second_selection = i
				break
			}

		}
		return doors[second_selection] == 'C'

	} else {
		return doors[first_selection] == 'C'
	}
}

func sumulate(n int, change_mind bool) {
	count := 0

	for i := 0; i < n; i++ {
		if true == monty_hall(change_mind) {
			count += 1
		}
	}

	fmt.Printf("Change mind %5t => %d / %d (%0.3f %%)\n", change_mind, count, n, float64(count*100)/float64(n))
}

func main() {
	fmt.Println("Monty Hall")

	sumulate(10000, false)
	sumulate(10000, true)
}
