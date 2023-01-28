package main

import (
	"fmt"
	"github.com/deckarep/gosx-notifier"
)

var (
	duration              = 25
	shortBreak            = 5
	longBreak             = 15
	pomodoros             int
	pomodoroSelectionDone bool
)

func main() {
	for {
		if !pomodoroSelectionDone {
			pomodoroSelection()
			n := gosxnotifier.NewNotification("Pomodoro session started! 🕑")
			n.Title = "Pomodoro Timer"
			n.Sound = gosxnotifier.Basso
			n.AppIcon = "pomodoro.png"
			n.Push()
		} else if pomodoroSelectionDone == true {

			fmt.Println("1. Start timer with default settings! 😊")
			fmt.Println("2. Start timer with custom settings! 🤓")
			fmt.Println("3. Exit! 😢")

			var input int
			fmt.Print("Enter your choice: ")
			fmt.Scan(&input)

			switch input {
			case 1:
				startTimer(duration, shortBreak, longBreak)

			case 2:
				setCustomTimer()
				startTimer(duration, shortBreak, longBreak)

			case 3:
				fmt.Println("die stupid! 😡")
				return
			default:
				fmt.Println("Invalid input, try again 😤")
			}
		}
	}
}

func sendNotificationEnd() {
	n := gosxnotifier.NewNotification("Timer ended! 🕑")
	n.Title = "Pomodoro Timer"
	n.Sound = gosxnotifier.Basso
	n.AppIcon = "pomodoro.png"
	n.Push()
}
