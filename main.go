package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("What is your name?")
	var firstName string
	fmt.Scanln(&firstName)
	currentTime := time.Now()
	dayPart := getDaypart(currentTime)
	fmt.Println("good", dayPart, firstName)
	fmt.Println("The current time is:", currentTime.Format(time.RFC822))

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
