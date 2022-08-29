package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	dayPart := getDaypart(currentTime)
	fmt.Println("good", dayPart)
	fmt.Println("The current time is:", currentTime)
}

func getDaypart(datetime time.Time) string {
	if datetime.Hour() < 12 {
		return "morning"
	}
	if datetime.Hour() < 18 {
		return "afternoon"
	}
	return "evening"
}
