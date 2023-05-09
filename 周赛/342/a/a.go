package main

// https://github.com/Ocyss
func findDelayedArrivalTime(arrivalTime int, delayedTime int) (ans int) {
	return (arrivalTime + delayedTime) % 24
}
