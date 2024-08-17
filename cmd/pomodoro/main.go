package main

import (
	"fmt"
	"github.com/deckarep/gosx-notifier"
	"pomodoro/internal"
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
			internal.PomodoroSelection()

			// TODO: fix the loop after selecting the amount of pomodoros

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
				internal.StartTimer(pomodoros, duration, shortBreak, longBreak, sendNotificationEnd)

			case 2:
				internal.SetCustomTimer()
				internal.StartTimer(pomodoros, duration, shortBreak, longBreak, sendNotificationEnd)

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
	err := n.Push()
	if err != nil {
		return
	}
}
