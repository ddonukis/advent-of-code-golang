package main

import "fmt"

const speedBoost int = 1

type Race struct {
	time     int
	distance int
}

func main() {
	races := [...]Race{
		// example
		// {7, 9},
		// {15, 40},
		// {30, 200},
		// part 1
		// {59, 597},
		// {79, 1234},
		// {65, 1032},
		// {75, 1328},
		// part 2
		{59_796_575, 597_123_410_321_328},
	}
	answer := 1
	for _, race := range races {
		winCount := 0
		for holdDuration := 1; holdDuration < race.time; holdDuration++ {
			leftoverTime := race.time - holdDuration
			distance := speedBoost * holdDuration * leftoverTime
			if distance > race.distance {
				winCount++
			}
		}
		fmt.Printf("%v: %d\n", race, winCount)
		answer *= winCount
	}
	fmt.Println(answer)
}
